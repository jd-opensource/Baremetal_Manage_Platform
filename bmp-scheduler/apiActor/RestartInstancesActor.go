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

type RestartInstancesActor struct {
	BaseActor
}

func NewRestartInstancesActor() RestartInstancesActor {
	return RestartInstancesActor{}
}

func (s RestartInstancesActor) Do(logger *log.Logger, msg string, task string) {

	message := rabbitIronicMsgApi.RestartInstancesMessage{}
	if err := json.Unmarshal([]byte(msg), &message); err != nil {
		logger.Warn("RestartInstancesActor Unmarshal error:", err.Error())
		return
	}
	message.RequestId = logger.GetPoint("logid").(string)
	instance_ids := message.InstanceIds
	for _, instance_id := range instance_ids {
		instance, err := instanceLogic.GetByInstanceId(logger, instance_id)
		if err != nil {
			logger.Warn("RestartInstancesActor_GetByInstanceId sql error:", instance_id, err.Error())
		}
		device, err := deviceLogic.GetById(logger, instance.DeviceID)
		if err != nil {
			logger.Warn("RestartInstancesActor_GetById sql error:", instance.DeviceID, err.Error())
		}

		auditLogLogic.SaveAuditLogs(logger, device.Sn, instance_id, AuditLogsType.AuditLogsRestartInstances)

		templater := template.RestartInstanceTemplate{}
		//拆分成子任务并写到库中
		templater.InitCommand(logger, message.RequestId, instance_id, device.Sn, task)
		GetAndStartFirstCommand(logger, device.Sn)
	}

}
