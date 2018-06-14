package helper

import (
	"errors"

	"github.com/use-go/goonvif"
)

//CallMethod device call mehod(request) with needed header(headerFiled)
//select body content(xmlTagInbody) as a unmarshaled result(response)
func CallMethod(device goonvif.IOnvif, headerFileds map[string]string, request interface{}, xmlTagInbody string, response interface{}) error {
	if device == nil {
		return errors.New("device is nil")
	}
	resp, err := device.CallMethod(request, headerFileds)

	if err != nil {
		return err
	}

	return Unmarshal(resp.Body, xmlTagInbody, response)
}
