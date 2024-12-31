package processor

import (
	commandDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/commandDao"
	rabbitMq "coding.jd.com/aidc-bmp/bmp-scheduler/service/rabbit_mq"
	agentTypes "coding.jd.com/aidc-bmp/bmp-scheduler/types/agent"
	"coding.jd.com/aidc-bmp/bmp-scheduler/util"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/TimeoutPolicy"
	ironicAgentEvent "git.jd.com/cps-golang/ironic-common/ironic/event/command"
	log "git.jd.com/cps-golang/log"
)

type SetUserCmdProcessor struct {
	BaseProcessor
}

func NewSetUserCmdProcessor() SetUserCmdProcessor {
	b := SetUserCmdProcessor{}
	b.BaseProcessor.doProcess = b.doProcess
	b.BaseProcessor.afterProcess = b.afterProcess
	return b
}

func (b SetUserCmdProcessor) Process(logger *log.Logger, command *commandDao.Command) {
	b.BaseProcessor.Process(logger, command)
}

func (b SetUserCmdProcessor) Callback(logger *log.Logger, command *commandDao.Command, msg ironicAgentEvent.CallbackCommandMessage) {
	b.BaseProcessor.CallBack(logger, command, msg)
}

func (b SetUserCmdProcessor) doProcess(logger *log.Logger, command *commandDao.Command) {

	logger.Info("SetUserCmdProcessor doProcess starting......", command.InstanceID)
	instance_extra := loadInstanceExtra(logger, command.InstanceID)
	if instance_extra == nil || instance_extra.UserData == "" {
		logger.Info("SetUserCmdCommand timeoutPolicy set to skip", command.InstanceID)
		command.TimeoutPolicy = TimeoutPolicy.SKIP
		return
	}
	user_cmd := agentTypes.SetUserCmd{
		Sn:      command.Sn,
		Content: instance_extra.UserData,
	}
	if err := rabbitMq.SendToAgent(command.Sn, user_cmd); err != nil {
		logger.Warnf("SetUserCmdProcessor SendToAgent error, msg:%s, error:%s", util.ObjToJson(user_cmd), err.Error())
		return
	}
	logger.Infof("SetUserCmdProcessor_SendToAgent_Msg success, msg:%s", util.ObjToJson(user_cmd))
}

func (b SetUserCmdProcessor) afterProcess(logger *log.Logger, command *commandDao.Command) {
	b.AfterSuccessExecute(logger, command)
}
