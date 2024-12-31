package raidLogic

import (
	"time"

	"coding.jd.com/aidc-bmp/bmp-openapi/constant"

	raidDao "coding.jd.com/aidc-bmp/bmp-openapi/dao/raidDao"
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
	responseTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/response"
	"coding.jd.com/aidc-bmp/bmp-openapi/util"
	log "coding.jd.com/aidc-bmp/bmp_log"
	namespacePrefix "git.jd.com/cps-golang/ironic-common/ironic/common/NamespacePrefixs"
	commonUtil "git.jd.com/cps-golang/ironic-common/ironic/util"
)

func DRaid2Raid(d *raidDao.DRaid) *responseTypes.Raid {
	return &responseTypes.Raid{
		//RaidID:               d.RaidId,
		//Name:                 d.Name,
		//DescriptionEn:        d.DescriptionEn,
		//DescriptionZh:        d.DescriptionZh,
		//DeviceType:           d.DeviceType,
		//VolumeType:           d.VolumeType,
		//VolumeDetail:         d.VolumeDetail,
		//SystemPartitionCount: d.SystemPartitionCount,
		//AvailableValue:       d.AvailableValue,
		//DiskType:             d.DiskType,
	}
}

func RaidEntity2Raid(d *raidDao.Raid) *responseTypes.Raid {
	return &responseTypes.Raid{
		RaidID:        d.RaidID,
		Name:          d.Name,
		DescriptionEn: d.DescriptionEn,
		DescriptionZh: d.DescriptionZh,
	}
}

func CreateRaid(logger *log.Logger, param *requestTypes.CreateRaidRequest) (string, error) {

	entity := &raidDao.Raid{
		Name:          param.Name,
		DescriptionEn: param.DescriptionEn,
		DescriptionZh: param.DescriptionZh,
		CreatedBy:     logger.GetPoint("username").(string),
		CreatedTime:   int(time.Now().Unix()),
		UpdatedBy:     logger.GetPoint("username").(string),
		UpdatedTime:   int(time.Now().Unix()),
	}
	entity.RaidID = commonUtil.GetRandString(namespacePrefix.RAID_ID_PREFIX, namespacePrefix.RAID_ID_LENGTH, false, true, true)
	if _, err := raidDao.AddRaid(logger, entity); err != nil {
		logger.Warnf("CreateRaid AddRaid sql error, entity:%s, error:%s", util.ObjToJson(entity), err.Error())
		return "", err
	}
	return entity.RaidID, nil
}

//func CreateDeviceTypeRaid(logger *log.Logger, param *requestTypes.CreateDeviceTypeRaidRequest) error {
//	entity := &rDeviceTypeRaidDao.RDeviceTypeRaid{
//		DeviceTypeID:         param.DeviceType,
//		RaidID:               param.RaidID,
//		VolumeDetail:         param.VolumeDetail,
//		VolumeType:           param.VolumeType,
//		AvailableValue:       param.AvailableValue,
//		SystemPartitionCount: param.SystemPartitionCount,
//		DiskType:             param.DiskType,
//		CreatedBy:            logger.GetPoint("username").(string),
//		CreatedTime:          int(time.Now().Unix()),
//		UpdatedBy:            logger.GetPoint("username").(string),
//		UpdatedTime:          int(time.Now().Unix()),
//	}
//	if _, err := rDeviceTypeRaidDao.AddRDeviceTypeRaid(logger, entity); err != nil {
//		logger.Warnf("CreateDeviceTypeRaid AddRDeviceTypeRaid sql error, entity:%s, error:%s", util.ObjToJson(entity), err.Error())
//		return err
//	}
//	return nil
//}
//
//func DeleteDeviceTypeRaid(logger *log.Logger, param *requestTypes.DeleteDeviceTypeRaidRequest) error {
//
//	query := map[string]interface{}{
//		"device_type": param.DeviceType,
//		"raid_id":     param.RaidId,
//		"volume_type": param.VolumeType,
//	}
//	entity, err := rDeviceTypeRaidDao.GetOneRDeviceTypeRaid(logger, query)
//	if err != nil {
//		logger.Warn("DeleteDeviceTypeRaid_GetOneRDeviceTypeRaid sql error:", param.DeviceType, param.RaidId, param.VolumeType, err.Error())
//		return err
//	}
//
//	entity.IsDel = 1
//	entity.UpdatedBy = logger.GetPoint("username").(string)
//	entity.UpdatedTime = int(time.Now().Unix())
//	entity.DeletedTime = int(time.Now().Unix())
//
//	if err := rDeviceTypeRaidDao.UpdateRDeviceTypeRaidById(logger, entity); err != nil {
//		logger.Warnf("DeleteDeviceTypeRaid UpdateRDeviceTypeRaidById sql error, entity:%s, error:%s", util.ObjToJson(entity), err.Error())
//		return err
//	}
//	return nil
//}

//func QueryByRequest(logger *log.Logger, param *requestTypes.QueryRaidsRequest) ([]*responseTypes.Raid, error) {
//
//	d_raids, err := raidDao.GetAllRraidDevice(logger, param.RaidID, param.DeviceType, param.VolumeType)
//	if err != nil {
//		logger.Warn("raidDao.GetAllRraidDevice sql error:", param.RaidID, param.DeviceType, param.VolumeType, err.Error())
//		return nil, err
//	}
//	raids := []*responseTypes.Raid{}
//	for _, raid := range d_raids {
//		raids = append(raids, DRaid2Raid(raid))
//	}
//	return raids, nil
//}

func QueryRaidsAll(logger *log.Logger) ([]*responseTypes.Raid, error) {

	param := map[string]interface{}{
		"is_del": 0,
	}
	entities, err := raidDao.GetAllRaid(logger, param)
	if err != nil {
		logger.Warn("QueryRaidsAll GetAllRaid sql error:", err.Error())
		return nil, err
	}
	raids := []*responseTypes.Raid{}
	for _, entity := range entities {
		//rDeviceTypeRaid, err := rDeviceTypeRaidDao.GetOneRDeviceTypeRaid(logger, map[string]interface{}{"raid_id": entity.RaidID})
		//if err != nil {
		//	logger.Warnf("QueryRDeviceTypeRaid sql error, raid_id:%s:, error:%s", entity.RaidID, err.Error())
		//	rDeviceTypeRaid = &rDeviceTypeRaidDao.RDeviceTypeRaid{}
		//}
		raids = append(raids, RaidEntity2Raid(entity))
	}
	return raids, nil
}

func GetByRaidId(logger *log.Logger, raidId string) (*responseTypes.Raid, error) {
	entity, err := raidDao.GetRaidByUuid(logger, raidId)
	if err != nil {
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}
	//rDeviceTypeRaid, err := rDeviceTypeRaidDao.GetOneRDeviceTypeRaid(logger, map[string]interface{}{"raid_id": entity.RaidID})
	//if err != nil {
	//	logger.Warnf("QueryRDeviceTypeRaid sql error, raid_id:%s, error:%s", entity.RaidID, err.Error())
	//	rDeviceTypeRaid = &rDeviceTypeRaidDao.RDeviceTypeRaid{}
	//}
	return RaidEntity2Raid(entity), nil
}

func GetRaidEntityByRaidId(logger *log.Logger, raidId string) (*raidDao.Raid, error) {
	entity, err := raidDao.GetRaidByUuid(logger, raidId)
	if err != nil {
		logger.Warn("GetRaidByUuid sql error:", raidId, err.Error())
		return nil, err
	}
	return entity, nil
}

func GetAndCheckByRaidId(logger *log.Logger, raidId string) (*responseTypes.Raid, error) {
	raid, err := GetByRaidId(logger, raidId)
	if err != nil {
		return nil, err
	}
	return raid, nil
}

func ModifyRaid(logger *log.Logger, param *requestTypes.ModifyRaidRequest, raidId string) error {
	entity, err := raidDao.GetRaidByUuid(logger, raidId) //GetRaidEntityByRaidId(logger, raid_id, true)
	if err != nil {
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}
	entity.Name = param.Name
	entity.DescriptionEn = param.DescriptionEn
	entity.DescriptionZh = param.DescriptionZh
	entity.UpdatedBy = logger.GetPoint("username").(string)
	entity.UpdatedTime = int(time.Now().Unix())
	if err := raidDao.UpdateRaidById(logger, entity); err != nil {
		logger.Warn("ModifyRaid UpdateRaidById sql error:", raidId, err.Error())
		return err
	}
	return nil
}

func DeleteRaid(logger *log.Logger, raidId string) error {
	entity, err := raidDao.GetRaidByUuid(logger, raidId)
	if err != nil {
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}
	entity.IsDel = 1
	entity.UpdatedBy = logger.GetPoint("username").(string)
	entity.UpdatedTime = int(time.Now().Unix())
	entity.DeletedTime = int(time.Now().Unix())
	if err := raidDao.UpdateRaidById(logger, entity); err != nil {
		logger.Warn("DeleteRaid UpdateRaidById sql error:", raidId, err.Error())
		return err
	}
	return nil
}
