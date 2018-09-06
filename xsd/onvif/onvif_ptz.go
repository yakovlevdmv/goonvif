package onvif

import (
	"github.com/use-go/goonvif/xsd"
)

type PTZCapabilities struct {
	XAddr xsd.AnyURI
}

type PTZConfiguration struct {
	ConfigurationEntity
	MoveRamp                               int                       `xml:"MoveRamp,attr"`
	PresetRamp                             int                       `xml:"PresetRamp,attr"`
	PresetTourRamp                         int                       `xml:"PresetTourRamp,attr"`
	NodeToken                              ReferenceToken            `xml:"NodeToken"`
	DefaultAbsolutePantTiltPositionSpace   xsd.AnyURI                `xml:"DefaultAbsolutePantTiltPositionSpace"`
	DefaultAbsoluteZoomPositionSpace       xsd.AnyURI                `xml:"DefaultAbsoluteZoomPositionSpace"`
	DefaultRelativePanTiltTranslationSpace xsd.AnyURI                `xml:"DefaultRelativePanTiltTranslationSpace"`
	DefaultRelativeZoomTranslationSpace    xsd.AnyURI                `xml:"DefaultRelativeZoomTranslationSpace"`
	DefaultContinuousPanTiltVelocitySpace  xsd.AnyURI                `xml:"DefaultContinuousPanTiltVelocitySpace"`
	DefaultContinuousZoomVelocitySpace     xsd.AnyURI                `xml:"DefaultContinuousZoomVelocitySpace"`
	DefaultPTZSpeed                        PTZSpeed                  `xml:"DefaultPTZSpeed"`
	DefaultPTZTimeout                      xsd.Duration              `xml:"DefaultPTZTimeout"`
	PanTiltLimits                          PanTiltLimits             `xml:"PanTiltLimits"`
	ZoomLimits                             ZoomLimits                `xml:"ZoomLimits"`
	Extension                              PTZConfigurationExtension `xml:"Extension"`
}

type ZoomLimits struct {
	Range Space1DDescription `xml:"Range"`
}

type Vector struct {
	X float64 `xml:"x,attr"`
	Y float64 `xml:"y,attr"`
}

type Space1DDescription struct {
	URI    xsd.AnyURI `xml:"URI"`
	XRange FloatRange `xml:"XRange"`
}
type FocusMove struct {
	Absolute   AbsoluteFocus   `xml:"onvif:Absolute"`
	Relative   RelativeFocus   `xml:"onvif:Relative"`
	Continuous ContinuousFocus `xml:"onvif:Continuous"`
}

type ContinuousFocus struct {
	Speed xsd.Float `xml:"onvif:Speed"`
}
type RelativeFocus struct {
	Distance xsd.Float `xml:"onvif:Distance"`
	Speed    xsd.Float `xml:"onvif:Speed"`
}
type AbsoluteFocus struct {
	Position xsd.Float `xml:"onvif:Position"`
	Speed    xsd.Float `xml:"onvif:Speed"`
}

type IntRectangle struct {
	X      int `xml:"x,attr"`
	Y      int `xml:"y,attr"`
	Width  int `xml:"width,attr"`
	Height int `xml:"height,attr"`
}

type IntRectangleRange struct {
	XRange      IntRange
	YRange      IntRange
	WidthRange  IntRange
	HeightRange IntRange
}
type IntRange struct {
	Min int
	Max int
}

type FloatRange struct {
	Min float64 `xml:"Min"`
	Max float64 `xml:"Max"`
}

type Vector2D struct {
	X     float64    `xml:"x,attr"`
	Y     float64    `xml:"y,attr"`
	Space xsd.AnyURI `xml:"space,attr"`
}

type Vector1D struct {
	X     float64    `xml:"x,attr"`
	Space xsd.AnyURI `xml:"space,attr"`
}
type PanTiltLimits struct {
	Range Space2DDescription `xml:"Range"`
}

type Space2DDescription struct {
	URI    xsd.AnyURI `xml:"URI"`
	XRange FloatRange `xml:"XRange"`
	YRange FloatRange `xml:"YRange"`
}

type Rectangle struct {
	Bottom float64 `xml:"bottom,attr"`
	Top    float64 `xml:"top,attr"`
	Right  float64 `xml:"right,attr"`
	Left   float64 `xml:"left,attr"`
}

type FocusConfiguration struct {
	AutoFocusMode AutoFocusMode
	DefaultSpeed  float64
	NearLimit     float64
	FarLimit      float64
}

type AutoFocusMode xsd.String

type IrCutFilterMode xsd.String

type WideDynamicRange struct {
	Mode  WideDynamicMode `xml:"onvif:Mode"`
	Level float64         `xml:"onvif:Level"`
}

type WideDynamicMode xsd.String

type WhiteBalance struct {
	Mode   WhiteBalanceMode
	CrGain float64
	CbGain float64
}

type WhiteBalanceMode xsd.String

type BacklightCompensation20 struct {
	Mode  BacklightCompensationMode `xml:"onvif:Mode"`
	Level float64                   `xml:"onvif:Level"`
}
type Exposure20 struct {
	Mode            ExposureMode     `xml:"onvif:Mode"`
	Priority        ExposurePriority `xml:"onvif:Priority"`
	Window          Rectangle        `xml:"onvif:Window"`
	MinExposureTime float64          `xml:"onvif:MinExposureTime"`
	MaxExposureTime float64          `xml:"onvif:MaxExposureTime"`
	MinGain         float64          `xml:"onvif:MinGain"`
	MaxGain         float64          `xml:"onvif:MaxGain"`
	MinIris         float64          `xml:"onvif:MinIris"`
	MaxIris         float64          `xml:"onvif:MaxIris"`
	ExposureTime    float64          `xml:"onvif:ExposureTime"`
	Gain            float64          `xml:"onvif:Gain"`
	Iris            float64          `xml:"onvif:Iris"`
}

type FocusConfiguration20 struct {
	AutoFocusMode AutoFocusMode                 `xml:"onvif:AutoFocusMode"`
	DefaultSpeed  float64                       `xml:"onvif:DefaultSpeed"`
	NearLimit     float64                       `xml:"onvif:NearLimit"`
	FarLimit      float64                       `xml:"onvif:FarLimit"`
	Extension     FocusConfiguration20Extension `xml:"onvif:Extension"`
}

type FocusConfiguration20Extension xsd.AnyType

type WideDynamicRange20 struct {
	Mode  WideDynamicMode `xml:"onvif:Mode"`
	Level float64         `xml:"onvif:Level"`
}

type WhiteBalance20 struct {
	Mode      WhiteBalanceMode        `xml:"onvif:Mode"`
	CrGain    float64                 `xml:"onvif:CrGain"`
	CbGain    float64                 `xml:"onvif:CbGain"`
	Extension WhiteBalance20Extension `xml:"onvif:Extension"`
}

type WhiteBalance20Extension xsd.AnyType

type IrCutFilterAutoAdjustment struct {
	BoundaryType   string                             `xml:"onvif:BoundaryType"`
	BoundaryOffset float64                            `xml:"onvif:BoundaryOffset"`
	ResponseTime   xsd.Duration                       `xml:"onvif:ResponseTime"`
	Extension      IrCutFilterAutoAdjustmentExtension `xml:"onvif:Extension"`
}

type IrCutFilterAutoAdjustmentExtension xsd.AnyType

type ToneCompensation struct {
	Mode      string                    `xml:"onvif:Mode"`
	Level     float64                   `xml:"onvif:Level"`
	Extension ToneCompensationExtension `xml:"onvif:Extension"`
}

type PTZSpeed struct {
	PanTilt Vector2D `xml:"onvif:PanTilt"`
	Zoom    Vector1D `xml:"onvif:Zoom"`
}

type PTZConfigurationExtension struct {
	PTControlDirection PTControlDirection         `xml:"onvif:PTControlDirection"`
	Extension          PTZConfigurationExtension2 `xml:"onvif:Extension"`
}

type PTControlDirection struct {
	EFlip     EFlip                       `xml:"onvif:EFlip"`
	Reverse   Reverse                     `xml:"onvif:Reverse"`
	Extension PTControlDirectionExtension `xml:"onvif:Extension"`
}

type PTControlDirectionExtension xsd.AnyType

type PTZConfigurationExtension2 xsd.AnyType

type MetadataConfiguration struct {
	ConfigurationEntity
	CompressionType              string                         `xml:"CompressionType,attr"`
	PTZStatus                    PTZFilter                      `xml:"onvif:PTZStatus"`
	Events                       EventSubscription              `xml:"onvif:Events"`
	Analytics                    xsd.Boolean                    `xml:"onvif:Analytics"`
	Multicast                    MulticastConfiguration         `xml:"onvif:Multicast"`
	SessionTimeout               xsd.Duration                   `xml:"onvif:SessionTimeout"`
	AnalyticsEngineConfiguration AnalyticsEngineConfiguration   `xml:"onvif:AnalyticsEngineConfiguration"`
	Extension                    MetadataConfigurationExtension `xml:"onvif:Extension"`
}

type PTZFilter struct {
	Status   bool `xml:"onvif:Status"`
	Position bool `xml:"onvif:Position"`
}

type MetadataConfigurationOptions struct {
	PTZStatusFilterOptions PTZStatusFilterOptions
	Extension              MetadataConfigurationOptionsExtension
}

type PTZStatusFilterOptions struct {
	PanTiltStatusSupported   bool
	ZoomStatusSupported      bool
	PanTiltPositionSupported bool
	ZoomPositionSupported    bool
	Extension                PTZStatusFilterOptionsExtension
}

type PTZStatusFilterOptionsExtension xsd.AnyType

//PTZ

type PTZNode struct {
	DeviceEntity
	FixedHomePosition      xsd.Boolean `xml:"FixedHomePosition,attr"`
	GeoMove                xsd.Boolean `xml:"GeoMove,attr"`
	Name                   Name
	SupportedPTZSpaces     PTZSpaces
	MaximumNumberOfPresets int
	HomeSupported          xsd.Boolean
	AuxiliaryCommands      AuxiliaryData
	Extension              PTZNodeExtension
}

type PTZSpaces struct {
	AbsolutePanTiltPositionSpace    Space2DDescription
	AbsoluteZoomPositionSpace       Space1DDescription
	RelativePanTiltTranslationSpace Space2DDescription
	RelativeZoomTranslationSpace    Space1DDescription
	ContinuousPanTiltVelocitySpace  Space2DDescription
	ContinuousZoomVelocitySpace     Space1DDescription
	PanTiltSpeedSpace               Space1DDescription
	ZoomSpeedSpace                  Space1DDescription
	Extension                       PTZSpacesExtension
}

type PTZSpacesExtension xsd.AnyType

//TODO: restriction
type AuxiliaryData xsd.String

type PTZNodeExtension struct {
	SupportedPresetTour PTZPresetTourSupported
	Extension           PTZNodeExtension2
}

type PTZPresetTourSupported struct {
	MaximumNumberOfPresetTours int
	PTZPresetTourOperation     PTZPresetTourOperation
	Extension                  PTZPresetTourSupportedExtension
}

type PTZPresetTourOperation xsd.String
type PTZPresetTourSupportedExtension xsd.AnyType

type PTZNodeExtension2 xsd.AnyType

type PTZConfigurationOptions struct {
	PTZRamps           IntAttrList `xml:"PTZRamps,attr"`
	Spaces             PTZSpaces
	PTZTimeout         DurationRange
	PTControlDirection PTControlDirectionOptions
	Extension          PTZConfigurationOptions2
}

//PTZConfigurationOptions2 PTZConfigurationOptions2
type PTZConfigurationOptions2 xsd.AnyType

//PTZPreset for ptz presets
type PTZPreset struct {
	Token       ReferenceToken `xml:"token,attr"`
	Name        []Name         `xml:"Name"`
	PTZPosition []PTZVector    `xml:"PTZPosition"`
}

//PTZVector for ptz presets
type PTZVector struct {
	PanTilt Vector2D `xml:"PanTilt"`
	Zoom    Vector1D `xml:"Zoom"`
}

//PTZStatus for ptz
type PTZStatus struct {
	Position   PTZVector     `xml:"Position"`
	MoveStatus PTZMoveStatus `xml:"MoveStatus"`
	Error      string
	UtcTime    xsd.DateTime `xml:"UtcTime"`
}

//PTZMoveStatus PTZMoveStatus
type PTZMoveStatus struct {
	PanTilt MoveStatus `xml:"PanTilt"`
	Zoom    MoveStatus `xml:"Zoom"`
}

//PresetTour PresetTour
type PresetTour struct {
	Token             ReferenceToken                 `xml:"token,attr"`
	Name              Name                           `xml:"onvif:Name"`
	Status            PTZPresetTourStatus            `xml:"onvif:Status"`
	AutoStart         xsd.Boolean                    `xml:"onvif:AutoStart"`
	StartingCondition PTZPresetTourStartingCondition `xml:"onvif:StartingCondition"`
	TourSpot          PTZPresetTourSpot              `xml:"onvif:TourSpot"`
	Extension         PTZPresetTourExtension         `xml:"onvif:Extension"`
}

//PTZPresetTourStatus PTZPresetTourStatus
type PTZPresetTourStatus struct {
	State           PTZPresetTourState           `xml:"onvif:State"`
	CurrentTourSpot PTZPresetTourSpot            `xml:"onvif:CurrentTourSpot"`
	Extension       PTZPresetTourStatusExtension `xml:"onvif:Extension"`
}

//PTZPresetTourState PTZPresetTourState
type PTZPresetTourState xsd.String

//PTZPresetTourSpot ..
type PTZPresetTourSpot struct {
	PresetDetail PTZPresetTourPresetDetail  `xml:"onvif:PresetDetail"`
	Speed        PTZSpeed                   `xml:"onvif:Speed"`
	StayTime     xsd.Duration               `xml:"onvif:StayTime"`
	Extension    PTZPresetTourSpotExtension `xml:"onvif:Extension"`
}

type PTZPresetTourPresetDetail struct {
	PresetToken   ReferenceToken             `xml:"onvif:PresetToken"`
	Home          xsd.Boolean                `xml:"onvif:Home"`
	PTZPosition   PTZVector                  `xml:"onvif:PTZPosition"`
	TypeExtension PTZPresetTourTypeExtension `xml:"onvif:TypeExtension"`
}

type PTZPresetTourTypeExtension xsd.AnyType

type PTZPresetTourSpotExtension xsd.AnyType

type PTZPresetTourStatusExtension xsd.AnyType

type PTZPresetTourStartingCondition struct {
	RandomPresetOrder xsd.Boolean                             `xml:"RandomPresetOrder,attr"`
	RecurringTime     xsd.Int                                 `xml:"onvif:RecurringTime"`
	RecurringDuration xsd.Duration                            `xml:"onvif:RecurringDuration"`
	Direction         PTZPresetTourDirection                  `xml:"onvif:Direction"`
	Extension         PTZPresetTourStartingConditionExtension `xml:"onvif:Extension"`
}

type PTZPresetTourDirection xsd.String

type PTZPresetTourStartingConditionExtension xsd.AnyType

type PTZPresetTourExtension xsd.AnyType

type PTZPresetTourOptions struct {
	AutoStart         xsd.Boolean
	StartingCondition PTZPresetTourStartingConditionOptions
	TourSpot          PTZPresetTourSpotOptions
}

type PTZPresetTourStartingConditionOptions struct {
	RecurringTime     IntRange
	RecurringDuration DurationRange
	Direction         PTZPresetTourDirection
	Extension         PTZPresetTourStartingConditionOptionsExtension
}

type PTZPresetTourStartingConditionOptionsExtension xsd.AnyType

type PTZPresetTourSpotOptions struct {
	PresetDetail PTZPresetTourPresetDetailOptions
	StayTime     DurationRange
}

type PTZPresetTourPresetDetailOptions struct {
	PresetToken          ReferenceToken
	Home                 xsd.Boolean
	PanTiltPositionSpace Space2DDescription
	ZoomPositionSpace    Space1DDescription
	Extension            PTZPresetTourPresetDetailOptionsExtension
}

type PTZPresetTourPresetDetailOptionsExtension xsd.AnyType
