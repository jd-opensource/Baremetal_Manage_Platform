package apiActor

import (
	"encoding/json"
	"fmt"

	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/deviceLogic"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/instanceLogic"
	"coding.jd.com/aidc-bmp/bmp-scheduler/service/redis"
	"coding.jd.com/aidc-bmp/bmp-scheduler/template"
	"git.jd.com/cps-golang/ironic-common/ironic/common/Constants"
	rabbitIronicMsgApi "git.jd.com/cps-golang/ironic-common/ironic/event/api"
	log "git.jd.com/cps-golang/log"
)

type ModifyBandwidthActor struct {
	BaseActor
}

func NewModifyBandwidthActor() ModifyBandwidthActor {
	return ModifyBandwidthActor{}
}

func (s ModifyBandwidthActor) Do(logger *log.Logger, msg string, task string) {

	message := rabbitIronicMsgApi.ModifyBandwidthMessage{}
	if err := json.Unmarshal([]byte(msg), &message); err != nil {
		logger.Warn("ModifyBandwidthActor Unmarshal error:", err.Error())
		return
	}
	message.RequestId = logger.GetPoint("logid").(string)
	instance_key := fmt.Sprintf("%s:%s", Constants.REDIS_KEY_STATUS, message.InstanceId)
	if err := redis.SetObjectToRedisWithExpire(instance_key, message.InstanceStatus, Constants.TIMEOUT_INSTANCE_STATUS); err != nil {
		logger.Warn("ModifyBandwidthActor SetObjectToRedis error:", err.Error())
	}
	instance, err := instanceLogic.GetByInstanceId(logger, message.InstanceId)
	if err != nil {
		logger.Warn("ModifyBandwidthActor_GetByInstanceId sql error:", err.Error())
	}
	device, err := deviceLogic.GetById(logger, instance.DeviceID)
	if err != nil {
		logger.Warn("ModifyBandwidthActor_GetById sql error:", err.Error())
	}

	templater := template.ModifyBandwidthTemplate{}
	//拆分成子任务并写到库中
	templater.InitCommand(logger, message.RequestId, message.InstanceId, device.Sn, task)
	GetAndStartFirstCommand(logger, device.Sn)
}
