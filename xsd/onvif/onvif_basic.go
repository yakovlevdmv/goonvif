package onvif

import (
	"github.com/use-go/goonvif/xsd"
)

type LocalOrientation struct {
	Lon       xsd.Double `xml:"lon,attr"`
	Lat       xsd.Double `xml:"lat,attr"`
	Elevation xsd.Float  `xml:"elevation,attr"`
}

type LocalLocation struct {
	X xsd.Float `xml:"x,attr"`
	Y xsd.Float `xml:"y,attr"`
	Z xsd.Float `xml:"z,attr"`
}

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

type Time struct {
	Hour   xsd.Int `xml:"onvif:Hour"`
	Minute xsd.Int `xml:"onvif:Minute"`
	Second xsd.Int `xml:"onvif:Second"`
}

type Date struct {
	Year  xsd.Int `xml:"onvif:Year"`
	Month xsd.Int `xml:"onvif:Month"`
	Day   xsd.Int `xml:"onvif:Day"`
}
