package processor

import (
	"fmt"

	commandDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/commandDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/instanceLogic"
	ironicAgentEvent "git.jd.com/cps-golang/ironic-common/ironic/event/command"
	log "git.jd.com/cps-golang/log"
)

type SetKeypairsProcessor struct {
	BaseProcessor
}

func NewSetKeypairsProcessor() SetKeypairsProcessor {
	b := SetKeypairsProcessor{}
	b.BaseProcessor.doProcess = b.doProcess
	b.BaseProcessor.afterProcess = b.afterProcess
	return b
}

func (b SetKeypairsProcessor) Process(logger *log.Logger, command *commandDao.Command) {
	b.BaseProcessor.Process(logger, command)
}

func (b SetKeypairsProcessor) Callback(logger *log.Logger, command *commandDao.Command, msg ironicAgentEvent.CallbackCommandMessage) {
	b.BaseProcessor.CallBack(logger, command, msg)
}

func (b SetKeypairsProcessor) doProcess(logger *log.Logger, command *commandDao.Command) {

	instance, err := instanceLogic.GetByInstanceId(logger, command.InstanceID)
	if err != nil {
		logger.Warn("SetKeypairsProcessor GetByInstanceId sql error:", command.InstanceID, err.Error())
	}
	if instance == nil {
		panic(ProcessAbortException{Msg: fmt.Sprintf("process SetKeypairs error, instance %s not found", command.InstanceID)})
	}
	//if instance.KeypairID == "" {
	//	logger.Info("SetKeypairsCommand timeoutPolicy set to skip", command.InstanceID)
	//	command.TimeoutPolicy = TimeoutPolicy.SKIP
	//	return
	//}
	//
	//keypair, err := keypairLogic.GetByUuidAndTenant(logger, instance.KeypairID, instance.Tenant)
	//if err != nil {
	//	logger.Warn("SetKeypairsProcessor GetByUuidAndTenant sql error:", instance.KeypairID, instance.Tenant, err.Error())
	//}
	//if keypair == nil {
	//	panic(ProcessAbortException{Msg: fmt.Sprintf("process SetKeypairs error, keypair %s not found", instance.KeypairID)})
	//}
	//
	//set_keypairs := agentTypes.SetKeypairs{
	//	Sn:        command.Sn,
	//	PublicKey: keypair.PublicKey,
	//}
	//if err := rabbitMq.SendToAgent(command.Sn, set_keypairs); err != nil {
	//	logger.Warn("SetKeypairsProcessor SendToAgent error, msg:%s, error:%s", util.ObjToJson(set_keypairs), err.Error())
	//	return
	//}
	//logger.Infof("SetKeypairsProcessor_SendToAgent_Msg success, msg:%s", util.ObjToJson(set_keypairs))
}

func (b SetKeypairsProcessor) afterProcess(logger *log.Logger, command *commandDao.Command) {
	b.AfterSuccessExecute(logger, command)

}
