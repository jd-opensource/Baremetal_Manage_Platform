package processor

import (
	commandDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/commandDao"
	rabbitMq "coding.jd.com/aidc-bmp/bmp-scheduler/service/rabbit_mq"
	agentTypes "coding.jd.com/aidc-bmp/bmp-scheduler/types/agent"
	"coding.jd.com/aidc-bmp/bmp-scheduler/util"
	ironicAgentEvent "git.jd.com/cps-golang/ironic-common/ironic/event/command"
	log "git.jd.com/cps-golang/log"
)

type PingProcessor struct {
	BaseProcessor
}

func NewPingProcessor() PingProcessor {
	b := PingProcessor{}
	b.BaseProcessor.doProcess = b.doProcess
	b.BaseProcessor.afterProcess = b.afterProcess
	return b
}

func (b PingProcessor) Process(logger *log.Logger, command *commandDao.Command) {
	b.BaseProcessor.Process(logger, command)
}

func (b PingProcessor) Callback(logger *log.Logger, command *commandDao.Command, msg ironicAgentEvent.CallbackCommandMessage) {
	b.BaseProcessor.CallBack(logger, command, msg)
}

func (b PingProcessor) doProcess(logger *log.Logger, command *commandDao.Command) {

	ping := agentTypes.Ping{
		Sn: command.Sn,
	}
	if err := rabbitMq.SendToAgent(command.Sn, ping); err != nil {
		logger.Warn("PingProcessor SendToAgent error:", command.Sn, err.Error())
		return
	}
	logger.Infof("PingProcessor_SendToAgent_Msg success, msg:%s", util.ObjToJson(ping))
}

func (b PingProcessor) afterProcess(logger *log.Logger, command *commandDao.Command) {
	//empty
}
