package jdBondInterfaceLogic

import (
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/jdBondInterfaceDao"
	log "git.jd.com/cps-golang/log"
)

func GetBySn(logger *log.Logger, sn string) (*jdBondInterfaceDao.JdBondInterface, error) {
	param := map[string]interface{}{
		"is_del": 0,
		"sn":     sn,
	}
	return jdBondInterfaceDao.GetOneInterface(logger, param)
}
