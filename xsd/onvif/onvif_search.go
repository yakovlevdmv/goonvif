package onvif

import (
	"github.com/use-go/goonvif/xsd"
)

type JobToken ReferenceToken

type SearchScopeExtension string

type SearchScope struct {
	IncludedSources   SourceReference      `xml:"tt:IncludedSources"`
	IncludeRecordings RecordingReference   `xml:"tt:IncludeRecordings"`
	Extension         SearchScopeExtension `xml:"tt:Extension"`
}

type SourceReference struct {
	Type  xsd.AnyURI     `xml:"tt:Type"`
	Token ReferenceToken `xml:"tt:Token"`
}
