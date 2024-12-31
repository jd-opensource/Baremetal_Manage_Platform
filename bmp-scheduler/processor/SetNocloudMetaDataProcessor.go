package processor

import (
	"fmt"

	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/commandDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/rInstanceSshkeyDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/sshkeyDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/instanceLogic"
	rabbitMq "coding.jd.com/aidc-bmp/bmp-scheduler/service/rabbit_mq"
	agentTypes "coding.jd.com/aidc-bmp/bmp-scheduler/types/agent"
	"coding.jd.com/aidc-bmp/bmp-scheduler/util"
	ironicAgentEvent "git.jd.com/cps-golang/ironic-common/ironic/event/command"
	log "git.jd.com/cps-golang/log"
)

type SetNocloudMetaDataProcessor struct {
	BaseProcessor
}

func NewSetNocloudMetaDataProcessor() SetNocloudMetaDataProcessor {
	b := SetNocloudMetaDataProcessor{}
	b.BaseProcessor.doProcess = b.doProcess
	b.BaseProcessor.afterProcess = b.afterProcess
	return b
}

func (b SetNocloudMetaDataProcessor) Process(logger *log.Logger, command *commandDao.Command) {
	b.BaseProcessor.Process(logger, command)
}

func (b SetNocloudMetaDataProcessor) Callback(logger *log.Logger, command *commandDao.Command, msg ironicAgentEvent.CallbackCommandMessage) {
	b.BaseProcessor.CallBack(logger, command, msg)
}

func (b SetNocloudMetaDataProcessor) doProcess(logger *log.Logger, command *commandDao.Command) {

	logger.Info("SetNocloudMetaDataProcessor doProcess......", command.Sn)
	instance, err := instanceLogic.GetByInstanceId(logger, command.InstanceID)
	if err != nil {
		logger.Warn("SetNocloudMetaDataProcessor GetByInstanceId sql error:", command.InstanceID, err.Error())
	}
	if instance == nil {
		panic(ProcessAbortException{Msg: fmt.Sprintf("process SetNocloudMetaData error, instance %s not found", command.InstanceID)})
	}
	SetNocloudMetaData := agentTypes.SetNocloudMetaData{
		Sn:      command.Sn,
		Version: "2.0",
	}
	instanceSshkeyList, err := rInstanceSshkeyDao.GetAllrInstanceSshkey(logger, map[string]interface{}{
		"instance_id": command.InstanceID,
		"is_del":      0,
	})
	//fmt.Println(instanceSshkeyList, len(instanceSshkeyList), err.Error())//不知为啥，err.Error()报错！！！
	if err != nil {
		panic(ProcessAbortException{Msg: fmt.Sprintf("process SetNocloudMetaData error, instance %s not found", command.InstanceID)})
		return
	}
	pubKeys := []string{}
	for _, v := range instanceSshkeyList {
		sshKeyInfo, err := sshkeyDao.GetSshkeyById(logger, v.SSHkeyID)
		if sshKeyInfo != nil {
			pubKeys = append(pubKeys, sshKeyInfo.Key)
		} else {
			logger.Warnf("GetInstanceSshkey sql error By sshkeyId:%s, error:%s", v.SSHkeyID, err.Error())
			return
		}
	}
	SetNocloudMetaData.MetaData.LocalHostname = instance.Hostname
	SetNocloudMetaData.MetaData.InstanceID = command.InstanceID
	if len(pubKeys) > 0 {
		SetNocloudMetaData.MetaData.PublicKeys = pubKeys
	}

	if err := rabbitMq.SendToAgent(command.Sn, SetNocloudMetaData); err != nil {
		logger.Warnf("SetNocloudMetaDataProcessor SendToAgent error, msg:%s, error:%s", util.ObjToJson(SetNocloudMetaData), err.Error())
		return
	}
	logger.Infof("SetNocloudMetaDataProcessor_SendToAgent_Msg success, msg:%s", util.ObjToJson(SetNocloudMetaData))

}

func (b SetNocloudMetaDataProcessor) afterProcess(logger *log.Logger, command *commandDao.Command) {
	//empty
}
