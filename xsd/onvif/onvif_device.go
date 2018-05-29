package onvif

import (
	"github.com/use-go/goonvif/xsd"
)

type DeviceEntity struct {
	Token ReferenceToken `xml:"token,attr"`
}
type SystemLogUriList struct {
	SystemLog SystemLogUri
}

type SystemLogUri struct {
	Type SystemLogType
	Uri  xsd.AnyURI
}

type LocationEntity struct {
	Entity           xsd.String       `xml:"Entity,attr"`
	Token            ReferenceToken   `xml:"Token,attr"`
	Fixed            xsd.Boolean      `xml:"Fixed,attr"`
	GeoSource        xsd.AnyURI       `xml:"GeoSource,attr"`
	AutoGeo          xsd.Boolean      `xml:"AutoGeo,attr"`
	GeoLocation      GeoLocation      `xml:"onvif:GeoLocation"`
	GeoOrientation   GeoOrientation   `xml:"onvif:GeoOrientation"`
	LocalLocation    LocalLocation    `xml:"onvif:LocalLocation"`
	LocalOrientation LocalOrientation `xml:"onvif:LocalOrientation"`
}

type GeoLocation struct {
	Lon       xsd.Double `xml:"lon,attr"`
	Lat       xsd.Double `xml:"lat,attr"`
	Elevation xsd.Float  `xml:"elevation,attr"`
}

//Device
type OnvifVersion struct {
	Major int
	Minor int
}

type SetDateTimeType xsd.String

type TimeZone struct {
	TZ xsd.Token `xml:"onvif:TZ"`
}

type SystemDateTime struct {
	DateTimeType    SetDateTimeType
	DaylightSavings xsd.Boolean
	TimeZone        TimeZone
	UTCDateTime     xsd.DateTime
	LocalDateTime   xsd.DateTime
	Extension       SystemDateTimeExtension
}
