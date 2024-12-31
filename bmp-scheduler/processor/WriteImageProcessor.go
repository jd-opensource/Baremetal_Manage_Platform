package processor

import (
	"encoding/json"
	"fmt"
	"strings"

	commandDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/commandDao"
	deviceDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/deviceDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/deviceHintsDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/volumeDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/imageLogic"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/instanceLogic"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/osLogic"
	rabbitMq "coding.jd.com/aidc-bmp/bmp-scheduler/service/rabbit_mq"
	agentTypes "coding.jd.com/aidc-bmp/bmp-scheduler/types/agent"
	"coding.jd.com/aidc-bmp/bmp-scheduler/util"
	ironicAgentEvent "git.jd.com/cps-golang/ironic-common/ironic/event/command"
	log "git.jd.com/cps-golang/log"
)

type WriteImageProcessor struct {
	BaseProcessor
}

func NewWriteImageProcessor() WriteImageProcessor {
	b := WriteImageProcessor{}
	b.BaseProcessor.doProcess = b.doProcess
	b.BaseProcessor.afterProcess = b.afterProcess
	return b
}

func (b WriteImageProcessor) Process(logger *log.Logger, command *commandDao.Command) {
	b.BaseProcessor.Process(logger, command)
}

func (b WriteImageProcessor) Callback(logger *log.Logger, command *commandDao.Command, msg ironicAgentEvent.CallbackCommandMessage) {

	//以下是基于重装系统的，需要存储hints
	if strings.EqualFold(msg.Status, OK) {

		v, err := json.Marshal(msg.Data)
		if err != nil {
			logger.Warn("WriteImageProcessor.data marshal error:", err.Error())
		} else {
			disk_hints := DiskHintsCallBackInfo{}
			if err := json.Unmarshal(v, &disk_hints); err != nil {
				logger.Warn("WriteImageProcessor.data Unmarshal error:", err.Error())
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

func synDiskHints(logger *log.Logger, sn string, root_disk_hints DiskHintsCallBackInfo) error {

	deviceEntity, err := deviceDao.GetBySn(logger, sn)
	if err != nil {
		logger.Warnf("synDiskHints.GetDeviceBySn error, sn:%s, error:%s", sn, err.Error())
		return err
	}

	disk_hints := root_disk_hints.RootDeviceHints
	entity, err := deviceHintsDao.GetSystemVolumnDeviceHintsBySn(logger, sn)
	if err != nil {
		logger.Warn("MakePartitionsProcessor.synDiskHints.GetSystemVolumnDeviceHintsBySn error:", sn, err.Error())
	}
	if entity != nil {
		entity.Size = disk_hints.Size
		entity.Rotational = util.Bool2Int8(disk_hints.Rotational)
		entity.Wwn = disk_hints.Wwn
		entity.Name = disk_hints.Name
		entity.WwnVendorExtension = disk_hints.WwnVendorExtension
		entity.WwnWithExtension = disk_hints.WwnWithExtension
		entity.Model = disk_hints.Model
		entity.Serial = disk_hints.Serial
		entity.Vendor = disk_hints.Vendor
		entity.Hctl = disk_hints.Hctl
		entity.ByPath = disk_hints.ByPath
		entity.IsDel = 0
		if err := deviceHintsDao.UpdateDeviceHintsById(logger, entity); err != nil {
			logger.Warn("MakePartitionsProcessor.synDiskHints.UpdateDeviceHintsById error:", sn, err.Error())
			return err
		}
	} else {
		entity = &deviceHintsDao.DeviceHints{
			Sn:                 sn,
			Size:               disk_hints.Size,
			Rotational:         util.Bool2Int8(disk_hints.Rotational),
			Wwn:                disk_hints.Wwn,
			Name:               disk_hints.Name,
			WwnVendorExtension: disk_hints.WwnVendorExtension,
			WwnWithExtension:   disk_hints.WwnWithExtension,
			Model:              disk_hints.Model,
			Serial:             disk_hints.Serial,
			Vendor:             disk_hints.Vendor,
			Hctl:               disk_hints.Hctl,
			ByPath:             disk_hints.ByPath,
			IsDel:              0,
		}
		volumeEntity, err := volumeDao.GetSystemVolumeByDeviceTypeId(logger, deviceEntity.DeviceTypeID)
		if err != nil {
			logger.Warnf("MakePartitionsProcessor.synDiskHints.GetSystemVolumeByDeviceTypeId error, device_type_id:%s, error:%s", deviceEntity.DeviceTypeID, err.Error())
		} else {
			entity.VolumeId = volumeEntity.VolumeID
			entity.VolumeType = "system"
			if _, err := deviceHintsDao.AddDeviceHints(logger, entity); err != nil {
				logger.Warn("MakePartitionsProcessor.synDiskHints.AddDeviceHints error:", sn, err.Error())
				return err
			}
		}

	}
	return nil

}

func (b WriteImageProcessor) doProcess(logger *log.Logger, command *commandDao.Command) {

	instance, err := instanceLogic.GetByInstanceId(logger, command.InstanceID)
	if err != nil {
		logger.Warn("WriteImageProcessor GetByInstanceId sql error:", command.InstanceID, err.Error())
	}
	if instance == nil {
		panic(ProcessAbortException{Msg: fmt.Sprintf("process WriteImage error, instance %s not found", command.InstanceID)})
	}

	image_entity, err := imageLogic.GetByImageId(logger, instance.ImageID)
	if err != nil {
		logger.Warn("WriteImageProcessor GetByImageId sql error:", instance.ImageID, err.Error())
	}
	if image_entity == nil {
		panic(ProcessAbortException{Msg: fmt.Sprintf("process WriteImage error, image %s not found", instance.ImageID)})
	}
	os_entity, err := osLogic.GetByOsId(logger, image_entity.OsID)
	if err != nil {
		logger.Warn("WriteImageProcessor GetByOsId sql error:", image_entity.OsID, err.Error())
	}
	if os_entity == nil {
		panic(ProcessAbortException{Msg: fmt.Sprintf("process WriteImage error, os %s not found", image_entity.OsID)})
	}

	write_image := agentTypes.WriteImage{
		Sn:        command.Sn,
		Url:       image_entity.URL,
		Format:    image_entity.Format,
		Filename:  image_entity.Filename,
		Hash:      image_entity.Hash,
		OsType:    os_entity.OsType,
		OsVersion: os_entity.OsVersion,
	}
	if err := rabbitMq.SendToAgent(command.Sn, write_image); err != nil {
		logger.Warnf("WriteImageProcessor SendToAgent error, msg:%s, error:%s", util.ObjToJson(write_image), err.Error())
		return
	}
	logger.Infof("WriteImageProcessor_SendToAgent_Msg success, msg:%s", util.ObjToJson(write_image))
}

func (b WriteImageProcessor) afterProcess(logger *log.Logger, command *commandDao.Command) {
	//empty
}
