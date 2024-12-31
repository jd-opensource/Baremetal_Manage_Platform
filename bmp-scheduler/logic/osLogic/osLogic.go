package osLogic

import (
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/osDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/imageLogic"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/instanceLogic"
	log "git.jd.com/cps-golang/log"
)

func GetByOsId(logger *log.Logger, os_id string) (*osDao.Os, error) {
	return osDao.GetOsByUuid(logger, os_id)
}

func GetOsByInstanceId(logger *log.Logger, instanceId string) (*osDao.Os, error) {
	instanceEntity, err := instanceLogic.GetByInstanceId(logger, instanceId)
	if err != nil {
		return nil, err
	}
	imageEntity, err := imageLogic.GetByImageId(logger, instanceEntity.ImageID)
	if err != nil {
		logger.Warn("CreateInstancesActor_GetByImageId sql error:", err.Error())
	}
	osEntity, err := GetByOsId(logger, imageEntity.OsID)
	if err != nil {
		logger.Warn("CreateInstancesActor GetByOsId sql error:", imageEntity.OsID, err.Error())
	}
	return osEntity, err

}
