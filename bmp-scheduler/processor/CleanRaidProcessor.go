package processor

import (
	"fmt"

	commandDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/commandDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/deviceHintsDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/diskDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/deviceLogic"
	rabbitMq "coding.jd.com/aidc-bmp/bmp-scheduler/service/rabbit_mq"
	agentTypes "coding.jd.com/aidc-bmp/bmp-scheduler/types/agent"
	"coding.jd.com/aidc-bmp/bmp-scheduler/util"
	ironicAgentEvent "git.jd.com/cps-golang/ironic-common/ironic/event/command"
	log "git.jd.com/cps-golang/log"
)

type CleanRaidProcessor struct {
	BaseProcessor
}

func NewCleanRaidProcessor() CleanRaidProcessor {
	b := CleanRaidProcessor{}
	b.BaseProcessor.doProcess = b.doProcess
	b.BaseProcessor.afterProcess = b.afterProcess
	return b
}

func (b CleanRaidProcessor) Process(logger *log.Logger, command *commandDao.Command) {
	b.BaseProcessor.Process(logger, command)
}

func (b CleanRaidProcessor) Callback(logger *log.Logger, command *commandDao.Command, msg ironicAgentEvent.CallbackCommandMessage) {
	b.BaseProcessor.CallBack(logger, command, msg)
}

func (b CleanRaidProcessor) doProcess(logger *log.Logger, command *commandDao.Command) {

	device, err := deviceLogic.GetBySn(logger, command.Sn)
	if err != nil {
		logger.Warnf("CleanRaidProcessor GetBySn sql error, sn:%s, error:%s", command.Sn, err.Error())
	}
	if device == nil {
		panic(ProcessAbortException{Msg: fmt.Sprintf("CleanRaidProcessor error, device %s not found", command.Sn)})
	}

	//需要删除device_hints中的盘符信息，但是device_hints不光有盘符，还有nvmes，所以要借助disk
	query := map[string]interface{}{
		"is_del":    0,
		"device_id": device.DeviceID,
		"types":     "panfu",
	}
	diskEntities, err := diskDao.GetDisks(logger, query)
	serials := []string{}
	for _, diskEntity := range diskEntities {
		serials = append(serials, diskEntity.Serial)
	}
	if len(serials) > 0 {
		q := map[string]interface{}{
			"is_del":    0,
			"serial.in": serials,
		}
		u := map[string]interface{}{
			"is_del": 1,
		}
		if err := deviceHintsDao.UpdateMultiDeviceHints(logger, q, u); err != nil {
			logger.Warnf("CleanRaidProcessor.UpdateMultiDeviceHints error, sn:%s, error:%s", command.Sn, err.Error())
		}

	}

	clean_block_device := agentTypes.CleanRaid{
		Sn: command.Sn,
		RaidDatas: []agentTypes.RaidData{
			{
				AdapterId:  device.AdapterID,
				RaidDriver: device.RaidDriver,
			},
		},
	}
	if err := rabbitMq.SendToAgent(command.Sn, clean_block_device); err != nil {
		logger.Warn("CleanRaidProcessor SendToAgent error:", command.Sn, err.Error())
		return
	}
	logger.Infof("CleanRaidProcessor_SendToAgent_Msg:%s", util.ObjToJson(clean_block_device))
}

func (b CleanRaidProcessor) afterProcess(logger *log.Logger, command *commandDao.Command) {
	//empty
}
