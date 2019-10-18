package goonvif

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	"github.com/beevik/etree"
	Device "github.com/larryhu/goonvif/device"
	"github.com/larryhu/goonvif/discover"
	"github.com/larryhu/goonvif/networking"
	"github.com/larryhu/goonvif/soap"
)

// Xlmns xlmns
var Xlmns = map[string]string{
	"onvif":   "http://www.onvif.org/ver10/schema",
	"tds":     "http://www.onvif.org/ver10/device/wsdl",
	"trt":     "http://www.onvif.org/ver10/media/wsdl",
	"tev":     "http://www.onvif.org/ver10/events/wsdl",
	"tptz":    "http://www.onvif.org/ver20/ptz/wsdl",
	"timg":    "http://www.onvif.org/ver20/imaging/wsdl",
	"tan":     "http://www.onvif.org/ver20/analytics/wsdl",
	"xmime":   "http://www.w3.org/2005/05/xmlmime",
	"wsnt":    "http://docs.oasis-open.org/wsn/b-2",
	"xop":     "http://www.w3.org/2004/08/xop/include",
	"wsa":     "http://www.w3.org/2005/08/addressing",
	"wstop":   "http://docs.oasis-open.org/wsn/t-1",
	"wsntw":   "http://docs.oasis-open.org/wsn/bw-2",
	"wsrf-rw": "http://docs.oasis-open.org/wsrf/rw-2",
	"wsaw":    "http://www.w3.org/2006/05/addressing/wsdl",
}

type DeviceType int

const (
	NVD DeviceType = iota
	NVS
	NVA
	NVT
)

func (devType DeviceType) String() string {
	stringRepresentation := []string{
		"NetworkVideoDisplay",
		"NetworkVideoStorage",
		"NetworkVideoAnalytics",
		"NetworkVideoTransmitter",
	}
	i := uint8(devType)
	switch {
	case i <= uint8(NVT):
		return stringRepresentation[i]
	default:
		return strconv.Itoa(int(i))
	}
}

//deviceInfo struct contains general information about ONVIF device
type deviceInfo struct {
	Manufacturer    string
	Model           string
	FirmwareVersion string
	SerialNumber    string
	HardwareId      string
}

//deviceInfo struct represents an abstract ONVIF device.
//It contains methods, which helps to communicate with ONVIF device
type device struct {
	xaddr    string
	login    string
	password string

	endpoints map[string]string
	info      deviceInfo
}

func (dev *device) GetServices() map[string]string {
	return dev.endpoints
}

func readResponse(resp *http.Response) string {
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(b)
}

// GetAvailableDevicesAtSpecificEthernetInterface available devices
func GetAvailableDevicesAtSpecificEthernetInterface(interfaceName string) ([]device, error) {
	/*
		Call an WS-Discovery Probe Message to Discover NVT type Devices
	*/
	devices := discover.SendProbe(interfaceName, nil, []string{"dn:" + NVT.String()}, map[string]string{"dn": "http://www.onvif.org/ver10/network/wsdl"})
	nvtDevices := make([]device, 0)
	for _, j := range devices {
		doc := etree.NewDocument()
		if err := doc.ReadFromString(j); err != nil {
			return nil, err
		}
		endpoints := doc.Root().FindElements("./Body/ProbeMatches/ProbeMatch/XAddrs")
		for _, xaddr := range endpoints {
			xaddr := strings.Split(strings.Split(xaddr.Text(), " ")[0], "/")[2]
			c := 0
			for c = 0; c < len(nvtDevices); c++ {
				if nvtDevices[c].xaddr == xaddr {
					break
				}
			}
			if c < len(nvtDevices) {
				continue
			}
			dev, err := NewDevice(strings.Split(xaddr, " ")[0])
			if err != nil {
				return nil, fmt.Errorf("error %s %s", xaddr, err)
			}
			nvtDevices = append(nvtDevices, *dev)
		}
	}
	return nvtDevices, nil
}

func (dev *device) getSupportedServices(resp *http.Response) {
	doc := etree.NewDocument()

	data, _ := ioutil.ReadAll(resp.Body)

	if err := doc.ReadFromBytes(data); err != nil {
		return
	}
	services := doc.FindElements("./Envelope/Body/GetCapabilitiesResponse/Capabilities/*/XAddr")
	for _, j := range services {
		// TODO suggest return error as result.
		URL, _ := url.Parse(j.Text())
		URL.Host = dev.xaddr
		dev.addEndpoint(j.Parent().Tag, URL.String())
	}
}

//NewDevice function construct a ONVIF Device entity
func NewDevice(xaddr string) (*device, error) {
	dev := new(device)
	dev.xaddr = xaddr
	dev.endpoints = make(map[string]string)
	dev.addEndpoint("Device", "http://"+xaddr+"/onvif/device_service")

	getCapabilities := Device.GetCapabilities{Category: "All"}

	resp, err := dev.CallMethod(getCapabilities)
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("camera is not available at %s or it does not support ONVIF services", xaddr)
	}

	dev.getSupportedServices(resp)
	return dev, nil
}

func (dev *device) addEndpoint(Key, Value string) {
	dev.endpoints[Key] = Value
}

//Authenticate function authenticate client in the ONVIF Device.
//Function takes <username> and <password> params.
//You should use this function to allow authorized requests to the ONVIF Device
//To change auth data call this function again.
func (dev *device) Authenticate(username, password string) {
	dev.login = username
	dev.password = password
}

//GetEndpoint returns specific ONVIF service endpoint address
func (dev *device) GetEndpoint(name string) string {
	return dev.endpoints[name]
}

func buildMethodSOAP(msg []byte) (soap.SoapMessage, error) {
	doc := etree.NewDocument()
	if err := doc.ReadFromBytes(msg); err != nil {

		return "", err
	}
	element := doc.Root()

	soap := soap.NewEmptySOAP()
	soap.AddBodyContent(element)
	//soap.AddRootNamespace("onvif", "http://www.onvif.org/ver10/device/wsdl")

	return soap, nil
}

//CallMethod functions call an method, defined <method> struct.
//You should use Authenticate method to call authorized requests.
func (dev device) CallMethod(method interface{}) (*http.Response, error) {
	pkgPath := strings.Split(reflect.TypeOf(method).PkgPath(), "/")
	pkg := pkgPath[len(pkgPath)-1]

	var endpoint string
	switch pkg {
	case "Device":
	case "Event":
	case "Imaging":
	case "Media":
	case "PTZ":
		endpoint = dev.endpoints[pkg]
	default:
		return nil, fmt.Errorf("Endpoint not found [%s]", pkg)
	}

	return dev.callMethod(endpoint, method)
}

//CallMethod functions call an method, defined <method> struct with authentication data
func (dev device) callMethod(endpoint string, method interface{}) (*http.Response, error) {
	/*
		Converting <method> struct to xml string representation
	*/
	output, err := xml.MarshalIndent(method, "  ", "    ")
	if err != nil {
		return nil, err
	}

	/*
		Build an SOAP request with <method>
	*/
	soap, err := buildMethodSOAP(output)
	if err != nil {
		return nil, err
	}

	/*
		Adding namespaces and WS-Security headers
	*/
	soap.AddRootNamespaces(Xlmns)

	if dev.login != "" && dev.password != "" {
		soap.AddWSSecurity(dev.login, dev.password)
	}

	/*
		Sending request and returns the response
	*/
	return networking.SendSoap(endpoint, soap.Bytes())
}
