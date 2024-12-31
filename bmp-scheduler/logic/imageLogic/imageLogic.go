package imageLogic

import (
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/imageDao"
	log "git.jd.com/cps-golang/log"
)

func GetByImageId(logger *log.Logger, image_id string) (*imageDao.Image, error) {
	return imageDao.GetImageByUuid(logger, image_id)
}
