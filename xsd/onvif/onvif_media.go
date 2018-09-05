package onvif

import (
	"github.com/use-go/goonvif/xsd"
)

type OSDPosConfigurationExtension xsd.AnyType

type OSDReference ReferenceToken

//OSDTextConfiguration for OSD
type OSDTextConfiguration struct {
	IsPersistentText xsd.Boolean                   `xml:"IsPersistentText,attr"`
	Type             xsd.String                    `xml:"onvif:Type"`
	DateFormat       xsd.String                    `xml:"onvif:DateFormat"`
	TimeFormat       xsd.String                    `xml:"onvif:TimeFormat"`
	FontSize         xsd.Int                       `xml:"onvif:FontSize"`
	FontColor        OSDColor                      `xml:"onvif:FontColor"`
	BackgroundColor  OSDColor                      `xml:"onvif:BackgroundColor"`
	PlainText        xsd.String                    `xml:"onvif:PlainText"`
	Extension        OSDTextConfigurationExtension `xml:"onvif:Extension"`
}

//OSDColor for OSD
type OSDColor struct {
	Transparent int   `xml:"Transparent,attr"`
	Color       Color `xml:"onvif:Color"`
}

type OSDTextConfigurationExtension xsd.AnyType

type OSDImgConfiguration struct {
	ImgPath   xsd.AnyURI                   `xml:"onvif:ImgPath"`
	Extension OSDImgConfigurationExtension `xml:"onvif:Extension"`
}

type OSDImgConfigurationExtension xsd.AnyType

type OSDConfigurationExtension xsd.AnyType

type OSDType xsd.String

type OSDConfiguration struct {
	DeviceEntity                  `xml:"token,attr"`
	VideoSourceConfigurationToken OSDReference              `xml:"onvif:VideoSourceConfigurationToken"`
	Type                          OSDType                   `xml:"onvif:Type"`
	Position                      OSDPosConfiguration       `xml:"onvif:Position"`
	TextString                    OSDTextConfiguration      `xml:"onvif:TextString"`
	Image                         OSDImgConfiguration       `xml:"onvif:Image"`
	Extension                     OSDConfigurationExtension `xml:"onvif:Extension"`
}

type OSDPosConfiguration struct {
	Type      string                       `xml:"onvif:Type"`
	Pos       Vector                       `xml:"onvif:Pos"`
	Extension OSDPosConfigurationExtension `xml:"onvif:Extension"`
}

type VideoSourceModeExtension xsd.AnyType

type OSDConfigurationOptions struct {
	MaximumNumberOfOSDs MaximumNumberOfOSDs
	Type                OSDType
	PositionOption      string
	TextOption          OSDTextOptions
	ImageOption         OSDImgOptions
	Extension           OSDConfigurationOptionsExtension
}

type MaximumNumberOfOSDs struct {
	Total       int `xml:"Total,attr"`
	Image       int `xml:"Image,attr"`
	PlainText   int `xml:"PlainText,attr"`
	Date        int `xml:"Date,attr"`
	Time        int `xml:"Time,attr"`
	DateAndTime int `xml:"DateAndTime,attr"`
}

type OSDTextOptions struct {
	Type            string
	FontSizeRange   IntRange
	DateFormat      string
	TimeFormat      string
	FontColor       OSDColorOptions
	BackgroundColor OSDColorOptions
	Extension       OSDTextOptionsExtension
}

type OSDColorOptions struct {
	Color       ColorOptions
	Transparent IntRange
	Extension   OSDColorOptionsExtension
}

type OSDColorOptionsExtension xsd.AnyType

type OSDTextOptionsExtension xsd.AnyType

type OSDImgOptions struct {
	FormatsSupported StringAttrList `xml:"FormatsSupported,attr"`
	MaxSize          int            `xml:"MaxSize,attr"`
	MaxWidth         int            `xml:"MaxWidth,attr"`
	MaxHeight        int            `xml:"MaxHeight,attr"`

	ImagePath xsd.AnyURI
	Extension OSDImgOptionsExtension
}

type OSDImgOptionsExtension xsd.AnyType

type OSDConfigurationOptionsExtension xsd.AnyType

type VideoSource struct {
	DeviceEntity
	Framerate  float64
	Resolution VideoResolution
	Imaging    ImagingSettings
	Extension  VideoSourceExtension
}

//VideoResolution of video
type VideoResolution struct {
	Width  xsd.Int `xml:"Width"`
	Height xsd.Int `xml:"Height"`
}

type VideoSourceExtension struct {
	Imaging   ImagingSettings20     `xml:"Imaging"`
	Extension VideoSourceExtension2 `xml:"Extension"`
}

type VideoSourceExtension2 xsd.AnyType

type AudioSource struct {
	DeviceEntity
	Channels int
}

type AudioOutput struct {
	DeviceEntity
}

//Profile for different stream
type Profile struct {
	Token                       ReferenceToken              `xml:"token,attr"`
	Fixed                       bool                        `xml:"fixed,attr"`
	Name                        Name                        `xml:"Name"`
	VideoSourceConfiguration    VideoSourceConfiguration    `xml:"VideoSourceConfiguration"`
	AudioSourceConfiguration    AudioSourceConfiguration    `xml:"AudioSourceConfiguration"`
	VideoEncoderConfiguration   VideoEncoderConfiguration   `xml:"VideoEncoderConfiguration"`
	AudioEncoderConfiguration   AudioEncoderConfiguration   `xml:"AudioEncoderConfiguration"`
	VideoAnalyticsConfiguration VideoAnalyticsConfiguration `xml:"VideoAnalyticsConfiguration"`
	PTZConfiguration            PTZConfiguration            `xml:"PTZConfiguration"`
	MetadataConfiguration       MetadataConfiguration       `xml:"MetadataConfiguration"`
	Extension                   ProfileExtension            `xml:"ProfileExtension"`
}

//VideoSourceConfiguration ...
type VideoSourceConfiguration struct {
	Token               ReferenceToken                    `xml:"token,attr"`
	ViewMode            xsd.String                        `xml:"ViewMode,attr"`
	ConfigurationEntity ConfigurationEntity               `xml:"ConfigurationEntity"`
	SourceToken         ReferenceToken                    `xml:"SourceToken"`
	Bounds              IntRectangle                      `xml:"Bounds"`
	Extension           VideoSourceConfigurationExtension `xml:"Extension"`
}

//ConfigurationEntity ...
type ConfigurationEntity struct {
	Token    ReferenceToken `xml:"token,attr"`
	Name     Name           `xml:"Name"`
	UseCount int            `xml:"UseCount"`
}

//VideoSourceConfigurationExtension ...
type VideoSourceConfigurationExtension struct {
	Rotate    Rotate                             `xml:"Rotate"`
	Extension VideoSourceConfigurationExtension2 `xml:"Extension"`
}

type VideoSourceConfigurationExtension2 struct {
	LensDescription  LensDescription  `xml:"LensDescription"`
	SceneOrientation SceneOrientation `xml:"SceneOrientation"`
}

//AudioSourceConfiguration ...
type AudioSourceConfiguration struct {
	ConfigurationEntity
	SourceToken ReferenceToken `xml:"SourceToken"`
}

//VideoEncoderConfiguration ..
type VideoEncoderConfiguration struct {
	ConfigurationEntity
	Encoding       VideoEncoding          `xml:"Encoding"`
	Resolution     VideoResolution        `xml:"Resolution"`
	Quality        xsd.Double             `xml:"Quality"`
	RateControl    VideoRateControl       `xml:"RateControl"`
	MPEG4          Mpeg4Configuration     `xml:"MPEG4"`
	H264           H264Configuration      `xml:"H264"`
	Multicast      MulticastConfiguration `xml:"Multicast"`
	SessionTimeout xsd.Duration           `xml:"SessionTimeout"`
}

type VideoEncoding xsd.String

type VideoRateControl struct {
	FrameRateLimit   xsd.Int `xml:"FrameRateLimit"`
	EncodingInterval xsd.Int `xml:"EncodingInterval"`
	BitrateLimit     xsd.Int `xml:"BitrateLimit"`
}

type Mpeg4Configuration struct {
	GovLength    xsd.Int      `xml:"GovLength"`
	Mpeg4Profile Mpeg4Profile `xml:"Mpeg4Profile"`
}

type Mpeg4Profile xsd.String

type H264Configuration struct {
	GovLength   xsd.Int     `xml:"GovLength"`
	H264Profile H264Profile `xml:"H264Profile"`
}

//H264Profile
type H264Profile xsd.String

type MulticastConfiguration struct {
	Address   IPAddress   `xml:"Address"`
	Port      int         `xml:"Port"`
	TTL       int         `xml:"TTL"`
	AutoStart xsd.Boolean `xml:"AutoStart"`
}

type AudioEncoderConfiguration struct {
	ConfigurationEntity
	Encoding       AudioEncoding          `xml:"Encoding"`
	Bitrate        int                    `xml:"Bitrate"`
	SampleRate     int                    `xml:"SampleRate"`
	Multicast      MulticastConfiguration `xml:"Multicast"`
	SessionTimeout xsd.Duration           `xml:"SessionTimeout"`
}

type AudioEncoding xsd.String

type VideoAnalyticsConfiguration struct {
	ConfigurationEntity
	AnalyticsEngineConfiguration AnalyticsEngineConfiguration `xml:"AnalyticsEngineConfiguration"`
	RuleEngineConfiguration      RuleEngineConfiguration      `xml:"RuleEngineConfiguration"`
}

type AnalyticsEngineConfiguration struct {
	AnalyticsModule Config                                `xml:"AnalyticsModule"`
	Extension       AnalyticsEngineConfigurationExtension `xml:"Extension"`
}

type MetadataConfigurationExtension xsd.AnyType

type ProfileExtension struct {
	AudioOutputConfiguration  AudioOutputConfiguration
	AudioDecoderConfiguration AudioDecoderConfiguration
	Extension                 ProfileExtension2
}

type AudioOutputConfiguration struct {
	ConfigurationEntity
	OutputToken ReferenceToken `xml:"OutputToken"`
	SendPrimacy xsd.AnyURI     `xml:"SendPrimacy"`
	OutputLevel int            `xml:"OutputLevel"`
}

type AudioDecoderConfiguration struct {
	ConfigurationEntity
}

type ProfileExtension2 xsd.AnyType

type VideoSourceConfigurationOptions struct {
	MaximumNumberOfProfiles    int `xml:"MaximumNumberOfProfiles,attr"`
	BoundsRange                IntRectangleRange
	VideoSourceTokensAvailable ReferenceToken
	Extension                  VideoSourceConfigurationOptionsExtension
}

type VideoSourceConfigurationOptionsExtension struct {
	Rotate    RotateOptions
	Extension VideoSourceConfigurationOptionsExtension2
}

type VideoSourceConfigurationOptionsExtension2 struct {
	SceneOrientationMode SceneOrientationMode
}

type VideoEncoderConfigurationOptions struct {
	QualityRange IntRange
	JPEG         JpegOptions
	MPEG4        Mpeg4Options
	H264         H264Options
	Extension    VideoEncoderOptionsExtension
}

type JpegOptions struct {
	ResolutionsAvailable  VideoResolution
	FrameRateRange        IntRange
	EncodingIntervalRange IntRange
}

type Mpeg4Options struct {
	ResolutionsAvailable   VideoResolution
	GovLengthRange         IntRange
	FrameRateRange         IntRange
	EncodingIntervalRange  IntRange
	Mpeg4ProfilesSupported Mpeg4Profile
}

type H264Options struct {
	ResolutionsAvailable  VideoResolution
	GovLengthRange        IntRange
	FrameRateRange        IntRange
	EncodingIntervalRange IntRange
	H264ProfilesSupported H264Profile
}

type VideoEncoderOptionsExtension struct {
	JPEG      JpegOptions2
	MPEG4     Mpeg4Options2
	H264      H264Options2
	Extension VideoEncoderOptionsExtension2
}

type JpegOptions2 struct {
	JpegOptions
	BitrateRange IntRange
}

type Mpeg4Options2 struct {
	Mpeg4Options
	BitrateRange IntRange
}

type H264Options2 struct {
	H264Options
	BitrateRange IntRange
}

type VideoEncoderOptionsExtension2 xsd.AnyType

type AudioSourceConfigurationOptions struct {
	InputTokensAvailable ReferenceToken
	Extension            AudioSourceOptionsExtension
}

type AudioSourceOptionsExtension xsd.AnyType

type AudioEncoderConfigurationOptions struct {
	Options AudioEncoderConfigurationOption
}

type AudioEncoderConfigurationOption struct {
	Encoding       AudioEncoding
	BitrateList    IntList
	SampleRateList IntList
}

type AudioOutputConfigurationOptions struct {
	OutputTokensAvailable ReferenceToken
	SendPrimacyOptions    xsd.AnyURI
	OutputLevelRange      IntRange
}

type AudioDecoderConfigurationOptions struct {
	AACDecOptions  AACDecOptions
	G711DecOptions G711DecOptions
	G726DecOptions G726DecOptions
	Extension      AudioDecoderConfigurationOptionsExtension
}

type AACDecOptions struct {
	Bitrate         IntList
	SampleRateRange IntList
}

type G711DecOptions struct {
	Bitrate         IntList
	SampleRateRange IntList
}

type G726DecOptions struct {
	Bitrate         IntList
	SampleRateRange IntList
}

type AudioDecoderConfigurationOptionsExtension xsd.AnyType

type StreamSetup struct {
	Stream    StreamType `xml:"onvif:Stream"`
	Transport Transport  `xml:"onvif:Transport"`
}

type StreamType xsd.String

type Transport struct {
	Protocol TransportProtocol `xml:"onvif:Protocol"`
	Tunnel   *Transport        `xml:"onvif:Tunnel"`
}

type VideoSourceMode struct {
	Token         ReferenceToken `xml:"token,attr"`
	Enabled       bool           `xml:"Enabled,attr"`
	MaxFramerate  float64
	MaxResolution VideoResolution
	Encodings     EncodingTypes
	Reboot        bool
	Description   Description
	Extension     VideoSourceModeExtension
}

//ReceiverCapabilities of media
type ReceiverCapabilities struct {
	XAddr                xsd.AnyURI
	RTP_Multicast        xsd.Boolean
	RTP_TCP              xsd.Boolean
	RTP_RTSP_TCP         xsd.Boolean
	SupportedReceivers   int
	MaximumRTSPURILength int
}
