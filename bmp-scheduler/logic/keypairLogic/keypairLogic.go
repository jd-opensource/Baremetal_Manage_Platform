package keypairLogic

import (
	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/keypairDao"
	log "git.jd.com/cps-golang/log"
)

func GetByUuidAndTenant(logger *log.Logger, keypair_id, tenant string) (*keypairDao.Keypair, error) {
	param := map[string]interface{}{
		"is_del": 0,
		"uuid":   keypair_id,
		"tenant": tenant,
	}
	return keypairDao.GetOneKeypair(logger, param)
}
