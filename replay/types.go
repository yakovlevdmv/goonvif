package replay

import (
	"github.com/use-go/goonvif/xsd"
	"github.com/use-go/goonvif/xsd/onvif"
)

type GetReplayUri struct {
	XMLName        string               `xml:"trp:GetReplayUri"`
	StreamSetup    onvif.StreamSetup    `xml:"trp:StreamSetup"`
	RecordingToken onvif.ReferenceToken `xml:"trp:RecordingToken"`
}

type GetReplayUriResponse struct {
	Uri xsd.AnyURI
}

type GetServiceCapabilities struct {
	XMLName string `xml:"trp:GetServiceCapabilities"`
}

type GetServiceCapabilitiesResponse struct {
	Capabilities struct {
		ReversePlayback     xsd.Boolean `xml:"ReversePlayback,attr"`
		SessionTimeoutRange xsd.Int     `xml:"SessionTimeoutRange,attr"`
		RTP_RTSP_TCP        xsd.Boolean `xml:"RTP_RTSP_TCP,attr"`
		RTSPWebSocketUri    xsd.AnyURI  `xml:"RTSPWebSocketUri,attr"`
	} `xml:"Capabilities"`
}
