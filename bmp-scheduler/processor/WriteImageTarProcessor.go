package processor

import (
	"encoding/json"
	"fmt"
	"strings"

	commandDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/commandDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/imageLogic"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/instanceLogic"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/osLogic"
	rabbitMq "coding.jd.com/aidc-bmp/bmp-scheduler/service/rabbit_mq"
	agentTypes "coding.jd.com/aidc-bmp/bmp-scheduler/types/agent"
	"coding.jd.com/aidc-bmp/bmp-scheduler/util"
	ironicAgentEvent "git.jd.com/cps-golang/ironic-common/ironic/event/command"
	log "git.jd.com/cps-golang/log"
)

type WriteImageTarProcessor struct {
	BaseProcessor
}

func NewWriteImageTarProcessor() WriteImageTarProcessor {
	b := WriteImageTarProcessor{}
	b.BaseProcessor.doProcess = b.doProcess
	b.BaseProcessor.afterProcess = b.afterProcess
	return b
}

func (b WriteImageTarProcessor) Process(logger *log.Logger, command *commandDao.Command) {
	b.BaseProcessor.Process(logger, command)
}

func (b WriteImageTarProcessor) Callback(logger *log.Logger, command *commandDao.Command, msg ironicAgentEvent.CallbackCommandMessage) {
	//以下是基于重装系统的，需要存储hints
	if strings.EqualFold(msg.Status, OK) {

		v, err := json.Marshal(msg.Data)
		if err != nil {
			logger.Warn("WriteImageTarProcessor.data marshal error:", err.Error())
		} else {
			disk_hints := DiskHintsCallBackInfo{}
			if err := json.Unmarshal(v, &disk_hints); err != nil {
				logger.Warn("WriteImageTarProcessor.data Unmarshal error:", err.Error())
			} else {
				if err := synDiskHints(logger, command.Sn, disk_hints); err != nil {
					msg.Status = ERROR
					msg.Message = err.Error()
				}
			}
		}
	}

	b.BaseProcessor.CallBack(logger, command, msg)
}

func (b WriteImageTarProcessor) doProcess(logger *log.Logger, command *commandDao.Command) {

	instance, err := instanceLogic.GetByInstanceId(logger, command.InstanceID)
	if err != nil {
		logger.Warn("WriteImageTarProcessor GetByInstanceId sql error:", command.InstanceID, err.Error())
	}
	if instance == nil {
		panic(ProcessAbortException{Msg: fmt.Sprintf("process WriteImageTar error, instance %s not found", command.InstanceID)})
	}

	image_entity, err := imageLogic.GetByImageId(logger, instance.ImageID)
	if err != nil {
		logger.Warn("WriteImageTarProcessor GetByImageId sql error:", instance.ImageID, err.Error())
	}
	if image_entity == nil {
		panic(ProcessAbortException{Msg: fmt.Sprintf("process WriteImageTar error, image %s not found", instance.ImageID)})
	}
	os_entity, err := osLogic.GetByOsId(logger, image_entity.OsID)
	if err != nil {
		logger.Warn("WriteImageTarProcessor GetByOsId sql error:", image_entity.OsID, err.Error())
	}
	if os_entity == nil {
		panic(ProcessAbortException{Msg: fmt.Sprintf("process WriteImageTar error, os %s not found", image_entity.OsID)})
	}

	write_image := agentTypes.WriteImageTar{
		Sn:        command.Sn,
		Url:       image_entity.URL,
		Format:    image_entity.Format,
		Filename:  image_entity.Filename,
		Hash:      image_entity.Hash,
		OsType:    os_entity.OsType,
		OsVersion: os_entity.OsVersion,
	}
	if err := rabbitMq.SendToAgent(command.Sn, write_image); err != nil {
		logger.Warnf("WriteImageTarProcessor SendToAgent error, msg:%s, error:%s", util.ObjToJson(write_image), err.Error())
		return
	}
	logger.Infof("WriteImageTarProcessor_SendToAgent_Msg success, msg:%s", util.ObjToJson(write_image))
}

func (b WriteImageTarProcessor) afterProcess(logger *log.Logger, command *commandDao.Command) {
	//empty
}
