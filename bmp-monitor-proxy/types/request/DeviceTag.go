package request

import (
	log "git.jd.com/cps-golang/log"
)

type GetDeviceTagsRequest struct {
	InstanceID string `json:"instanceId" validate:"required"`
	Metric     string `json:"metric" validate:"required"`
}

func (req *GetDeviceTagsRequest) Validate(logger *log.Logger) {

	validate(req, logger)

}
