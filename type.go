package goonvif

import "net/http"

//deviceInfo struct contains general information about ONVIF device
type deviceInfo struct {
	Manufacturer    string
	Model           string
	FirmwareVersion string
	SerialNumber    string
	HardwareID      string
}

//IOnvif Exported Service
type IOnvif interface {
	Authenticate(username, password string)
	GetEndpoint(name string) string
	CallMethod(method interface{}, headerFields map[string]string) (*http.Response, error)
	GetServices() map[string]string
}

//Device for a new device of onvif and deviceInfo
//struct represents an abstract ONVIF device.
//It contains methods, which helps to communicate with ONVIF device
type Device struct {
	xaddr     string
	login     string
	password  string
	endpoints map[string]string
	info      deviceInfo
}
