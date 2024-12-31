package apiActor

import (
	"encoding/json"
	"fmt"

	auditLogLogic "coding.jd.com/aidc-bmp/bmp-scheduler/logic/auditLogLogic"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/deviceLogic"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/instanceLogic"
	"coding.jd.com/aidc-bmp/bmp-scheduler/service/redis"
	"coding.jd.com/aidc-bmp/bmp-scheduler/template"
	"coding.jd.com/aidc-bmp/bmp-scheduler/types"
	constants "git.jd.com/cps-golang/ironic-common/ironic/common/Constants"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/AuditLogsType"
	rabbitIronicMsgApi "git.jd.com/cps-golang/ironic-common/ironic/event/api"
	log "git.jd.com/cps-golang/log"
)

type InstanceResetPasswordActor struct {
	BaseActor
}

func NewInstanceResetPasswordActor() InstanceResetPasswordActor {
	return InstanceResetPasswordActor{}
}

func (s InstanceResetPasswordActor) Do(logger *log.Logger, msg string, task string) {

	message := rabbitIronicMsgApi.InstanceResetPasswordMessage{}
	if err := json.Unmarshal([]byte(msg), &message); err != nil {
		logger.Warn("InstanceResetPasswordActor Unmarshal error:", err.Error())
		return
	}
	message.RequestId = logger.GetPoint("logid").(string)
	instance_id := message.InstanceId
	instance, err := instanceLogic.GetByInstanceId(logger, instance_id)
	if err != nil {
		logger.Warn("InstanceResetPasswordActor GetByInstanceId sql error:", instance_id, err.Error())
	}
	device, err := deviceLogic.GetById(logger, instance.DeviceID)
	if err != nil {
		logger.Warn("InstanceResetPasswordActor GetById sql error:", instance.DeviceID, err.Error())
	}
	if message.Action == "Monitor" {
		sendResetPasswordMessage()
		return
	}

	auditLogLogic.SaveAuditLogs(logger, device.Sn, message.InstanceId, AuditLogsType.AuditLogsInstanceResetPassword)

	instance_extra := types.InstanceExtra{
		InstanceId: instance_id,
		Password:   message.Password,
	}
	val, _ := json.Marshal(instance_extra)
	if err = redis.SetObjectToRedisWithExpire(fmt.Sprintf(constants.INSTANCE_EXTRA_KEY, instance_id), string(val), 60*60*24*2); err != nil {
		logger.Warn("InstanceResetPasswordActor SetObjectToRedis error:", instance_id, err.Error())
	}
	templater := template.ResetPasswordTemplateComposite{}
	//拆分成子任务并写到库中
	templater.InitCommand(logger, "instance.NetworkType", message.RequestId, instance.InstanceID, device.Sn, task)
	GetAndStartFirstCommand(logger, device.Sn)

}

func sendResetPasswordMessage() {
	//TODO
}
