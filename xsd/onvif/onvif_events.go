package onvif

import (
	"github.com/use-go/goonvif/xsd"
)

//EventSubscription ... for ptz use only
type EventSubscription struct {
	Filter             FilterType `xml:"onvif:Filter"`
	SubscriptionPolicy `xml:"onvif:SubscriptionPolicy"`
}

//FilterType ... for ptz use only
type FilterType xsd.AnyType

//SubscriptionPolicy ... for ptz use only
type SubscriptionPolicy xsd.AnyType

//EventCapabilities ...
type EventCapabilities struct {
	XAddr                                         xsd.AnyURI
	WSSubscriptionPolicySupport                   xsd.Boolean
	WSPullPointSupport                            xsd.Boolean
	WSPausableSubscriptionManagerInterfaceSupport xsd.Boolean
}

//PropertyOperationType ...
/*
	<xs:enumeration value="Initialized"/>
	<xs:enumeration value="Deleted"/>
	<xs:enumeration value="Changed"/>
*/
type PropertyOperationType xsd.String

//ItemList ...
type ItemList struct {
	SimpleItem  []SimpleItem      `xml:"SimpleItem"`
	ElementItem ElementItem       `xml:"ElementItem"`
	Extension   ItemListExtension `xml:"Extension"`
}

//SimpleItem for data
type SimpleItem struct {
	Name  string            `xml:"Name,attr"`
	Value xsd.AnySimpleType `xml:"Value,attr"`
}

//ElementItem for data
type ElementItem struct {
	Name string `xml:"Name,attr"`
}

//ItemListExtension for data
type ItemListExtension xsd.AnyType
