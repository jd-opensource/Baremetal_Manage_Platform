package processor

import (
	"fmt"

	commandDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/commandDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/deviceHintsDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/imageLogic"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/instanceLogic"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/osLogic"
	rabbitMq "coding.jd.com/aidc-bmp/bmp-scheduler/service/rabbit_mq"
	agentTypes "coding.jd.com/aidc-bmp/bmp-scheduler/types/agent"
	domainTypes "coding.jd.com/aidc-bmp/bmp-scheduler/types/domain"
	"coding.jd.com/aidc-bmp/bmp-scheduler/util"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/InstanceStatus"
	ironicAgentEvent "git.jd.com/cps-golang/ironic-common/ironic/event/command"
	log "git.jd.com/cps-golang/log"
)

type InitRootDeviceProcessorV2 struct {
	BaseProcessor
}

func NewInitRootDeviceProcessorV2() InitRootDeviceProcessorV2 {
	b := InitRootDeviceProcessorV2{}
	b.BaseProcessor.doProcess = b.doProcess
	b.BaseProcessor.afterProcess = b.afterProcess
	return b
}

func (b InitRootDeviceProcessorV2) Process(logger *log.Logger, command *commandDao.Command) {
	b.BaseProcessor.Process(logger, command)
}

func (b InitRootDeviceProcessorV2) Callback(logger *log.Logger, command *commandDao.Command, msg ironicAgentEvent.CallbackCommandMessage) {
	b.BaseProcessor.CallBack(logger, command, msg)
}

func (b InitRootDeviceProcessorV2) doProcess(logger *log.Logger, command *commandDao.Command) {

	instance, err := instanceLogic.GetByInstanceId(logger, command.InstanceID)
	if err != nil {
		logger.Warn("InitRootDeviceProcessorV2 GetByInstanceId sql error:", command.InstanceID, err.Error())
	}
	if instance == nil {
		panic(ProcessAbortException{Msg: fmt.Sprintf("process InitRootDeviceProcessorV2 error, instance %s not found", command.InstanceID)})
	}
	image_entity, err := imageLogic.GetByImageId(logger, instance.ImageID)
	if err != nil {
		logger.Warn("InitRootDeviceProcessorV2 GetByImageId sql error:", instance.ImageID, err.Error())
	}
	if image_entity == nil {
		panic(ProcessAbortException{Msg: fmt.Sprintf("process InitRootDeviceProcessorV2 error, image %s not found", instance.ImageID)})
	}
	os_entity, err := osLogic.GetByOsId(logger, image_entity.OsID)
	if err != nil {
		logger.Warn("InitRootDeviceProcessorV2 GetByOsId sql error:", image_entity.OsID, err.Error())
	}
	if os_entity == nil {
		panic(ProcessAbortException{Msg: fmt.Sprintf("process InitRootDeviceProcessorV2 error, os %s not found", image_entity.OsID)})
	}

	init_root_device := agentTypes.InitRootDeviceV2{
		Sn:        command.Sn,
		OsType:    os_entity.OsType,
		OsVersion: os_entity.OsVersion,
	}

	entity, err := deviceHintsDao.GetSystemVolumnDeviceHintsBySn(logger, command.Sn)
	if err != nil {
		logger.Warn("InitRootDeviceProcessorV2.GetSystemVolumnDeviceHintsBySn error:", command.Sn, err.Error())
	}
	if instance.Status != InstanceStatus.CREATING && entity != nil {
		rootDeviceHints := domainTypes.RootDeviceHints{
			Size:               entity.Size,
			Rotational:         util.Int82Bool(entity.Rotational),
			Wwn:                entity.Wwn,
			Name:               entity.Name,
			WwnVendorExtension: entity.WwnVendorExtension,
			WwnWithExtension:   entity.WwnWithExtension,
			Model:              entity.Model,
			Serial:             entity.Serial,
			Vendor:             entity.Vendor,
			Hctl:               entity.Hctl,
			ByPath:             entity.ByPath,
		}
		init_root_device.RootDeviceHints = rootDeviceHints
	}

	if err := rabbitMq.SendToAgent(command.Sn, init_root_device); err != nil {
		logger.Warnf("InitRootDeviceProcessorV2 SendToAgent error, sn:%s,msg:%s, error:%s", command.Sn, util.ObjToJson(init_root_device), err.Error())
		return
	}
	logger.Infof("InitRootDeviceProcessorV2_SendToAgent_Msg:%s", util.ObjToJson(init_root_device))
}

func (b InitRootDeviceProcessorV2) afterProcess(logger *log.Logger, command *commandDao.Command) {
	//empty
}
