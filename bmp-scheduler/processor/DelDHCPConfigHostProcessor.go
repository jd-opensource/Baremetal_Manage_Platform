package processor

import (
	"time"

	commandDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/commandDao"
	deviceDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/deviceDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/interfaceDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/jdBondInterfaceDao"
	rabbitMq "coding.jd.com/aidc-bmp/bmp-scheduler/service/rabbit_mq"
	driverTypes "coding.jd.com/aidc-bmp/bmp-scheduler/types/driver"
	"coding.jd.com/aidc-bmp/bmp-scheduler/util"
	ironicAgentEvent "git.jd.com/cps-golang/ironic-common/ironic/event/command"
	log "git.jd.com/cps-golang/log"
)

type DelDHCPConfigHostProcessor struct {
	BaseProcessor
}

func NewDelDHCPConfigHostProcessor() DelDHCPConfigHostProcessor {
	b := DelDHCPConfigHostProcessor{}
	b.BaseProcessor.doProcess = b.doProcess
	b.BaseProcessor.afterProcess = b.afterProcess
	return b
}

func (b DelDHCPConfigHostProcessor) Process(logger *log.Logger, command *commandDao.Command) {
	b.BaseProcessor.Process(logger, command)
}

func (b DelDHCPConfigHostProcessor) Callback(logger *log.Logger, command *commandDao.Command, msg ironicAgentEvent.CallbackCommandMessage) {
	b.BaseProcessor.CallBack(logger, command, msg)
}

func (b DelDHCPConfigHostProcessor) doProcess(logger *log.Logger, command *commandDao.Command) {

	var entity *deviceDao.Device
	var err error
	entity, err = deviceDao.GetBySn(logger, command.Sn)

	if err != nil || entity == nil {
		logger.Warn("DelDHCPConfigHostProcessor GetBySn error:", command.Sn, err.Error())
	}
	//if entity == nil {
	//	panic(ProcessAbortException{Msg: fmt.Sprintf("process DelDHCPConfigHost error, sn %s device not found", command.Sn)})
	//}
	//entities, err := interfaceLogic.GetBySn(logger, command.Sn)
	//if err != nil {
	//	logger.Warn("DelDHCPConfigHostProcessor GetBySn sql error:", command.Sn, err.Error())
	//}
	//if len(entities) == 0 {
	//	panic(ProcessAbortException{Msg: fmt.Sprintf("process DelDHCPConfigHost error, sn %s interface not found", command.Sn)})
	//}
	//jd_bond_interface, err := jdBondInterfaceLogic.GetBySn(logger, command.Sn)
	//if err != nil {
	//	logger.Warn("DelDHCPConfigHostProcessor GetBySn sql error:", command.Sn, err.Error())
	//}
	//if jd_bond_interface != nil {
	//
	//	// delJdBondInterface(logger, jd_bond_interface)
	//}
	config_host := driverTypes.DelDHCPConfigHost{
		Sn:  command.Sn,
		Ip:  entity.PrivateIPv4, //.PrivateIP,
		Mac: entity.Mac1,
	}
	if err := rabbitMq.SendToAgent4Topic(entity.IDcID, config_host); err != nil {
		logger.Warnf("DelDHCPConfigHostProcessor SendToAgent error, routingkey:%s, msg:%s, error:%s", entity.IDcID, util.ObjToJson(config_host), err.Error())
		return
	}
	logger.Infof("DelDHCPConfigHostProcessor_SendToAgent_Msg:%s, routingkey:%s", util.ObjToJson(config_host), entity.IDcID)
	// delInterface(logger, entities)
}

func (b DelDHCPConfigHostProcessor) afterProcess(logger *log.Logger, command *commandDao.Command) {
	b.AfterSuccessExecute(logger, command)
}

func delJdBondInterface(logger *log.Logger, jd_bond_interface *jdBondInterfaceDao.JdBondInterface) {
	jd_bond_interface.UpdateTime = time.Now()
	jd_bond_interface.IsDel = 1
	if err := jdBondInterfaceDao.UpdateJdBondInterfaceById(logger, jd_bond_interface); err != nil {
		logger.Warn("delJdBondInterface sql error:", err.Error())
	}
}

func delInterface(logger *log.Logger, entities []*interfaceDao.Interface) {
	for _, entity := range entities {
		entity.UpdateTime = time.Now()
		entity.IsDel = 1
		if err := interfaceDao.UpdateInterfaceById(logger, entity); err != nil {
			logger.Warn("delInterface sql error:", err.Error())
		}
	}
}
