package processor

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/deviceHintsDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/diskDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/rInstanceVolumeRaidDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/volumeDao"
	rabbitMq "coding.jd.com/aidc-bmp/bmp-scheduler/service/rabbit_mq"
	agentTypes "coding.jd.com/aidc-bmp/bmp-scheduler/types/agent"
	"coding.jd.com/aidc-bmp/bmp-scheduler/util"

	commandDao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/commandDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/deviceLogic"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/instanceLogic"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/raidLogic"
	ironicAgentEvent "git.jd.com/cps-golang/ironic-common/ironic/event/command"
	log "git.jd.com/cps-golang/log"
)

type MakeRaidProcessor struct {
	BaseProcessor
}

func NewMakeRaidProcessor() MakeRaidProcessor {
	b := MakeRaidProcessor{}
	b.BaseProcessor.doProcess = b.doProcess
	b.BaseProcessor.afterProcess = b.afterProcess
	return b
}

func (b MakeRaidProcessor) Process(logger *log.Logger, command *commandDao.Command) {
	b.BaseProcessor.Process(logger, command)
}

func (b MakeRaidProcessor) Callback(logger *log.Logger, command *commandDao.Command, msg ironicAgentEvent.CallbackCommandMessage) {
	logger.Info("MakeRaidProcessor Callback starting...", msg.Sn)
	defer logger.Info("MakeRaidProcessor Callback ending...", msg.Sn)
	if strings.ToUpper(msg.Status) == OK {
		logger.Infof("MakeRaidProcessor callback content is:%s", util.ObjToJson(msg.Data))
		//以后再优化成更优雅的方法
		b, err := json.Marshal(msg.Data)
		if err != nil {
			logger.Warn("MakeRaidProcessor Callback unmarshal error:", msg.Sn, err.Error())
		} else {
			saveDeviceHints(logger, msg.Sn, b)
		}
	}
	b.BaseProcessor.CallBack(logger, command, msg)
}

func (b MakeRaidProcessor) doProcess(logger *log.Logger, command *commandDao.Command) {

	logger.Info("MakeRaidProcessor doProcess......", command.Sn)
	instance, err := instanceLogic.GetByInstanceId(logger, command.InstanceID)
	if err != nil {
		logger.Warn("MakeRaidProcessor GetByInstanceId sql error:", command.InstanceID, err.Error())
	}
	if instance == nil {
		panic(ProcessAbortException{Msg: fmt.Sprintf("process MakeRaid error, instance %s not found", command.InstanceID)})
	}

	sys_raid, err := raidLogic.GetByRaidId(logger, instance.SystemVolumeRaidID)
	if err != nil {
		logger.Warn("MakeRaidProcessor GetByRaidId sql error:", instance.SystemVolumeRaidID, err.Error())
	}
	if sys_raid == nil {
		panic(ProcessAbortException{Msg: fmt.Sprintf("process MakeRaid error, sys_raid %s not found", instance.SystemVolumeRaidID)})
	}

	//data_raid, err := raidLogic.GetByRaidId(logger, instance.DataVolumeRaidID)
	//if err != nil {
	//	logger.Warn("MakeRaidProcessor GetByRaidId sql error:", instance.DataVolumeRaidID, err.Error())
	//} else {
	//	if command.Sn != "CFC7HP2" { //TODO trick逻辑，记得放开。因为机型不匹配，此机器的数据盘为nvme无法做raid
	//		raid_levels = append(raid_levels, data_raid.Name)
	//	}
	//
	//}数据盘暂时不做raid
	device, err := deviceLogic.GetById(logger, instance.DeviceID)
	if err != nil {
		logger.Warn("MakeRaidProcessor GetById sql error:", instance.DeviceID, err.Error())
	}
	if device == nil {
		panic(ProcessAbortException{Msg: fmt.Sprintf("process MakeRaid error, device %s not found", fmt.Sprint(instance.DeviceID))})
	}

	var systemVolumeItem agentTypes.Volume
	var dataVolumeItem = []agentTypes.Volume{}
	volumeEntities, err := rInstanceVolumeRaidDao.GetRInstanceVolumeRaids(logger, instance.InstanceID)
	if err != nil {
		logger.Warnf("MakeRaidProcessor.GetRInstanceVolumeRaids error, instance_id:%s, err:%s", instance.InstanceID, err.Error())
	}
	for _, volumeEntity := range volumeEntities {
		diskEntities, err := diskDao.GetVolumeDisks(logger, instance.DeviceID, volumeEntity.VolumeID)
		if err != nil {
			logger.Warnf("MakeRaidProcessor.GetAllRDeviceVolumeDisks error, device_id:%s, volume_id:%s, error:%s", instance.DeviceID, volumeEntity.VolumeID, err.Error())
		}
		if volumeEntity.VolumeType == "system" {
			systemVolumeItem = agentTypes.Volume{
				VolumeID:      volumeEntity.VolumeID,
				IsRootDevice:  true,
				RaidLevel:     volumeEntity.RaidName,
				PhysicalDisks: getPhysicalDisk(diskEntities),
			}
		} else {
			dataVolumeItem = append(dataVolumeItem, agentTypes.Volume{
				VolumeID:      volumeEntity.VolumeID,
				IsDataDevice:  true,
				RaidLevel:     volumeEntity.RaidName,
				PhysicalDisks: getPhysicalDisk(diskEntities),
			})
		}
	}

	volumes := append([]agentTypes.Volume{systemVolumeItem}, dataVolumeItem...)

	makeRaidData := agentTypes.MakeRaidData{
		RaidDriver: device.RaidDriver,
		AdapterID:  device.AdapterID,
		Volumes:    volumes,
	}

	make_data_raid := agentTypes.MakeRaid{
		Sn:            command.Sn,
		MakeRaidDatas: []agentTypes.MakeRaidData{makeRaidData},
	}
	if err := rabbitMq.SendToAgent(command.Sn, make_data_raid); err != nil {
		logger.Warnf("MakeRaidProcessor SendToAgent error, msg:%s, error:%s", util.ObjToJson(make_data_raid), err.Error())
		return
	}
	logger.Infof("MakeRaidProcessor_SendToAgent_Msg success, msg:%s", util.ObjToJson(make_data_raid))

}

func getPhysicalDisk(disks []*diskDao.Disk) []agentTypes.PhysicalDisk {
	res := []agentTypes.PhysicalDisk{}
	for _, v := range disks {
		ec, _ := strconv.Atoi(v.Enclosure)
		res = append(res, agentTypes.PhysicalDisk{
			Enclosure: ec,
			Slot:      v.Slot,
		})
	}
	return res
}

func (b MakeRaidProcessor) afterProcess(logger *log.Logger, command *commandDao.Command) {
	//empty
}

func saveDeviceHints(logger *log.Logger, sn string, content []byte) error {

	originContent := string(content)
	logger.Infof("saveDeviceHints sn: %s, origin content:%s", sn, originContent)
	//controllers.disks[0]的struct

	var err error
	c := agentTypes.MakeRaidAgentRespData{}
	if err := json.Unmarshal(content, &c); err != nil {
		logger.Warnf("SaveDiskLocationInfo parse error, sn:%s, content:%s, error:%s", sn, originContent, err.Error())
		return err
	}

	if len(c.MakeRaidDatas) > 0 {
		for _, v := range c.MakeRaidDatas[0].Volumes {
			deviceHintsEntity, _ := deviceHintsDao.GetDeviceHintsBySnAndVolumeId(logger, sn, v.VolumeId)
			if deviceHintsEntity == nil { //新建
				entity := &deviceHintsDao.DeviceHints{
					Sn:                 sn,
					VolumeId:           v.VolumeId,
					Size:               v.RootDeviceHints.Size,
					Rotational:         util.Bool2Int8(v.RootDeviceHints.Rotational),
					Wwn:                v.RootDeviceHints.Wwn,
					Name:               v.RootDeviceHints.Name,
					WwnVendorExtension: v.RootDeviceHints.WwnVendorExtension,
					WwnWithExtension:   v.RootDeviceHints.WwnWithExtension,
					Model:              v.RootDeviceHints.Model,
					Serial:             v.RootDeviceHints.Serial,
					Hctl:               v.RootDeviceHints.Hctl,
					ByPath:             v.RootDeviceHints.ByPath,
					Vendor:             v.RootDeviceHints.Vendor,
				}
				volumeEntity, err := volumeDao.GetVolumeById(logger, v.VolumeId)
				if err != nil {
					logger.Warnf("makeRaidProcessor.Callback.GetVolumeById error, sn:%s, volume_id:%s, error:%s", sn, v.VolumeId, err.Error())
				}
				if volumeEntity != nil {
					entity.VolumeType = volumeEntity.VolumeType
				}

				if _, err := deviceHintsDao.AddDeviceHints(logger, entity); err != nil {
					logger.Warnf("makeRaidProcessor.Callback.saveDeviceHints error, sn:%s, error:%s", sn, err.Error())
				}
			} else { //更新
				deviceHintsEntity.Size = v.RootDeviceHints.Size
				deviceHintsEntity.Rotational = util.Bool2Int8(v.RootDeviceHints.Rotational)
				deviceHintsEntity.Wwn = v.RootDeviceHints.Wwn
				deviceHintsEntity.Name = v.RootDeviceHints.Name
				deviceHintsEntity.WwnVendorExtension = v.RootDeviceHints.WwnVendorExtension
				deviceHintsEntity.WwnWithExtension = v.RootDeviceHints.WwnWithExtension
				deviceHintsEntity.Model = v.RootDeviceHints.Model
				deviceHintsEntity.Serial = v.RootDeviceHints.Serial
				deviceHintsEntity.Hctl = v.RootDeviceHints.Hctl
				deviceHintsEntity.ByPath = v.RootDeviceHints.ByPath
				deviceHintsEntity.Vendor = v.RootDeviceHints.Vendor
				if err := deviceHintsDao.UpdateDeviceHintsById(logger, deviceHintsEntity); err != nil {
					logger.Warnf("makeRaidProcessor.Callback.UpdateDeviceHintsById error, sn:%s, error:%s", sn, err.Error())
				}
			}

		}
	}

	return err

}
