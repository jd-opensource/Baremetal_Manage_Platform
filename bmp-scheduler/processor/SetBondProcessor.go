package processor

import (
	commandDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/commandDao"
	deviceDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/deviceDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/instanceDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/interfaceDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/jdSwitchDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/deviceLogic"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/instanceLogic"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/interfaceLogic"
	rabbitMq "coding.jd.com/aidc-bmp/bmp-scheduler/service/rabbit_mq"
	agentTypes "coding.jd.com/aidc-bmp/bmp-scheduler/types/agent"
	"coding.jd.com/aidc-bmp/bmp-scheduler/util"
	"fmt"
	ironicAgentEvent "git.jd.com/cps-golang/ironic-common/ironic/event/command"
	log "git.jd.com/cps-golang/log"
	beego "github.com/beego/beego/v2/server/web"
)

type SetBondProcessor struct {
	BaseProcessor
}

func NewSetBondProcessor() SetBondProcessor {
	b := SetBondProcessor{}
	b.BaseProcessor.doProcess = b.doProcess
	b.BaseProcessor.afterProcess = b.afterProcess
	return b
}

func (b SetBondProcessor) Process(logger *log.Logger, command *commandDao.Command) {
	b.BaseProcessor.Process(logger, command)
}

func (b SetBondProcessor) Callback(logger *log.Logger, command *commandDao.Command, msg ironicAgentEvent.CallbackCommandMessage) {
	b.BaseProcessor.CallBack(logger, command, msg)
}

func (b SetBondProcessor) doProcess(logger *log.Logger, command *commandDao.Command) {
	instance, err := instanceLogic.GetByInstanceId(logger, command.InstanceID)
	if err != nil {
		logger.Warn("SetBondProcessor GetByInstanceId sql error:", command.InstanceID, err.Error())
	}
	if instance == nil {
		panic(ProcessAbortException{Msg: fmt.Sprintf("process SetBondProcessor error, instance %s not found", command.InstanceID)})
	}

	device, err := deviceLogic.GetById(logger, instance.DeviceID)
	if err != nil {
		logger.Warn("SetBondProcessor GetById sql error:", instance.DeviceID, err.Error())
	}
	if device == nil {
		panic(ProcessAbortException{Msg: fmt.Sprintf("process SetBondProcessor error, device %s not found", fmt.Sprint(instance.DeviceID))})
	}

	interfaces, err := interfaceLogic.GetBySn(logger, command.Sn)
	if err != nil || len(interfaces) < 2 {
		logger.Warn("SetBondProcessor GetBySn sql error:", command.Sn, err.Error())
	}

	/*if strings.EqualFold(instance.NetworkType, NetworkType.RETAIL) {
		setBondOfRetail(logger, command, instance, device, interfaces)
	}*/

}

func setBondOfRetail(logger *log.Logger, command *commandDao.Command, instance *instanceDao.Instance, device *deviceDao.Device, interfaces []*interfaceDao.Interface) {
	switch_ip := interfaces[0].SwitchIP
	jd_switch, err := jdSwitchDao.GetBySwitchIp(logger, switch_ip)
	if err != nil {
		logger.Warnf("SetBondProcessor.setBondOfRetail.GetBySwitchIp error, switch_ip:%s, error:%s", switch_ip, err.Error())
		panic("process SetBond is error, switchIp %s JdSwitchEntity not found")
	}
	primary_interface := interfaces[0]
	second_interface := interfaces[1]

	nameservers, _ := beego.AppConfig.Strings("sdn.configs.nameservers")
	if len(nameservers) == 0 {
		nameservers = []string{"10.10.10.10"}
	}

	set_bond := agentTypes.SetBond{
		Sn:                     command.Sn,
		PrimaryInterfaceName:   primary_interface.InterfaceName,
		PrimaryMacAddr:         primary_interface.Mac,
		SecondaryInterfaceName: second_interface.InterfaceName,
		SecondaryMacAddr:       second_interface.Mac,
		IpAddr:                 "instance.PrivateIP",
		Netmask:                jd_switch.Mask,
		Gateway:                jd_switch.Gateway,
		Nameservers:            nameservers,
		Mode:                   "4",
	}

	if err := rabbitMq.SendToAgent(command.Sn, set_bond); err != nil {
		logger.Warnf("SetBondProcessor SendToAgent error, msg:%s, error:%s", util.ObjToJson(set_bond), err.Error())
		return
	}
	logger.Infof("SetBondProcessor_SendToAgent_Msg success, msg:%s", util.ObjToJson(set_bond))

}

func (b SetBondProcessor) afterProcess(logger *log.Logger, command *commandDao.Command) {
	// b.AfterSuccessExecute(logger, command)
}
