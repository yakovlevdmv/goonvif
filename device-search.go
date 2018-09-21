package goonvif

import (
	"fmt"
	"net"
	"strings"

	"github.com/beevik/etree"
	wsdiscovery "github.com/use-go/goonvif/ws-discovery"
)

//GetAvailableDevicesAtSpecificEthernetInterface ...
func GetAvailableDevicesAtSpecificEthernetInterface(interfaceName string) []Device {
	/*
		Call an ws-discovery Probe Message to Discover NVT type Devices
	*/
	devices, err := wsdiscovery.SendProbe(interfaceName, nil, []string{"dn:" + NVT.String()}, map[string]string{"dn": "http://www.onvif.org/ver10/network/wsdl"})

	if err != nil {
		return []Device{}
	}
	nvtDevices := make([]Device, 0)
	////fmt.Println(devices)
	for _, j := range devices {
		doc := etree.NewDocument()
		if err := doc.ReadFromString(j); err != nil {
			//fmt.Errorf("%s", err.Error())
			return nil
		}
		////fmt.Println(j)
		endpoints := doc.Root().FindElements("./Body/ProbeMatches/ProbeMatch/XAddrs")
		for _, xaddr := range endpoints {
			//fmt.Println(xaddr.Tag,strings.Split(strings.Split(xaddr.Text(), " ")[0], "/")[2] )
			xaddr := strings.Split(strings.Split(xaddr.Text(), " ")[0], "/")[2]
			fmt.Println(xaddr)
			c := 0
			for c = 0; c < len(nvtDevices); c++ {
				if nvtDevices[c].xaddr == xaddr {
					fmt.Println(nvtDevices[c].xaddr, "==", xaddr)
					break
				}
			}
			if c < len(nvtDevices) {
				continue
			}
			dev, err := NewDevice(strings.Split(xaddr, " ")[0])
			//fmt.Println(dev)
			if err != nil {
				fmt.Println("Error", xaddr)
				fmt.Println(err)
				continue
			} else {
				nvtDevices = append(nvtDevices, *dev)
			}
		}
	}
	return nvtDevices
}

//GetAvailableDevicesBySpecificEthernetInterface ...
func GetAvailableDevicesBySpecificEthernetInterface(requestID string, netInterface *net.Interface) []Device {
	/*
		Call an ws-discovery Probe Message to Discover NVT type Devices
	*/
	devices, err := wsdiscovery.SendProbeByInterface(requestID, netInterface, nil, []string{"dn:" + NVT.String()}, nil)

	if err != nil {
		return []Device{}
	}
	nvtDevices := make([]Device, 0)
	for _, j := range devices {
		doc := etree.NewDocument()
		if err := doc.ReadFromString(j); err != nil {
			return nil
		}
		endpoints := doc.Root().FindElements("./Body/ProbeMatches/ProbeMatch/XAddrs")
		for _, xaddr := range endpoints {

			xaddress := xaddr.Text()
			c := 0
			for c = 0; c < len(nvtDevices); c++ {
				if nvtDevices[c].xaddr == xaddress {
					break
				}
			}
			if c > 0 && c < len(nvtDevices) {
				continue
			}

			devIPAddress, devPort := wsdiscovery.ExtractPortAndIP(xaddress)
			dev := Device{
				ipaddress: devIPAddress,
				port:      devPort,
				xaddr:     xaddress,
				login:     "admin",  //only for default
				password:  "123456", //only for default
			}
			nvtDevices = append(nvtDevices, dev)

		}

	}
	return nvtDevices
}
