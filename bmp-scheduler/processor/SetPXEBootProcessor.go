package processor

import (
	"fmt"

	commandDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/commandDao"
	deviceDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/deviceDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/idcDao"
	rabbitMq "coding.jd.com/aidc-bmp/bmp-scheduler/service/rabbit_mq"
	driverTypes "coding.jd.com/aidc-bmp/bmp-scheduler/types/driver"
	"coding.jd.com/aidc-bmp/bmp-scheduler/util"
	ironicAgentEvent "git.jd.com/cps-golang/ironic-common/ironic/event/command"
	log "git.jd.com/cps-golang/log"
)

type SetPXEBootProcessor struct {
	BaseProcessor
}

func NewSetPXEBootProcessor() SetPXEBootProcessor {
	b := SetPXEBootProcessor{}
	b.BaseProcessor.doProcess = b.doProcess
	b.BaseProcessor.afterProcess = b.afterProcess
	return b
}

func (b SetPXEBootProcessor) Process(logger *log.Logger, command *commandDao.Command) {
	b.BaseProcessor.Process(logger, command)
}

func (b SetPXEBootProcessor) Callback(logger *log.Logger, command *commandDao.Command, msg ironicAgentEvent.CallbackCommandMessage) {
	b.BaseProcessor.CallBack(logger, command, msg)
}

func (b SetPXEBootProcessor) doProcess(logger *log.Logger, command *commandDao.Command) {
	logger.Info("SetPXEBootProcessor doProcess begin...", command.Sn)
	var device *deviceDao.Device
	var err error
	device, _ = deviceDao.GetBySn(logger, command.Sn)

	if device == nil {
		logger.Warn("SetPXEBootProcessor GetBySn sql error:", err.Error())
		panic(ProcessAbortException{Msg: fmt.Sprintf("process SetPXEBootProcessor error, sn %s device not found", command.Sn)})
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
	set_PXE_boot := driverTypes.SetPXEBoot{
		Sn:       command.Sn,
		IloIp:    device.IloIP,
		Username: iloUser,
		Password: iloPassword,
	}
	if err := rabbitMq.SendToAgent(device.IDcID, set_PXE_boot); err != nil {
		logger.Warnf("SetPXEBootProcessor SendToDriver error, routekey: %s, msg:%s, error:%s", device.IDcID, util.ObjToJson(set_PXE_boot), err.Error())
		return
	}
	logger.Infof("SetPXEBootProcessor_SendToDriver_Msg success, routekey: %s, msg:%s", device.IDcID, util.ObjToJson(set_PXE_boot))
}

func (b SetPXEBootProcessor) afterProcess(logger *log.Logger, command *commandDao.Command) {
	logger.Info("SetPXEBootProcessor afterProcess done!", command.Sn)
	defer logger.Info("SetPXEBootProcessor afterProcess done!", command.Sn)
	//empty
}
