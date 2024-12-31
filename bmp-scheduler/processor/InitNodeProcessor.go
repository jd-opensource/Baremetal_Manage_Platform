package processor

import (
	"fmt"

	commandDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/commandDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/deviceHintsDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/instanceDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/imageLogic"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/osLogic"
	rabbitMq "coding.jd.com/aidc-bmp/bmp-scheduler/service/rabbit_mq"
	agentTypes "coding.jd.com/aidc-bmp/bmp-scheduler/types/agent"
	"coding.jd.com/aidc-bmp/bmp-scheduler/types/domain"
	"coding.jd.com/aidc-bmp/bmp-scheduler/util"
	ironicAgentEvent "git.jd.com/cps-golang/ironic-common/ironic/event/command"
	log "git.jd.com/cps-golang/log"
)

type InitNodeProcessor struct {
	BaseProcessor
}

func NewInitNodeProcessor() InitNodeProcessor {
	b := InitNodeProcessor{}
	b.BaseProcessor.doProcess = b.doProcess
	b.BaseProcessor.afterProcess = b.afterProcess
	return b
}

func (b InitNodeProcessor) Process(logger *log.Logger, command *commandDao.Command) {
	b.BaseProcessor.Process(logger, command)
}

func (b InitNodeProcessor) Callback(logger *log.Logger, command *commandDao.Command, msg ironicAgentEvent.CallbackCommandMessage) {
	b.BaseProcessor.CallBack(logger, command, msg)
}

func (b InitNodeProcessor) doProcess(logger *log.Logger, command *commandDao.Command) {

	logger.Info("InitNodeProcessor doProcess......", command.Sn)
	instanceEntity, err := instanceDao.GetInstanceByUuid(logger, command.InstanceID)
	if instanceEntity == nil {
		logger.Warnf("InitNodeProcessor GetInstanceByUuid error, sn:%s, instance_id:%s, err:%+v", command.Sn, command.InstanceID, err)
		panic(ProcessAbortException{Msg: fmt.Sprintf("InitNodeProcessor error, instance %s not found", command.InstanceID)})
	}
	imageEntity, err := imageLogic.GetByImageId(logger, instanceEntity.ImageID)
	if imageEntity == nil {
		logger.Warnf("InitNodeProcessor GetByImageId error, instance_id:%s, image_id:%s, err:%+v", command.InstanceID, instanceEntity.ImageID, err)
		panic(ProcessAbortException{Msg: fmt.Sprintf("InitNodeProcessor error, image %s not found", instanceEntity.ImageID)})
	}
	osEntity, err := osLogic.GetByOsId(logger, imageEntity.OsID)
	if osEntity == nil {
		logger.Warnf("InitNodeProcessor GetOsByuuid error, instance_id:%s, os_id:%s, err:%+v", command.InstanceID, imageEntity.OsID, err)
		panic(ProcessAbortException{Msg: fmt.Sprintf("InitNodeProcessor error, os %s not found", imageEntity.OsID)})
	}

	rootDeviceHintsEntity, err := deviceHintsDao.GetSystemVolumnDeviceHintsBySn(logger, command.Sn)

	initNode := agentTypes.InitNode{
		Sn: command.Sn,
		NodeData: agentTypes.NodeData{
			OsType:    osEntity.OsType,
			OsVersion: osEntity.OsVersion,
		},
	}
	if rootDeviceHintsEntity == nil {
		logger.Warnf("InitNodeProcessor GetSystemVolumnDeviceHintsBySn error, sn:%s, err:%+v", command.Sn, err)
	} else {
		initNode.NodeData.RootDeviceHints = domain.RootDeviceHints{
			Serial: rootDeviceHintsEntity.Serial,
		}
	}

	if err := rabbitMq.SendToAgent(command.Sn, initNode); err != nil {
		logger.Warnf("InitNodeProcessor SendToAgent error, msg:%s, error:%s", util.ObjToJson(initNode), err.Error())
		return
	}
	logger.Infof("InitNodeProcessor_SendToAgent_Msg success, msg:%s", util.ObjToJson(initNode))

}

func (b InitNodeProcessor) afterProcess(logger *log.Logger, command *commandDao.Command) {
	//empty
}
