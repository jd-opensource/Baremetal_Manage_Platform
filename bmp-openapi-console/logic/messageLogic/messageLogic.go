package messageLogic

import (
	"errors"
	"strconv"
	"strings"

	"coding.jd.com/aidc-bmp/bmp-openapi/constant"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/deviceDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/idcDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/instanceDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/messageDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/projectDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/baseLogic"
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
	"coding.jd.com/aidc-bmp/bmp-openapi/util"
	log "coding.jd.com/aidc-bmp/bmp_log"
	messageType "git.jd.com/cps-golang/ironic-common/ironic/enums/MessageType"
)

var Types2ShowEn = map[string]string{
	//操作消息
	messageType.MessageTypeOperation: "operation",
	//系统消息
	messageType.MessageTypeSystem: "system",
	//监控消息
	messageType.MessageTypeOobMonitor: "oob-monitor",
	//带内报警消息
	messageType.MessageTypeInBondAlert: "inbond-alert",
	//操作消息类型
	messageType.MessageSubTypeCollectHardwareInfo:   "CollectHardwareInfo",
	messageType.MessageSubTypeCreateDevices:         "CreateDevices",
	messageType.MessageSubTypeCreateInstances:       "CreateInstances",
	messageType.MessageSubTypeCutDeviceStock:        "CutDeviceStock",
	messageType.MessageSubTypeDeleteInstance:        "DeleteInstance",
	messageType.MessageSubTypeInstanceResetPassword: "InstanceResetPassword",
	messageType.MessageSubTypePutawayDevice:         "PutawayDevice",
	messageType.MessageSubTypeReinstallInstance:     "ReinstallInstance",
	messageType.MessageSubTypeRestartInstances:      "RestartInstances",
	messageType.MessageSubTypeStartInstances:        "StartInstances",
	messageType.MessageSubTypeStopInstances:         "StopInstances",
	//系统消息类型
	messageType.MessageSubTypeLicenseNearTimeout: "LicenseNearTimeout",
	//监控消息类型
	messageType.MessageSubTypePowerFault:       "Power Fault",
	messageType.MessageSubTypeTemperatureFault: "Temp Fault",
	messageType.MessageSubTypeCPUFault:         "CPU Fault",
	messageType.MessageSubTypeMemFault:         "Mem Fault",
	messageType.MessageSubTypeDiskFault:        "Disk Fault",
	messageType.MessageSubTypeOtherFault:       "Other Fault",

	//带内告警
	messageType.MessageSubTypeCPUUsage:                  "CPU Usage",
	messageType.MessageSubTypeMemoryUsage:               "Memory Usage",
	messageType.MessageSubTypeMemoryUsed:                "Memory Used",
	messageType.MessageSubTypeDiskUsed:                  "Disk Used",
	messageType.MessageSubTypeDiskUsage:                 "Disk Usage",
	messageType.MessageSubTypeDiskReadTraffic:           "Disk Read Traffic",
	messageType.MessageSubTypeDiskWriteTraffic:          "Disk Write Traffic",
	messageType.MessageSubTypeDiskReadIOPS:              "Disk Read IOPS",
	messageType.MessageSubTypeDiskWriteIOPS:             "Disk Write IOPS",
	messageType.MessageSubTypeNetworkIngressTraffic:     "Network Ingress Traffic",
	messageType.MessageSubTypeNetworkEgressTraffic:      "Network Egress Traffic",
	messageType.MessageSubTypeNetworkIngressPackets:     "Network Ingress Packets",
	messageType.MessageSubTypeNetworkEgressPackets:      "Network Egress Packets",
	messageType.MessageSubTypeLoadAverage1min:           "Load Average 1min",
	messageType.MessageSubTypeLoadAverage5min:           "Load Average 5min",
	messageType.MessageSubTypeLoadAverage15min:          "Load Average 15min",
	messageType.MessageSubTypeTotalTCPConnections:       "Total TCP Connections",
	messageType.MessageSubTypeEstablishedTCPConnections: "Established TCP Connections",
	messageType.MessageSubTypeTotalProcessCount:         "Total Process Count",
}

var Types2ShowZh = map[string]string{
	//操作消息
	messageType.MessageTypeOperation: "操作消息",
	//系统消息
	messageType.MessageTypeSystem: "系统消息",
	//监控消息
	messageType.MessageTypeOobMonitor: "故障消息",
	//带内告警消息
	messageType.MessageTypeInBondAlert: "报警消息",

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
	messageType.MessageSubTypeLicenseNearTimeout: "许可证到期提醒",
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

var ShowEn2Types = map[string]string{
	//操作消息
	"operation": messageType.MessageTypeOperation,
	//系统消息
	"system": messageType.MessageTypeSystem,
	//监控消息
	"oob-monitor": messageType.MessageTypeOobMonitor,
	//带内告警消息
	"inbond-alert": messageType.MessageTypeInBondAlert,

	//操作消息类型
	"CollectHardwareInfo":   messageType.MessageSubTypeCollectHardwareInfo,
	"CreateDevices":         messageType.MessageSubTypeCreateDevices,
	"CreateInstances":       messageType.MessageSubTypeCreateInstances,
	"CutDeviceStock":        messageType.MessageSubTypeCutDeviceStock,
	"DeleteInstance":        messageType.MessageSubTypeDeleteInstance,
	"InstanceResetPassword": messageType.MessageSubTypeInstanceResetPassword,
	"PutawayDevice":         messageType.MessageSubTypePutawayDevice,
	"ReinstallInstance":     messageType.MessageSubTypeReinstallInstance,
	"RestartInstances":      messageType.MessageSubTypeRestartInstances,
	"StartInstances":        messageType.MessageSubTypeStartInstances,
	"StopInstances":         messageType.MessageSubTypeStopInstances,
	//系统消息类型
	"LicenseNearTimeout": messageType.MessageSubTypeLicenseNearTimeout,
	//监控消息类型
	"Power Fault": messageType.MessageSubTypePowerFault,
	"Temp Fault":  messageType.MessageSubTypeTemperatureFault,
	"CPU Fault":   messageType.MessageSubTypeCPUFault,
	"Mem Fault":   messageType.MessageSubTypeMemFault,
	"Disk Fault":  messageType.MessageSubTypeDiskFault,
	"Other Fault": messageType.MessageSubTypeOtherFault,
	//带内监控告警
	"CPU Usage":                   messageType.MessageSubTypeCPUUsage,
	"Memory Usage":                messageType.MessageSubTypeMemoryUsage,
	"Memory Used":                 messageType.MessageSubTypeMemoryUsed,
	"Disk Used":                   messageType.MessageSubTypeDiskUsed,
	"Disk Usage":                  messageType.MessageSubTypeDiskUsage,
	"Disk Read Traffic":           messageType.MessageSubTypeDiskReadTraffic,
	"Disk Write Traffic":          messageType.MessageSubTypeDiskWriteTraffic,
	"Disk Read IOPS":              messageType.MessageSubTypeDiskReadIOPS,
	"Disk Write IOPS":             messageType.MessageSubTypeDiskWriteIOPS,
	"Network Ingress Traffic":     messageType.MessageSubTypeNetworkIngressTraffic,
	"Network Egress Traffic":      messageType.MessageSubTypeNetworkEgressTraffic,
	"Network Ingress Packets":     messageType.MessageSubTypeNetworkIngressPackets,
	"Network Egress Packets":      messageType.MessageSubTypeNetworkEgressPackets,
	"Load Average 1min":           messageType.MessageSubTypeLoadAverage1min,
	"Load Average 5min":           messageType.MessageSubTypeLoadAverage5min,
	"Load Average 15min":          messageType.MessageSubTypeLoadAverage15min,
	"Total TCP Connections":       messageType.MessageSubTypeTotalTCPConnections,
	"Established TCP Connections": messageType.MessageSubTypeEstablishedTCPConnections,
	"Total Process Count":         messageType.MessageSubTypeTotalProcessCount,
}

var ShowZh2Types = map[string]string{
	//操作消息
	"操作消息": messageType.MessageTypeOperation,
	//系统消息
	"系统消息": messageType.MessageTypeSystem,
	//监控消息
	"故障消息": messageType.MessageTypeOobMonitor,
	//带内告警消息
	"报警消息": messageType.MessageTypeInBondAlert,

	//操作消息类型
	"采集":   messageType.MessageSubTypeCollectHardwareInfo,
	"上传设备": messageType.MessageSubTypeCreateDevices,
	"创建实例": messageType.MessageSubTypeCreateInstances,
	"削减库存": messageType.MessageSubTypeCutDeviceStock,
	"删除实例": messageType.MessageSubTypeDeleteInstance,
	"重置密码": messageType.MessageSubTypeInstanceResetPassword,
	"设备上架": messageType.MessageSubTypePutawayDevice,
	"重装系统": messageType.MessageSubTypeReinstallInstance,
	"重启实例": messageType.MessageSubTypeRestartInstances,
	"开机":   messageType.MessageSubTypeStartInstances,
	"关机":   messageType.MessageSubTypeStopInstances,
	//系统消息类型
	"许可证到期提醒": messageType.MessageSubTypeLicenseNearTimeout,
	//监控消息类型
	"电力故障":  messageType.MessageSubTypePowerFault,
	"温控故障":  messageType.MessageSubTypeTemperatureFault,
	"CPU故障": messageType.MessageSubTypeCPUFault,
	"内存故障":  messageType.MessageSubTypeMemFault,
	"硬盘故障":  messageType.MessageSubTypeDiskFault,
	"其他故障":  messageType.MessageSubTypeOtherFault,

	//带内告警
	"CPU使用率":      messageType.MessageSubTypeCPUUsage,
	"内存使用率":       messageType.MessageSubTypeMemoryUsage,
	"内存使用量":       messageType.MessageSubTypeMemoryUsed,
	"磁盘使用量":       messageType.MessageSubTypeDiskUsed,
	"磁盘使用率":       messageType.MessageSubTypeDiskUsage,
	"磁盘读流量":       messageType.MessageSubTypeDiskReadTraffic,
	"磁盘写流量":       messageType.MessageSubTypeDiskWriteTraffic,
	"磁盘读IOPS":     messageType.MessageSubTypeDiskReadIOPS,
	"磁盘写IOPS":     messageType.MessageSubTypeDiskWriteIOPS,
	"网卡进流量":       messageType.MessageSubTypeNetworkIngressTraffic,
	"网卡出流量":       messageType.MessageSubTypeNetworkEgressTraffic,
	"网络进包数量":      messageType.MessageSubTypeNetworkIngressPackets,
	"网络出包数量":      messageType.MessageSubTypeNetworkEgressPackets,
	"平均负载均衡1min":  messageType.MessageSubTypeLoadAverage1min,
	"平均负载均衡5min":  messageType.MessageSubTypeLoadAverage5min,
	"平均负载均衡15min": messageType.MessageSubTypeLoadAverage15min,
	"TCP总连接数":     messageType.MessageSubTypeTotalTCPConnections,
	"TCP正常接数":     messageType.MessageSubTypeEstablishedTCPConnections,
	"总进程数":        messageType.MessageSubTypeTotalProcessCount,
}

var IncludeMessagetypes []string = []string{

	messageType.MessageSubTypeCreateInstances,

	messageType.MessageSubTypeDeleteInstance,
	messageType.MessageSubTypeInstanceResetPassword,
	messageType.MessageSubTypeReinstallInstance,
	messageType.MessageSubTypeRestartInstances,
	messageType.MessageSubTypeStartInstances,
	messageType.MessageSubTypeStopInstances,
	messageType.MessageSubTypePowerFault,
	messageType.MessageSubTypeTemperatureFault,
	messageType.MessageSubTypeCPUFault,
	messageType.MessageSubTypeMemFault,
	messageType.MessageSubTypeDiskFault,
	messageType.MessageSubTypeOtherFault,

	messageType.MessageSubTypeCPUUsage,
	messageType.MessageSubTypeMemoryUsed,
	messageType.MessageSubTypeMemoryUsage,
	messageType.MessageSubTypeDiskUsed,
	messageType.MessageSubTypeDiskUsage,
	messageType.MessageSubTypeDiskReadTraffic,
	messageType.MessageSubTypeDiskWriteTraffic,
	messageType.MessageSubTypeDiskReadIOPS,
	messageType.MessageSubTypeDiskWriteIOPS,
	messageType.MessageSubTypeNetworkIngressTraffic,
	messageType.MessageSubTypeNetworkEgressTraffic,
	messageType.MessageSubTypeNetworkIngressPackets,
	messageType.MessageSubTypeNetworkEgressPackets,
	messageType.MessageSubTypeLoadAverage1min,
	messageType.MessageSubTypeLoadAverage5min,
	messageType.MessageSubTypeLoadAverage15min,
	messageType.MessageSubTypeTotalTCPConnections,
	messageType.MessageSubTypeEstablishedTCPConnections,
	messageType.MessageSubTypeTotalProcessCount,
}

func MessageEntity2Message(logger *log.Logger, msgEntity *messageDao.WebMessage, isDetail bool) *response.Message {
	if msgEntity == nil {
		return nil
	}
	res := &response.Message{
		WebMessage: *msgEntity,
	}

	if isDetail {

		if msgEntity.InstanceID != "" {
			instanceEntity, err := instanceDao.GetInstanceById(logger, msgEntity.InstanceID)
			if err != nil {
				logger.Warnf("MessageEntity2Message.GetInstanceById error, instanceId:%s, error:%s", msgEntity.InstanceID, err.Error())
			} else {
				res.IDcID = instanceEntity.IDcID
				res.DeviceID = instanceEntity.DeviceID
				idcEntity, err := idcDao.GetIdcById(logger, instanceEntity.IDcID)
				if err != nil {
					logger.Warnf("MessageEntity2Message.GetIdcById error, idcId:%s, error:%s", instanceEntity.IDcID, err.Error())
				} else {
					res.IDcName = idcEntity.Name
					if logger.GetPoint("language").(string) == baseLogic.EN_US {
						res.IDcName = idcEntity.NameEn
					}
				}
				res.ProjectID = instanceEntity.ProjectID
				projectEntity, err := projectDao.GetProjectById(logger, instanceEntity.ProjectID)
				if err != nil {
					logger.Warnf("MessageEntity2Message.GetProjectById error, ProjectId:%s, error:%s", instanceEntity.ProjectID, err.Error())
				} else {
					res.ProjectName = projectEntity.ProjectName
				}
			}
		} else if msgEntity.SN != "" {
			deviceEntity, err := deviceDao.GetDeviceBySn(logger, msgEntity.SN)
			if err != nil {
				logger.Warnf("MessageEntity2Message.GetDeviceBySn error, sn:%s, error:%s", msgEntity.SN, err.Error())
			} else {
				res.IDcID = deviceEntity.IDcID
				res.DeviceID = deviceEntity.DeviceID
				idcEntity, err := idcDao.GetIdcById(logger, deviceEntity.IDcID)
				if err != nil {
					logger.Warnf("MessageEntity2Message.GetIdcById error, idcId:%s, error:%s", deviceEntity.IDcID, err.Error())
				} else {
					res.IDcName = idcEntity.Name
					if logger.GetPoint("language").(string) == baseLogic.EN_US {
						res.IDcName = idcEntity.NameEn
					}
				}

			}
		}

	}
	return res

}

func HasUnreadMessage(logger *log.Logger) (bool, error) {
	userId := logger.GetPoint("userId").(string)
	q := map[string]interface{}{
		"is_del":   0,
		"user_id":  userId,
		"has_read": 0,
	}
	cnt, err := messageDao.GetWebMessagesByUserIdCount(logger, q)
	if err != nil {
		logger.Warnf("HasUnreadMessage.GetWebMessagesByUserIdCount error, userid:%s, error:%s", userId, err.Error())
		return false, err
	}
	return cnt != 0, nil
}

func GetPageMessages(logger *log.Logger, param *requestTypes.QueryMessagesRequest, p util.Pageable) ([]*response.Message, int64, error) {

	userId := logger.GetPoint("userId").(string)
	offset, limit := p.Pageable2offset()

	q := map[string]interface{}{
		"is_del":  0,
		"user_id": userId,
	}
	if param.HasRead != "" {
		r, err := strconv.Atoi(param.HasRead)
		if err != nil {
			logger.Warnf("GetPageMessages.HasRead param error, user_id:%s, error:%s", userId, err.Error())
			return nil, 0, errors.New("param hasRead error")
		}
		q["has_read"] = r
	}
	if param.MessageType != "" {
		mt := ShowEn2Types[param.MessageType]
		if mt == "" {
			mt = ShowZh2Types[param.MessageType]
		}
		if mt == "" {
			mt = param.MessageType
		}
		q["message_type"] = mt
	}
	if param.MessageSubType != "" {
		mts := []string{}
		items := strings.Split(param.MessageSubType, ",")
		for _, item := range items {
			mt := ShowEn2Types[item]
			if mt == "" {
				mt = ShowZh2Types[item]
			}
			if mt == "" {
				mt = item
			}
			mts = append(mts, mt)
		}

		q["message_sub_type.in"] = mts
	} else {
		q["message_sub_type.in"] = IncludeMessagetypes
	}

	if param.Detail != "" {
		q["detail.like"] = "%" + strings.TrimSpace(param.Detail) + "%"
	}
	cnt, err := messageDao.GetWebMessagesByUserIdCount(logger, q)
	if err != nil {
		logger.Warnf("GetPageMessages.GetWebMessagesByUserIdCount error, user_id:%s, error:%s", userId, err.Error())
		return nil, 0, err
	}

	var entityList []*messageDao.WebMessage
	if param.IsAll == baseLogic.IS_ALL {
		entityList, err = messageDao.GetMultiWebMessagesByUserId(logger, q, []string{}, []string{"id"}, []string{"desc"}, 0, 0)
	} else {
		entityList, err = messageDao.GetMultiWebMessagesByUserId(logger, q, []string{}, []string{"id"}, []string{"desc"}, offset, limit)
	}
	if err != nil {
		logger.Warn("GetPageMessages GetMultiWebMessagesByUserId sql error:", err.Error())
		return nil, 0, err
	}
	res := []*response.Message{}
	for _, v := range entityList {
		var mt, mst string
		if logger.GetPoint("language").(string) == baseLogic.EN_US {
			mt = Types2ShowEn[v.MessageType]
			mst = Types2ShowEn[v.MessageSubType]
		} else {
			mt = Types2ShowZh[v.MessageType]
			mst = Types2ShowZh[v.MessageSubType]
		}
		if mt == "" {
			mt = v.MessageType
		}
		if mst == "" {
			mst = v.MessageSubType
		}
		v.MessageType = mt
		v.MessageSubType = mst
		v.FaultType = v.MessageSubType
		res = append(res, MessageEntity2Message(logger, v, false))
	}

	return res, cnt, nil

}

func GetMessageStatistic(logger *log.Logger) (int64, int64, error) {

	userId := logger.GetPoint("userId").(string)

	q1 := map[string]interface{}{
		"is_del":              0,
		"user_id":             userId,
		"message_sub_type.in": IncludeMessagetypes,
	}
	q2 := map[string]interface{}{
		"is_del":              0,
		"user_id":             userId,
		"has_read":            0,
		"message_sub_type.in": IncludeMessagetypes,
	}

	all, err := messageDao.GetWebMessagesByUserIdCount(logger, q1)
	if err != nil {
		logger.Warnf("GetMessageStatistic.GetWebMessagesByUserIdCount error, user_id:%s, error:%s", userId, err.Error())
		return 0, 0, err
	}
	unread, err := messageDao.GetWebMessagesByUserIdCount(logger, q2)
	if err != nil {
		logger.Warnf("GetMessageStatistic.GetWebMessagesByUserIdCount error, user_id:%s, error:%s", userId, err.Error())
		return 0, 0, err
	}
	return all, unread, nil

}

func ReadMessages(logger *log.Logger, param *requestTypes.ReadMessagesRequest) (bool, error) {
	userId := logger.GetPoint("userId").(string)
	//校验messages和userid的归属关系
	q := map[string]interface{}{
		"is_del":        0,
		"user_id":       userId,
		"message_id.in": param.MessageIds,
	}

	msgs, err := messageDao.GetAllWebMessages(logger, q)
	if err != nil {
		logger.Warnf("ReadMessages.GetAllWebMessages error, query:%v, error:%s", q, err.Error())
	}
	if len(msgs) < len(param.MessageIds) {
		logger.Warnf("ReadMessages.message_ids invalided, user_id:%s, message_ids:%s", userId, param.MessageIds)
		panic(constant.BuildInvalidArgumentWithArgs("参数messageIds非法", "param messageIds invalid"))
	}

	u := map[string]interface{}{
		"has_read": 1,
	}
	if err := messageDao.UpdateWebMessages(logger, q, u); err != nil {
		logger.Warnf("ReadMessages.UpdateWebMessages error, query:%v, error:%s", q, err.Error())
		return false, err
	}
	return true, nil
}

func GetMessageById(logger *log.Logger, param *requestTypes.GetMessageByIdRequest) (*response.Message, string, string, error) {

	var nextMessageId, prevMessageId string

	userId := logger.GetPoint("userId").(string)
	//校验messages和userid的归属关系
	q := map[string]interface{}{
		"is_del":     0,
		"user_id":    userId,
		"message_id": param.MessageID,
	}

	msg, err := messageDao.GetOneWebMessage(logger, q, "asc")
	if err != nil {
		logger.Warnf("GetOneMoreMessage.GetOneWebMessage error, user_id:%s, message_id:%s, error:%s", userId, param.MessageID, err.Error())
		return nil, "", "", err
	}
	if msg != nil {

		if logger.GetPoint("language").(string) == baseLogic.EN_US {
			msg.MessageType = Types2ShowEn[msg.MessageType]
			msg.MessageSubType = Types2ShowEn[msg.MessageSubType]
		} else {
			msg.MessageType = Types2ShowZh[msg.MessageType]
			msg.MessageSubType = Types2ShowZh[msg.MessageSubType]
		}
		msg.FaultType = msg.MessageSubType
		if msg.UserName == "" {
			msg.UserName = "admin"
		}

		mid := msg.ID
		nextQ := map[string]interface{}{
			"is_del":  0,
			"user_id": userId,
			"id.gt":   mid,
		}
		nextMsg, err := messageDao.GetOneWebMessage(logger, nextQ, "asc")
		if err != nil {
			logger.Warnf("GetOneMoreMessage.getNextmsg error:%s", err.Error())
		}

		prevQ := map[string]interface{}{
			"is_del":  0,
			"user_id": userId,
			"id.lt":   mid,
		}
		prevMsg, err := messageDao.GetOneWebMessage(logger, prevQ, "desc")
		if err != nil {
			logger.Warnf("GetOneMoreMessage.getPrevmsg error:%s", err.Error())
		}

		if nextMsg != nil {
			nextMessageId = nextMsg.MessageID
		}
		if prevMsg != nil {
			prevMessageId = prevMsg.MessageID
		}
		return MessageEntity2Message(logger, msg, true), nextMessageId, prevMessageId, nil

	}
	return nil, "", "", nil

}

func GetMessageTypes(logger *log.Logger) response.MessageTypesResp {
	if logger.GetPoint("language").(string) == baseLogic.EN_US {
		return response.MessageTypesResp{
			Types2ShowEn[messageType.MessageTypeOperation]: []string{
				// Types2ShowEn[messageType.MessageSubTypeCollectHardwareInfo],
				// Types2ShowEn[messageType.MessageSubTypeCreateDevices],
				Types2ShowEn[messageType.MessageSubTypeCreateInstances],
				// Types2ShowEn[messageType.MessageSubTypeCutDeviceStock],
				Types2ShowEn[messageType.MessageSubTypeDeleteInstance],
				Types2ShowEn[messageType.MessageSubTypeInstanceResetPassword],
				// Types2ShowEn[messageType.MessageSubTypePutawayDevice],
				Types2ShowEn[messageType.MessageSubTypeReinstallInstance],
				Types2ShowEn[messageType.MessageSubTypeRestartInstances],
				Types2ShowEn[messageType.MessageSubTypeStartInstances],
				Types2ShowEn[messageType.MessageSubTypeStopInstances],
			},
			// Types2ShowEn[messageType.MessageTypeSystem]: []string{
			// 	Types2ShowEn[messageType.MessageSubTypeLicenseNearTimeout],
			// },
			Types2ShowEn[messageType.MessageTypeOobMonitor]: []string{
				Types2ShowEn[messageType.MessageSubTypePowerFault],
				Types2ShowEn[messageType.MessageSubTypeTemperatureFault],
				Types2ShowEn[messageType.MessageSubTypeCPUFault],
				Types2ShowEn[messageType.MessageSubTypeMemFault],
				Types2ShowEn[messageType.MessageSubTypeDiskFault],
				Types2ShowEn[messageType.MessageSubTypeOtherFault],
			},
			Types2ShowEn[messageType.MessageTypeInBondAlert]: []string{
				Types2ShowEn[messageType.MessageSubTypeCPUUsage],
				Types2ShowEn[messageType.MessageSubTypeMemoryUsed],
				Types2ShowEn[messageType.MessageSubTypeDiskUsed],
				Types2ShowEn[messageType.MessageSubTypeDiskUsage],
				Types2ShowEn[messageType.MessageSubTypeDiskReadTraffic],
				Types2ShowEn[messageType.MessageSubTypeDiskWriteTraffic],
				Types2ShowEn[messageType.MessageSubTypeDiskReadIOPS],
				Types2ShowEn[messageType.MessageSubTypeDiskWriteIOPS],
				Types2ShowEn[messageType.MessageSubTypeNetworkIngressTraffic],
				Types2ShowEn[messageType.MessageSubTypeNetworkEgressTraffic],
				Types2ShowEn[messageType.MessageSubTypeNetworkIngressPackets],
				Types2ShowEn[messageType.MessageSubTypeNetworkEgressPackets],
				Types2ShowEn[messageType.MessageSubTypeLoadAverage1min],
				Types2ShowEn[messageType.MessageSubTypeLoadAverage5min],
				Types2ShowEn[messageType.MessageSubTypeLoadAverage15min],
				Types2ShowEn[messageType.MessageSubTypeTotalTCPConnections],
				Types2ShowEn[messageType.MessageSubTypeEstablishedTCPConnections],
				Types2ShowEn[messageType.MessageSubTypeTotalProcessCount],
			},
		}
	} else {
		return response.MessageTypesResp{
			Types2ShowZh[messageType.MessageTypeOperation]: []string{
				// Types2ShowZh[messageType.MessageSubTypeCollectHardwareInfo],
				// Types2ShowZh[messageType.MessageSubTypeCreateDevices],
				Types2ShowZh[messageType.MessageSubTypeCreateInstances],
				// Types2ShowZh[messageType.MessageSubTypeCutDeviceStock],
				Types2ShowZh[messageType.MessageSubTypeDeleteInstance],
				Types2ShowZh[messageType.MessageSubTypeInstanceResetPassword],
				// Types2ShowZh[messageType.MessageSubTypePutawayDevice],
				Types2ShowZh[messageType.MessageSubTypeReinstallInstance],
				Types2ShowZh[messageType.MessageSubTypeRestartInstances],
				Types2ShowZh[messageType.MessageSubTypeStartInstances],
				Types2ShowZh[messageType.MessageSubTypeStopInstances],
			},
			// Types2ShowZh[messageType.MessageTypeSystem]: []string{
			// 	Types2ShowZh[messageType.MessageSubTypeLicenseNearTimeout],
			// },
			Types2ShowZh[messageType.MessageTypeOobMonitor]: []string{
				Types2ShowZh[messageType.MessageSubTypePowerFault],
				Types2ShowZh[messageType.MessageSubTypeTemperatureFault],
				Types2ShowZh[messageType.MessageSubTypeCPUFault],
				Types2ShowZh[messageType.MessageSubTypeMemFault],
				Types2ShowZh[messageType.MessageSubTypeDiskFault],
				Types2ShowZh[messageType.MessageSubTypeOtherFault],
			},
			Types2ShowZh[messageType.MessageTypeInBondAlert]: []string{
				Types2ShowZh[messageType.MessageSubTypeCPUUsage],
				Types2ShowZh[messageType.MessageSubTypeMemoryUsed],
				Types2ShowZh[messageType.MessageSubTypeDiskUsed],
				Types2ShowZh[messageType.MessageSubTypeDiskUsage],
				Types2ShowZh[messageType.MessageSubTypeDiskReadTraffic],
				Types2ShowZh[messageType.MessageSubTypeDiskWriteTraffic],
				Types2ShowZh[messageType.MessageSubTypeDiskReadIOPS],
				Types2ShowZh[messageType.MessageSubTypeDiskWriteIOPS],
				Types2ShowZh[messageType.MessageSubTypeNetworkIngressTraffic],
				Types2ShowZh[messageType.MessageSubTypeNetworkEgressTraffic],
				Types2ShowZh[messageType.MessageSubTypeNetworkIngressPackets],
				Types2ShowZh[messageType.MessageSubTypeNetworkEgressPackets],
				Types2ShowZh[messageType.MessageSubTypeLoadAverage1min],
				Types2ShowZh[messageType.MessageSubTypeLoadAverage5min],
				Types2ShowZh[messageType.MessageSubTypeLoadAverage15min],
				Types2ShowZh[messageType.MessageSubTypeTotalTCPConnections],
				Types2ShowZh[messageType.MessageSubTypeEstablishedTCPConnections],
				Types2ShowZh[messageType.MessageSubTypeTotalProcessCount],
			},
		}
	}

}

func DeleteMessages(logger *log.Logger, param *requestTypes.DeleteMessagesRequest) (bool, error) {
	userId := logger.GetPoint("userId").(string)
	//校验messages和userid的归属关系
	q := map[string]interface{}{
		"is_del":        0,
		"user_id":       userId,
		"message_id.in": param.MessageIds,
	}

	msgs, err := messageDao.GetAllWebMessages(logger, q)
	if err != nil {
		logger.Warnf("DeleteMessages.GetAllWebMessages error, query:%v, error:%s", q, err.Error())
	}
	if len(msgs) < len(param.MessageIds) {
		logger.Warnf("DeleteMessages.message_ids invalided, user_id:%s, message_ids:%s", userId, param.MessageIds)
		panic(constant.BuildInvalidArgumentWithArgs("参数messageIds非法", "param messageIds invalid"))
	}

	u := map[string]interface{}{
		"is_del": 1,
	}
	if err := messageDao.UpdateWebMessages(logger, q, u); err != nil {
		logger.Warnf("DeleteMessages.UpdateWebMessages error, query:%v, error:%s", q, err.Error())
		return false, err
	}
	return true, nil
}

func AddWebMessage(logger *log.Logger, webMsg *messageDao.WebMessage) error {

	if _, err := messageDao.AddWebMessage(logger, webMsg); err != nil {
		logger.Warnf("AddWebMessage error, msg:%s, error:%s", util.InterfaceToJson(webMsg), err.Error())
		return err
	}
	logger.Infof("AddWebMessage success, msg:%s", util.InterfaceToJson(webMsg))
	return nil

}
