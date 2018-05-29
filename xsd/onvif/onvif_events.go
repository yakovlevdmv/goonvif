package onvif

import (
	"github.com/use-go/goonvif/xsd"
)

type RemoteUser struct {
	Username           string      `xml:"onvif:Username"`
	Password           string      `xml:"onvif:Password"`
	UseDerivedPassword xsd.Boolean `xml:"onvif:UseDerivedPassword"`
}

type User struct {
	Username  string        `xml:"onvif:Username"`
	Password  string        `xml:"onvif:Password"`
	UserLevel UserLevel     `xml:"onvif:UserLevel"`
	Extension UserExtension `xml:"onvif:Extension"`
}

type UserLevel xsd.String

type UserExtension xsd.String

type EventSubscription struct {
	Filter             FilterType `xml:"onvif:Filter"`
	SubscriptionPolicy `xml:"onvif:SubscriptionPolicy"`
}

type FilterType xsd.AnyType

type SubscriptionPolicy xsd.AnyType

type EventCapabilities struct {
	XAddr                                         xsd.AnyURI
	WSSubscriptionPolicySupport                   xsd.Boolean
	WSPullPointSupport                            xsd.Boolean
	WSPausableSubscriptionManagerInterfaceSupport xsd.Boolean
}
