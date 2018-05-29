package onvif

import (
	"github.com/use-go/goonvif/xsd"
)

type SecurityCapabilities struct {
	TLS1_1               xsd.Boolean
	TLS1_2               xsd.Boolean
	OnboardKeyGeneration xsd.Boolean
	AccessPolicyConfig   xsd.Boolean
	X_509Token           xsd.Boolean
	SAMLToken            xsd.Boolean
	KerberosToken        xsd.Boolean
	RELToken             xsd.Boolean
	Extension            SecurityCapabilitiesExtension
}

type Certificate struct {
	CertificateID xsd.Token  `xml:"onvif:CertificateID"`
	Certificate   BinaryData `xml:"onvif:Certificate"`
}

type CertificateStatus struct {
	CertificateID xsd.Token   `xml:"onvif:CertificateID"`
	Status        xsd.Boolean `xml:"onvif:Status"`
}

type SecurityCapabilitiesExtension struct {
	TLS1_0    xsd.Boolean
	Extension SecurityCapabilitiesExtension2
}

type SecurityCapabilitiesExtension2 struct {
	Dot1X              xsd.Boolean
	SupportedEAPMethod int
	RemoteUserHandling xsd.Boolean
}

type Dot11Capabilities struct {
	TKIP                  xsd.Boolean
	ScanAvailableNetworks xsd.Boolean
	MultipleConfiguration xsd.Boolean
	AdHocStationMode      xsd.Boolean
	WEP                   xsd.Boolean
}

type Dot11Status struct {
	SSID              Dot11SSIDType
	BSSID             xsd.String
	PairCipher        Dot11Cipher
	GroupCipher       Dot11Cipher
	SignalStrength    Dot11SignalStrength
	ActiveConfigAlias ReferenceToken
}

//TODO: enumeration
type Dot11SignalStrength xsd.String
type Dot11AvailableNetworks struct {
	SSID                  Dot11SSIDType
	BSSID                 xsd.String
	AuthAndMangementSuite Dot11AuthAndMangementSuite
	PairCipher            Dot11Cipher
	GroupCipher           Dot11Cipher
	SignalStrength        Dot11SignalStrength
	Extension             Dot11AvailableNetworksExtension
}

type Dot11AvailableNetworksExtension xsd.AnyType

type CertificateWithPrivateKey struct {
	CertificateID xsd.Token  `xml:"onvif:CertificateID"`
	Certificate   BinaryData `xml:"onvif:Certificate"`
	PrivateKey    BinaryData `xml:"onvif:PrivateKey"`
}

type CertificateInformation struct {
	CertificateID      xsd.Token
	IssuerDN           xsd.String
	SubjectDN          xsd.String
	KeyUsage           CertificateUsage
	ExtendedKeyUsage   CertificateUsage
	KeyLength          xsd.Int
	Version            xsd.String
	SerialNum          xsd.String
	SignatureAlgorithm xsd.String
	Validity           DateTimeRange
	Extension          CertificateInformationExtension
}

type CertificateInformationExtension xsd.AnyType

type DateTimeRange struct {
	From  xsd.DateTime
	Until xsd.DateTime
}

type CertificateUsage struct {
	Critical         xsd.Boolean `xml:"Critical,attr"`
	CertificateUsage xsd.String
}

type Dot1XConfiguration struct {
	Dot1XConfigurationToken ReferenceToken              `xml:"onvif:Dot1XConfigurationToken"`
	Identity                xsd.String                  `xml:"onvif:Identity"`
	AnonymousID             xsd.String                  `xml:"onvif:AnonymousID,omitempty"`
	EAPMethod               xsd.Int                     `xml:"onvif:EAPMethod"`
	CACertificateID         xsd.Token                   `xml:"onvif:CACertificateID,omitempty"`
	EAPMethodConfiguration  EAPMethodConfiguration      `xml:"onvif:EAPMethodConfiguration,omitempty"`
	Extension               Dot1XConfigurationExtension `xml:"onvif:Extension,omitempty"`
}

type Dot1XConfigurationExtension xsd.AnyType

type EAPMethodConfiguration struct {
	TLSConfiguration TLSConfiguration   `xml:"onvif:TLSConfiguration,omitempty"`
	Password         xsd.String         `xml:"onvif:Password,omitempty"`
	Extension        EapMethodExtension `xml:"onvif:Extension,omitempty"`
}

type EapMethodExtension xsd.AnyType

type TLSConfiguration struct {
	CertificateID xsd.Token `xml:"onvif:CertificateID,omitempty"`
}

//TODO: enumeration
type Dot11AuthAndMangementSuite xsd.String
