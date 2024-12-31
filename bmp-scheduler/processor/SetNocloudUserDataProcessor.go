package processor

import (
	commandDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/commandDao"
	rabbitMq "coding.jd.com/aidc-bmp/bmp-scheduler/service/rabbit_mq"
	agentTypes "coding.jd.com/aidc-bmp/bmp-scheduler/types/agent"
	"coding.jd.com/aidc-bmp/bmp-scheduler/util"
	ironicAgentEvent "git.jd.com/cps-golang/ironic-common/ironic/event/command"
	log "git.jd.com/cps-golang/log"
)

type SetNocloudUserDataProcessor struct {
	BaseProcessor
}

func NewSetNocloudUserDataProcessor() SetNocloudUserDataProcessor {
	b := SetNocloudUserDataProcessor{}
	b.BaseProcessor.doProcess = b.doProcess
	b.BaseProcessor.afterProcess = b.afterProcess
	return b
}

func (b SetNocloudUserDataProcessor) Process(logger *log.Logger, command *commandDao.Command) {
	b.BaseProcessor.Process(logger, command)
}

func (b SetNocloudUserDataProcessor) Callback(logger *log.Logger, command *commandDao.Command, msg ironicAgentEvent.CallbackCommandMessage) {
	b.BaseProcessor.CallBack(logger, command, msg)
}

func (b SetNocloudUserDataProcessor) doProcess(logger *log.Logger, command *commandDao.Command) {

	logger.Info("SetNocloudUserDataProcessor doProcess......", command.Sn)
	instance_extra := loadInstanceExtra(logger, command.InstanceID)

	SetNocloudUserData := agentTypes.SetNocloudUserData{
		Sn:       command.Sn,
		Version:  "2.0",
		UserData: "", // 用户自定义脚本
	}
	if instance_extra != nil && instance_extra.UserData != "" {
		SetNocloudUserData.UserData = instance_extra.UserData
	}

	if err := rabbitMq.SendToAgent(command.Sn, SetNocloudUserData); err != nil {
		logger.Warnf("SetNocloudUserDataProcessor SendToAgent error, msg:%s, error:%s", util.ObjToJson(SetNocloudUserData), err.Error())
		return
	}
	logger.Infof("SetNocloudUserDataProcessor_SendToAgent_Msg success, msg:%s", util.ObjToJson(SetNocloudUserData))

}

func (b SetNocloudUserDataProcessor) afterProcess(logger *log.Logger, command *commandDao.Command) {
	//empty
}
