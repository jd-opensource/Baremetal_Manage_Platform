package request

import (
	log "coding.jd.com/aidc-bmp/bmp_log"
)

type QueryInstancesSshkeyRequest struct {
	// 实例Id
	InstanceId string `json:"instanceId"`
	// 秘钥Id
	SshkeyId string `json:"sshkeyId"`
	// 是否显示全部
	IsAll string `json:"isAll"`
}

func (req *QueryInstancesSshkeyRequest) Validate(logger *log.Logger) {
	validate(req, logger)
}
