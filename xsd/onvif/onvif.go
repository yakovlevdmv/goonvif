package onvif

import (
	"github.com/use-go/goonvif/xsd"
)

// BUG(r): Enum types implemented as simple string

//TODO: enumerations
//TODO: type <typeName> struct {Any string} convert to type <typeName> AnyType
//TODO: process restrictions

//TODO: Name and similar ones should be examined for the presence of "List of ..." errors
//TODO: view all Extensions (Any string)
//TODO: what to do with xs: any = Any
//TODO: IntList and the like. Check whether you need a slice. Change to slice
//TODO: to see if it is possible to replace StreamType and similar types with the types that are typed in
//TODO: to test the type of VideoSourceMode because of the Description-a
//TODO: in the documentation to describe that Capabilities is repeated for each service, so each has its Capabilities MediaCapabilities)
//TODO: AuxiliaryData and other simpleTypes, how to implement the restriction
//TODO: Name and similar ones should be examined for the presence of "List of ..." errors
//TODO: Add in buit-in

//ContentType from string
type ContentType string // minLength value="3"

//ReferenceToken from xsd.string
type ReferenceToken xsd.String

//Name from xsd.string
type Name xsd.String

type Color struct {
	X          float64    `xml:"X,attr"`
	Y          float64    `xml:"Y,attr"`
	Z          float64    `xml:"Z,attr"`
	Colorspace xsd.AnyURI `xml:"Colorspace,attr"`
}

type BacklightCompensation struct {
	Mode  BacklightCompensationMode
	Level float64
}

type BacklightCompensationMode xsd.String

type Exposure struct {
	Mode            ExposureMode
	Priority        ExposurePriority
	Window          Rectangle
	MinExposureTime float64
	MaxExposureTime float64
	MinGain         float64
	MaxGain         float64
	MinIris         float64
	MaxIris         float64
	ExposureTime    float64
	Gain            float64
	Iris            float64
}

type ExposureMode xsd.String

type ExposurePriority xsd.String

type ToneCompensationExtension xsd.AnyType

type Defogging struct {
	Mode      string
	Level     float64
	Extension DefoggingExtension
}

type DefoggingExtension xsd.AnyType

type NoiseReduction struct {
	Level float64 `xml:"onvif:Level"`
}

type Rotate struct {
	Mode      RotateMode      `xml:"onvif:Mode"`
	Degree    xsd.Int         `xml:"onvif:Degree"`
	Extension RotateExtension `xml:"onvif:Extension"`
}

type RotateMode xsd.String

type RotateExtension xsd.AnyType

type LensDescription struct {
	FocalLength float64        `xml:"FocalLength,attr"`
	Offset      LensOffset     `xml:"onvif:Offset"`
	Projection  LensProjection `xml:"onvif:Projection"`
	XFactor     float64        `xml:"onvif:XFactor"`
}

type LensOffset struct {
	X float64 `xml:"x,attr"`
	Y float64 `xml:"y,attr"`
}

type LensProjection struct {
	Angle         float64 `xml:"onvif:Angle"`
	Radius        float64 `xml:"onvif:Radius"`
	Transmittance float64 `xml:"onvif:Transmittance"`
}
type SceneOrientation struct {
	Mode        SceneOrientationMode `xml:"onvif:Mode"`
	Orientation xsd.String           `xml:"onvif:Orientation"`
}

type SceneOrientationMode xsd.String

type Config struct {
	Name       string    `xml:"Name,attr"`
	Type       xsd.QName `xml:"Type,attr"`
	Parameters ItemList  `xml:"onvif:Parameters"`
}

type ItemList struct {
	SimpleItem  SimpleItem        `xml:"onvif:SimpleItem"`
	ElementItem ElementItem       `xml:"onvif:ElementItem"`
	Extension   ItemListExtension `xml:"onvif:Extension"`
}

type SimpleItem struct {
	Name  string            `xml:"onvif:Name,attr"`
	Value xsd.AnySimpleType `xml:"onvif:Value,attr"`
}

type ElementItem struct {
	Name string `xml:"Name,attr"`
}
type ItemListExtension xsd.AnyType

type RuleEngineConfiguration struct {
	Rule      Config                           `xml:"onvif:Rule"`
	Extension RuleEngineConfigurationExtension `xml:"onvif:Extension"`
}
type RuleEngineConfigurationExtension xsd.AnyType

type EFlip struct {
	Mode EFlipMode `xml:"onvif:Mode"`
}
type EFlipMode xsd.String

type Reverse struct {
	Mode ReverseMode `xml:"onvif:Mode"`
}
type ReverseMode xsd.String

type RotateOptions struct {
	Mode       RotateMode
	DegreeList IntList
	Extension  RotateOptionsExtension
}

type IntList struct {
	Items []int
}

type RotateOptionsExtension xsd.AnyType

type MetadataConfigurationOptionsExtension struct {
	CompressionType string
	Extension       MetadataConfigurationOptionsExtension2
}

type MetadataConfigurationOptionsExtension2 xsd.AnyType

//enum
type TransportProtocol xsd.String

type MediaUri struct {
	Uri                 xsd.AnyURI
	InvalidAfterConnect bool
	InvalidAfterReboot  bool
	Timeout             xsd.Duration
}

type EncodingTypes struct {
	EncodingTypes []string
}

type Description struct {
	Description string
}
type ColorOptions struct {
	ColorList       Color
	ColorspaceRange ColorspaceRange
}
type ColorspaceRange struct {
	X          FloatRange
	Y          FloatRange
	Z          FloatRange
	Colorspace xsd.AnyURI
}

type StringAttrList struct {
	AttrList []string
}

type IntAttrList struct {
	IntAttrList []int
}

type DurationRange struct {
	Min xsd.Duration
	Max xsd.Duration
}
type PTControlDirectionOptions struct {
	EFlip     EFlipOptions
	Reverse   ReverseOptions
	Extension PTControlDirectionOptionsExtension
}
type EFlipOptions struct {
	Mode      EFlipMode
	Extension EFlipOptionsExtension
}

type EFlipOptionsExtension xsd.AnyType

type ReverseOptions struct {
	Mode      ReverseMode
	Extension ReverseOptionsExtension
}

type ReverseOptionsExtension xsd.AnyType

type PTControlDirectionOptionsExtension xsd.AnyType
type MoveStatus struct {
	Status string
}

type SystemDateTimeExtension xsd.AnyType

type FactoryDefaultType xsd.String

type AttachmentData struct {
	ContentType ContentType `xml:"contentType,attr"`
	Include     Include     `xml:"inc:Include"`
}
type Include struct {
	Href xsd.AnyURI `xml:"href,attr"`
}
type BackupFile struct {
	Name string         `xml:"onvif:Name"`
	Data AttachmentData `xml:"onvif:Data"`
}
type SystemLogType xsd.String
type SystemLog struct {
	Binary AttachmentData
	String string
}
type SupportInformation struct {
	Binary AttachmentData
	String string
}
type Scope struct {
	ScopeDef  ScopeDefinition
	ScopeItem xsd.AnyURI
}
type ScopeDefinition xsd.String

type CapabilityCategory xsd.String

//Capabilities of device
type Capabilities struct {
	Analytics AnalyticsCapabilities
	Device    DeviceCapabilities
	Events    EventCapabilities
	Imaging   ImagingCapabilities
	Media     MediaCapabilities
	PTZ       PTZCapabilities
	Extension CapabilitiesExtension
}

//DeviceCapabilities Check
type DeviceCapabilities struct {
	XAddr     xsd.AnyURI
	Network   NetworkCapabilities
	System    SystemCapabilities
	IO        IOCapabilities
	Security  SecurityCapabilities
	Extension DeviceCapabilitiesExtension
}

//SystemCapabilities check
type SystemCapabilities struct {
	DiscoveryResolve  xsd.Boolean
	DiscoveryBye      xsd.Boolean
	RemoteDiscovery   xsd.Boolean
	SystemBackup      xsd.Boolean
	SystemLogging     xsd.Boolean
	FirmwareUpgrade   xsd.Boolean
	SupportedVersions OnvifVersion
	Extension         SystemCapabilitiesExtension
}

type SystemCapabilitiesExtension struct {
	HttpFirmwareUpgrade    xsd.Boolean
	HttpSystemBackup       xsd.Boolean
	HttpSystemLogging      xsd.Boolean
	HttpSupportInformation xsd.Boolean
	Extension              SystemCapabilitiesExtension2
}

type SystemCapabilitiesExtension2 xsd.AnyType

type IOCapabilities struct {
	InputConnectors int
	RelayOutputs    int
	Extension       IOCapabilitiesExtension
}

type IOCapabilitiesExtension struct {
	Auxiliary         xsd.Boolean
	AuxiliaryCommands AuxiliaryData
	Extension         IOCapabilitiesExtension2
}

type IOCapabilitiesExtension2 xsd.AnyType

type DeviceCapabilitiesExtension xsd.AnyType

type MediaCapabilities struct {
	XAddr                 xsd.AnyURI
	StreamingCapabilities RealTimeStreamingCapabilities
	Extension             MediaCapabilitiesExtension
}

type RealTimeStreamingCapabilities struct {
	RTPMulticast xsd.Boolean
	RTP_TCP      xsd.Boolean
	RTP_RTSP_TCP xsd.Boolean
	Extension    RealTimeStreamingCapabilitiesExtension
}

type RealTimeStreamingCapabilitiesExtension xsd.AnyType

type MediaCapabilitiesExtension struct {
	ProfileCapabilities ProfileCapabilities
}

type ProfileCapabilities struct {
	MaximumNumberOfProfiles int
}

type CapabilitiesExtension struct {
	DeviceIO        DeviceIOCapabilities
	Display         DisplayCapabilities
	Recording       RecordingCapabilities
	Search          SearchCapabilities
	Replay          ReplayCapabilities
	Receiver        ReceiverCapabilities
	AnalyticsDevice AnalyticsDeviceCapabilities
	Extensions      CapabilitiesExtension2
}

type DeviceIOCapabilities struct {
	XAddr        xsd.AnyURI
	VideoSources int
	VideoOutputs int
	AudioSources int
	AudioOutputs int
	RelayOutputs int
}

type DisplayCapabilities struct {
	XAddr       xsd.AnyURI
	FixedLayout xsd.Boolean
}

type RecordingCapabilities struct {
	XAddr              xsd.AnyURI
	ReceiverSource     xsd.Boolean
	MediaProfileSource xsd.Boolean
	DynamicRecordings  xsd.Boolean
	DynamicTracks      xsd.Boolean
	MaxStringLength    int
}

type SearchCapabilities struct {
	XAddr          xsd.AnyURI
	MetadataSearch xsd.Boolean
}

type ReplayCapabilities struct {
	XAddr xsd.AnyURI
}

type CapabilitiesExtension2 xsd.AnyType

//TODO: attribite <xs:attribute ref="xmime:contentType" use="optional"/>
type BinaryData struct {
	X    ContentType      `xml:"xmime:contentType,attr"`
	Data xsd.Base64Binary `xml:"onvif:Data"`
}

type RelayOutput struct {
	DeviceEntity
	Properties RelayOutputSettings
}

type RelayOutputSettings struct {
	Mode      RelayMode      `xml:"onvif:Mode"`
	DelayTime xsd.Duration   `xml:"onvif:DelayTime"`
	IdleState RelayIdleState `xml:"onvif:IdleState"`
}

//TODO:enumeration
type RelayIdleState xsd.String

//TODO: enumeration
type RelayMode xsd.String

//TODO: enumeration
type RelayLogicalState xsd.String
