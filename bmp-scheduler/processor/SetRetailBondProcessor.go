package processor

import (
	commandDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/commandDao"
	ironicAgentEvent "git.jd.com/cps-golang/ironic-common/ironic/event/command"
	log "git.jd.com/cps-golang/log"
)

type SetRetailBondProcessor struct {
	BaseProcessor
}

func NewSetRetailBondProcessor() SetRetailBondProcessor {
	b := SetRetailBondProcessor{}
	b.BaseProcessor.doProcess = b.doProcess
	b.BaseProcessor.afterProcess = b.afterProcess
	return b
}

func (b SetRetailBondProcessor) Process(logger *log.Logger, command *commandDao.Command) {
	b.BaseProcessor.Process(logger, command)
}

func (b SetRetailBondProcessor) Callback(logger *log.Logger, command *commandDao.Command, msg ironicAgentEvent.CallbackCommandMessage) {
	b.BaseProcessor.CallBack(logger, command, msg)
}

func (b SetRetailBondProcessor) doProcess(logger *log.Logger, command *commandDao.Command) {
	//TODO
}

func (b SetRetailBondProcessor) afterProcess(logger *log.Logger, command *commandDao.Command) {
	//empty
}
