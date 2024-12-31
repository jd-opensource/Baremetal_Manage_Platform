package processor

import (
	"net"
	"strconv"
	"time"

	commandDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/commandDao"
	deviceDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/deviceDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/idcLogic"
	rabbitMq "coding.jd.com/aidc-bmp/bmp-scheduler/service/rabbit_mq"
	"coding.jd.com/aidc-bmp/bmp-scheduler/service/redis"
	driverTypes "coding.jd.com/aidc-bmp/bmp-scheduler/types/driver"
	"coding.jd.com/aidc-bmp/bmp-scheduler/util"
	ironicAgentEvent "git.jd.com/cps-golang/ironic-common/ironic/event/command"
	log "git.jd.com/cps-golang/log"
)

type DHCPConfigAddSubnetProcessor struct {
	BaseProcessor
}

func NewDHCPConfigAddSubnetProcessor() DHCPConfigAddSubnetProcessor {
	b := DHCPConfigAddSubnetProcessor{}
	b.BaseProcessor.doProcess = b.doProcess
	b.BaseProcessor.afterProcess = b.afterProcess
	return b
}
func (b DHCPConfigAddSubnetProcessor) Process(logger *log.Logger, command *commandDao.Command) {
	b.BaseProcessor.Process(logger, command)
}

func (b DHCPConfigAddSubnetProcessor) Callback(logger *log.Logger, command *commandDao.Command, msg ironicAgentEvent.CallbackCommandMessage) {
	b.BaseProcessor.CallBack(logger, command, msg)
}

func (b DHCPConfigAddSubnetProcessor) doProcess(logger *log.Logger, command *commandDao.Command) {

	var _ipv4, _mask, _gateway string

	v, _ := deviceDao.GetBySn(logger, command.Sn)
	if v != nil {
		_ipv4 = v.PrivateIPv4
		_mask = v.Mask
		_gateway = v.Gateway
	}

	_, _, err := net.ParseCIDR(_ipv4)
	ipv4 := _ipv4
	if err != nil { //如果不是cird地址，先根据掩码转成cird,假如给的地址是这种格式，10.208.14.81
		mask := net.IPMask(net.ParseIP(_mask).To4())
		prefixSize, _ := mask.Size()
		ipv4 = _ipv4 + "/" + strconv.Itoa(prefixSize)
	}
	subnet := util.NetworkIp(ipv4)
	msg := driverTypes.DHCPConfigAddSubnet{
		Subnet:     subnet,
		Sn:         command.Sn,
		SubnetMask: _mask,
		Routes:     _gateway,
	}
	idc, err := idcLogic.GetIdcID(logger)
	if err != nil {
		logger.Warnf("GetOneIdc error:%s", err.Error())
		panic(ProcessAbortException{Msg: "process DHCPConfigAddSubnetProcessor GetIdcID error"})

	}
	res, err := redis.RedisUtil.SetNX("set_subenet_"+subnet, true, time.Second*60).Result()
	logger.Infof("set_subenet_", command.Sn, res, err, util.ObjToJson(msg))
	if res {
		if err := rabbitMq.SendToAgent4Topic(idc.IDcID, msg); err != nil {
			logger.Warnf("DHCPConfigAddSubnetProcessor.SendToAgent4Topic error, sn:%s, idc_id:%s, msg:%s", command.Sn, idc.IDcID, util.ObjToJson(msg))
			return
		} else {
			logger.Infof("DHCPConfigAddSubnetProcessor.SendToAgent4Topic success, sn:%s, idc_id:%s, msg:%s", command.Sn, idc.IDcID, util.ObjToJson(msg))
		}
	}

	logger.Infof("DHCPConfigAddSubnetProcessor_SendToAgent_Msg: %s, routingkey:%s", util.ObjToJson(msg), idc.IDcID)
}

func (b DHCPConfigAddSubnetProcessor) afterProcess(logger *log.Logger, command *commandDao.Command) {
	//不管是否返回成功，都默认成功，执行下一条指令
	b.AfterSuccessExecute(logger, command)
}
