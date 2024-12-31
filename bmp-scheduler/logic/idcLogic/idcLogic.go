package idcLogic

import (
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/idcDao"
	log "git.jd.com/cps-golang/log"
)

func GetIdcID(logger *log.Logger) (*idcDao.Idc, error) {

	query := map[string]interface{}{
		"is_del": 0,
	}
	idc, err := idcDao.GetOneIdc(logger, query)
	if err != nil {
		return nil, err
	}
	return idc, nil
}
