package raidLogic

import (
	raidDao "coding.jd.com/aidc-bmp/bmp-openapi/dao/raidDao"
	responseTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/response"
	log "coding.jd.com/aidc-bmp/bmp_log"
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

func GetRaidEntityByRaidId(logger *log.Logger, raidId string) (*raidDao.Raid, error) {
	entity, err := raidDao.GetRaidByUuid(logger, raidId)
	if err != nil {
		logger.Warn("GetRaidByUuid sql error:", raidId, err.Error())
		return nil, err
	}
	return entity, nil
}
