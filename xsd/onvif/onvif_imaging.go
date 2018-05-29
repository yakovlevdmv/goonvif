package onvif

import (
	"github.com/use-go/goonvif/xsd"
)

type ImagingSettingsExtension xsd.AnyType

type ImagingSettings20 struct {
	BacklightCompensation BacklightCompensation20    `xml:"onvif:BacklightCompensation"`
	Brightness            float64                    `xml:"onvif:Brightness"`
	ColorSaturation       float64                    `xml:"onvif:ColorSaturation"`
	Contrast              float64                    `xml:"onvif:Contrast"`
	Exposure              Exposure20                 `xml:"onvif:Exposure"`
	Focus                 FocusConfiguration20       `xml:"onvif:Focus"`
	IrCutFilter           IrCutFilterMode            `xml:"onvif:IrCutFilter"`
	Sharpness             float64                    `xml:"onvif:Sharpness"`
	WideDynamicRange      WideDynamicRange20         `xml:"onvif:WideDynamicRange"`
	WhiteBalance          WhiteBalance20             `xml:"onvif:WhiteBalance"`
	Extension             ImagingSettingsExtension20 `xml:"onvif:Extension"`
}

type ImagingSettingsExtension20 struct {
	ImageStabilization ImageStabilization          `xml:"onvif:ImageStabilization"`
	Extension          ImagingSettingsExtension202 `xml:"onvif:Extension"`
}

type ImageStabilization struct {
	Mode      ImageStabilizationMode      `xml:"onvif:Mode"`
	Level     float64                     `xml:"onvif:Level"`
	Extension ImageStabilizationExtension `xml:"onvif:Extension"`
}

type ImageStabilizationMode xsd.String

type ImageStabilizationExtension xsd.AnyType

type ImagingSettingsExtension202 struct {
	IrCutFilterAutoAdjustment IrCutFilterAutoAdjustment   `xml:"onvif:IrCutFilterAutoAdjustment"`
	Extension                 ImagingSettingsExtension203 `xml:"onvif:Extension"`
}

type ImagingSettingsExtension203 struct {
	ToneCompensation ToneCompensation            `xml:"onvif:ToneCompensation"`
	Defogging        Defogging                   `xml:"onvif:Defogging"`
	NoiseReduction   NoiseReduction              `xml:"onvif:NoiseReduction"`
	Extension        ImagingSettingsExtension204 `xml:"onvif:Extension"`
}

type ImagingSettingsExtension204 xsd.AnyType

type ImagingCapabilities struct {
	XAddr xsd.AnyURI
}

type ImagingSettings struct {
	BacklightCompensation BacklightCompensation
	Brightness            float64
	ColorSaturation       float64
	Contrast              float64
	Exposure              Exposure
	Focus                 FocusConfiguration
	IrCutFilter           IrCutFilterMode
	Sharpness             float64
	WideDynamicRange      WideDynamicRange
	WhiteBalance          WhiteBalance
	Extension             ImagingSettingsExtension
}
