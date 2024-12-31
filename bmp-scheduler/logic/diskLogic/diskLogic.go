package diskLogic

import (
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/diskDao"
	log "git.jd.com/cps-golang/log"
)

func QueryByManufacturerId(logger *log.Logger, manufacturer_id int64, disk_type string) ([]*diskDao.Disk, error) {
	param := map[string]interface{}{
		"manufacturer_id": manufacturer_id,
		"disk_type":       disk_type,
		"is_del":          0,
	}
	return diskDao.GetAllDisk(logger, param, nil, []string{"slot"}, []string{"asc"})
}
