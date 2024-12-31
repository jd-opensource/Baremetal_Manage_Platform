package auditLogLogic

import (
	"strings"
	"time"

	"coding.jd.com/aidc-bmp/bmp-openapi/dao/auditLogsDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/deviceDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/userDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/baseLogic"
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
	responseTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/response"
	"coding.jd.com/aidc-bmp/bmp-openapi/util"
	log "coding.jd.com/aidc-bmp/bmp_log"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/AuditLogsType"
)

var auditLogsTypeEnMap = map[string]string{
	AuditLogsType.AuditLogsPutaway:               "putaway",
	AuditLogsType.AuditLogsUnPutaway:             "unPutaway",
	AuditLogsType.AuditLogsStartInstances:        "startInstance",
	AuditLogsType.AuditLogsStopInstances:         "stopInstance",
	AuditLogsType.AuditLogsRestartInstances:      "restartInstance",
	AuditLogsType.AuditLogsResetInstanceStatus:   "resetInstanceStatus",
	AuditLogsType.AuditLogsRemoveDevice:          "removeDevice",
	AuditLogsType.AuditLogsInstanceResetPassword: "resetInstancePassword",
	AuditLogsType.AuditLogsReinstallInstance:     "reinstallInstance",
	AuditLogsType.AuditLogsDeleteDevice:          "deleteDevice",
	AuditLogsType.AuditLogsEditInstanceName:      "editInstanceName",
	AuditLogsType.AuditLogsCollectHardwareInfo:   "collectDevice",
	AuditLogsType.AuditLogsTypeCreateInstances:   "createInstance",
	AuditLogsType.AuditLogsDeleteInstance:        "deleteInstance",
	AuditLogsType.AuditLogsTypeLockInstances:     "lockInstance",
}

//
var auditLogsTypeZhMap = map[string]string{
	AuditLogsType.AuditLogsPutaway:               "设备上架",
	AuditLogsType.AuditLogsUnPutaway:             "设备下架",
	AuditLogsType.AuditLogsStartInstances:        "设备开机",
	AuditLogsType.AuditLogsStopInstances:         "设备关机",
	AuditLogsType.AuditLogsRestartInstances:      "设备重启",
	AuditLogsType.AuditLogsResetInstanceStatus:   "重置实例状态",
	AuditLogsType.AuditLogsRemoveDevice:          "设备移除",
	AuditLogsType.AuditLogsInstanceResetPassword: "重置密码",
	AuditLogsType.AuditLogsReinstallInstance:     "重装系统",
	AuditLogsType.AuditLogsDeleteDevice:          "设备删除",
	AuditLogsType.AuditLogsEditInstanceName:      "编辑实例名称",
	AuditLogsType.AuditLogsCollectHardwareInfo:   "设备采集",
	AuditLogsType.AuditLogsTypeCreateInstances:   "创建实例",
	AuditLogsType.AuditLogsDeleteInstance:        "实例删除",
	AuditLogsType.AuditLogsTypeLockInstances:     "实例锁定",
}

func SaveAuditLogs(logger *log.Logger, deviceId string, instanceId string, msgType string) error {

	userID := logger.GetPoint("userId").(string)
	var username string
	if userID != "" {
		userEntity, err := userDao.GetUserById(logger, userID)
		if err != nil {
			logger.Warnf("SaveAuditLogs.GetUserById error, user_id:%s, error:%s", userID, err.Error())
		} else {
			username = userEntity.UserName
		}
	}

	deviceEntity, err := deviceDao.GetDeviceById(logger, deviceId)
	if err != nil {
		logger.Warnf("SaveAuditLogs.GetBySn error, device_id:%s, error:%s", deviceId, err.Error())
	}

	al := &auditLogsDao.AuditLogs{
		Logid:           util.GetUuid("al-"),
		Sn:              deviceEntity.Sn,
		DeviceID:        deviceEntity.DeviceID,
		Operation:       msgType,
		OperateUserID:   userID,
		OperateUserName: username,
		CreatedTime:     int(time.Now().Unix()),
		Result:          "success",
	}
	_, err = auditLogsDao.AddAuditLogs(logger, al)
	if err != nil {
		logger.Warnf("SaveAuditLogs.AddAuditLogs error, al:%s, error:%s", util.ObjToJson(al), err.Error())
	} else {
		logger.Warnf("SaveAuditLogs.AddAuditLogs success, al:%s", util.ObjToJson(al))
	}
	return err
}

func DeleteAuditLogsByDeviceId(logger *log.Logger, deviceId string) error {

	q := map[string]interface{}{
		"is_del":    0,
		"device_id": deviceId,
	}
	u := map[string]interface{}{
		"is_del": 1,
	}
	err := auditLogsDao.UpdateMultiAuditLogs(logger, q, u)
	if err != nil {
		logger.Warnf("DeleteAuditLogsByDeviceId.UpdateMultiAuditLogs error, deviceId:%s, error:%s", deviceId, err.Error())
	} else {
		logger.Warnf("DeleteAuditLogsByDeviceId.UpdateMultiAuditLogs success, deviceId:%s", deviceId)
	}
	return err
}

func QueryAuditLogs(logger *log.Logger, param *requestTypes.DescribeAuditLogsRequest, p util.Pageable) ([]*responseTypes.AuditLog, int64, error) {
	offset, limit := p.Pageable2offset()
	query := map[string]interface{}{
		"sn":     param.Sn,
		"is_del": 0,
	}
	if param.Operation != "" {
		query["operation.in"] = strings.Split(param.Operation, ",")
	}
	if param.UserName != "" {
		query["operate_user_name"] = param.UserName
	}
	if param.Result != "" {
		query["result.in"] = strings.Split(param.Result, ",")
	}
	if param.StartTime != 0 {
		query["created_time.gte"] = param.StartTime
	}
	if param.EndTime != 0 {
		query["updated_time.lte"] = param.EndTime
	}

	count, err := auditLogsDao.GetAuditLogsCount(logger, query)
	if err != nil {
		logger.Warnf("QueryAuditLogs.GetAuditLogsCount sql error, query:%s, error:%s", util.ObjToJson(query), err.Error())
		return nil, 0, err
	}
	dts := []*auditLogsDao.AuditLogs{}
	if param.IsAll == baseLogic.IS_ALL {
		dts, err = auditLogsDao.GetAllAuditLogs(logger, query)
	} else {
		dts, err = auditLogsDao.QueryAuditLogs(logger, query, offset, limit)
	}
	if err != nil {
		logger.Warnf("QueryAuditLogs sql err, param:%s, error:%s", util.ObjToJson(param), err.Error())
		return nil, 0, err
	}
	res := []*responseTypes.AuditLog{}

	for _, value := range dts {
		v := AuditLogEntity2AuditLog(logger, value)

		res = append(res, v)
	}
	return res, count, err
}

func AuditLogEntity2AuditLog(logger *log.Logger, param *auditLogsDao.AuditLogs) *responseTypes.AuditLog {
	res := &responseTypes.AuditLog{
		ID:          param.ID,
		LogID:       param.Logid,
		Operation:   param.Operation,
		UserName:    param.OperateUserName,
		UserID:      param.OperateUserID,
		Sn:          param.Sn,
		DeviceID:    param.DeviceID,
		InstanceID:  param.InstanceID,
		Result:      param.Result,
		FailReason:  param.FailReason,
		OperateTime: param.CreatedTime,
		FinishTime:  param.UpdatedTime,
	}
	var reason string
	if logger.GetPoint("language").(string) == baseLogic.EN_US {
		res.OperationName = auditLogsTypeEnMap[param.Operation]
		reason = baseLogic.InstanceReasonEn[res.FailReason]
	} else {
		res.OperationName = auditLogsTypeZhMap[param.Operation]
		reason = baseLogic.InstanceReason[res.FailReason]
	}
	if reason == "" {
		reason = res.FailReason
	}
	res.FailReason = reason

	return res
}

/*
AuditLogsPutaway               = "PutawayDevice"         //上架
	AuditLogsUnPutaway             = "UnPutawayDevice"       //下架
	AuditLogsStartInstances        = "StartInstances"        //开机
	AuditLogsStopInstances         = "StopInstances"         //关机
	AuditLogsRestartInstances      = "RestartInstances"      //重启
	AuditLogsResetInstanceStatus   = "ResetInstances"        //重置实例状态（新增）
	AuditLogsRemoveDevice          = "RemoveDevice"          //设备移除（新增）
	AuditLogsDeleteDevice          = "DeleteDevice"          //设备删除（新增）
	AuditLogsDeleteInstance        = "DeleteInstance"        //回收实例
	AuditLogsInstanceResetPassword = "InstanceResetPassword" //重置密码
	AuditLogsReinstallInstance     = "ReinstallInstance"     //重装系统
	AuditLogsEditInstanceName      = "EditInstanceName"      //编辑实例名称（新增）
	AuditLogsCollectHardwareInfo   = "CollectHardwareInfo"   //设备采集
	AuditLogsTypeCreateInstances   = "CreateInstances"       //创建实例
	AuditLogsTypeLockInstances     = "LockInstances"         //锁定实例(新增)
*/

func DescribeAuditLogsTypes(logger *log.Logger) []responseTypes.AuditLogsType {
	if logger.GetPoint("language").(string) == baseLogic.EN_US {
		return []responseTypes.AuditLogsType{
			{
				Operation: AuditLogsType.AuditLogsPutaway,
				Name:      auditLogsTypeEnMap[AuditLogsType.AuditLogsPutaway],
			},
			{
				Operation: AuditLogsType.AuditLogsUnPutaway,
				Name:      auditLogsTypeEnMap[AuditLogsType.AuditLogsUnPutaway],
			},
			{
				Operation: AuditLogsType.AuditLogsStartInstances,
				Name:      auditLogsTypeEnMap[AuditLogsType.AuditLogsStartInstances],
			},
			{
				Operation: AuditLogsType.AuditLogsStopInstances,
				Name:      auditLogsTypeEnMap[AuditLogsType.AuditLogsStopInstances],
			},
			{
				Operation: AuditLogsType.AuditLogsRestartInstances,
				Name:      auditLogsTypeEnMap[AuditLogsType.AuditLogsRestartInstances],
			},
			{
				Operation: AuditLogsType.AuditLogsResetInstanceStatus,
				Name:      auditLogsTypeEnMap[AuditLogsType.AuditLogsResetInstanceStatus],
			},
			{
				Operation: AuditLogsType.AuditLogsRemoveDevice,
				Name:      auditLogsTypeEnMap[AuditLogsType.AuditLogsRemoveDevice],
			},
			{
				Operation: AuditLogsType.AuditLogsInstanceResetPassword,
				Name:      auditLogsTypeEnMap[AuditLogsType.AuditLogsInstanceResetPassword],
			},
			{
				Operation: AuditLogsType.AuditLogsReinstallInstance,
				Name:      auditLogsTypeEnMap[AuditLogsType.AuditLogsReinstallInstance],
			},
			{
				Operation: AuditLogsType.AuditLogsDeleteDevice,
				Name:      auditLogsTypeEnMap[AuditLogsType.AuditLogsDeleteDevice],
			},
			{
				Operation: AuditLogsType.AuditLogsEditInstanceName,
				Name:      auditLogsTypeEnMap[AuditLogsType.AuditLogsEditInstanceName],
			},
			{
				Operation: AuditLogsType.AuditLogsCollectHardwareInfo,
				Name:      auditLogsTypeEnMap[AuditLogsType.AuditLogsCollectHardwareInfo],
			},
			{
				Operation: AuditLogsType.AuditLogsTypeCreateInstances,
				Name:      auditLogsTypeEnMap[AuditLogsType.AuditLogsTypeCreateInstances],
			},
			{
				Operation: AuditLogsType.AuditLogsDeleteInstance,
				Name:      auditLogsTypeEnMap[AuditLogsType.AuditLogsDeleteInstance],
			},
		}
	} else {

		return []responseTypes.AuditLogsType{
			{
				Operation: AuditLogsType.AuditLogsPutaway,
				Name:      auditLogsTypeZhMap[AuditLogsType.AuditLogsPutaway],
			},
			{
				Operation: AuditLogsType.AuditLogsUnPutaway,
				Name:      auditLogsTypeZhMap[AuditLogsType.AuditLogsUnPutaway],
			},
			{
				Operation: AuditLogsType.AuditLogsStartInstances,
				Name:      auditLogsTypeZhMap[AuditLogsType.AuditLogsStartInstances],
			},
			{
				Operation: AuditLogsType.AuditLogsStopInstances,
				Name:      auditLogsTypeZhMap[AuditLogsType.AuditLogsStopInstances],
			},
			{
				Operation: AuditLogsType.AuditLogsRestartInstances,
				Name:      auditLogsTypeZhMap[AuditLogsType.AuditLogsRestartInstances],
			},
			{
				Operation: AuditLogsType.AuditLogsResetInstanceStatus,
				Name:      auditLogsTypeZhMap[AuditLogsType.AuditLogsResetInstanceStatus],
			},
			{
				Operation: AuditLogsType.AuditLogsRemoveDevice,
				Name:      auditLogsTypeZhMap[AuditLogsType.AuditLogsRemoveDevice],
			},
			{
				Operation: AuditLogsType.AuditLogsInstanceResetPassword,
				Name:      auditLogsTypeZhMap[AuditLogsType.AuditLogsInstanceResetPassword],
			},
			{
				Operation: AuditLogsType.AuditLogsReinstallInstance,
				Name:      auditLogsTypeZhMap[AuditLogsType.AuditLogsReinstallInstance],
			},
			{
				Operation: AuditLogsType.AuditLogsDeleteDevice,
				Name:      auditLogsTypeZhMap[AuditLogsType.AuditLogsDeleteDevice],
			},
			{
				Operation: AuditLogsType.AuditLogsEditInstanceName,
				Name:      auditLogsTypeZhMap[AuditLogsType.AuditLogsEditInstanceName],
			},
			{
				Operation: AuditLogsType.AuditLogsCollectHardwareInfo,
				Name:      auditLogsTypeZhMap[AuditLogsType.AuditLogsCollectHardwareInfo],
			},
			{
				Operation: AuditLogsType.AuditLogsTypeCreateInstances,
				Name:      auditLogsTypeZhMap[AuditLogsType.AuditLogsTypeCreateInstances],
			},
			{
				Operation: AuditLogsType.AuditLogsDeleteInstance,
				Name:      auditLogsTypeZhMap[AuditLogsType.AuditLogsDeleteInstance],
			},
		}
	}
}
