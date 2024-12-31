package processor

import (
	"fmt"
	"strings"

	commandDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/commandDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/deviceLogic"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/imageLogic"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/instanceLogic"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/osLogic"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/partitionLogic"
	rabbitMq "coding.jd.com/aidc-bmp/bmp-scheduler/service/rabbit_mq"
	agentTypes "coding.jd.com/aidc-bmp/bmp-scheduler/types/agent"
	domainTypes "coding.jd.com/aidc-bmp/bmp-scheduler/types/domain"
	"coding.jd.com/aidc-bmp/bmp-scheduler/util"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/InstanceStatus"
	ironicAgentEvent "git.jd.com/cps-golang/ironic-common/ironic/event/command"
	log "git.jd.com/cps-golang/log"
)

type Qcow2MakePartitionsProcessor struct {
	BaseProcessor
}

func NewQcow2MakePartitionsProcessor() Qcow2MakePartitionsProcessor {
	b := Qcow2MakePartitionsProcessor{}
	b.BaseProcessor.doProcess = b.doProcess
	b.BaseProcessor.afterProcess = b.afterProcess
	return b
}

func (b Qcow2MakePartitionsProcessor) Process(logger *log.Logger, command *commandDao.Command) {
	b.BaseProcessor.Process(logger, command)
}

func (b Qcow2MakePartitionsProcessor) Callback(logger *log.Logger, command *commandDao.Command, msg ironicAgentEvent.CallbackCommandMessage) {
	b.BaseProcessor.CallBack(logger, command, msg)
}

func (b Qcow2MakePartitionsProcessor) doProcess(logger *log.Logger, command *commandDao.Command) {

	instance, err := instanceLogic.GetByInstanceId(logger, command.InstanceID)
	if err != nil {
		logger.Warn("Qcow2MakePartitionsProcessor GetByInstanceId sql error:", command.InstanceID, err.Error())
	}
	if instance == nil {
		panic(ProcessAbortException{Msg: fmt.Sprintf("process Qcow2MakePartitions error, instance %s not found", command.InstanceID)})
	}
	image_entity, err := imageLogic.GetByImageId(logger, instance.ImageID)
	if err != nil {
		logger.Warn("Qcow2MakePartitionsProcessor GetByImageId sql error:", instance.ImageID, err.Error())
	}
	if image_entity == nil {
		panic(ProcessAbortException{Msg: fmt.Sprintf("process Qcow2MakePartitions error, image %s not found", instance.ImageID)})
	}
	os_entity, err := osLogic.GetByOsId(logger, image_entity.OsID)
	if err != nil {
		logger.Warn("Qcow2MakePartitionsProcessor GetByOsId sql error:", image_entity.OsID, err.Error())
	}
	device_entity, err := deviceLogic.GetById(logger, instance.DeviceID)
	if err != nil {
		logger.Warn("Qcow2MakePartitionsProcessor GetById sql error:", instance.DeviceID, err.Error())
	}
	if device_entity == nil {
		panic(ProcessAbortException{Msg: fmt.Sprintf("process Qcow2MakePartitions error, device %s not found", fmt.Sprint(instance.DeviceID))})
	}
	partitions_entity, err := partitionLogic.GetByDeviceTypeAndImageId(logger, device_entity.DeviceTypeID, instance.ImageID)
	if err != nil {
		logger.Warn("Qcow2MakePartitionsProcessor GetByDeviceTypeAndImageId sql error:", device_entity.DeviceTypeID, instance.ImageID, err.Error())
	}
	if len(partitions_entity) == 0 {
		sendDefaultQcow2MakePartitionsCommand(logger, command)
		return
	}
	make_partition := agentTypes.Qcow2MakePartitions{
		Sn:        command.Sn,
		OsType:    os_entity.OsType,
		OsVersion: os_entity.OsVersion,
	}

	make_partition.SystemPartitions = []domainTypes.Partition{}
	make_partition.DataPartitions = []domainTypes.Partition{}
	for _, partition := range partitions_entity {
		make_partition.DataDiskLabel = partition.DataDiskLabel
		p := domainTypes.Partition{
			Size:       partition.PartitionSize,
			FsType:     partition.PartitionFsType,
			MountPoint: partition.PartitionMountPoint,
			Label:      partition.PartitionLabel,
		}
		switch strings.ToLower(partition.PartitionType) {
		case SYSTEM:
			make_partition.SystemPartitions = append(make_partition.SystemPartitions, p)
		case DATA:
			make_partition.DataPartitions = append(make_partition.DataPartitions, p)
		}
	}
	make_partition.KeepData = false
	instance_extra := loadInstanceExtra(logger, command.InstanceID)
	if instance_extra != nil && instance.Status == InstanceStatus.REINSTALLING {
		if instance_extra.KeepData {
			make_partition.KeepData = true
		}
	}
	if err := rabbitMq.SendToAgent(command.Sn, make_partition); err != nil {
		logger.Warn("Qcow2MakePartitionsProcessor SendToAgent error, msg:%s, error:%s", util.ObjToJson(make_partition), err.Error())
		return
	}
	logger.Infof("Qcow2MakePartitionsProcessor_SendToAgent_Msg success, msg:%s", util.ObjToJson(make_partition))

}

func (b Qcow2MakePartitionsProcessor) afterProcess(logger *log.Logger, command *commandDao.Command) {
	//empty
}

func sendDefaultQcow2MakePartitionsCommand(logger *log.Logger, command *commandDao.Command) {
	make_partition := agentTypes.Qcow2MakePartitions{
		Sn: command.Sn,
	}
	if err := rabbitMq.SendToAgent(command.Sn, make_partition); err != nil {
		logger.Warnf("Qcow2MakePartitionsProcessor sendDefaultCommand error, msg:%s, error:%s", util.ObjToJson(make_partition), err.Error())
		return
	}

}
