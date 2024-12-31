package request

import log "coding.jd.com/aidc-bmp/bmp_log"

type DesrcibeAgentStatusRequest struct {
	// InstanceId,多个用逗号分隔
	// required: true
	InstanceID string `json:"instanceId" validate:"required"`
}

func (req *DesrcibeAgentStatusRequest) Validate(logger *log.Logger) {
	validate(req, logger)
}

type DesrcibeTagsRequest struct {
	// instanceId
	// required: true
	InstanceID string `json:"instanceId" validate:"required"`
	// tagName [disk mountpoint nic]
	// required: true
	TagName string `json:"tagName" validate:"required,oneof=disk mountpoint nic"`
}

func (req *DesrcibeTagsRequest) Validate(logger *log.Logger) {
	validate(req, logger)
}
