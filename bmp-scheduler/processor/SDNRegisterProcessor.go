package processor

import (
	commandDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/commandDao"
	ironicAgentEvent "git.jd.com/cps-golang/ironic-common/ironic/event/command"
	log "git.jd.com/cps-golang/log"
)

type SDNRegisterProcessor struct {
	BaseProcessor
}

func NewSDNRegisterProcessor() SDNRegisterProcessor {
	b := SDNRegisterProcessor{}
	b.BaseProcessor.doProcess = b.doProcess
	b.BaseProcessor.afterProcess = b.afterProcess
	return b
}

func (b SDNRegisterProcessor) Process(logger *log.Logger, command *commandDao.Command) {
	b.BaseProcessor.Process(logger, command)
}

func (b SDNRegisterProcessor) Callback(logger *log.Logger, command *commandDao.Command, msg ironicAgentEvent.CallbackCommandMessage) {
	b.BaseProcessor.CallBack(logger, command, msg)
}

func (b SDNRegisterProcessor) doProcess(logger *log.Logger, command *commandDao.Command) {
	//TODO
}

func (b SDNRegisterProcessor) afterProcess(logger *log.Logger, command *commandDao.Command) {
	b.AfterSuccessExecute(logger, command)
}
