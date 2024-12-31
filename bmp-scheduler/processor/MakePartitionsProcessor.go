package processor

import (
	"fmt"
	"strings"

	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/deviceHintsDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/imageDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/instanceLogic"

	commandDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/commandDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/instancePartitionDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/osDao"
	rabbitMq "coding.jd.com/aidc-bmp/bmp-scheduler/service/rabbit_mq"
	agentTypes "coding.jd.com/aidc-bmp/bmp-scheduler/types/agent"
	domainTypes "coding.jd.com/aidc-bmp/bmp-scheduler/types/domain"
	util "coding.jd.com/aidc-bmp/bmp-scheduler/util"
	ironicAgentEvent "git.jd.com/cps-golang/ironic-common/ironic/event/command"
	log "git.jd.com/cps-golang/log"
)

const (
	BOOT   = "boot"
	ROOT   = "root"
	SYSTEM = "system"
	DATA   = "data"
)

type MakePartitionsProcessor struct {
	BaseProcessor
}

func NewMakePartitionsProcessor() MakePartitionsProcessor {
	b := MakePartitionsProcessor{}
	b.BaseProcessor.doProcess = b.doProcess
	b.BaseProcessor.afterProcess = b.afterProcess
	return b
}

type DiskHintsCallBackInfo struct {
	RootDeviceHints domainTypes.RootDeviceHints `json:"root_device_hints"`
}

func (b MakePartitionsProcessor) Process(logger *log.Logger, command *commandDao.Command) {
	b.BaseProcessor.Process(logger, command)
}

func (b MakePartitionsProcessor) Callback(logger *log.Logger, command *commandDao.Command, msg ironicAgentEvent.CallbackCommandMessage) {

	b.BaseProcessor.CallBack(logger, command, msg)
}

func (b MakePartitionsProcessor) doProcess(logger *log.Logger, command *commandDao.Command) {
	fmt.Println(command.InstanceID)
	instance, err := instanceLogic.GetByInstanceId(logger, command.InstanceID)
	if err != nil {
		logger.Warn("MakePartitionsProcessor GetByInstanceId sql error:", command.InstanceID, err.Error())
	}
	if instance == nil {
		panic(ProcessAbortException{Msg: fmt.Sprintf("process MakePartitions error, instance %s not found", command.InstanceID)})
	}
	image, _ := imageDao.GetImageByUuid(logger, instance.ImageID)
	os, _ := osDao.GetOsByUuid(logger, image.OsID)
	fmt.Println(os, os.OsType, os.OsVersion)
	list, _ := instancePartitionDao.GetAllInstancePartition(logger, map[string]interface{}{
		"instance_id":       command.InstanceID,
		"image_id":          instance.ImageID,
		"partition_type.in": []string{"root", "boot", "system"},
		"is_del":            0,
	})

	make_partition := agentTypes.MakePartitions{
		Sn:                 command.Sn,
		Version:            "2.0",
		AutoCreateEfi:      true,
		AutoCreateBiosGrub: true,
	}
	ps := []agentTypes.MPartition{}
	for _, partition := range list {

		m := agentTypes.MPartition{
			ID:       partition.InstancePartitionID,
			PartType: partition.PartitionType,
			FsType:   partition.PartitionFsType,
			Number:   partition.Number,
			Size:     partition.PartitionSize,
		}

		ps = append(ps, m)
	}
	make_partition.Disks = []agentTypes.Disk{
		{
			IsRootDevice: true,
			DiskLabel:    "gpt",
			Partitions:   ps,
			KeepData:     false,
		},
	}
	entity, err := deviceHintsDao.GetSystemVolumnDeviceHintsBySn(logger, command.Sn)
	if err != nil {
		logger.Warn("MakePartitionsProcessor.GetSystemVolumnDeviceHintsBySn error:", command.Sn, err.Error())
	}

	if instance.Status != "creating" && entity != nil {
		rootDeviceHints := domainTypes.RootDeviceHints{
			// Size:               entity.Size,
			// Rotational:         util.Int82Bool(entity.Rotational),
			// Wwn:                entity.Wwn,
			// Name: entity.Name,
			// WwnVendorExtension: entity.WwnVendorExtension,
			// WwnWithExtension:   entity.WwnWithExtension,
			// Model:              entity.Model,
			Serial: entity.Serial,
			// Vendor:             entity.Vendor,
			// Hctl:               entity.Hctl,
			// ByPath:             entity.ByPath,
		}
		make_partition.Disks[0].DiskHints = &rootDeviceHints
	} else if instance.Status == "reinstalling" {
		logger.Warnf("rootDeviceHints deal error, sn:%s", command.Sn)
	}

	make_partition.BootMode = instance.BootMode                    //以实例的bootmode为准
	if strings.ToUpper(make_partition.BootMode) == "Legacy/BIOS" { //文案转换
		make_partition.BootMode = "bios"
	} else if strings.ToUpper(make_partition.BootMode) == "UEFI" {
		make_partition.BootMode = "uefi"
	} else { //bios兜底
		make_partition.BootMode = "bios"
	}

	if err := rabbitMq.SendToAgent(command.Sn, make_partition); err != nil {
		logger.Warnf("MakePartitionsProcessor SendToAgent error, sn:%s, msg:%s, error:%s", command.Sn, util.ObjToJson(make_partition), err.Error())
		return
	}
	logger.Infof("MakePartitionsProcessor_SendToAgent_Msg success, sn:%s, msg:%s", command.Sn, util.ObjToJson(make_partition))
	/*以下是基于重装系统的，需要存储hits
	instance, err := instanceLogic.GetByInstanceId(logger, command.InstanceID)
	if err != nil {
		logger.Warn("MakePartitionsProcessor GetByInstanceId sql error:", command.InstanceID, err.Error())
	}

	if instance == nil {
		panic(ProcessAbortException{Msg: fmt.Sprintf("process MakePartitions error, instance %s not found", command.InstanceID)})
	}
	image_entity, err := imageLogic.GetByImageId(logger, instance.ImageID)
	if err != nil {
		logger.Warn("MakePartitionsProcessor GetByImageId sql error:", instance.ImageID, err.Error())
	}
	if image_entity == nil {
		panic(ProcessAbortException{Msg: fmt.Sprintf("process MakePartitions error, image %s not found", instance.ImageID)})
	}
	os_entity, err := osLogic.GetByOsId(logger, image_entity.OsID)
	if err != nil {
		logger.Warn("MakePartitionsProcessor GetByOsId sql error:", image_entity.OsID, err.Error())
	}
	device_entity, err := deviceLogic.GetById(logger, instance.DeviceID)
	if err != nil {
		logger.Warn("MakePartitionsProcessor GetById sql error:", instance.DeviceID, err.Error())
	}
	if device_entity == nil {
		panic(ProcessAbortException{Msg: fmt.Sprintf("process MakePartitions error, device %s not found", fmt.Sprint(instance.DeviceID))})
	}
	partitions_entity, err := partitionLogic.GetByDeviceTypeAndImageId(logger, device_entity.DeviceTypeID, instance.ImageID)
	if err != nil {
		logger.Warn("MakePartitionsProcessor GetByDeviceTypeAndImageId sql error:", device_entity.DeviceTypeID, instance.ImageID, err.Error())
	}
	if len(partitions_entity) == 0 {
		sendDefaultMakePartitionsCommand(logger, command)
		return
	}

	make_partition := mergePartitions(logger, command, os_entity, instance, partitions_entity, instance.ImageID)
	make_partition.KeepData = false

	instance_extra := loadInstanceExtra(logger, command.InstanceID)
	if instance_extra != nil && instance.Status == InstanceStatus.REINSTALLING {
		if instance_extra.KeepData {
			make_partition.KeepData = true
		}
	}

	entity, err := rootDeviceHintsDao.GetRootDeviceHintsBySn(logger, command.Sn)
	if err != nil {
		logger.Warn("MakePartitionsProcessor.GetRootDeviceHintsBySn error:", command.Sn, err.Error())
	}

	if instance.Status != InstanceStatus.CREATING && entity != nil {
		rootDeviceHints := domainTypes.RootDeviceHints{
			Size:               entity.Size,
			Rotational:         util.Int82Bool(entity.Rotational),
			Wwn:                entity.Wwn,
			Name:               entity.Name,
			WwnVendorExtension: entity.WwnVendorExtension,
			WwnWithExtension:   entity.WwnWithExtension,
			Model:              entity.Model,
			Serial:             entity.Serial,
			Vendor:             entity.Vendor,
			Hctl:               entity.Hctl,
			ByPath:             entity.ByPath,
		}
		make_partition.RootDeviceHints = &rootDeviceHints
	}

	if err := rabbitMq.SendToAgent(command.Sn, make_partition); err != nil {
		logger.Warnf("MakePartitionsProcessor SendToAgent error, sn:%s, msg:%s, error:%s", command.Sn, util.ObjToJson(make_partition), err.Error())
		return
	}
	logger.Infof("MakePartitionsProcessor_SendToAgent_Msg success, sn:%s, msg:%s", command.Sn, util.ObjToJson(make_partition))
	*/
}

func (b MakePartitionsProcessor) afterProcess(logger *log.Logger, command *commandDao.Command) {
	//empty
}

func sendDefaultMakePartitionsCommand(logger *log.Logger, command *commandDao.Command) {
	make_partition := agentTypes.MakePartitions{
		Sn: command.Sn,
	}
	if err := rabbitMq.SendToAgent(command.Sn, make_partition); err != nil {
		logger.Warnf("MakePartitionsProcessor sendDefaultCommand error, msg:%s, error:%s", util.ObjToJson(make_partition), err.Error())
		return
	}

}
