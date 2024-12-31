package processor

import (
	commandDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/commandDao"
	deviceDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/deviceDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/deviceLogic"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/instanceLogic"
	rabbitMq "coding.jd.com/aidc-bmp/bmp-scheduler/service/rabbit_mq"
	agentTypes "coding.jd.com/aidc-bmp/bmp-scheduler/types/agent"
	"coding.jd.com/aidc-bmp/bmp-scheduler/util"
	"fmt"
	ironicAgentEvent "git.jd.com/cps-golang/ironic-common/ironic/event/command"
	log "git.jd.com/cps-golang/log"
	beego "github.com/beego/beego/v2/server/web"
)

type SetNetworkProcessor struct {
	BaseProcessor
}

func NewSetNetworkProcessor() SetNetworkProcessor {
	b := SetNetworkProcessor{}
	b.BaseProcessor.doProcess = b.doProcess
	b.BaseProcessor.afterProcess = b.afterProcess
	return b
}

func (b SetNetworkProcessor) Process(logger *log.Logger, command *commandDao.Command) {
	b.BaseProcessor.Process(logger, command)
}

func (b SetNetworkProcessor) Callback(logger *log.Logger, command *commandDao.Command, msg ironicAgentEvent.CallbackCommandMessage) {
	b.BaseProcessor.CallBack(logger, command, msg)
}

func (b SetNetworkProcessor) doProcess(logger *log.Logger, command *commandDao.Command) {
	instance, err := instanceLogic.GetByInstanceId(logger, command.InstanceID)
	if err != nil {
		logger.Warn("SetNetworkProcessor GetByInstanceId sql error:", command.InstanceID, err.Error())
	}
	if instance == nil {
		panic(ProcessAbortException{Msg: fmt.Sprintf("process SetBondProcessor error, instance %s not found", command.InstanceID)})
	}

	device, err := deviceLogic.GetById(logger, instance.DeviceID)
	if err != nil {
		logger.Warn("SetNetworkProcessor GetById sql error:", instance.DeviceID, err.Error())
	}
	if device == nil {
		panic(ProcessAbortException{Msg: fmt.Sprintf("process SetBondProcessor error, device %s not found", fmt.Sprint(instance.DeviceID))})
	}
	setNetworkOfRetail(logger, command, device)
	//interfaces, err := interfaceLogic.GetBySn(logger, command.Sn)
	//if err != nil || len(interfaces) < 2 {
	//	logger.Warn("SetBondProcessor GetBySn sql error,", command.Sn, err.Error())
	//}

	//if strings.EqualFold(instance.NetworkType, NetworkType.RETAIL) {
	//setNetworkOfRetail(logger, command, device)
	//}

}

func setNetworkOfRetail(logger *log.Logger, command *commandDao.Command, device *deviceDao.Device) {
	//switch_ip := interfaces[0].SwitchIP
	//jd_switch, err := jdSwitchDao.GetBySwitchIp(logger, switch_ip)
	//if err != nil {
	//	logger.Warn("SetNetworkProcessor.setNetworkOfRetail.GetBySwitchIp error:", switch_ip, err.Error())
	//	panic("process SetBond is error, switchIp %s JdSwitchEntity not found")
	//}
	//primary_interface := interfaces[0]

	nameservers, _ := beego.AppConfig.Strings("sdn.configs.nameservers")
	if len(nameservers) == 0 {
		nameservers = []string{"10.10.10.10"}
	}
	set_network := agentTypes.SetNetwork{
		Sn:                   command.Sn,
		PrivateInterfaceName: "eth0",         //写死？？primary_interface.InterfaceName,
		PrivateMacAddr:       device.Mac1,    //primary_interface.Mac,
		PrivateGateway:       device.Gateway, //jd_switch.Gateway,
		PrivateNetmask:       device.Mask,    //jd_switch.Mask,
		PrivateIpAddr:        device.PrivateIPv4,
		// EnableNet:            instance.EnableInternet,
		Nameservers: nameservers,
		IsSetDns:    true,
	}

	if err := rabbitMq.SendToAgent(command.Sn, set_network); err != nil {
		logger.Warnf("SetNetworkProcessor SendToAgent error, msg:%s, error:%s", util.ObjToJson(set_network), err.Error())
		return
	}
	logger.Infof("SetNetworkProcessor_SendToAgent_Msg success, msg:%s", util.ObjToJson(set_network))

}

func (b SetNetworkProcessor) afterProcess(logger *log.Logger, command *commandDao.Command) {
	// b.AfterSuccessExecute(logger, command)
}
