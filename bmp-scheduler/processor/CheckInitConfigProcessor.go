package processor

import (
	commandDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/commandDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/deviceDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/idcLogic"
	rabbitMq "coding.jd.com/aidc-bmp/bmp-scheduler/service/rabbit_mq"
	driverTypes "coding.jd.com/aidc-bmp/bmp-scheduler/types/driver"
	"coding.jd.com/aidc-bmp/bmp-scheduler/util"
	ironicAgentEvent "git.jd.com/cps-golang/ironic-common/ironic/event/command"
	log "git.jd.com/cps-golang/log"
)

type CheckInitConfigProcessor struct {
	BaseProcessor
}

func NewCheckInitConfigProcessor() CheckInitConfigProcessor {
	b := CheckInitConfigProcessor{}
	b.BaseProcessor.doProcess = b.doProcess
	b.BaseProcessor.afterProcess = b.afterProcess
	return b
}
func (b CheckInitConfigProcessor) Process(logger *log.Logger, command *commandDao.Command) {
	b.BaseProcessor.Process(logger, command)
}

func (b CheckInitConfigProcessor) Callback(logger *log.Logger, command *commandDao.Command, msg ironicAgentEvent.CallbackCommandMessage) {
	b.BaseProcessor.CallBack(logger, command, msg)
}

func (b CheckInitConfigProcessor) doProcess(logger *log.Logger, command *commandDao.Command) {
	var err error

	var iloIP, iloUser, iloPassword string

	v, _ := deviceDao.GetBySn(logger, command.Sn)
	if v != nil {
		iloIP = v.IloIP
		iloUser = v.IloUser
		iloPassword = v.IloPassword
	}

	//再查找idcid
	idc, err := idcLogic.GetIdcID(logger)
	if err != nil {
		logger.Warnf("GetOneIdc error:%s", err.Error())
		panic(ProcessAbortException{Msg: "process CheckInitConfigProcessor GetIdc error"})

	}

	msgCheck := driverTypes.CheckInitConfig{
		IloIp:     iloIP,
		Sn:        command.Sn,
		Username:  iloUser,
		Password:  iloPassword,
		Privilege: "Administrator",
	}
	//优先从设备取带外user/passwd，没有的话用idc的
	if msgCheck.Username == "" {
		msgCheck.Username = idc.IloUser
	}
	if msgCheck.Password == "" {
		msgCheck.Password = idc.IloPassword
	}

	if err := rabbitMq.SendToAgent(idc.IDcID, msgCheck); err != nil {
		logger.Warn("CheckInitConfigProcessor SendToAgent error:", idc.IDcID, err.Error())
		return
	}

	logger.Infof("CheckInitConfigProcessor_SendToAgent_Msg: %s, routingkey:%s", util.ObjToJson(msgCheck), idc.IDcID)
}

func (b CheckInitConfigProcessor) afterProcess(logger *log.Logger, command *commandDao.Command) {
	//默认成功，执行下一条指令
	b.AfterSuccessExecute(logger, command)
}
