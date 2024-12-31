package processor

import (
	"fmt"

	commandDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/commandDao"
	deviceDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/deviceDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/idcDao"
	rabbitMq "coding.jd.com/aidc-bmp/bmp-scheduler/service/rabbit_mq"
	dirverTypes "coding.jd.com/aidc-bmp/bmp-scheduler/types/driver"
	"coding.jd.com/aidc-bmp/bmp-scheduler/util"
	ironicAgentEvent "git.jd.com/cps-golang/ironic-common/ironic/event/command"
	log "git.jd.com/cps-golang/log"
)

type PowerOffProcessor struct {
	BaseProcessor
}

func NewPowerOffProcessor() PowerOffProcessor {
	b := PowerOffProcessor{}
	b.BaseProcessor.doProcess = b.doProcess
	b.BaseProcessor.afterProcess = b.afterProcess
	return b
}

func (b PowerOffProcessor) Process(logger *log.Logger, command *commandDao.Command) {
	b.BaseProcessor.Process(logger, command)
}

func (b PowerOffProcessor) Callback(logger *log.Logger, command *commandDao.Command, msg ironicAgentEvent.CallbackCommandMessage) {
	b.BaseProcessor.CallBack(logger, command, msg)
}

func (b PowerOffProcessor) doProcess(logger *log.Logger, command *commandDao.Command) {

	var device *deviceDao.Device
	var err error
	device, err = deviceDao.GetBySn(logger, command.Sn)
	if err != nil {
		logger.Warn("PowerOffProcessor GetById sql error:", command.Sn, err.Error())
	}
	if device == nil {
		panic(ProcessAbortException{Msg: fmt.Sprintf("process PowerOff error,sn %s device not found", command.Sn)})
	}
	idc, err := idcDao.GetIdcById(logger, device.IDcID)
	if err != nil {
		logger.Warn("机房ID不存在", device.Sn, device.IDcID, err.Error())
	}
	var iloUser, iloPassword string
	//如果设备的带外账号，密码没有录入，默认取机房填写的带外账号密码
	if device.IloUser == "" || device.IloPassword == "" {
		if idc == nil {
			logger.Warn("设备和机房iloip信息都缺失", device.Sn, device.IDcID)
			return
		}
	}
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
	power_off := dirverTypes.PowerOff{
		Sn:       command.Sn,
		IloIp:    device.IloIP,
		Username: iloUser,
		Password: iloPassword,
	}
	if err := rabbitMq.SendToAgent(device.IDcID, power_off); err != nil {
		logger.Warnf("PowerOffProcessor SendToAgent error, routingkey:%s, msg:%s, error:%s", device.IDcID, util.ObjToJson(power_off), err.Error())
		return
	}
	logger.Infof("PowerOffProcessor_SendToAgent_Msg success, routingkey:%s, msg:%s", device.IDcID, util.ObjToJson(power_off))
}

func (b PowerOffProcessor) afterProcess(logger *log.Logger, command *commandDao.Command) {
	//empty
}
