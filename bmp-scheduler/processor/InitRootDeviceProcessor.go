package processor

import (
	"fmt"

	commandDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/commandDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/imageLogic"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/instanceLogic"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/osLogic"
	rabbitMq "coding.jd.com/aidc-bmp/bmp-scheduler/service/rabbit_mq"
	agentTypes "coding.jd.com/aidc-bmp/bmp-scheduler/types/agent"
	"coding.jd.com/aidc-bmp/bmp-scheduler/util"
	ironicAgentEvent "git.jd.com/cps-golang/ironic-common/ironic/event/command"
	log "git.jd.com/cps-golang/log"
)

type InitRootDeviceProcessor struct {
	BaseProcessor
}

func NewInitRootDeviceProcessor() InitRootDeviceProcessor {
	b := InitRootDeviceProcessor{}
	b.BaseProcessor.doProcess = b.doProcess
	b.BaseProcessor.afterProcess = b.afterProcess
	return b
}

func (b InitRootDeviceProcessor) Process(logger *log.Logger, command *commandDao.Command) {
	b.BaseProcessor.Process(logger, command)
}

func (b InitRootDeviceProcessor) Callback(logger *log.Logger, command *commandDao.Command, msg ironicAgentEvent.CallbackCommandMessage) {
	b.BaseProcessor.CallBack(logger, command, msg)
}

func (b InitRootDeviceProcessor) doProcess(logger *log.Logger, command *commandDao.Command) {

	instance, err := instanceLogic.GetByInstanceId(logger, command.InstanceID)
	if err != nil {
		logger.Warn("InitRootDeviceProcessor GetByInstanceId sql error:", command.InstanceID, err.Error())
	}
	if instance == nil {
		panic(ProcessAbortException{Msg: fmt.Sprintf("process InitRootDevice error, instance %s not found", command.InstanceID)})
	}
	image_entity, err := imageLogic.GetByImageId(logger, instance.ImageID)
	if err != nil {
		logger.Warn("InitRootDeviceProcessor GetByImageId sql error:", instance.ImageID, err.Error())
	}
	if image_entity == nil {
		panic(ProcessAbortException{Msg: fmt.Sprintf("process InitRootDevice error, image %s not found", instance.ImageID)})
	}
	os_entity, err := osLogic.GetByOsId(logger, image_entity.OsID)
	if err != nil {
		logger.Warn("InitRootDeviceProcessor GetByOsId sql error:", image_entity.OsID, err.Error())
	}
	if os_entity == nil {
		panic(ProcessAbortException{Msg: fmt.Sprintf("process InitRootDevice error, os %s not found", image_entity.OsID)})
	}

	init_root_device := agentTypes.InitRootDevice{
		Sn:        command.Sn,
		OsType:    os_entity.OsType,
		OsVersion: os_entity.OsVersion,
	}
	if err := rabbitMq.SendToAgent(command.Sn, init_root_device); err != nil {
		logger.Warnf("InitRootDeviceProcessor SendToAgent error, msg:%s, error:%s", util.ObjToJson(init_root_device), err.Error())
		return
	}
	logger.Infof("InitRootDeviceProcessor_SendToAgent_Msg:%s", util.ObjToJson(init_root_device))
}

func (b InitRootDeviceProcessor) afterProcess(logger *log.Logger, command *commandDao.Command) {
	//empty
}
