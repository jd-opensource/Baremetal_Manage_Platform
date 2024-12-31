package processor

import (
	commandDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/commandDao"
	rabbitMq "coding.jd.com/aidc-bmp/bmp-scheduler/service/rabbit_mq"
	agentTypes "coding.jd.com/aidc-bmp/bmp-scheduler/types/agent"
	ironicAgentEvent "git.jd.com/cps-golang/ironic-common/ironic/event/command"
	log "git.jd.com/cps-golang/log"
)

type UploadSystemLogProcessor struct {
	BaseProcessor
}

func NewUploadSystemLogProcessor() UploadSystemLogProcessor {
	b := UploadSystemLogProcessor{}
	b.BaseProcessor.doProcess = b.doProcess
	b.BaseProcessor.afterProcess = b.afterProcess
	return b
}

func (b UploadSystemLogProcessor) Process(logger *log.Logger, command *commandDao.Command) {
	b.BaseProcessor.Process(logger, command)
}

func (b UploadSystemLogProcessor) Callback(logger *log.Logger, command *commandDao.Command, msg ironicAgentEvent.CallbackCommandMessage) {
	b.BaseProcessor.CallBack(logger, command, msg)
}

func (b UploadSystemLogProcessor) doProcess(logger *log.Logger, command *commandDao.Command) {

	upload_system_log := agentTypes.UploadSystemLog{
		Sn: command.Sn,
	}
	if err := rabbitMq.SendToAgent(command.Sn, upload_system_log); err != nil {
		logger.Warn("UploadSystemLogProcessor SendToAgent error:", command.Sn, err.Error())
		return
	}
	logger.Infof("UploadSystemLogProcessor_SendToAgent_Msg success, sn:", command.Sn)
}

func (b UploadSystemLogProcessor) afterProcess(logger *log.Logger, command *commandDao.Command) {
	//empty
}
