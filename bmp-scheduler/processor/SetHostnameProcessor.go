package processor

import (
	"fmt"

	commandDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/commandDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/instanceLogic"
	rabbitMq "coding.jd.com/aidc-bmp/bmp-scheduler/service/rabbit_mq"
	agentTypes "coding.jd.com/aidc-bmp/bmp-scheduler/types/agent"
	"coding.jd.com/aidc-bmp/bmp-scheduler/util"
	ironicAgentEvent "git.jd.com/cps-golang/ironic-common/ironic/event/command"
	log "git.jd.com/cps-golang/log"
)

type SetHostnameProcessor struct {
	BaseProcessor
}

func NewSetHostnameProcessor() SetHostnameProcessor {
	b := SetHostnameProcessor{}
	b.BaseProcessor.doProcess = b.doProcess
	b.BaseProcessor.afterProcess = b.afterProcess
	return b
}

func (b SetHostnameProcessor) Process(logger *log.Logger, command *commandDao.Command) {
	b.BaseProcessor.Process(logger, command)
}

func (b SetHostnameProcessor) Callback(logger *log.Logger, command *commandDao.Command, msg ironicAgentEvent.CallbackCommandMessage) {
	b.BaseProcessor.CallBack(logger, command, msg)
}

func (b SetHostnameProcessor) doProcess(logger *log.Logger, command *commandDao.Command) {

	instance, err := instanceLogic.GetByInstanceId(logger, command.InstanceID)
	if err != nil {
		logger.Warn("SetHostnameProcessor GetByInstanceId sql error:", command.InstanceID, err.Error())
	}
	if instance == nil {
		panic(ProcessAbortException{Msg: fmt.Sprintf("process SetHostname error, instance %s not found", command.InstanceID)})
	}

	set_hostname := agentTypes.SetHostname{
		Sn:       command.Sn,
		Hostname: instance.Hostname,
	}
	if err := rabbitMq.SendToAgent(command.Sn, set_hostname); err != nil {
		logger.Warnf("SetHostnameProcessor SendToAgent error, msg:%s, error:%s", util.ObjToJson(set_hostname), err.Error())
		return
	}
	logger.Infof("SetHostnameProcessor_SendToAgent_Msg success, msg:%s", util.ObjToJson(set_hostname))
}

func (b SetHostnameProcessor) afterProcess(logger *log.Logger, command *commandDao.Command) {
	//empty
}
