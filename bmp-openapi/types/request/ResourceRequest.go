package request

import (
	log "coding.jd.com/aidc-bmp/bmp_log"
)

type QueryResourcesRequest struct {
	/*// 机房名称
	IdcName string `json:"idcName"`
	// 机房英文名称
	IdcNameEn string `json:"idcNameEn"`*/
	// 机型名称
	Name string `json:"name"`
	// 机型规格
	DeviceType string `json:"deviceType"`
	// 镜像名称
	ImageName string `json:"imageName"`
	// 用户名称
	UserName string `json:"userName"`
}

func (req *QueryResourcesRequest) Validate(logger *log.Logger) {
	validate(req, logger)
}
