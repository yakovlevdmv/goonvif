package event

//Alarm Enum Type
const (
	DeviceHeartbeat = iota
	DeviceTriggerDigitalInput
	DeviceTriggerDigitalRelay
	VideoSourceMotionAlarm
	VideoSourceSignalLoss
	VideoSourceShading
	VideoSourceImageTooBrightAnalyticsService
	VideoSourceImageTooBrightImagingService
	VideoSourceImageTooBrightRecordingService
	RuleEngineLineDetectorCrossed
	RuleEngineFieldDetectorObjectsInside
	RuleEngineFieldDetectorObjectsOutside
	RuleEngineFieldDetectorObjectsInsideOrOutside
	RuleEngineFieldDetectorObjectsLoitering
	RuleEngineObjectDetectorToken
	RuleEngineObjectDetectorLeave
	RuleEngineObjectDetectorObjectsTokenOrLeave
	RuleEngineCountAggregationCounter
	RuleEngineFireDetector
	RuleEnginePlateDetector
	RuleEngineFaceAnalysisFaceDetector
	RuleEngineCellMotionDetectorMotion
	MediaProfileChanged
	MediaConfigurationChanged
	MonitoringProfileActiveConnections
)

//Topic Content No: 25 in total
const (
	DeviceHeartbeatTopic                 = "tns1:Device/tnshoneywell:Heartbeat"
	DeviceTriggerDigitalInputTopic       = "tns1:Device/Trigger/DigitalInput"
	DeviceTriggerDigitalRelayTopic       = "tns1:Device/Trigger/DigitalRelay"
	VideoSourceMotionAlarmTopic          = "tns1:VideoSource/MotionAlarm"
	VideoSourceSignalLossTopic           = "tns1:VideoSource/SignalLoss"
	VideoSourceShadingTopic              = "tns1:VideoSource/Shading"
	VideoSourceImageTooBrightATopic      = "tns1:VideoSource/AnalyticsService"
	VideoSourceImageTooBrightITopic      = "tns1:VideoSource/ImagingService"
	VideoSourceImageTooBrightRTopic      = "tns1:VideoSource/RecordingService"
	RuleEngineLineDetectorCrossedTopic   = "tns1:RuleEngine/LineDetector/Crossed"
	RuleEngineFieldDetectorOInTopic      = "tns1:RuleEngine/FieldDetector/ObjectsInside"
	RuleEngineFieldDetectorOOutTopic     = "tns1:RuleEngine/FieldDetector/ObjectsOutside"
	RuleEngineFieldDetectorOInOrOutTopic = "tns1:RuleEngine/FieldDetector/ObjectsInsideOrOutside"
	RuleEngineFieldDetectorOLTopic       = "tns1:RuleEngine/FieldDetector/ObjectsLoitering"
	RuleEngineObjectDetectorTokenTopic   = "tns1:RuleEngine/ObjectDetector/ObjectsToken"
	RuleEngineObjectDetectorLeaveTopic   = "tns1:RuleEngine/ObjectDetector/ObjectsLeave"
	RuleEngineObjectDetectorTLTopic      = "tns1:RuleEngine/ObjectDetector/ObjectsTokenOrLeave"
	RuleEngineCountAggregationCTopic     = "tns1:RuleEngine/CountAggregation/Counter"
	RuleEngineFireDetectorTopic          = "tns1:RuleEngine/FireDetector"
	RuleEnginePlateDetectorTopic         = "tns1:RuleEngine/PlateDetector"
	RuleEngineFaceAnalysisFDTopic        = "tns1:RuleEngine/FaceAnalysis/FaceDetector"
	RuleEngineCellMotionDetectorMTopic   = "tns1:RuleEngine/CellMotionDetector/Motion"
	MediaProfileChangedTopic             = "tns1:Media/ProfileChanged"
	MediaConfigurationChangedTopic       = "tns1:Media/ConfigurationChanged"
	MonitoringProfileActiveCTopic        = "tns1:Monitoring/Profile/ActiveConnections"
)
