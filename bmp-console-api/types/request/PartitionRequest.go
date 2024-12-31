package request

import log "coding.jd.com/aidc-bmp/bmp_log"

type QueryDefaultSystemPartitionRequest struct {
	ImageID      string `json:"imageId" validate:"required"`
	DeviceTypeID string `json:"deviceTypeId" validate:"required"`
}

func (req *QueryDefaultSystemPartitionRequest) Validate(logger *log.Logger) error {

	return validate(req, logger)
}
