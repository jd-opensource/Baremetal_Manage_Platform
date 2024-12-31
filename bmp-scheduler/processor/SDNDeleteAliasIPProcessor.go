package processor

import (
	commandDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/commandDao"
	ironicAgentEvent "git.jd.com/cps-golang/ironic-common/ironic/event/command"
	log "git.jd.com/cps-golang/log"
)

type SDNDeleteAliasIPProcessor struct {
	BaseProcessor
}

func NewSDNDeleteAliasIPProcessor() SDNDeleteAliasIPProcessor {
	b := SDNDeleteAliasIPProcessor{}
	b.BaseProcessor.doProcess = b.doProcess
	b.BaseProcessor.afterProcess = b.afterProcess
	return b
}

func (b SDNDeleteAliasIPProcessor) Process(logger *log.Logger, command *commandDao.Command) {
	b.BaseProcessor.Process(logger, command)
}

func (b SDNDeleteAliasIPProcessor) Callback(logger *log.Logger, command *commandDao.Command, msg ironicAgentEvent.CallbackCommandMessage) {
	b.BaseProcessor.CallBack(logger, command, msg)
}

func (b SDNDeleteAliasIPProcessor) doProcess(logger *log.Logger, command *commandDao.Command) {
	//TODO
}

func (b SDNDeleteAliasIPProcessor) afterProcess(logger *log.Logger, command *commandDao.Command) {
	b.AfterSuccessExecute(logger, command)
}
