package partitionLogic

import (
	dao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/rDeviceTypeImagePartitionDao"
	log "git.jd.com/cps-golang/log"
)

func GetByDeviceTypeAndImageId(logger *log.Logger, device_type_id, image_id string) ([]*dao.RDeviceTypeImagePartition, error) {
	param := map[string]interface{}{
		"is_del":         0,
		"device_type_id": device_type_id, // []string{device_type, "common"},
		"image_id":       image_id,
	}
	return dao.GetAllRDeviceTypeImagePartition(logger, param)
}
