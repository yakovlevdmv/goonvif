package gosoap

import (
	"encoding/xml"
)

//Xlmns XML Scheam
var actionHeaders = map[string]string{
	"Subscribe":          "http://docs.oasis-open.org/wsn/bw-2/NotificationProducer/SubscribeRequest",
	"ResumeSubscription": "http://docs.oasis-open.org/wsn/bw-2/PausableSubscriptionManager/ResumeSubscriptionRequest",
	"PauseSubscription":  "http://docs.oasis-open.org/wsn/bw-2/PausableSubscriptionManager/PauseSubscriptionRequest",
	"Unsubscribe":        "http://docs.oasis-open.org/wsn/bw-2/PausableSubscriptionManager/UnsubscribeRequest",
	"RenewRequest":       "http://docs.oasis-open.org/wsn/bw-2/PausableSubscriptionManager/RenewRequest",
}

/*************************
	Action Type in Header
*************************/

//Action type
type Action struct {
	//XMLName xml.Name  `xml:"wsse:Security"`
	XMLName xml.Name `xml:"wsa5:Action"`
	Value   string   `xml:",chardata"`
}

/*
   <wsa:Action>
     http://docs.oasis-open.org/wsn/bw-2/NotificationProducer/SubscribeRequest
   </wsa:Action>
*/

//NewAction get a new Head Action Section
func NewAction(actionString string) Action {

	return Action{
		Value: actionString,
		//	Created: time.Now().UTC().Format(time.RFC3339Nano),
	}

}
