package gosoap

import (
	"encoding/xml"
)

//const string for event action
const (
	SubscribeAction          = "Subscribe"
	UnsubscribeAction        = "Unsubscribe"
	ResumeSubscriptionAction = "ResumeSubscription"
	PauseSubscriptionAction  = "PauseSubscription"
	RenewRequestAction       = "RenewRequest"

	SubscribeActionValue          = "http://docs.oasis-open.org/wsn/bw-2/NotificationProducer/SubscribeRequest"
	ResumeSubscriptionActionValue = "http://docs.oasis-open.org/wsn/bw-2/PausableSubscriptionManager/ResumeSubscriptionRequest"
	PauseSubscriptionActionValue  = "http://docs.oasis-open.org/wsn/bw-2/PausableSubscriptionManager/PauseSubscriptionRequest"
	UnsubscribeActionValue        = "http://docs.oasis-open.org/wsn/bw-2/SubscriptionManager/UnsubscribeRequest"
	RenewRequestActionValue       = "http://docs.oasis-open.org/wsn/bw-2/SubscriptionManager/RenewRequest"
)

//Xlmns XML Scheam
var (
	ActionHeaders = map[string]string{
		SubscribeAction:          SubscribeActionValue,
		ResumeSubscriptionAction: ResumeSubscriptionActionValue,
		PauseSubscriptionAction:  PauseSubscriptionActionValue,
		UnsubscribeAction:        UnsubscribeActionValue,
		RenewRequestAction:       RenewRequestActionValue,
	}
)

/*************************
	Action Type in Header
*************************/

//Action type
type Action struct {
	//XMLName xml.Name  `xml:"wsse:Security"`
	XMLName xml.Name `xml:"wsa:Action"`
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
