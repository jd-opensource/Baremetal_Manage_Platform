package interfaceLogic

import (
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/interfaceDao"
	log "git.jd.com/cps-golang/log"
)

func GetBySn(logger *log.Logger, sn string) ([]*interfaceDao.Interface, error) {
	return interfaceDao.GetBySn(logger, sn)
}
