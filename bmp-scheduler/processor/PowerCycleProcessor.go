package processor

import (
	commandDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/commandDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/idcDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/deviceLogic"
	rabbitMq "coding.jd.com/aidc-bmp/bmp-scheduler/service/rabbit_mq"
	dirverTypes "coding.jd.com/aidc-bmp/bmp-scheduler/types/driver"
	"coding.jd.com/aidc-bmp/bmp-scheduler/util"
	"fmt"
	ironicAgentEvent "git.jd.com/cps-golang/ironic-common/ironic/event/command"
	log "git.jd.com/cps-golang/log"
)

type PowerCycleProcessor struct {
	BaseProcessor
}

func NewPowerCycleProcessor() PowerCycleProcessor {
	b := PowerCycleProcessor{}
	b.BaseProcessor.doProcess = b.doProcess
	b.BaseProcessor.afterProcess = b.afterProcess
	return b
}

func (b PowerCycleProcessor) Process(logger *log.Logger, command *commandDao.Command) {
	b.BaseProcessor.Process(logger, command)
}

func (b PowerCycleProcessor) Callback(logger *log.Logger, command *commandDao.Command, msg ironicAgentEvent.CallbackCommandMessage) {
	b.BaseProcessor.CallBack(logger, command, msg)
}

func (b PowerCycleProcessor) doProcess(logger *log.Logger, command *commandDao.Command) {

	device, err := deviceLogic.GetBySn(logger, command.Sn)
	if err != nil {
		logger.Warn("PowerCycleProcessor GetById sql error:", command.Sn, err.Error())
	}
	if device == nil {
		panic(ProcessAbortException{Msg: fmt.Sprintf("process PowerCycle error,sn %s device not found", command.Sn)})
	}
	idc, err := idcDao.GetIdcById(logger, device.IDcID)
	if err != nil {
		logger.Warn("机房ID不存在", device.IDcID, err.Error())
		return
	}
	var iloUser, iloPassword string
	//如果设备的带外账号，密码没有录入，默认取机房填写的带外账号密码
	if device.IloUser == "" {
		iloUser = idc.IloUser
	} else {
		iloUser = device.IloUser
	}
	if device.IloPassword == "" {
		iloPassword = idc.IloPassword
	} else {
		iloPassword = device.IloPassword
	}
	power_cycle := dirverTypes.PowerCycle{
		Sn:       command.Sn,
		IloIp:    device.IloIP,
		Username: iloUser,
		Password: iloPassword,
	}
	if err := rabbitMq.SendToAgent(device.IDcID, power_cycle); err != nil {
		logger.Warnf("PowerCycleProcessor SendToAgent error, msg:%s, error:%s", util.ObjToJson(power_cycle), err.Error())
		return
	}
	logger.Infof("PowerCycleProcessor_SendToAgent_Msg success, msg:%s", util.ObjToJson(power_cycle))
}

func (b PowerCycleProcessor) afterProcess(logger *log.Logger, command *commandDao.Command) {
	//empty
}