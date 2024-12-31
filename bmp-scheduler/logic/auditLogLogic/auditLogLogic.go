package auditLogLogic

import (
	"time"

	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/auditLogsDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/deviceDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/userDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/util"
	log "git.jd.com/cps-golang/log"
)

func SaveAuditLogs(logger *log.Logger, sn string, instanceId string, msgType string) error {

	userID := logger.GetPoint("user_id").(string)
	var username string
	if userID != "" {
		userEntity, err := userDao.GetUserById(logger, userID)
		if err != nil {
			logger.Warnf("SaveAuditLogs.GetUserById error, user_id:%s, error:%s", userID, err.Error())
		} else {
			username = userEntity.UserName
		}
	}

	deviceEntity, err := deviceDao.GetBySn(logger, sn)
	if err != nil {
		logger.Warnf("SaveAuditLogs.GetBySn error, sn:%s, error:%s", sn, err.Error())
	}

	al := &auditLogsDao.AuditLogs{
		Logid:           util.GetUuid("al-"),
		Sn:              sn,
		InstanceID:      instanceId,
		DeviceID:        deviceEntity.DeviceID,
		Operation:       msgType,
		OperateUserID:   userID,
		OperateUserName: username,
		CreatedTime:     int(time.Now().Unix()),
		Result:          "doing",
	}
	_, err = auditLogsDao.AddAuditLogs(logger, al)
	if err != nil {
		logger.Warnf("SaveAuditLogs.AddAuditLogs error, al:%s, error:%s", util.ObjToJson(al), err.Error())
	} else {
		logger.Warnf("SaveAuditLogs.AddAuditLogs success, al:%s", util.ObjToJson(al))
	}
	return err
}

func SetAuditLogToSuccess(logger *log.Logger, sn string, instanceId string, msgType string) error {

	q := map[string]interface{}{
		"is_del":      0,
		"sn":          sn,
		"instance_id": instanceId,
		"operation":   msgType,
	}
	al, err := auditLogsDao.GetOneUnfinishedAuditLogs(logger, sn, instanceId, msgType)
	if err != nil {
		logger.Warnf("SetAuditLogToSuccess.GetOneAuditLogs error, query:%s, error:%s", util.ObjToJson(q), err.Error())
		return err
	}
	al.Result = "success"
	al.UpdatedTime = int(time.Now().Unix())

	if err := auditLogsDao.UpdateAuditLogsById(logger, al); err != nil {
		logger.Warnf("SetAuditLogToSuccess.UpdateAuditLogsById error, query:%s, error:%s", util.ObjToJson(q), err.Error())
		return err
	}
	return nil
}

func SetAuditLogToFail(logger *log.Logger, sn string, instanceId string, msgType string, reason string) error {

	q := map[string]interface{}{
		"is_del":      0,
		"sn":          sn,
		"instance_id": instanceId,
		"operation":   msgType,
	}
	al, err := auditLogsDao.GetOneUnfinishedAuditLogs(logger, sn, instanceId, msgType)
	if err != nil {
		logger.Warnf("SetAuditLogToSuccess.GetOneAuditLogs error, query:%s, error:%s", util.ObjToJson(q), err.Error())
		return err
	}
	al.Result = "fail"
	al.FailReason = reason
	al.UpdatedTime = int(time.Now().Unix())

	if err := auditLogsDao.UpdateAuditLogsById(logger, al); err != nil {
		logger.Warnf("SetAuditLogToFail.UpdateAuditLogsById error, query:%s, error:%s", util.ObjToJson(q), err.Error())
		return err
	}
	return nil
}
