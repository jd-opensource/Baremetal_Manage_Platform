package processor

import (
	commandDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/commandDao"
	ironicAgentEvent "git.jd.com/cps-golang/ironic-common/ironic/event/command"
	log "git.jd.com/cps-golang/log"
)

type SDNAddAliasIPProcessor struct {
	BaseProcessor
}

func NewSDNAddAliasIPProcessor() SDNAddAliasIPProcessor {
	b := SDNAddAliasIPProcessor{}
	b.BaseProcessor.doProcess = b.doProcess
	b.BaseProcessor.afterProcess = b.afterProcess
	return b
}

func (b SDNAddAliasIPProcessor) Process(logger *log.Logger, command *commandDao.Command) {
	b.BaseProcessor.Process(logger, command)
}

func (b SDNAddAliasIPProcessor) Callback(logger *log.Logger, command *commandDao.Command, msg ironicAgentEvent.CallbackCommandMessage) {
	b.BaseProcessor.CallBack(logger, command, msg)
}

func (b SDNAddAliasIPProcessor) doProcess(logger *log.Logger, command *commandDao.Command) {
	//TODO
}

func (b SDNAddAliasIPProcessor) afterProcess(logger *log.Logger, command *commandDao.Command) {
	b.AfterSuccessExecute(logger, command)
}
