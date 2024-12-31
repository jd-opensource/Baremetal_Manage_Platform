package processor

import (
	"fmt"

	commandDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/commandDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/interfaceDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/instanceLogic"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/interfaceLogic"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/InterfaceType"
	ironicAgentEvent "git.jd.com/cps-golang/ironic-common/ironic/event/command"
	log "git.jd.com/cps-golang/log"
)

type AddArpStaticProcessor struct {
	BaseProcessor
}

func NewAddArpStaticProcessor() AddArpStaticProcessor {
	b := AddArpStaticProcessor{}
	b.BaseProcessor.doProcess = b.doProcess
	b.BaseProcessor.afterProcess = b.afterProcess
	return b
}

func (b AddArpStaticProcessor) Process(logger *log.Logger, command *commandDao.Command) {
	b.BaseProcessor.Process(logger, command)
}

func (b AddArpStaticProcessor) Callback(logger *log.Logger, command *commandDao.Command, msg ironicAgentEvent.CallbackCommandMessage) {
	b.BaseProcessor.CallBack(logger, command, msg)
}

func (b AddArpStaticProcessor) doProcess(logger *log.Logger, command *commandDao.Command) {
	instance, err := instanceLogic.GetByInstanceId(logger, command.InstanceID)
	if err != nil || instance == nil {
		panic(ProcessAbortException{Msg: fmt.Sprintf("process AddArpStatic error, instance %s not found", command.InstanceID)})
	}
	//if !util.Int82Bool(instance.EnableInternet) {
	//	logger.Warn(command.InstanceID, "instance not enable internet, skip process AddArpStatic")
	//	return
	//}
	entities, err := interfaceLogic.GetBySn(logger, command.Sn)
	if err != nil || len(entities) == 0 {
		panic(ProcessAbortException{Msg: fmt.Sprintf("process AddArpStatic error, instance %s interface not found", command.InstanceID)})
	}
	interface_map := map[string]*interfaceDao.Interface{}
	for _, entity := range entities {
		interface_map[entity.InterfaceType] = entity
	}
	wan := interface_map[InterfaceType.WAN]
	if wan == nil {
		panic(ProcessAbortException{Msg: fmt.Sprintf("process AddArpStatic error, instance %s wan not found", command.InstanceID)})
	}
	//TODO
	// joypaw_api.createArpStatic()
}

func (b AddArpStaticProcessor) afterProcess(logger *log.Logger, command *commandDao.Command) {
	b.AfterSuccessExecute(logger, command)
}
