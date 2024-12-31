package processor

import (
	"fmt"
	"strings"

	commandDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/commandDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/deviceHintsDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/instancePartitionDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/instanceLogic"
	rabbitMq "coding.jd.com/aidc-bmp/bmp-scheduler/service/rabbit_mq"
	agentTypes "coding.jd.com/aidc-bmp/bmp-scheduler/types/agent"
	"coding.jd.com/aidc-bmp/bmp-scheduler/types/domain"
	"coding.jd.com/aidc-bmp/bmp-scheduler/util"
	ironicAgentEvent "git.jd.com/cps-golang/ironic-common/ironic/event/command"
	log "git.jd.com/cps-golang/log"
)

type MountPartitionsProcessor struct {
	BaseProcessor
}

func NewMountPartitionsProcessor() MountPartitionsProcessor {
	b := MountPartitionsProcessor{}
	b.BaseProcessor.doProcess = b.doProcess
	b.BaseProcessor.afterProcess = b.afterProcess
	return b
}

func (b MountPartitionsProcessor) Process(logger *log.Logger, command *commandDao.Command) {
	b.BaseProcessor.Process(logger, command)
}

func (b MountPartitionsProcessor) Callback(logger *log.Logger, command *commandDao.Command, msg ironicAgentEvent.CallbackCommandMessage) {
	b.BaseProcessor.CallBack(logger, command, msg)
}

func (b MountPartitionsProcessor) doProcess(logger *log.Logger, command *commandDao.Command) {

	logger.Info("MountPartitionsProcessor doProcess......", command.Sn)
	instancePartitionsEntity, err := instancePartitionDao.GetInstancePartitionsByInstanceId(logger, command.InstanceID)
	if instancePartitionsEntity == nil {
		logger.Warnf("FormatPartitionsProcessor GetInstancePartitionsByInstanceId error, sn:%s, instance_id:%s, err:%+v", command.Sn, command.InstanceID, err)
	}

	instance, err := instanceLogic.GetByInstanceId(logger, command.InstanceID)
	if err != nil {
		logger.Warn("MountPartitionsProcessor GetByInstanceId sql error:", command.InstanceID, err.Error())
	}
	if instance == nil {
		panic(ProcessAbortException{Msg: fmt.Sprintf("process MountPartitionsProcessor error, instance %s not found", command.InstanceID)})
	}
	rootDeviceHintsEntity, err := deviceHintsDao.GetSystemVolumnDeviceHintsBySn(logger, command.Sn)
	if err != nil {
		logger.Warn("MountPartitionsProcessor.GetSystemVolumnDeviceHintsBySn error:", command.Sn, err.Error())
	}
	mountPartitions := agentTypes.MonutPartitions{
		Sn:           command.Sn,
		Version:      "2.0",
		AutoMountEfi: false, //bios的时候为false，uefi的时候为true
	}
	if strings.ToUpper(instance.BootMode) == "UEFI" {
		mountPartitions.AutoMountEfi = true
	}

	ms := []agentTypes.Mount{}
	for _, v := range instancePartitionsEntity {
		if v.PartitionType == "data" { //数据盘
			continue
		}
		m := agentTypes.Mount{
			IsRootDevice: v.PartitionType != "data",
			IsDataDevice: v.PartitionType == "data",
			Label:        v.PartitionLabel,
			Mountpoint:   v.PartitionMountPoint,
			// Options:      v.Options,
		}
		if rootDeviceHintsEntity != nil {
			m.DiskHints = &domain.RootDeviceHints{
				// Name: rootDeviceHintsEntity.Name,
				Serial: rootDeviceHintsEntity.Serial,
			}
		}

		ms = append(ms, m)
	}
	mountPartitions.Mounts = ms

	if err := rabbitMq.SendToAgent(command.Sn, mountPartitions); err != nil {
		logger.Warnf("MountPartitionsProcessor SendToAgent error, msg:%s, error:%s", util.ObjToJson(mountPartitions), err.Error())
		return
	}
	logger.Infof("MountPartitionsProcessor_SendToAgent_Msg success, msg:%s", util.ObjToJson(mountPartitions))

}

func (b MountPartitionsProcessor) afterProcess(logger *log.Logger, command *commandDao.Command) {
	//empty
}
