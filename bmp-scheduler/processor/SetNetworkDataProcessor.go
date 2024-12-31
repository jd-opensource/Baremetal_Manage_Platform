package processor

import (
	"fmt"
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

type SetNetworkDataProcessor struct {
	BaseProcessor
}

func NewSetNetworkDataProcessor() SetNetworkDataProcessor {
	b := SetNetworkDataProcessor{}
	b.BaseProcessor.doProcess = b.doProcess
	b.BaseProcessor.afterProcess = b.afterProcess
	return b
}

func (b SetNetworkDataProcessor) Process(logger *log.Logger, command *commandDao.Command) {
	b.BaseProcessor.Process(logger, command)
}

func (b SetNetworkDataProcessor) Callback(logger *log.Logger, command *commandDao.Command, msg ironicAgentEvent.CallbackCommandMessage) {
	b.BaseProcessor.CallBack(logger, command, msg)
}

func (b SetNetworkDataProcessor) doProcess(logger *log.Logger, command *commandDao.Command) {

	logger.Info("SetNetworkDataProcessor doProcess......", command.Sn)
	instance, err := instanceLogic.GetByInstanceId(logger, command.InstanceID)
	if err != nil {
		logger.Warn("SetNetworkDataProcessor GetByInstanceId sql error:", command.InstanceID, err.Error())
	}
	if instance == nil {
		panic(ProcessAbortException{Msg: fmt.Sprintf("process MakeRaid error, instance %s not found", command.InstanceID)})
	}
	device, err := deviceDao.GetDeviceById(logger, instance.DeviceID)
	if err != nil {
		logger.Warn("SetNetworkDataProcessor GetByInstanceId sql error:", command.InstanceID, err.Error())
	}
	if device == nil {
		panic(ProcessAbortException{Msg: fmt.Sprintf("process SetNetworkData error, device %s not found", instance.DeviceID)})
	}

	//address := make([]string, 0)
	//mask := net.IPMask(net.ParseIP(device.Mask).To4())
	//prefixSize, _ := mask.Size()
	//ipv4 := device.PrivateIPv4 + "/" + strconv.Itoa(prefixSize)
	//address = append(address, ipv4)
	//if device.PrivateIPv6 != "" {
	//	address = append(address, device.PrivateIPv6)
	//}

	//ipv4andv6 := make([]string, 0)
	//ipv4andv6 = append(ipv4andv6, device.PrivateIPv4)
	//if device.PrivateIPv6 != "" {
	//	ipv4andv6 = append(ipv4andv6, device.PrivateIPv6)
	//}
	deviceType, err := deviceTypeDao.GetDeviceTypeById(logger, instance.DeviceTypeID)
	if err != nil {
		logger.Warn("SetNetworkDataProcessor GetByDeviceTypeId sql error:", command.InstanceID, err.Error())
	}
	if deviceType == nil {
		panic(ProcessAbortException{Msg: fmt.Sprintf("process SetNetworkData error, deviceType %s not found", instance.DeviceTypeID)})
	}
	SetNetworkData := agentTypes.SetNetworkData{
		Sn: command.Sn,
	}
	SetNetworkData.NetworkData.Services = append(SetNetworkData.NetworkData.Services, agentTypes.Service{
		Type:    "dns",
		Address: "114.114.114.114",
	})
	SetNetworkData.NetworkData.Services = append(SetNetworkData.NetworkData.Services, agentTypes.Service{
		Type:    "dns",
		Address: "8.8.8.8",
	})
	if deviceType.InterfaceMode == "bond" { //如果是bond模式
		NetWork := agentTypes.NetWork{
			Type:      "ipv4",
			IPAddress: device.PrivateIPv4,
			Netmask:   device.Mask,
			Link:      "bond0",
		}
		NetWork.Routes = append(NetWork.Routes, agentTypes.Route{
			Netmask: "0.0.0.0",
			Network: "0.0.0.0",
			Gateway: device.Gateway,
		})
		SetNetworkData.NetworkData.Networks = append(SetNetworkData.NetworkData.Networks, NetWork) //ipv4
		if device.PrivateIPv6 != "" {
			mask := strings.Split(device.PrivateIPv6, "/")
			maskStr := "64"
			if len(mask) > 1 {
				maskStr = mask[1]
			}
			NetWork := agentTypes.NetWork{
				Type:      "ipv6",
				IPAddress: mask[0],
				Netmask:   maskStr,
				Link:      "bond0",
			}
			NetWork.Routes = append(NetWork.Routes, agentTypes.Route{
				Netmask: "::",
				Network: "::",
				Gateway: device.Gateway6,
			})
			SetNetworkData.NetworkData.Networks = append(SetNetworkData.NetworkData.Networks, NetWork) //ipv6
		}

		SetNetworkData.NetworkData.Links = append(SetNetworkData.NetworkData.Links, agentTypes.Link{
			EthernetMacAddress: device.Mac1,
			Type:               "phy",
			ID:                 "eth0",
			Mtu:                1500,
		})
		SetNetworkData.NetworkData.Links = append(SetNetworkData.NetworkData.Links, agentTypes.Link{
			EthernetMacAddress: device.Mac2,
			Type:               "phy",
			ID:                 "eth1",
			Mtu:                1500,
		})
		SetNetworkData.NetworkData.Links = append(SetNetworkData.NetworkData.Links, agentTypes.Link{
			EthernetMacAddress: device.Mac1,
			Type:               "bond",
			ID:                 "bond0",
			BondLinks:          []string{"eth0", "eth1"},
			BondMiimon:         100,
			BondMode:           "802.3ad",
		})
	} else if deviceType.InterfaceMode == "dual" { //双网口不做聚合
		//eth0
		NetWork1 := agentTypes.NetWork{
			Type:      "ipv4",
			IPAddress: device.PrivateIPv4,
			Netmask:   device.Mask,
			Link:      "eth0",
		}
		NetWork1.Routes = append(NetWork1.Routes, agentTypes.Route{
			Netmask: "0.0.0.0",
			Network: "0.0.0.0",
			Gateway: device.Gateway,
		})
		SetNetworkData.NetworkData.Networks = append(SetNetworkData.NetworkData.Networks, NetWork1)
		if device.PrivateIPv6 != "" {
			mask := strings.Split(device.PrivateIPv6, "/")
			maskStr := "64"
			if len(mask) > 1 {
				maskStr = mask[1]
			}
			NetWork := agentTypes.NetWork{
				Type:      "ipv6",
				IPAddress: mask[0],
				Netmask:   maskStr,
				Link:      "eth0",
			}
			NetWork.Routes = append(NetWork.Routes, agentTypes.Route{
				Netmask: "::",
				Network: "::",
				Gateway: device.Gateway6,
			})
			SetNetworkData.NetworkData.Networks = append(SetNetworkData.NetworkData.Networks, NetWork) //ipv6
		}
		//eth1
		NetWork2 := agentTypes.NetWork{
			Type:      "ipv4",
			IPAddress: device.PrivateEth1IPv4,
			Netmask:   device.MaskEth1,
			Link:      "eth1",
		}
		NetWork2.Routes = append(NetWork2.Routes, agentTypes.Route{
			Netmask: "0.0.0.0",
			Network: "0.0.0.0",
			Gateway: device.Gateway,
		})
		SetNetworkData.NetworkData.Networks = append(SetNetworkData.NetworkData.Networks, NetWork2) //ipv4
		if device.PrivateEth1IPv6 != "" {
			mask := strings.Split(device.PrivateEth1IPv6, "/")
			maskStr := "64"
			if len(mask) > 1 {
				maskStr = mask[1]
			}
			NetWork := agentTypes.NetWork{
				Type:      "ipv6",
				IPAddress: mask[0],
				Netmask:   maskStr,
				Link:      "eth1",
			}
			NetWork.Routes = append(NetWork.Routes, agentTypes.Route{
				Netmask: "::",
				Network: "::",
				Gateway: device.Gateway6,
			})
			SetNetworkData.NetworkData.Networks = append(SetNetworkData.NetworkData.Networks, NetWork) //ipv6
		}

		SetNetworkData.NetworkData.Links = append(SetNetworkData.NetworkData.Links,
			agentTypes.Link{
				EthernetMacAddress: device.Mac1,
				Type:               "phy",
				ID:                 "eth0",
				Mtu:                1500,
			}, agentTypes.Link{
				EthernetMacAddress: device.Mac2,
				Type:               "phy",
				ID:                 "eth1",
				Mtu:                1500,
			})
	} else { //如果是单网口，只需要eth0即可
		NetWork := agentTypes.NetWork{
			Type:      "ipv4",
			IPAddress: device.PrivateIPv4,
			Netmask:   device.Mask,
			Link:      "eth0",
		}
		NetWork.Routes = append(NetWork.Routes, agentTypes.Route{
			Netmask: "0.0.0.0",
			Network: "0.0.0.0",
			Gateway: device.Gateway,
		})
		SetNetworkData.NetworkData.Networks = append(SetNetworkData.NetworkData.Networks, NetWork) //ipv4
		if device.PrivateIPv6 != "" {
			mask := strings.Split(device.PrivateIPv6, "/")
			maskStr := "64"
			if len(mask) > 1 {
				maskStr = mask[1]
			}
			NetWork := agentTypes.NetWork{
				Type:      "ipv6",
				IPAddress: mask[0],
				Netmask:   maskStr,
				Link:      "eth0",
			}
			NetWork.Routes = append(NetWork.Routes, agentTypes.Route{
				Netmask: "::",
				Network: "::",
				Gateway: device.Gateway6,
			})
			SetNetworkData.NetworkData.Networks = append(SetNetworkData.NetworkData.Networks, NetWork) //ipv6
		}

		SetNetworkData.NetworkData.Links = append(SetNetworkData.NetworkData.Links, agentTypes.Link{
			EthernetMacAddress: device.Mac1,
			Type:               "phy",
			ID:                 "eth0",
			Mtu:                1500,
		})
	}

	if err := rabbitMq.SendToAgent(command.Sn, SetNetworkData); err != nil {
		logger.Warnf("SetNetworkDataProcessor SendToAgent error, msg:%s, error:%s", util.ObjToJson(SetNetworkData), err.Error())
		return
	}
	logger.Infof("SetNetworkDataProcessor_SendToAgent_Msg success, msg:%s", util.ObjToJson(SetNetworkData))

}

func (b SetNetworkDataProcessor) afterProcess(logger *log.Logger, command *commandDao.Command) {
	//empty
}
