package raidLogic

import (
	dao "coding.jd.com/aidc-bmp/bmp-scheduler/dao/raidDao"
	log "git.jd.com/cps-golang/log"
)

func GetByRaidId(logger *log.Logger, raid_id string) (*dao.Raid, error) {

	return dao.GetRaidByUuid(logger, raid_id)
}
