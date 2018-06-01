package onvif

import (
	"github.com/use-go/goonvif/xsd"
)

//LocalOrientation type
type LocalOrientation struct {
	Lon       xsd.Double `xml:"lon,attr"`
	Lat       xsd.Double `xml:"lat,attr"`
	Elevation xsd.Float  `xml:"elevation,attr"`
}

//LocalLocation type
type LocalLocation struct {
	X xsd.Float `xml:"x,attr"`
	Y xsd.Float `xml:"y,attr"`
	Z xsd.Float `xml:"z,attr"`
}

//GeoOrientation type
type GeoOrientation struct {
	Roll  xsd.Float `xml:"roll,attr"`
	Pitch xsd.Float `xml:"pitch,attr"`
	Yaw   xsd.Float `xml:"yaw,attr"`
}

//DateTime ...
type DateTime struct {
	Time Time `xml:"onvif:Time"`
	Date Date `xml:"onvif:Date"`
}

//Time type
type Time struct {
	Hour   xsd.Int `xml:"onvif:Hour"`
	Minute xsd.Int `xml:"onvif:Minute"`
	Second xsd.Int `xml:"onvif:Second"`
}

//Date type
type Date struct {
	Year  xsd.Int `xml:"onvif:Year"`
	Month xsd.Int `xml:"onvif:Month"`
	Day   xsd.Int `xml:"onvif:Day"`
}

//ErrorType type
type ErrorType struct {
	Dialect xsd.AnyURI  `xml:"dialect,attr"`
	Value   xsd.AnyType `xml:",chardata"`
}

//BaseFaultType type
type BaseFaultType struct {
	Timestamp  xsd.DateTime          `xml:"Timestamp"`
	Originator EndpointReferenceType `xml:"Originator"` //type="wsa:EndpointReferenceType"
	ErrorCode  ErrorType             `xml:"ErrorCode"`
}

//AttributedURIType in ws-addr
type AttributedURIType xsd.AnyURI //wsa https://www.w3.org/2005/08/addressing/ws-addr.xsd

//EndpointReferenceType in ws-addr
type EndpointReferenceType struct { //wsa http://www.w3.org/2005/08/addressing/ws-addr.xsd
	Address AttributedURIType `xml:"wsa5:Address"`
	//	ReferenceParameters ReferenceParametersType `xml:"wsa5:ReferenceParameters,omit"`
	//	Metadata            MetadataType            `xml:"wsa5:Metadata,omit"`
}
