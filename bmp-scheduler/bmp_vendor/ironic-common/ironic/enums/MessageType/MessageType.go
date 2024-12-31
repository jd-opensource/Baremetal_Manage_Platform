package MessageType

const (
	//操作消息
	MessageTypeOperation = "operation"
	//系统消息
	MessageTypeSystem = "system"
	//监控消息
	MessageTypeOobMonitor = "oob-monitor"
	//带内告警
	MessageTypeInBondAlert = "inbond-alert"
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
	//带外监控消息类型
	MessageSubTypePowerFault       = "电力故障"
	MessageSubTypeTemperatureFault = "温控故障"
	MessageSubTypeCPUFault         = "CPU故障"
	MessageSubTypeMemFault         = "内存故障"
	MessageSubTypeDiskFault        = "硬盘故障"
	MessageSubTypeOtherFault       = "其他故障"

	//带内告警类型
	MessageSubTypeCPUUsage                  = "bmp.cpu.util"
	MessageSubTypeMemoryUsage               = "bmp.memory.util"
	MessageSubTypeMemoryUsed                = "bmp.memory.used"
	MessageSubTypeDiskUsed                  = "bmp.disk.used"
	MessageSubTypeDiskUsage                 = "bmp.disk.util"
	MessageSubTypeDiskReadTraffic           = "bmp.disk.bytes.read"
	MessageSubTypeDiskWriteTraffic          = "bmp.disk.bytes.write"
	MessageSubTypeDiskReadIOPS              = "bmp.disk.counts.read"
	MessageSubTypeDiskWriteIOPS             = "bmp.disk.counts.write"
	MessageSubTypeNetworkIngressTraffic     = "bmp.network.bytes.ingress"
	MessageSubTypeNetworkEgressTraffic      = "bmp.network.bytes.egress"
	MessageSubTypeNetworkIngressPackets     = "bmp.network.packets.ingress"
	MessageSubTypeNetworkEgressPackets      = "bmp.network.packets.egress"
	MessageSubTypeLoadAverage1min           = "bmp.avg.load1"
	MessageSubTypeLoadAverage5min           = "bmp.avg.load5"
	MessageSubTypeLoadAverage15min          = "bmp.avg.load15"
	MessageSubTypeTotalTCPConnections       = "bmp.tcp.connect.total"
	MessageSubTypeEstablishedTCPConnections = "bmp.tcp.connect.established"
	MessageSubTypeTotalProcessCount         = "bmp.process.total"
)
