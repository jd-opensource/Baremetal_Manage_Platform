package service

import (
	log "coding.jd.com/aidc-bmp/bmp_log"
	"coding.jd.com/aidc-bmp/oob-log-alert/dao"
)

func GetUserIdByName(logger *log.Logger, username string) (string, error) {

	user, err := dao.GetUserByName(logger, username)
	if err != nil {
		logger.Warnf("GetUserIdByName error, username:%s, err:%s", username, err.Error())
		return "", err
	}
	logger.Infof("GetUserIdByName succ, username:%s, userid:%s", username, user.UserID)
	return user.UserID, nil

}
