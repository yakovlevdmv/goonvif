package recording

import (
	"github.com/use-go/goonvif/xsd"
	"github.com/use-go/goonvif/xsd/onvif"
)

type GetRecordings struct {
	XMLName string `xml:"trc:GetRecordings"`
}

type GetRecordingsResponse struct {
	RecordingItem []struct {
		RecordingToken onvif.RecordingReference
		Configuration  struct {
			Source struct {
				SourceId    xsd.AnyURI
				Name        xsd.Name
				Location    onvif.Description
				Description onvif.Description
				Address     xsd.AnyURI
			}
			Content              onvif.Description
			MaximumRetentionTime xsd.Duration
		}
		Tracks []struct {
			Track struct {
				TrackToken    string
				Configuration struct {
					TrackType   string
					Description string
				}
			}
		}
	}
}
