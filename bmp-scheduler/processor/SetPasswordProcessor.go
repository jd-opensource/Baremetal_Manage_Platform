package processor

import (
	commandDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/commandDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/rInstanceSshkeyDao"
	rabbitMq "coding.jd.com/aidc-bmp/bmp-scheduler/service/rabbit_mq"
	agentTypes "coding.jd.com/aidc-bmp/bmp-scheduler/types/agent"
	"coding.jd.com/aidc-bmp/bmp-scheduler/util"
	"fmt"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/TimeoutPolicy"
	ironicAgentEvent "git.jd.com/cps-golang/ironic-common/ironic/event/command"
	log "git.jd.com/cps-golang/log"
)

type SetPasswordProcessor struct {
	BaseProcessor
}

func NewSetPasswordProcessor() SetPasswordProcessor {
	b := SetPasswordProcessor{}
	b.BaseProcessor.doProcess = b.doProcess
	b.BaseProcessor.afterProcess = b.afterProcess
	return b
}

func (b SetPasswordProcessor) Process(logger *log.Logger, command *commandDao.Command) {
	b.BaseProcessor.Process(logger, command)
}

func (b SetPasswordProcessor) Callback(logger *log.Logger, command *commandDao.Command, msg ironicAgentEvent.CallbackCommandMessage) {
	b.BaseProcessor.CallBack(logger, command, msg)
}

func (b SetPasswordProcessor) doProcess(logger *log.Logger, command *commandDao.Command) {
	instance_extra := loadInstanceExtra(logger, command.InstanceID)
	fmt.Println(instance_extra, instance_extra.Password)
	if instance_extra == nil || instance_extra.Password == "" {
		command.TimeoutPolicy = TimeoutPolicy.SKIP
		logger.Info("SetPasswordCommand timeoutPolicy set to skip", command.InstanceID)
		return
	}
	set_password := agentTypes.SetPassword{
		Sn:       command.Sn,
		Password: instance_extra.Password,
	}
	if err := rabbitMq.SendToAgent(command.Sn, set_password); err != nil {
		logger.Warnf("SetPasswordProcessor SendToAgent error, msg:%s, error:%s", util.ObjToJson(set_password), err.Error())
		return
	}
	logger.Infof("SetPasswordProcessor_SendToAgent_Msg success, msg:%s", util.ObjToJson(set_password))

}

func (b SetPasswordProcessor) afterProcess(logger *log.Logger, command *commandDao.Command) {
	instanceSshkeyList, _ := rInstanceSshkeyDao.GetAllrInstanceSshkey(logger, map[string]interface{}{
		"instance_id": command.InstanceID,
		"is_del":      0,
	})
	if len(instanceSshkeyList) != 0 { //如果使用了ssh登录方式,直接执行成功，执行下一条指令
		b.AfterSuccessExecute(logger, command) //无论是否成功，指令都没往下走?
	}

}
