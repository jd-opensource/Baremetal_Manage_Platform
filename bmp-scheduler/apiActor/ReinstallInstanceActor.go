package apiActor

import (
	"encoding/json"
	"fmt"

	auditLogLogic "coding.jd.com/aidc-bmp/bmp-scheduler/logic/auditLogLogic"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/deviceLogic"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/imageLogic"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/instanceLogic"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/partitionLogic"
	"coding.jd.com/aidc-bmp/bmp-scheduler/service/redis"
	"coding.jd.com/aidc-bmp/bmp-scheduler/template"
	"coding.jd.com/aidc-bmp/bmp-scheduler/types"
	constants "git.jd.com/cps-golang/ironic-common/ironic/common/Constants"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/AuditLogsType"
	rabbitIronicMsgApi "git.jd.com/cps-golang/ironic-common/ironic/event/api"
	log "git.jd.com/cps-golang/log"
)

type ReinstallInstanceActor struct {
	BaseActor
}

func NewReinstallInstanceActor() ReinstallInstanceActor {
	return ReinstallInstanceActor{}
}

func (s ReinstallInstanceActor) Do(logger *log.Logger, msg string, task string) {

	message := rabbitIronicMsgApi.ReinstallInstanceMessage{}
	if err := json.Unmarshal([]byte(msg), &message); err != nil {
		logger.Warn("ReinstallInstanceActor Unmarshal error:", err.Error())
		return
	}
	message.RequestId = logger.GetPoint("logid").(string)
	instance_id := message.InstanceId
	instance, err := instanceLogic.GetByInstanceId(logger, instance_id)
	if err != nil {
		logger.Warn("ReinstallInstanceActor GetByInstanceId sql error:", instance_id, err.Error())
	}
	device, err := deviceLogic.GetById(logger, instance.DeviceID)
	if err != nil {
		logger.Warn("ReinstallInstanceActor GetById sql error:", instance.DeviceID, err.Error())
	}

	auditLogLogic.SaveAuditLogs(logger, device.Sn, message.InstanceId, AuditLogsType.AuditLogsReinstallInstance)

	instance_extra := types.InstanceExtra{
		InstanceId: instance_id,
		Password:   message.Password,
		UserData:   message.UserData,
		KeepData:   message.KeepData,
	}
	val, _ := json.Marshal(instance_extra)
	if err = redis.SetObjectToRedisWithExpire(fmt.Sprintf(constants.INSTANCE_EXTRA_KEY, instance_id), string(val), 60*60*24*2); err != nil {
		logger.Warn("ReinstallInstanceActor SetObjectToRedis error:", instance_id, err.Error())
	}
	partitions_entity, err := partitionLogic.GetByDeviceTypeAndImageId(logger, device.DeviceTypeID, instance.ImageID)
	if err != nil {
		logger.Warn("ReinstallInstanceActor GetByDeviceTypeAndImageId sql error:", device.DeviceTypeID, instance.ImageID, err.Error())
	}
	image_entity, err := imageLogic.GetByImageId(logger, instance.ImageID)
	if err != nil {
		logger.Warn("ReinstallInstanceActor GetByImageId sql error:", instance.ImageID, err.Error())
	}
	isMakePat := false
	if len(partitions_entity) > 0 {
		isMakePat = true
	}

	templater := template.ReinstallInstanceTemplateComposite{}
	//拆分成子任务并写到库中
	templater.InitCommand(logger, "instance.NetworkType", "instance.InterfaceMode", message.RequestId, instance.InstanceID, device.Sn, message.KeepData, isMakePat, image_entity.Format, task)
	s.sendEmail()
	GetAndStartFirstCommand(logger, device.Sn)
}

func (s *ReinstallInstanceActor) sendEmail() {
	//TODO
}
