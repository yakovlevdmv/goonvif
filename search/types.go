package search

import (
	"github.com/use-go/goonvif/xsd"
	"github.com/use-go/goonvif/xsd/onvif"
)

type EventFilter onvif.FilterType

type FindEvents struct {
	XMLName           string            `xml:"tse:FindEvents"`
	StartPoint        xsd.DateTime      `xml:"tse:StartPoint"`
	EndPoint          xsd.DateTime      `xml:"tse:EndPoint"`
	Scope             onvif.SearchScope `xml:"tse:Scope"`
	SearchFilter      EventFilter       `xml:"tse:SearchFilter"`
	IncludeStartState xsd.Boolean       `xml:"tse:IncludeStartState"`
	MaxMatches        xsd.Int           `xml:"tse:MaxMatches"`
	KeepAliveTime     xsd.Duration      `xml:"tse:KeepAliveTime"`
}

type FindEventsResponse struct {
	SearchToken string
}

type GetEventSearchResults struct {
	XMLName     string `xml:"tse:GetEventSearchResults"`
	SearchToken string `xml:"tse:SearchToken"`
}

type GetEventSearchResultsResponse struct {
	ResultList struct {
		SearchState string
		Result      []struct {
			RecordingToken  string
			TrackToken      string
			Time            xsd.Duration
			Event           string
			StartStateEvent xsd.Boolean
		}
	}
}

type FindRecordings struct {
	XMLName string `xml:"tse:FindRecordings"`
}

type FindRecordingsResponse struct{}
