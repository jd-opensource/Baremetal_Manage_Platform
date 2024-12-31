package processor

import (
	commandDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/commandDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/deviceHintsDao"
	rabbitMq "coding.jd.com/aidc-bmp/bmp-scheduler/service/rabbit_mq"
	agentTypes "coding.jd.com/aidc-bmp/bmp-scheduler/types/agent"
	"coding.jd.com/aidc-bmp/bmp-scheduler/util"
	ironicAgentEvent "git.jd.com/cps-golang/ironic-common/ironic/event/command"
	log "git.jd.com/cps-golang/log"
)

type CleanBlockDeviceProcessor struct {
	BaseProcessor
}

func NewCleanBlockDeviceProcessor() CleanBlockDeviceProcessor {
	b := CleanBlockDeviceProcessor{}
	b.BaseProcessor.doProcess = b.doProcess
	b.BaseProcessor.afterProcess = b.afterProcess
	return b
}

func (b CleanBlockDeviceProcessor) Process(logger *log.Logger, command *commandDao.Command) {
	b.BaseProcessor.Process(logger, command)
}

func (b CleanBlockDeviceProcessor) Callback(logger *log.Logger, command *commandDao.Command, msg ironicAgentEvent.CallbackCommandMessage) {
	b.BaseProcessor.CallBack(logger, command, msg)
}

func (b CleanBlockDeviceProcessor) doProcess(logger *log.Logger, command *commandDao.Command) {

	clean_block_device := agentTypes.CleanBlockDevice{
		Sn: command.Sn,
	}
	if err := rabbitMq.SendToAgent(command.Sn, clean_block_device); err != nil {
		logger.Warn("CleanBlockDeviceProcessor SendToAgent error:", command.Sn, err.Error())
		return
	}
	logger.Infof("CleanBlockDeviceProcessor_SendToAgent_Msg:%s", util.ObjToJson(clean_block_device))

	// 重装时再使用

	deviceHintEntity, err := deviceHintsDao.GetSystemVolumnDeviceHintsBySn(logger, command.Sn)
	if err != nil {
		logger.Warn("CleanBlockDeviceProcessor.GetAllRootDeviceHints error:", command.Sn, err.Error())
		return
	}
	// for _, entity := range entities {
	deviceHintEntity.IsDel = 1
	if err := deviceHintsDao.UpdateDeviceHintsById(logger, deviceHintEntity); err != nil {
		logger.Warn("CleanBlockDeviceProcessor.UpdateRootDeviceHintsById error:", err.Error())
	}
	// }

}

func (b CleanBlockDeviceProcessor) afterProcess(logger *log.Logger, command *commandDao.Command) {
	//empty
}
