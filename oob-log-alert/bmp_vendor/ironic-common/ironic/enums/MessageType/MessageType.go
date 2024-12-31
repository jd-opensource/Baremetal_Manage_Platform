package MessageType

const (
	//操作消息
	MessageTypeOperation = "operation"
	//系统消息
	MessageTypeSystem = "system"
	//监控消息
	MessageTypeOobMonitor = "oob-monitor"
)

const (
	//操作消息类型
	MessageSubTypeCollectHardwareInfo   = "CollectHardwareInfo"
	MessageSubTypeCreateDevices         = "CreateDevices"
	MessageSubTypeCreateInstances       = "CreateInstances"
	MessageSubTypeCutDeviceStock        = "CutDeviceStock"
	MessageSubTypeDeleteInstance        = "DeleteInstance"
	MessageSubTypeInstanceResetPassword = "InstanceResetPassword"
	MessageSubTypePutawayDevice         = "PutawayDevice"
	MessageSubTypeReinstallInstance     = "ReinstallInstance"
	MessageSubTypeRestartInstances      = "RestartInstances"
	MessageSubTypeStartInstances        = "StartInstances"
	MessageSubTypeStopInstances         = "StopInstances"
	//系统消息类型
	MessageSubTypeLicenseNearTimeout = "LicenseNearTimeout"
	//监控消息类型
	MessageSubTypePowerFault       = "电力故障"
	MessageSubTypeTemperatureFault = "温控故障"
	MessageSubTypeCPUFault         = "CPU故障"
	MessageSubTypeMemFault         = "内存故障"
	MessageSubTypeDiskFault        = "硬盘故障"
	MessageSubTypeOtherFault       = "其他故障"
)
