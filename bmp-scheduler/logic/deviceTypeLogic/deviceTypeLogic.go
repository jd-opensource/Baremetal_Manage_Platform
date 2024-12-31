package deviceTypeLogic

import (
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/deviceTypeDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/instanceLogic"
	log "git.jd.com/cps-golang/log"
)

func GetDeviceTypeByID(logger *log.Logger, deviceTypeId string) (*deviceTypeDao.DeviceType, error) {
	return deviceTypeDao.GetDeviceTypeById(logger, deviceTypeId)
}

func GetDeviceTypeByInstanceID(logger *log.Logger, instanceId string) (*deviceTypeDao.DeviceType, error) {
	instanceEntity, err := instanceLogic.GetByInstanceId(logger, instanceId)
	if err != nil {
		return nil, err
	}
	return deviceTypeDao.GetDeviceTypeById(logger, instanceEntity.DeviceTypeID)
}
