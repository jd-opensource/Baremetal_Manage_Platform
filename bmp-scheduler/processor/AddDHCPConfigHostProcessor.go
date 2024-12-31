package processor

import (
	"fmt"

	commandDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/commandDao"
	deviceDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/deviceDao"
	rabbitMq "coding.jd.com/aidc-bmp/bmp-scheduler/service/rabbit_mq"
	driverTypes "coding.jd.com/aidc-bmp/bmp-scheduler/types/driver"
	"coding.jd.com/aidc-bmp/bmp-scheduler/util"
	ironicAgentEvent "git.jd.com/cps-golang/ironic-common/ironic/event/command"
	log "git.jd.com/cps-golang/log"
)

type AddDHCPConfigHostProcessor struct {
	BaseProcessor
}

func NewAddDHCPConfigHostProcessor() AddDHCPConfigHostProcessor {
	b := AddDHCPConfigHostProcessor{}
	b.BaseProcessor.doProcess = b.doProcess
	b.BaseProcessor.afterProcess = b.afterProcess
	return b
}
func (b AddDHCPConfigHostProcessor) Process(logger *log.Logger, command *commandDao.Command) {
	b.BaseProcessor.Process(logger, command)
}

func (b AddDHCPConfigHostProcessor) Callback(logger *log.Logger, command *commandDao.Command, msg ironicAgentEvent.CallbackCommandMessage) {
	b.BaseProcessor.CallBack(logger, command, msg)
}

func (b AddDHCPConfigHostProcessor) doProcess(logger *log.Logger, command *commandDao.Command) {
	var entity *deviceDao.Device
	var err error
	entity, err = deviceDao.GetBySn(logger, command.Sn)

	if err != nil || entity == nil {

		logger.Warn("AddDHCPConfigHostProcessor GetBySn sql error:", err.Error())
		panic(ProcessAbortException{Msg: fmt.Sprintf("process AddDHCPConfigHostProcessor error, sn %s device not found", command.Sn)})
	}
	//entities, err := interfaceLogic.GetBySn(logger, command.Sn)
	//if err != nil {
	//	logger.Warn("AddDHCPConfigHostProcessor GetBySn sql error:", command.Sn, err.Error())
	//}
	//if len(entities) == 0 {
	//	panic(ProcessAbortException{Msg: fmt.Sprintf("process AddDHCPConfigHost error, sn %s interface not found", command.Sn)})
	//}
	//jd_bond_interface, err := jdBondInterfaceLogic.GetBySn(logger, command.Sn)
	//if err != nil {
	//	logger.Warn("AddDHCPConfigHostProcessor GetBySn sql error:", command.Sn, err.Error())
	//}
	//if jd_bond_interface == nil || jd_bond_interface.PrivateIP == "" {
	//	panic(ProcessAbortException{Msg: fmt.Sprintf("process AddDHCPConfigHost error, sn %s jd_bond_interface not found", command.Sn)})
	//}

	config_host := driverTypes.AddDHCPConfigHost{
		Sn:  command.Sn,
		Ip:  entity.PrivateIPv4,
		Mac: entity.Mac1,
	}
	if err := rabbitMq.SendToAgent4Topic(entity.IDcID, config_host); err != nil {
		logger.Warn("AddDHCPConfigHostProcessor SendToAgent error:", entity.IDcID, err.Error())
		return
	}

	logger.Infof("AddDHCPConfigHostProcessor_SendToAgent_Msg: %s, routingkey:%s", util.ObjToJson(config_host), entity.IDcID)
}

func (b AddDHCPConfigHostProcessor) afterProcess(logger *log.Logger, command *commandDao.Command) {
	//不管agent，driver是否返回成功，都默认成功，执行下一条指令
	b.AfterSuccessExecute(logger, command)
}
