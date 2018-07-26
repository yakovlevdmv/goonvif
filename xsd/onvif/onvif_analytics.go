package onvif

import (
	"github.com/use-go/goonvif/xsd"
)

//AnalyticsDeviceCapabilities ...
type AnalyticsDeviceCapabilities struct {
	XAddr       xsd.AnyURI
	RuleSupport xsd.Boolean
	Extension   AnalyticsDeviceExtension
}

type AnalyticsDeviceExtension xsd.AnyType

//AnalyticsCapabilities Check
type AnalyticsCapabilities struct {
	XAddr                  xsd.AnyURI
	RuleSupport            xsd.Boolean
	AnalyticsModuleSupport xsd.Boolean
}

//AnalyticsEngineConfigurationExtension ...
type AnalyticsEngineConfigurationExtension xsd.AnyType
