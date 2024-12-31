package processor

import (
	commandDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/commandDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/rInstanceSshkeyDao"
	rabbitMq "coding.jd.com/aidc-bmp/bmp-scheduler/service/rabbit_mq"
	agentTypes "coding.jd.com/aidc-bmp/bmp-scheduler/types/agent"
	"coding.jd.com/aidc-bmp/bmp-scheduler/util"
	ironicAgentEvent "git.jd.com/cps-golang/ironic-common/ironic/event/command"
	log "git.jd.com/cps-golang/log"
)

type SetCloudinitConfProcessor struct {
	BaseProcessor
}

func NewSetCloudinitConfProcessor() SetCloudinitConfProcessor {
	b := SetCloudinitConfProcessor{}
	b.BaseProcessor.doProcess = b.doProcess
	b.BaseProcessor.afterProcess = b.afterProcess
	return b
}

func (b SetCloudinitConfProcessor) Process(logger *log.Logger, command *commandDao.Command) {
	b.BaseProcessor.Process(logger, command)
}

func (b SetCloudinitConfProcessor) Callback(logger *log.Logger, command *commandDao.Command, msg ironicAgentEvent.CallbackCommandMessage) {
	b.BaseProcessor.CallBack(logger, command, msg)
}

func (b SetCloudinitConfProcessor) doProcess(logger *log.Logger, command *commandDao.Command) {

	logger.Info("SetCloudinitConfProcessor doProcess......", command.Sn)
	instanceSshkeyList, _ := rInstanceSshkeyDao.GetAllrInstanceSshkey(logger, map[string]interface{}{
		"instance_id": command.InstanceID,
		"is_del":      0,
	})
	sshPwauth := "yes"
	if len(instanceSshkeyList) != 0 { //如果使用了ssh登录方式,那么就将密码登录禁用
		sshPwauth = "no"
	}
	setCloudinitConf := agentTypes.SetCloudinitConf{
		Sn:        command.Sn,
		Version:   "2.0",
		SshPwauth: sshPwauth, // //yes:启用密码登陆  //no禁用密码登陆
	}
	if err := rabbitMq.SendToAgent(command.Sn, setCloudinitConf); err != nil {
		logger.Warnf("SetCloudinitConfProcessor SendToAgent error, msg:%s, error:%s", util.ObjToJson(setCloudinitConf), err.Error())
		return
	}
	logger.Infof("SetCloudinitConfProcessor_SendToAgent_Msg success, msg:%s", util.ObjToJson(setCloudinitConf))

}

func (b SetCloudinitConfProcessor) afterProcess(logger *log.Logger, command *commandDao.Command) {
	//empty
}
