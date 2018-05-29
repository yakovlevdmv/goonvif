package onvif

import (
	"github.com/use-go/goonvif/xsd"
)

type IPAddress struct {
	Type        IPType      `xml:"onvif:Type"`
	IPv4Address IPv4Address `xml:"onvif:IPv4Address"`
	IPv6Address IPv6Address `xml:"onvif:IPv6Address"`
}

type IPType xsd.String

//IPv4 address
type IPv4Address xsd.Token

//IPv6 address
type IPv6Address xsd.Token

//DNSName from token
type DNSName xsd.Token

type HostnameInformation struct {
	FromDHCP  xsd.Boolean
	Name      xsd.Token
	Extension HostnameInformationExtension
}

type HostnameInformationExtension xsd.AnyType

type DNSInformation struct {
	FromDHCP     xsd.Boolean
	SearchDomain xsd.Token
	DNSFromDHCP  IPAddress
	DNSManual    IPAddress
	Extension    DNSInformationExtension
}

type DNSInformationExtension xsd.AnyType

type NTPInformation struct {
	FromDHCP    xsd.Boolean
	NTPFromDHCP NetworkHost
	NTPManual   NetworkHost
	Extension   NTPInformationExtension
}

type NTPInformationExtension xsd.AnyType

type DynamicDNSInformation struct {
	Type      DynamicDNSType
	Name      DNSName
	TTL       xsd.Duration
	Extension DynamicDNSInformationExtension
}

//TODO: enumeration
type DynamicDNSType xsd.String

type DynamicDNSInformationExtension xsd.AnyType

type DiscoveryMode xsd.String

type NetworkHost struct {
	Type        NetworkHostType      `xml:"onvif:Type"`
	IPv4Address IPv4Address          `xml:"onvif:IPv4Address"`
	IPv6Address IPv6Address          `xml:"onvif:IPv6Address"`
	DNSname     DNSName              `xml:"onvif:DNSname"`
	Extension   NetworkHostExtension `xml:"onvif:Extension"`
}

type NetworkHostType xsd.String

type NetworkHostExtension xsd.String

//NetworkCapabilities Check
type NetworkCapabilities struct {
	IPFilter          xsd.Boolean
	ZeroConfiguration xsd.Boolean
	IPVersion6        xsd.Boolean
	DynDNS            xsd.Boolean
	Extension         NetworkCapabilitiesExtension
}

//NetworkCapabilitiesExtension Check
type NetworkCapabilitiesExtension struct {
	Dot11Configuration xsd.Boolean
	Extension          NetworkCapabilitiesExtension2
}

//NetworkCapabilitiesExtension2 Extension2
type NetworkCapabilitiesExtension2 xsd.AnyType

type NetworkInterface struct {
	DeviceEntity
	Enabled   xsd.Boolean
	Info      NetworkInterfaceInfo
	Link      NetworkInterfaceLink
	IPv4      IPv4NetworkInterface
	IPv6      IPv6NetworkInterface
	Extension NetworkInterfaceExtension
}

type NetworkInterfaceInfo struct {
	Name      xsd.String
	HwAddress HwAddress
	MTU       xsd.Int
}

type HwAddress xsd.Token

type NetworkInterfaceLink struct {
	AdminSettings NetworkInterfaceConnectionSetting
	OperSettings  NetworkInterfaceConnectionSetting
	InterfaceType IANA_IfTypes `xml:"IANA-IfTypes"`
}

type IANA_IfTypes xsd.Int

type NetworkInterfaceConnectionSetting struct {
	AutoNegotiation xsd.Boolean `xml:"onvif:AutoNegotiation"`
	Speed           xsd.Int     `xml:"onvif:Speed"`
	Duplex          Duplex      `xml:"onvif:Duplex"`
}

//TODO: enum
type Duplex xsd.String

type NetworkInterfaceExtension struct {
	InterfaceType IANA_IfTypes
	Dot3          Dot3Configuration
	Dot11         Dot11Configuration
	Extension     NetworkInterfaceExtension2
}

type NetworkInterfaceExtension2 xsd.AnyType

type Dot11Configuration struct {
	SSID     Dot11SSIDType                  `xml:"onvif:SSID"`
	Mode     Dot11StationMode               `xml:"onvif:Mode"`
	Alias    Name                           `xml:"onvif:Alias"`
	Priority NetworkInterfaceConfigPriority `xml:"onvif:Priority"`
	Security Dot11SecurityConfiguration     `xml:"onvif:Security"`
}

type Dot11SecurityConfiguration struct {
	Mode      Dot11SecurityMode                   `xml:"onvif:Mode"`
	Algorithm Dot11Cipher                         `xml:"onvif:Algorithm"`
	PSK       Dot11PSKSet                         `xml:"onvif:PSK"`
	Dot1X     ReferenceToken                      `xml:"onvif:Dot1X"`
	Extension Dot11SecurityConfigurationExtension `xml:"onvif:Extension"`
}

type Dot11SecurityConfigurationExtension xsd.AnyType

type Dot11PSKSet struct {
	Key        Dot11PSK             `xml:"onvif:Key"`
	Passphrase Dot11PSKPassphrase   `xml:"onvif:Passphrase"`
	Extension  Dot11PSKSetExtension `xml:"onvif:Extension"`
}

type Dot11PSKSetExtension xsd.AnyType

type Dot11PSKPassphrase xsd.String

type Dot11PSK xsd.HexBinary

//TODO: enumeration
type Dot11Cipher xsd.String

//TODO: enumeration
type Dot11SecurityMode xsd.String

//TODO: restrictions
type NetworkInterfaceConfigPriority xsd.Integer

//TODO: enumeration
type Dot11StationMode xsd.String

//TODO: restrictions
type Dot11SSIDType xsd.HexBinary

type Dot3Configuration xsd.String

type IPv6NetworkInterface struct {
	Enabled xsd.Boolean
	Config  IPv6Configuration
}

type IPv6Configuration struct {
	AcceptRouterAdvert xsd.Boolean
	DHCP               IPv6DHCPConfiguration
	Manual             PrefixedIPv6Address
	LinkLocal          PrefixedIPv6Address
	FromDHCP           PrefixedIPv6Address
	FromRA             PrefixedIPv6Address
	Extension          IPv6ConfigurationExtension
}

type IPv6ConfigurationExtension xsd.AnyType

//PrefixedIPv6Address ...
type PrefixedIPv6Address struct {
	Address      IPv6Address `xml:"onvif:Address"`
	PrefixLength xsd.Int     `xml:"onvif:PrefixLength"`
}

//TODO: enumeration
type IPv6DHCPConfiguration xsd.String

type IPv4NetworkInterface struct {
	Enabled xsd.Boolean
	Config  IPv4Configuration
}

type IPv4Configuration struct {
	Manual    PrefixedIPv4Address
	LinkLocal PrefixedIPv4Address
	FromDHCP  PrefixedIPv4Address
	DHCP      xsd.Boolean
}

//optional, unbounded
type PrefixedIPv4Address struct {
	Address      IPv4Address `xml:"onvif:Address"`
	PrefixLength xsd.Int     `xml:"onvif:PrefixLength"`
}

type NetworkInterfaceSetConfiguration struct {
	Enabled   xsd.Boolean                               `xml:"onvif:Enabled"`
	Link      NetworkInterfaceConnectionSetting         `xml:"onvif:Link"`
	MTU       xsd.Int                                   `xml:"onvif:MTU"`
	IPv4      IPv4NetworkInterfaceSetConfiguration      `xml:"onvif:IPv4"`
	IPv6      IPv6NetworkInterfaceSetConfiguration      `xml:"onvif:IPv6"`
	Extension NetworkInterfaceSetConfigurationExtension `xml:"onvif:Extension"`
}

type NetworkInterfaceSetConfigurationExtension struct {
	Dot3      Dot3Configuration                          `xml:"onvif:Dot3"`
	Dot11     Dot11Configuration                         `xml:"onvif:Dot11"`
	Extension NetworkInterfaceSetConfigurationExtension2 `xml:"onvif:Extension"`
}

type NetworkInterfaceSetConfigurationExtension2 xsd.AnyType

type IPv6NetworkInterfaceSetConfiguration struct {
	Enabled            xsd.Boolean           `xml:"onvif:Enabled"`
	AcceptRouterAdvert xsd.Boolean           `xml:"onvif:AcceptRouterAdvert"`
	Manual             PrefixedIPv6Address   `xml:"onvif:Manual"`
	DHCP               IPv6DHCPConfiguration `xml:"onvif:DHCP"`
}

type IPv4NetworkInterfaceSetConfiguration struct {
	Enabled xsd.Boolean         `xml:"onvif:Enabled"`
	Manual  PrefixedIPv4Address `xml:"onvif:Manual"`
	DHCP    xsd.Boolean         `xml:"onvif:DHCP"`
}

type NetworkProtocol struct {
	Name      NetworkProtocolType      `xml:"onvif:Name"`
	Enabled   xsd.Boolean              `xml:"onvif:Enabled"`
	Port      xsd.Int                  `xml:"onvif:Port"`
	Extension NetworkProtocolExtension `xml:"onvif:Extension"`
}

type NetworkProtocolExtension xsd.AnyType

//TODO: enumeration
type NetworkProtocolType xsd.String

type NetworkGateway struct {
	IPv4Address IPv4Address
	IPv6Address IPv6Address
}

type NetworkZeroConfiguration struct {
	InterfaceToken ReferenceToken
	Enabled        xsd.Boolean
	Addresses      IPv4Address
	Extension      NetworkZeroConfigurationExtension
}

type NetworkZeroConfigurationExtension struct {
	Additional *NetworkZeroConfiguration
	Extension  NetworkZeroConfigurationExtension2
}

type NetworkZeroConfigurationExtension2 xsd.AnyType

type IPAddressFilter struct {
	Type        IPAddressFilterType      `xml:"onvif:Type"`
	IPv4Address PrefixedIPv4Address      `xml:"onvif:IPv4Address,omitempty"`
	IPv6Address PrefixedIPv6Address      `xml:"onvif:IPv6Address,omitempty"`
	Extension   IPAddressFilterExtension `xml:"onvif:Extension,omitempty"`
}

type IPAddressFilterExtension xsd.AnyType

//enum { 'Allow', 'Deny' }
//TODO: enumeration
type IPAddressFilterType xsd.String
