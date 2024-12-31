package apiActor

import (
	"encoding/json"

	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/deviceLogic"
	"coding.jd.com/aidc-bmp/bmp-scheduler/template"
	rabbitIronicMsgApi "git.jd.com/cps-golang/ironic-common/ironic/event/api"
	log "git.jd.com/cps-golang/log"
)

type CutDeviceStockActor struct {
	BaseActor
}

func NewCutDeviceStockActor() CutDeviceStockActor {
	return CutDeviceStockActor{}
}

func (s CutDeviceStockActor) Do(logger *log.Logger, msg string, task string) {

	message := rabbitIronicMsgApi.CutDeviceStockMessage{}
	if err := json.Unmarshal([]byte(msg), &message); err != nil {
		logger.Warn("CutDeviceStockActor Unmarshal error:", err.Error())
		return
	}
	message.RequestId = logger.GetPoint("logid").(string)
	sns := message.Sns
	for i := 0; i < len(sns); i++ {
		if sns[i] == "" {
			continue
		}
		device_entity, err := deviceLogic.GetBySn(logger, sns[i])
		if err != nil {
			logger.Warn("CutDeviceStockActor GetBySn sql error:", sns[i], err.Error())
		}
		if device_entity != nil {
			logger.Warn(sns[i], " device_entity isNot del")
			continue
		}
		templater := template.CutDeviceStockTemplate{}
		//拆分成子任务并写到库中
		templater.InitCommand(logger, message.RequestId, sns[i], task)
		//执行第一条任务(发送到agent或者driver等下游，等下游执行完成后发送callback消息到mq，scheduler获取callback消息后再从库中取下一条指令执行)
		GetAndStartFirstCommand(logger, sns[i])

	}
}
