package processor

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	commandDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/commandDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/deviceDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/deviceTypeDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/instanceLogic"
	rabbitMq "coding.jd.com/aidc-bmp/bmp-scheduler/service/rabbit_mq"
	agentTypes "coding.jd.com/aidc-bmp/bmp-scheduler/types/agent"
	"coding.jd.com/aidc-bmp/bmp-scheduler/util"
	ironicAgentEvent "git.jd.com/cps-golang/ironic-common/ironic/event/command"
	log "git.jd.com/cps-golang/log"
)

type SetNocloudNetworkProcessor struct {
	BaseProcessor
}

//linux设置network

func NewSetNocloudNetworkProcessor() SetNocloudNetworkProcessor {
	b := SetNocloudNetworkProcessor{}
	b.BaseProcessor.doProcess = b.doProcess
	b.BaseProcessor.afterProcess = b.afterProcess
	return b
}

func (b SetNocloudNetworkProcessor) Process(logger *log.Logger, command *commandDao.Command) {
	b.BaseProcessor.Process(logger, command)
}

func (b SetNocloudNetworkProcessor) Callback(logger *log.Logger, command *commandDao.Command, msg ironicAgentEvent.CallbackCommandMessage) {
	b.BaseProcessor.CallBack(logger, command, msg)
}

func (b SetNocloudNetworkProcessor) doProcess(logger *log.Logger, command *commandDao.Command) {

	logger.Info("SetNocloudNetworkProcessor doProcess......", command.Sn)
	instance, err := instanceLogic.GetByInstanceId(logger, command.InstanceID)
	if err != nil {
		logger.Warn("SetNocloudNetworkProcessor GetByInstanceId sql error:", command.InstanceID, err.Error())
	}
	fmt.Println("shili", instance)
	if instance == nil {
		panic(ProcessAbortException{Msg: fmt.Sprintf("process MakeRaid error, instance %s not found", command.InstanceID)})
	}
	device, err := deviceDao.GetDeviceById(logger, instance.DeviceID)
	if err != nil {
		logger.Warn("SetNocloudNetworkProcessor GetByInstanceId sql error:", command.InstanceID, err.Error())
	}
	fmt.Println("device", device)
	if device == nil {
		panic(ProcessAbortException{Msg: fmt.Sprintf("process SetNocloudNetwork error, device %s not found", instance.DeviceID)})
	}

	address := make([]string, 0)
	mask := net.IPMask(net.ParseIP(device.Mask).To4())
	prefixSize, _ := mask.Size()
	ipv4 := device.PrivateIPv4 + "/" + strconv.Itoa(prefixSize)
	address = append(address, ipv4)
	if device.PrivateIPv6 != "" {
		address = append(address, device.PrivateIPv6)
	}

	addressEth1 := make([]string, 0)
	mask2 := net.IPMask(net.ParseIP(device.MaskEth1).To4())
	prefixSize2, _ := mask2.Size()
	ipv4eth1 := device.PrivateEth1IPv4 + "/" + strconv.Itoa(prefixSize2)
	addressEth1 = append(addressEth1, ipv4eth1)
	if device.PrivateIPv6 != "" {
		addressEth1 = append(addressEth1, device.PrivateEth1IPv6)
	}

	//ipv4andv6 := make([]string, 0)
	//ipv4andv6 = append(ipv4andv6, device.PrivateIPv4)
	//if device.PrivateIPv6 != "" {
	//	ipv4andv6 = append(ipv4andv6, device.PrivateIPv6)
	//}
	deviceType, err := deviceTypeDao.GetDeviceTypeById(logger, instance.DeviceTypeID)
	if err != nil {
		logger.Warn("SetNocloudNetworkProcessor GetByDeviceTypeId sql error:", command.InstanceID, err.Error())
	}
	if deviceType == nil {
		panic(ProcessAbortException{Msg: fmt.Sprintf("process SetNocloudNetwork error, deviceType %s not found", instance.DeviceTypeID)})
	}
	if strings.ToLower(deviceType.InterfaceMode) == "bond" { //如果是bond模式
		SetNocloudNetwork := agentTypes.SetNocloudNetwork{
			Sn:      command.Sn,
			Version: "2.0",
		}
		SetNocloudNetwork.Network.Bonds.Bond0.Macaddress = device.Mac1
		SetNocloudNetwork.Network.Bonds.Bond0.Addresses = address
		SetNocloudNetwork.Network.Bonds.Bond0.Gateway4 = device.Gateway
		if device.PrivateIPv6 != "" {
			SetNocloudNetwork.Network.Bonds.Bond0.Gateway6 = device.Gateway6
		}
		SetNocloudNetwork.Network.Bonds.Bond0.Nameservers.Addresses = []string{"114.114.114.114", "8.8.8.8"}

		SetNocloudNetwork.Network.Bonds.Bond0.Interfaces = []string{"eth0", "eth1"}
		SetNocloudNetwork.Network.Bonds.Bond0.Parameters.Mode = "802.3ad"
		SetNocloudNetwork.Network.Bonds.Bond0.Parameters.MiiMonitorInterval = 100

		SetNocloudNetwork.Network.Ethernets.Eth0.Match.Macaddress = device.Mac1
		SetNocloudNetwork.Network.Ethernets.Eth0.SetName = "eth0"
		SetNocloudNetwork.Network.Ethernets.Eth1.Match.Macaddress = device.Mac2
		SetNocloudNetwork.Network.Ethernets.Eth1.SetName = "eth1"
		SetNocloudNetwork.Network.Version = 2 //写死

		fmt.Println("最终的json:", util.ObjToJson(SetNocloudNetwork))
		//os.Exit(1)
		if err := rabbitMq.SendToAgent(command.Sn, SetNocloudNetwork); err != nil {
			logger.Warnf("SetNocloudNetworkProcessor SendToAgent error, msg:%s, error:%s", util.ObjToJson(SetNocloudNetwork), err.Error())
			return
		}
		logger.Infof("SetNocloudNetworkProcessor_SendToAgent_Msg success, msg:%s", util.ObjToJson(SetNocloudNetwork))

	} else if deviceType.NicAmount == 1 { //如果是单网口，只需要eth0即可
		SetNocloudNetworkNoBond := agentTypes.SetNocloudNetworkNoBond{
			Sn:      command.Sn,
			Version: "2.0",
		}
		SetNocloudNetworkNoBond.Network.Ethernets.Eth0.Match.Macaddress = device.Mac1
		SetNocloudNetworkNoBond.Network.Ethernets.Eth0.SetName = "eth0" //默认写死
		SetNocloudNetworkNoBond.Network.Ethernets.Eth0.Addresses = address
		SetNocloudNetworkNoBond.Network.Ethernets.Eth0.Netmask = device.Mask
		SetNocloudNetworkNoBond.Network.Ethernets.Eth0.Gateway4 = device.Gateway
		if device.PrivateIPv6 != "" {
			SetNocloudNetworkNoBond.Network.Ethernets.Eth0.Gateway6 = device.Gateway6
		}
		//fmt.Println("设置前", SetNocloudNetworkNoBond, util.ObjToJson(SetNocloudNetworkNoBond))
		//双网卡基础网络环境，开启外网时(eth1是外网)，此时eth0不设置gateway4，eth1设置gateway4
		SetNocloudNetworkNoBond.Network.Ethernets.Eth0.Nameservers.Addresses = []string{"103.224.222.222", "103.224.222.223"} //"114.114.114.114", "8.8.8.8"
		//fmt.Println("设置后", SetNocloudNetworkNoBond, util.ObjToJson(SetNocloudNetworkNoBond))
		ri := []agentTypes.RouteItem{
			{
				To:  "0.0.0.0/0",
				Via: device.Gateway,
			},
		}
		SetNocloudNetworkNoBond.Network.Ethernets.Eth0.Routes = ri
		SetNocloudNetworkNoBond.Network.Version = 2 //写死

		//os.Exit(1)
		if err := rabbitMq.SendToAgent(command.Sn, SetNocloudNetworkNoBond); err != nil {
			logger.Warnf("SetNocloudNetworkProcessor SendToAgent error, msg:%s, error:%s", util.ObjToJson(SetNocloudNetworkNoBond), err.Error())
			return
		}
		logger.Infof("SetNocloudNetworkProcessor_SendToAgent_Msg success, msg:%s", util.ObjToJson(SetNocloudNetworkNoBond))
	} else if deviceType.NicAmount == 2 { //如果是双网口非bond，需要eth0,eth1
		SetNocloudNetworkNoBond := agentTypes.SetNocloudNetworkNoBond{
			Sn:      command.Sn,
			Version: "2.0",
		}
		SetNocloudNetworkNoBond.Network.Ethernets.Eth0.Match.Macaddress = device.Mac1
		SetNocloudNetworkNoBond.Network.Ethernets.Eth0.SetName = "eth0" //默认写死
		SetNocloudNetworkNoBond.Network.Ethernets.Eth0.Addresses = address
		SetNocloudNetworkNoBond.Network.Ethernets.Eth0.Netmask = device.Mask
		SetNocloudNetworkNoBond.Network.Ethernets.Eth0.Gateway4 = device.Gateway
		if device.PrivateIPv6 != "" {
			SetNocloudNetworkNoBond.Network.Ethernets.Eth0.Gateway6 = device.Gateway6
		}
		//fmt.Println("设置前", SetNocloudNetworkNoBond, util.ObjToJson(SetNocloudNetworkNoBond))
		//双网卡基础网络环境，开启外网时(eth1是外网)，此时eth0不设置gateway4，eth1设置gateway4
		SetNocloudNetworkNoBond.Network.Ethernets.Eth0.Nameservers.Addresses = []string{"103.224.222.222", "103.224.222.223"}
		ri := []agentTypes.RouteItem{
			{
				To:  "0.0.0.0/0",
				Via: device.Gateway,
			},
		}
		SetNocloudNetworkNoBond.Network.Ethernets.Eth0.Routes = ri
		//fmt.Println("设置后", SetNocloudNetworkNoBond, util.ObjToJson(SetNocloudNetworkNoBond))
		SetNocloudNetworkNoBond.Network.Version = 2 //写死

		SetNocloudNetworkNoBond.Network.Ethernets.Eth1.Match.Macaddress = device.Mac2
		SetNocloudNetworkNoBond.Network.Ethernets.Eth1.SetName = "eth1"
		SetNocloudNetworkNoBond.Network.Ethernets.Eth1.Addresses = addressEth1
		SetNocloudNetworkNoBond.Network.Ethernets.Eth1.Netmask = device.MaskEth1
		SetNocloudNetworkNoBond.Network.Ethernets.Eth1.Nameservers.Addresses = []string{"103.224.222.222", "103.224.222.223"}

		//os.Exit(1)
		if err := rabbitMq.SendToAgent(command.Sn, SetNocloudNetworkNoBond); err != nil {
			logger.Warnf("SetNocloudNetworkProcessor SendToAgent error, msg:%s, error:%s", util.ObjToJson(SetNocloudNetworkNoBond), err.Error())
			return
		}
		logger.Infof("SetNocloudNetworkProcessor_SendToAgent_Msg success, msg:%s", util.ObjToJson(SetNocloudNetworkNoBond))
	}

}

func (b SetNocloudNetworkProcessor) afterProcess(logger *log.Logger, command *commandDao.Command) {
	//empty
}
