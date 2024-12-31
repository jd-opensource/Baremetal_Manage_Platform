package processor

import (
	commandDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/commandDao"
	ironicAgentEvent "git.jd.com/cps-golang/ironic-common/ironic/event/command"
	log "git.jd.com/cps-golang/log"
)

type SaveConfigToFtpServerProcessor struct {
	BaseProcessor
}

func NewSaveConfigToFtpServerProcessor() SaveConfigToFtpServerProcessor {
	b := SaveConfigToFtpServerProcessor{}
	b.BaseProcessor.doProcess = b.doProcess
	b.BaseProcessor.afterProcess = b.afterProcess
	return b
}

func (b SaveConfigToFtpServerProcessor) Process(logger *log.Logger, command *commandDao.Command) {
	b.BaseProcessor.Process(logger, command)
}

func (b SaveConfigToFtpServerProcessor) Callback(logger *log.Logger, command *commandDao.Command, msg ironicAgentEvent.CallbackCommandMessage) {
	b.BaseProcessor.CallBack(logger, command, msg)
}

func (b SaveConfigToFtpServerProcessor) doProcess(logger *log.Logger, command *commandDao.Command) {
	//TODO
}

func (b SaveConfigToFtpServerProcessor) afterProcess(logger *log.Logger, command *commandDao.Command) {
	b.AfterSuccessExecute(logger, command)
}
