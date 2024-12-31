package apiActor

import (
	"encoding/json"

	auditLogLogic "coding.jd.com/aidc-bmp/bmp-scheduler/logic/auditLogLogic"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/deviceLogic"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/instanceLogic"
	"coding.jd.com/aidc-bmp/bmp-scheduler/template"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/AuditLogsType"
	rabbitIronicMsgApi "git.jd.com/cps-golang/ironic-common/ironic/event/api"
	log "git.jd.com/cps-golang/log"
)

type DeleteInstanceActor struct {
	BaseActor
}

func NewDeleteInstanceActor() DeleteInstanceActor {
	return DeleteInstanceActor{}
}

func (s DeleteInstanceActor) Do(logger *log.Logger, msg string, task string) {

	message := rabbitIronicMsgApi.DeleteInstanceMessage{}
	if err := json.Unmarshal([]byte(msg), &message); err != nil {
		logger.Warn("DeleteInstanceActor Unmarshal error:", err.Error())
		return
	}
	message.RequestId = logger.GetPoint("logid").(string)
	instance, err := instanceLogic.GetByInstanceId(logger, message.InstanceId)
	if err != nil {
		logger.Warn("DeleteInstanceActor GetByInstanceId sql error:", message.InstanceId, err.Error())
	}
	device, err := deviceLogic.GetById(logger, instance.DeviceID)
	if err != nil {
		logger.Warn("DeleteInstanceActor GetById sql error:", instance.DeviceID, err.Error())
	}

	auditLogLogic.SaveAuditLogs(logger, device.Sn, message.InstanceId, AuditLogsType.AuditLogsDeleteInstance)

	templater := template.DeleteInstanceTemplateComposite{}
	//拆分成子任务并写到库中
	templater.InitCommand(logger, "instance.NetworkType", message.RequestId, instance.InstanceID, device.Sn, task)
	s.sendEmail()
	GetAndStartFirstCommand(logger, device.Sn)
}

func (s *DeleteInstanceActor) sendEmail() {
	//TODO
}
