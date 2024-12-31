package apiActor

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/deviceDao"
	auditLogLogic "coding.jd.com/aidc-bmp/bmp-scheduler/logic/auditLogLogic"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/deviceTypeLogic"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/imageLogic"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/instanceLogic"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/osLogic"

	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/partitionLogic"
	"coding.jd.com/aidc-bmp/bmp-scheduler/service/redis"
	"coding.jd.com/aidc-bmp/bmp-scheduler/template"
	"coding.jd.com/aidc-bmp/bmp-scheduler/types"
	constants "git.jd.com/cps-golang/ironic-common/ironic/common/Constants"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/AuditLogsType"
	rabbitIronicMsgApi "git.jd.com/cps-golang/ironic-common/ironic/event/api"
	log "git.jd.com/cps-golang/log"
)

type CreateInstancesActor struct {
	BaseActor
}

func NewCreateInstancesActor() CreateInstancesActor {
	return CreateInstancesActor{}
}

func (s CreateInstancesActor) Do(logger *log.Logger, msg string, task string) {
	logger.Info("CreateInstancesActor do starting...", msg)
	defer logger.Info("CreateInstancesActor do ending...", msg)
	var message = rabbitIronicMsgApi.CreateInstancesMessage{}
	if err := json.Unmarshal([]byte(msg), &message); err != nil {
		logger.Warn("CreateInstancesActor Unmarshal error:", err.Error())
		return
	}
	message.RequestId = logger.GetPoint("logid").(string)
	for _, instance_id := range message.InstanceIds {
		instance, err := instanceLogic.GetByInstanceId(logger, instance_id)
		if err != nil {
			logger.Warn("CreateInstancesActor.GetByInstanceId error:", instance_id, err.Error())
			return
		}
		device, _ := deviceDao.GetDeviceById(logger, instance.DeviceID)
		v, _ := json.Marshal(device)
		logger.Info("the device is:", device.DeviceID, string(v))

		auditLogLogic.SaveAuditLogs(logger, device.Sn, instance_id, AuditLogsType.AuditLogsTypeCreateInstances)

		instance_extra := types.InstanceExtra{
			InstanceId: instance.InstanceID,
			Password:   message.Password,
			UserData:   message.UserData,
		}

		//如果要下载监控agent的话，在这里才确定了用哪个device，根据device的architecture来决定下载哪个版本的agent
		ud := instance_extra.UserData
		if ud != "" && device.Architecture == "aarch64" { //默认是x86_64的，如果是arm的需替换,后面需要把openapi-console的userdata逻辑移动过来
			decodeUd, err := base64.StdEncoding.DecodeString(ud)
			if err != nil {
				logger.Warnf("CreateInstancesActor.Base64Decode userData error:%s", err.Error())
			} else {
				decodeUdStr := string(decodeUd)
				decodeUdStr = strings.Replace(decodeUdStr, "bmp-agent.bin", "bmp-agent.bin.arm", -1)
				ud = base64.StdEncoding.EncodeToString([]byte(decodeUdStr))
			}
		}
		instance_extra.UserData = ud

		if message.AliasIps != nil {
			instance_extra.AliasIps = message.AliasIps.([]types.AliasIP)
		}
		if message.ExtensionAliasIps != nil {
			instance_extra.ExtensionAliasIps = message.ExtensionAliasIps.([]types.AliasIP)
		}
		val, _ := json.Marshal(instance_extra)
		err = redis.SetObjectToRedisWithExpire(fmt.Sprintf(constants.INSTANCE_EXTRA_KEY, instance_id), string(val), 60*60*24*2)
		if err != nil {
			logger.Warn("CreateInstancesActor SetObjectToRedis error:", instance_id, err.Error())
		}

		partitions_entity, err := partitionLogic.GetByDeviceTypeAndImageId(logger, instance.DeviceTypeID, instance.ImageID)
		if err != nil {
			logger.Warn("CreateInstancesActor_GetByDeviceTypeAndImageId sql error:", instance.DeviceTypeID, instance.ImageID, err.Error())
		}
		isMakePat := false
		if len(partitions_entity) > 0 {
			isMakePat = true
		}
		v1, _ := json.Marshal(partitions_entity)
		logger.Info("partitions_entity res:", string(v1))
		image_entity, err := imageLogic.GetByImageId(logger, instance.ImageID)
		if err != nil {
			logger.Warn("CreateInstancesActor_GetByImageId sql error:", err.Error())
		}
		os_entity, err := osLogic.GetByOsId(logger, image_entity.OsID)
		if err != nil {
			logger.Warn("CreateInstancesActor GetByOsId sql error:", image_entity.OsID, err.Error())
		}
		deviceTypeEntity, err := deviceTypeLogic.GetDeviceTypeByID(logger, device.DeviceTypeID)
		if err != nil {
			logger.Warn("CreateInstancesActor GetDeviceTypeByID sql error:", device.DeviceTypeID, err.Error())
		}
		if os_entity.OsType != "Windows" {
			templater := template.CreateInstanceTemplateComposite{}
			templater.InitCommand(logger, "basic", deviceTypeEntity.InterfaceMode, message.RequestId, instance_id, device.Sn, isMakePat, image_entity.Format, task)
		} else {
			templater := template.CreateInstanceTemplateCompositeForWindows{}
			templater.InitCommand(logger, "basic", deviceTypeEntity.InterfaceMode, message.RequestId, instance_id, device.Sn, isMakePat, image_entity.Format, task)
		}
		//拆分成子任务并写到库中
		GetAndStartFirstCommand(logger, device.Sn)
	}

}
