package consumer

import (
	"regexp"
	"testing"
)

func TestGetDeviceIDAndName(t *testing.T) {

	mockedbytesArray := []byte("/onvifAlarmDevice/a58a3c16-8a4d-47a6-ab0c-d43e6fd89255/")

	deviceName, deviceID, err := getDeviceIDAndName(mockedbytesArray)

	if err != nil {
		t.Log(err)
	}
	t.Log(deviceName)
	t.Log(deviceID)

}

func TestHttpRequestCtxPah(t *testing.T) {

	//校验uri格式，需要是系列格式之一,如果格式不符，则为匿名消息
	str1Byte := []byte("/camera-in-gate-onvif-device/a1340608-a8a1-4495-8132-95303580e588")
	str2Byte := []byte("/camera-三楼4#-gate-onvif-device/a1340608-a8a1-4495-8132-95303580e588/")
	str3Byte := []byte("camera-in-gate-onvif-device/a1340608-a8a1-4495-8132-95303580e588/")

	reg := regexp.MustCompile(`(/?)(?P<DeviceName>\S+?)/(?P<DeviceID>[0-9A-Za-z]{8}(-[0-9A-Za-z]{4}){3}-[0-9A-Za-z]{12})(/?)`)
	result := reg.Match(str1Byte)
	t.Log(result)
	result = reg.Match(str2Byte)
	t.Log(result)
	result = reg.Match(str3Byte)
	t.Log(result)

}
