package apiActor

import (
	"encoding/json"

	auditLogLogic "coding.jd.com/aidc-bmp/bmp-scheduler/logic/auditLogLogic"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/deviceLogic"
	"coding.jd.com/aidc-bmp/bmp-scheduler/template"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/AuditLogsType"
	rabbitIronicMsgApi "git.jd.com/cps-golang/ironic-common/ironic/event/api"
	log "git.jd.com/cps-golang/log"
)

type PutawayDeviceActor struct {
	BaseActor
}

func NewPutawayDeviceActor() PutawayDeviceActor {
	return PutawayDeviceActor{}
}

func (s PutawayDeviceActor) Do(logger *log.Logger, msg string, task string) {

	message := rabbitIronicMsgApi.PutawayDeviceMessage{}
	if err := json.Unmarshal([]byte(msg), &message); err != nil {
		logger.Warn("StopInstancesActor Unmarshal error:", err.Error())
		return
	}
	message.RequestId = logger.GetPoint("logid").(string)

	device, err := deviceLogic.GetBySn(logger, message.Sn)
	if err != nil {
		logger.Warn("PutawayDeviceActor deviceLogic.GetBySn sql error:", message.Sn, err.Error())
	}

	auditLogLogic.SaveAuditLogs(logger, device.Sn, "", AuditLogsType.AuditLogsPutaway)

	templater := template.PutawayDeviceTemplate{}
	//拆分成子任务并写到库中
	templater.InitCommand(logger, message.RequestId, "", message.Sn, task)
	GetAndStartFirstCommand(logger, device.Sn)
}
