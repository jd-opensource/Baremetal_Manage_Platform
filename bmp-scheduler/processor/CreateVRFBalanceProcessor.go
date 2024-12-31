package processor

import (
	commandDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/commandDao"
	ironicAgentEvent "git.jd.com/cps-golang/ironic-common/ironic/event/command"
	log "git.jd.com/cps-golang/log"
)

type CreateVRFBalanceProcessor struct {
	BaseProcessor
}

func NewCreateVRFBalanceProcessor() CreateVRFBalanceProcessor {
	b := CreateVRFBalanceProcessor{}
	b.BaseProcessor.doProcess = b.doProcess
	b.BaseProcessor.afterProcess = b.afterProcess
	return b
}

func (b CreateVRFBalanceProcessor) Process(logger *log.Logger, command *commandDao.Command) {
	b.BaseProcessor.Process(logger, command)
}

func (b CreateVRFBalanceProcessor) Callback(logger *log.Logger, command *commandDao.Command, msg ironicAgentEvent.CallbackCommandMessage) {
	b.BaseProcessor.CallBack(logger, command, msg)
}

func (b CreateVRFBalanceProcessor) doProcess(logger *log.Logger, command *commandDao.Command) {
	//TODO
}

func (b CreateVRFBalanceProcessor) afterProcess(logger *log.Logger, command *commandDao.Command) {
	b.AfterSuccessExecute(logger, command)
}
