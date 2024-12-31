package object

const (
	// ConfigKey store configration, each field store different config
	ConfigKey string = "Config"

	// FilterKey store filters, each field store different manufacurer filter.
	FilterKey string = "Filter"

	// LogChannel log alert pubsub channel namespace
	LogChannel string = "Logs"

	// OOBErrorChannel OOB access error pubsub channel namespace
	OOBErrorChannel string = "OOBError"

	// PDisksField redis store field name
	PDisksField string = "PDisks"

	// CPUsField redis store field name
	CPUsField string = "CPUs"

	// MemsField redis store field name
	MemsField string = "Mems"

	// NicsField redis store field name
	NicsField string = "Nics"

	// CodeField code store monitor result
	CodeField string = "Code"

	// ErrorCountField count for oob access error
	ErrorCountField string = "OOBErrorCount"

	// MessageField code store monitor result message
	MessageField string = "Message"

	// LogSelInfoField store last ipmi system event log info
	LogSelInfoField string = "LogSelInfo"

	//LastCollectOOBTimeField 最后一次oob收集的时间，是目标机器的oob带外时间，不是目标机器的系统时间，也不是agent服务器的时间
	LastCollectOOBTimeField string = "LastCollectOobTime"
)
