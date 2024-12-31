package messagelogic

import (
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/messageDao"
	messageType "git.jd.com/cps-golang/ironic-common/ironic/enums/MessageType"
	log "git.jd.com/cps-golang/log"
)

var Types2ShowZh = map[string]string{
	//操作消息
	messageType.MessageTypeOperation: "操作消息",
	//系统消息
	messageType.MessageTypeSystem: "系统消息",
	//监控消息
	messageType.MessageTypeOobMonitor: "故障消息",
	//带内告警消息
	messageType.MessageTypeInBondAlert: "告警消息",
	//操作消息类型
	messageType.MessageSubTypeCollectHardwareInfo:   "采集",
	messageType.MessageSubTypeCreateDevices:         "上传设备",
	messageType.MessageSubTypeCreateInstances:       "创建实例",
	messageType.MessageSubTypeCutDeviceStock:        "削减库存",
	messageType.MessageSubTypeDeleteInstance:        "删除实例",
	messageType.MessageSubTypeInstanceResetPassword: "重置密码",
	messageType.MessageSubTypePutawayDevice:         "设备上架",
	messageType.MessageSubTypeReinstallInstance:     "重装系统",
	messageType.MessageSubTypeRestartInstances:      "重启实例",
	messageType.MessageSubTypeStartInstances:        "开机",
	messageType.MessageSubTypeStopInstances:         "关机",
	//系统消息类型
	messageType.MessageSubTypeLicenseNearTimeout: "许可证到期",
	//监控消息类型
	messageType.MessageSubTypePowerFault:       "电力故障",
	messageType.MessageSubTypeTemperatureFault: "温控故障",
	messageType.MessageSubTypeCPUFault:         "CPU故障",
	messageType.MessageSubTypeMemFault:         "内存故障",
	messageType.MessageSubTypeDiskFault:        "硬盘故障",
	messageType.MessageSubTypeOtherFault:       "其他故障",

	//带内告警
	messageType.MessageSubTypeCPUUsage:                  "CPU使用率",
	messageType.MessageSubTypeMemoryUsage:               "内存使用率",
	messageType.MessageSubTypeMemoryUsed:                "内存使用量",
	messageType.MessageSubTypeDiskUsed:                  "磁盘使用量",
	messageType.MessageSubTypeDiskUsage:                 "磁盘使用率",
	messageType.MessageSubTypeDiskReadTraffic:           "磁盘读流量",
	messageType.MessageSubTypeDiskWriteTraffic:          "磁盘写流量",
	messageType.MessageSubTypeDiskReadIOPS:              "磁盘读IOPS",
	messageType.MessageSubTypeDiskWriteIOPS:             "磁盘写IOPS",
	messageType.MessageSubTypeNetworkIngressTraffic:     "网卡进流量",
	messageType.MessageSubTypeNetworkEgressTraffic:      "网卡出流量",
	messageType.MessageSubTypeNetworkIngressPackets:     "网络进包数量",
	messageType.MessageSubTypeNetworkEgressPackets:      "网络出包数量",
	messageType.MessageSubTypeLoadAverage1min:           "平均负载均衡1min",
	messageType.MessageSubTypeLoadAverage5min:           "平均负载均衡5min",
	messageType.MessageSubTypeLoadAverage15min:          "平均负载均衡15min",
	messageType.MessageSubTypeTotalTCPConnections:       "TCP总连接数",
	messageType.MessageSubTypeEstablishedTCPConnections: "TCP正常接数",
	messageType.MessageSubTypeTotalProcessCount:         "总进程数",
}

func SendWebNotice(logger *log.Logger, msg messageDao.WebMessage) error {
	if _, err := messageDao.AddWebMessage(logger, &msg); err != nil {
		logger.Warnf("SendWebNotice error, msg:%v, error:%s", msg, err.Error())
		return err
	}
	logger.Infof("SendWebNotice success, msg:%v", msg)
	return nil
}
