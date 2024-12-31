package apiActor

import (
	"encoding/json"
	"fmt"

	"coding.jd.com/aidc-bmp/bmp-scheduler/service/redis"
	"coding.jd.com/aidc-bmp/bmp-scheduler/types"
	constants "git.jd.com/cps-golang/ironic-common/ironic/common/Constants"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/AuditLogsType"

	auditLogLogic "coding.jd.com/aidc-bmp/bmp-scheduler/logic/auditLogLogic"
	"coding.jd.com/aidc-bmp/bmp-scheduler/template"
	rabbitIronicMsgApi "git.jd.com/cps-golang/ironic-common/ironic/event/api"
	log "git.jd.com/cps-golang/log"
)

type CollectHardwareInfoActor struct {
	BaseActor
}

func NewCollectHardwareInfoActor() CollectHardwareInfoActor {
	return CollectHardwareInfoActor{}
}

func (s CollectHardwareInfoActor) Do(logger *log.Logger, msg string, task string) {
	logger.Info("CollectHardwareInfoActor do starting......", msg, task)
	defer logger.Info("CollectHardwareInfoActor do ending......", msg, task)
	message := rabbitIronicMsgApi.CollectHardwareInfoMessage{}
	if err := json.Unmarshal([]byte(msg), &message); err != nil {
		logger.Warn("CollectHardwareInfoActor Unmarshal error:", err.Error())
		return
	}
	message.RequestId = logger.GetPoint("logid").(string)
	saveCollectActionInfo(logger, message)
	sns := message.Sns
	c := make(chan int, 100) //并发度=100
	for i := 0; i < len(sns); i++ {
		c <- 1
		go func(i int) {
			defer func() {
				<-c
			}()

			auditLogLogic.SaveAuditLogs(logger, sns[i], "", AuditLogsType.AuditLogsCollectHardwareInfo)

			templater := template.CollectHardwareInfoTemplate{}
			//拆分成子任务并写到库中 status状态为wait
			templater.InitCommand(logger, message.RequestId, sns[i], message.NetworkType, message.AllowOverride, task)

			//执行第一条任务(发送到agent或者driver等下游，等下游执行完成后发送callback消息到mq，scheduler获取callback消息后再从库中取下一条指令执行)
			GetAndStartFirstCommand(logger, sns[i])

		}(i)
	}
}

// 保存信息到redis
func saveCollectActionInfo(logger *log.Logger, message rabbitIronicMsgApi.CollectHardwareInfoMessage) {
	for _, sn := range message.Sns {
		instance_extra := types.CollectInfoExtra{
			NetworkType: message.NetworkType,
			Sn:          sn,
		}
		val, err := json.Marshal(instance_extra)
		if err != nil {
			logger.Warn("saveCollectActionInfo Marshal error:", err.Error())
			return
		}
		if err := redis.SetObjectToRedisWithExpire(fmt.Sprintf(constants.COLLECT_EXTRA_KEY, sn), string(val), 60*60*24*15); err != nil {
			logger.Warn("saveCollectActionInfo SetObjectToRedisWithExpire error:", err.Error())
		}
	}
}
