package onvif

import (
	"github.com/use-go/goonvif/xsd"
)

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
