package event

import (
	"github.com/use-go/goonvif/xsd"
	"github.com/use-go/goonvif/xsd/onvif"
)

//Address Alias
type Address xsd.String

//CurrentTime alias
type CurrentTime xsd.DateTime //wsnt http://docs.oasis-open.org/wsn/b-2.xsd
//TerminationTime alias
type TerminationTime xsd.DateTime //wsnt http://docs.oasis-open.org/wsn/b-2.xsd
//FixedTopicSet alias
type FixedTopicSet xsd.Boolean //wsnt http://docs.oasis-open.org/wsn/b-2.xsd

//Documentation alias
type Documentation xsd.AnyType //wstop http://docs.oasis-open.org/wsn/t-1.xsd

//TopicExpressionDialect alias
type TopicExpressionDialect xsd.AnyURI

//Message alias
type Message xsd.AnyType

//ActionType for AttributedURIType
type ActionType AttributedURIType

//AttributedURIType in ws-addr
type AttributedURIType xsd.AnyURI //wsa https://www.w3.org/2005/08/addressing/ws-addr.xsd

//AbsoluteOrRelativeTimeType <xsd:union memberTypes="xsd:dateTime xsd:duration"/>
/*
e.g.
   <wsnt:TerminationTime>
      [xsd:dateTime | xsd:duration]
   </wsnt:TerminationTime>
*/
type AbsoluteOrRelativeTimeType xsd.DateTime //wsnt http://docs.oasis-open.org/wsn/b-2.xsd

//EndpointReferenceType in ws-addr
type EndpointReferenceType struct { //wsa http://www.w3.org/2005/08/addressing/ws-addr.xsd
	Address AttributedURIType `xml:"wsa5:Address"`
	//	ReferenceParameters ReferenceParametersType `xml:"wsa5:ReferenceParameters,omit"`
	//	Metadata            MetadataType            `xml:"wsa5:Metadata,omit"`
}

//SubscriptionReference in ws-addr
type SubscriptionReference struct { //wsa http://www.w3.org/2005/08/addressing/ws-addr.xsd
	Address AttributedURIType `xml:"Address"`
	//	ReferenceParameters ReferenceParametersType `xml:"wsa5:ReferenceParameters,omit"`
	//	Metadata            MetadataType            `xml:"wsa5:Metadata,omit"`
}

// FilterType struct
type FilterType struct {
	TopicExpression TopicExpressionType `xml:"wsnt:TopicExpression,omitempty"`
	MessageContent  QueryExpressionType `xml:"wsnt:MessageContent,omitempty"`
}

//ReferenceParametersType in ws-addr
type ReferenceParametersType struct { //wsa https://www.w3.org/2005/08/addressing/ws-addr.xsd
	//Here can be anyAttribute
}

//Metadata in ws-addr
type Metadata MetadataType //wsa https://www.w3.org/2005/08/addressing/ws-addr.xsd

//MetadataType in ws-addr
type MetadataType struct { //wsa https://www.w3.org/2005/08/addressing/ws-addr.xsd

	//Here can be anyAttribute
}

//HeartbeatType ...
type HeartbeatType struct {
	MessageDescription onvif.MessageDescriptionType `xml:"MessageDescription"`
}

//DigitalInputType ...
type DigitalInputType struct {
	Topic              xsd.String                   `xml:"topic,attr"`
	MessageDescription onvif.MessageDescriptionType `xml:"MessageDescription"`
}

//RelayType ...
type RelayType struct {
	Topic              xsd.String                   `xml:"topic,attr"`
	MessageDescription onvif.MessageDescriptionType `xml:"MessageDescription"`
}

//TriggerType ...
type TriggerType struct {
	Topic        xsd.String       `xml:"topic,attr"`
	DigitalInput DigitalInputType `xml:"DigitalInput"`
	Relay        RelayType        `xml:"Relay"`
}

//DeviceType for event
type DeviceType struct { //wstop http://docs.oasis-open.org/wsn/t-1.xsd
	Topic     xsd.String    `xml:"topic,attr"`
	Heartbeat HeartbeatType `xml:"Heartbeat"`
	Trigger   TriggerType   `xml:"Trigger"`
}

//MotionAlarmType ...
type MotionAlarmType struct {
	MessageDescription onvif.MessageDescriptionType `xml:"MessageDescription"`
}

//SignalLossType ...
type SignalLossType struct {
	MessageDescription onvif.MessageDescriptionType `xml:"MessageDescription"`
}

//ShadingType ...
type ShadingType struct {
	MessageDescription onvif.MessageDescriptionType `xml:"MessageDescription"`
}

//AnalyticsServiceType ...
type AnalyticsServiceType struct {
	MessageDescription onvif.MessageDescriptionType `xml:"MessageDescription"`
}

//ImagingServiceType ...
type ImagingServiceType struct {
	MessageDescription onvif.MessageDescriptionType `xml:"MessageDescription"`
}

//RecordingServiceType ...
type RecordingServiceType struct {
	MessageDescription onvif.MessageDescriptionType `xml:"MessageDescription"`
}

//ImageTooBrightType ...
type ImageTooBrightType struct {
	AnalyticsService AnalyticsServiceType `xml:"AnalyticsService"`
	ImagingService   ImagingServiceType   `xml:"ImagingService"`
	RecordingService RecordingServiceType `xml:"RecordingService"`
}

//ImageTooBlurryType ...
type ImageTooBlurryType struct {
	AnalyticsService AnalyticsServiceType `xml:"AnalyticsService"`
	ImagingService   ImagingServiceType   `xml:"ImagingService"`
	RecordingService RecordingServiceType `xml:"RecordingService"`
}

//ImageTooDarkType ...
type ImageTooDarkType struct {
	AnalyticsService AnalyticsServiceType `xml:"AnalyticsService"`
	ImagingService   ImagingServiceType   `xml:"ImagingService"`
	RecordingService RecordingServiceType `xml:"RecordingService"`
}

//ImageColorCastType ...
type ImageColorCastType struct {
	AnalyticsService AnalyticsServiceType `xml:"AnalyticsService"`
	ImagingService   ImagingServiceType   `xml:"ImagingService"`
	RecordingService RecordingServiceType `xml:"RecordingService"`
}

//VideoSourceType for event
type VideoSourceType struct { //wstop http://docs.oasis-open.org/wsn/t-1.xsd
	Topic          xsd.String         `xml:"topic,attr"`
	MotionAlarm    MotionAlarmType    `xml:"MotionAlarm"`
	SignalLoss     SignalLossType     `xml:"SignalLoss"`
	Shading        ShadingType        `xml:"Shading"`
	ImageTooBlurry ImageTooBlurryType `xml:"ImageTooBlurry"`
	ImageTooDark   ImageTooDarkType   `xml:"ImageTooDark"`
	ImageColorCast ImageColorCastType `xml:"ImageColorCast"`
}

//CrossedType ...
type CrossedType struct {
	MessageDescription onvif.MessageDescriptionType `xml:"MessageDescription"`
}

//LineDetectorType ...
type LineDetectorType struct {
	Crossed CrossedType `xml:"Crossed"`
}

//ObjectsInsideType ...
type ObjectsInsideType struct {
	Topic              xsd.String                   `xml:"topic,attr"`
	MessageDescription onvif.MessageDescriptionType `xml:"MessageDescription"`
}

//ObjectsOutsideType ...
type ObjectsOutsideType struct {
	Topic              xsd.String                   `xml:"topic,attr"`
	MessageDescription onvif.MessageDescriptionType `xml:"MessageDescription"`
}

//ObjectsInsideOrOutsideType ...
type ObjectsInsideOrOutsideType struct {
	Topic              xsd.String                   `xml:"topic,attr"`
	MessageDescription onvif.MessageDescriptionType `xml:"MessageDescription"`
}

//ObjectsLoiteringType ...
type ObjectsLoiteringType struct {
	Topic              xsd.String                   `xml:"topic,attr"`
	MessageDescription onvif.MessageDescriptionType `xml:"MessageDescription"`
}

//FieldDetectorType ...
type FieldDetectorType struct {
	ObjectsInside          ObjectsInsideType          `xml:"ObjectsInside"`
	ObjectsOutside         ObjectsOutsideType         `xml:"ObjectsOutside"`
	ObjectsInsideOrOutside ObjectsInsideOrOutsideType `xml:"ObjectsInsideOrOutside"`
	ObjectsLoitering       ObjectsLoiteringType       `xml:"ObjectsLoitering"`
}

//ObjectsTokenType ...
type ObjectsTokenType struct {
	Topic              xsd.String                   `xml:"topic,attr"`
	MessageDescription onvif.MessageDescriptionType `xml:"MessageDescription"`
}

//ObjectsLeaveType ...
type ObjectsLeaveType struct {
	Topic              xsd.String                   `xml:"topic,attr"`
	MessageDescription onvif.MessageDescriptionType `xml:"MessageDescription"`
}

//ObjectsTokenOrLeaveType ...
type ObjectsTokenOrLeaveType struct {
	Topic              xsd.String                   `xml:"topic,attr"`
	MessageDescription onvif.MessageDescriptionType `xml:"MessageDescription"`
}

//ObjectDetectorType ...
type ObjectDetectorType struct {
	ObjectsToken        ObjectsTokenType        `xml:"ObjectsToken"`
	ObjectsLeave        ObjectsLeaveType        `xml:"ObjectsLeave"`
	ObjectsTokenOrLeave ObjectsTokenOrLeaveType `xml:"ObjectsTokenOrLeave"`
}

//CounterType ...
type CounterType struct {
	Topic              xsd.String                   `xml:"topic,attr"`
	MessageDescription onvif.MessageDescriptionType `xml:"MessageDescription"`
}

//CountAggregationType ...
type CountAggregationType struct {
	Counter CounterType `xml:"Counter"`
}

//FireDetectorType ...
type FireDetectorType struct {
	Topic              xsd.String                   `xml:"topic,attr"`
	MessageDescription onvif.MessageDescriptionType `xml:"MessageDescription"`
}

//PlateDetectorType ...
type PlateDetectorType struct {
	Topic              xsd.String                   `xml:"topic,attr"`
	MessageDescription onvif.MessageDescriptionType `xml:"MessageDescription"`
}

//FaceDetectorType ...
type FaceDetectorType struct {
	Topic              xsd.String                   `xml:"topic,attr"`
	MessageDescription onvif.MessageDescriptionType `xml:"MessageDescription"`
}

//FaceAnalysisType ...
type FaceAnalysisType struct {
	FaceDetector FaceDetectorType `xml:"FaceDetector"`
}

//MotionType ...
type MotionType struct {
	Topic              xsd.String                   `xml:"topic,attr"`
	MessageDescription onvif.MessageDescriptionType `xml:"MessageDescription"`
}

//CellMotionDetectorType ...
type CellMotionDetectorType struct {
	FaceDetector MotionType `xml:"FaceDetector"`
}

//RuleEngineType for event
type RuleEngineType struct {
	LineDetector       LineDetectorType       `xml:"LineDetector"`
	FieldDetector      FieldDetectorType      `xml:"FieldDetector"`
	ObjectDetector     ObjectDetectorType     `xml:"ObjectDetector"`
	CountAggregation   CountAggregationType   `xml:"CountAggregation"`
	FireDetector       FireDetectorType       `xml:"FireDetector"`
	PlateDetector      PlateDetectorType      `xml:"PlateDetector"`
	FaceAnalysis       FaceAnalysisType       `xml:"FaceAnalysis"`
	CellMotionDetector CellMotionDetectorType `xml:"CellMotionDetector"`
}

//ProfileChangedType ...
type ProfileChangedType struct {
	Topic              xsd.String                   `xml:"topic,attr"`
	MessageDescription onvif.MessageDescriptionType `xml:"MessageDescription"`
}

//ConfigurationChangedType ...
type ConfigurationChangedType struct {
	Topic              xsd.String                   `xml:"topic,attr"`
	MessageDescription onvif.MessageDescriptionType `xml:"MessageDescription"`
}

//MediaType for event
type MediaType struct {
	ProfileChanged       ProfileChangedType       `xml:"ProfileChanged"`
	ConfigurationChanged ConfigurationChangedType `xml:"ConfigurationChanged"`
}

//ActiveConnectionsType ...
type ActiveConnectionsType struct {
	Topic              xsd.String                   `xml:"topic,attr"`
	MessageDescription onvif.MessageDescriptionType `xml:"MessageDescription"`
}

//ProfileType ...
type ProfileType struct {
	ActiveConnections ActiveConnectionsType `xml:"ActiveConnections"`
}

//MonitoringType for event
type MonitoringType struct {
	Profile ProfileType `xml:"Profile"`
}

//TopicSetType alias
type TopicSetType struct {
	Device      DeviceType      `xml:"Device"`
	VideoSource VideoSourceType `xml:"VideoSource"`
	RuleEngine  RuleEngineType  `xml:"RuleEngine"`
	Media       MediaType       `xml:"Media"`
	Monitoring  MonitoringType  `xml:"Monitoring"`
	//ExtensibleDocumented
	//here can be any element
}

//ExtensibleDocumented struct
type ExtensibleDocumented struct {
	Documentation Documentation //к xsd-документе documentation с маленькой буквы начинается
	//here can be anyAttribute
}

//ProducerReference Alias
type ProducerReference SubscriptionReference

//NotificationMessageHolderType Alias
type NotificationMessageHolderType struct { //wsnt http://docs.oasis-open.org/wsn/b-2.xsd
	SubscriptionReference SubscriptionReference `xml:"SubscriptionReference"`
	Topic                 Topic                 `xml:"Topic"`
	ProducerReference     ProducerReference     `xml:"ProducerReference"`
	Message               MessageType           `xml:"Message>Message"`
}

//MessageType ...
type MessageType struct {
	UtcTime           xsd.DateTime                `xml:"UtcTime,attr"`
	PropertyOperation onvif.PropertyOperationType `xml:"PropertyOperation,attr"`
	Source            onvif.ItemList              `xml:"Source"`
	Key               onvif.ItemList              `xml:"Key"`
	Data              onvif.ItemList              `xml:"Data"`
}

//Notify Message in Body
type Notify struct {
	XMLName                  struct{}              `xml:"Notify"`
	NotificationMessagesList []NotificationMessage `xml:"NotificationMessage"`
}

//NotificationMessage Alias
type NotificationMessage NotificationMessageHolderType //wsnt http://docs.oasis-open.org/wsn/b-2.xsd

//QueryExpressionType struct for wsnt:MessageContent
type QueryExpressionType struct { //wsnt http://docs.oasis-open.org/wsn/b-2.xsd
	Dialect xsd.AnyURI `xml:"Dialect,attr"`
	Value   xsd.String `xml:",chardata"` // boolean(ncex:Producer="15")
}

//MessageContentType Alias
type MessageContentType QueryExpressionType

//QueryExpression Alias
type QueryExpression QueryExpressionType

//TopicExpressionType struct for wsnt:TopicExpression
type TopicExpressionType struct { //wsnt http://docs.oasis-open.org/wsn/b-2.xsd
	Dialect xsd.AnyURI `xml:"Dialect,attr"`
	Value   xsd.String `xml:",chardata"`
}

//Topic Alias
type Topic TopicExpressionType

// Capabilities of event
type Capabilities struct { //tev
	WSSubscriptionPolicySupport                   xsd.Boolean `xml:"WSSubscriptionPolicySupport,attr"`
	WSPullPointSupport                            xsd.Boolean `xml:"WSPullPointSupport,attr"`
	WSPausableSubscriptionManagerInterfaceSupport xsd.Boolean `xml:"WSPausableSubscriptionManagerInterfaceSupport,attr"`
	MaxNotificationProducers                      xsd.Int     `xml:"MaxNotificationProducers,attr"`
	MaxPullPoints                                 xsd.Int     `xml:"MaxPullPoints,attr"`
	PersistentNotificationStorage                 xsd.Boolean `xml:"PersistentNotificationStorage,attr"`
}

//ResourceUnknownFault response type
type ResourceUnknownFault struct {
}

//InvalidFilterFault response type
type InvalidFilterFault struct {
}

//TopicExpressionDialectUnknownFault response type
type TopicExpressionDialectUnknownFault struct {
}

//InvalidTopicExpressionFault response type
type InvalidTopicExpressionFault struct {
}

//TopicNotSupportedFault response type
type TopicNotSupportedFault struct {
}

//InvalidProducerPropertiesExpressionFault response type
type InvalidProducerPropertiesExpressionFault struct {
}

//InvalidMessageContentExpressionFault response type
type InvalidMessageContentExpressionFault struct {
}

//UnacceptableInitialTerminationTimeFault response type
type UnacceptableInitialTerminationTimeFault struct {
}

//UnrecognizedPolicyRequestFault response type
type UnrecognizedPolicyRequestFault struct {
}

//UnsupportedPolicyRequestFault response type
type UnsupportedPolicyRequestFault struct {
}

//NotifyMessageNotSupportedFault response type
type NotifyMessageNotSupportedFault onvif.BaseFaultType

//SubscribeCreationFailedFault response type
type SubscribeCreationFailedFault onvif.BaseFaultType

//ResumeFailedFault type
type ResumeFailedFault onvif.BaseFaultType

//PauseFailedFault type
type PauseFailedFault onvif.BaseFaultType

//UnableToDestroySubscriptionFault type
type UnableToDestroySubscriptionFault onvif.BaseFaultType

//UnacceptableTerminationTimeFault type
type UnacceptableTerminationTimeFault onvif.BaseFaultType

//UnableToCreatePullPointFault type
type UnableToCreatePullPointFault onvif.BaseFaultType
