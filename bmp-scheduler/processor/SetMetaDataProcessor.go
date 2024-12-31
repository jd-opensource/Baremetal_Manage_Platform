package processor

import (
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/commandDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/instanceLogic"
	rabbitMq "coding.jd.com/aidc-bmp/bmp-scheduler/service/rabbit_mq"
	agentTypes "coding.jd.com/aidc-bmp/bmp-scheduler/types/agent"
	"coding.jd.com/aidc-bmp/bmp-scheduler/util"
	"fmt"
	ironicAgentEvent "git.jd.com/cps-golang/ironic-common/ironic/event/command"
	log "git.jd.com/cps-golang/log"
)

type SetMetaDataProcessor struct {
	BaseProcessor
}

func NewSetMetaDataProcessor() SetMetaDataProcessor {
	b := SetMetaDataProcessor{}
	b.BaseProcessor.doProcess = b.doProcess
	b.BaseProcessor.afterProcess = b.afterProcess
	return b
}

func (b SetMetaDataProcessor) Process(logger *log.Logger, command *commandDao.Command) {
	b.BaseProcessor.Process(logger, command)
}

func (b SetMetaDataProcessor) Callback(logger *log.Logger, command *commandDao.Command, msg ironicAgentEvent.CallbackCommandMessage) {
	b.BaseProcessor.CallBack(logger, command, msg)
}

func (b SetMetaDataProcessor) doProcess(logger *log.Logger, command *commandDao.Command) {

	logger.Info("SetMetaDataProcessor doProcess......", command.Sn)
	instance, err := instanceLogic.GetByInstanceId(logger, command.InstanceID)
	if err != nil {
		logger.Warn("SetMetaDataProcessor GetByInstanceId sql error:", command.InstanceID, err.Error())
	}
	if instance == nil {
		panic(ProcessAbortException{Msg: fmt.Sprintf("process SetMetaData error, instance %s not found", command.InstanceID)})
	}
	SetMetaData := agentTypes.SetMetaData{
		Sn: command.Sn,
	}
	SetMetaData.MetaData.HostName = instance.Hostname
	SetMetaData.MetaData.Uuid = instance.InstanceID

	if err := rabbitMq.SendToAgent(command.Sn, SetMetaData); err != nil {
		logger.Warnf("SetMetaDataProcessor SendToAgent error, msg:%s, error:%s", util.ObjToJson(SetMetaData), err.Error())
		return
	}
	logger.Infof("SetMetaDataProcessor_SendToAgent_Msg success, msg:%s", util.ObjToJson(SetMetaData))
}

func (b SetMetaDataProcessor) afterProcess(logger *log.Logger, command *commandDao.Command) {
	//empty
}
