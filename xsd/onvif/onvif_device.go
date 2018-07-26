package onvif

import (
	"github.com/use-go/goonvif/xsd"
)

//RemoteUser ...
type RemoteUser struct {
	Username           string      `xml:"Username"`
	Password           string      `xml:"Password"`
	UseDerivedPassword xsd.Boolean `xml:"UseDerivedPassword"`
}

//User ...
type User struct {
	Username  string        `xml:"Username"`
	Password  string        `xml:"Password"`
	UserLevel UserLevel     `xml:"UserLevel"`
	Extension UserExtension `xml:"Extension"`
}

type UserLevel xsd.String

type UserExtension xsd.String

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
	GeoLocation      GeoLocation      `xml:"GeoLocation"`
	GeoOrientation   GeoOrientation   `xml:"GeoOrientation"`
	LocalLocation    LocalLocation    `xml:"LocalLocation"`
	LocalOrientation LocalOrientation `xml:"LocalOrientation"`
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

//TimeZone ..
type TimeZone struct {
	TZ xsd.Token `xml:"TZ"`
}

//SystemDateTime ...
type SystemDateTime struct {
	DateTimeType    SetDateTimeType
	DaylightSavings xsd.Boolean
	TimeZone        TimeZone
	UTCDateTime     xsd.DateTime
	LocalDateTime   xsd.DateTime
	Extension       SystemDateTimeExtension
}
