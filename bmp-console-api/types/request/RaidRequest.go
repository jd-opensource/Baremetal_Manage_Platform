package request

import log "coding.jd.com/aidc-bmp/bmp_log"

type QueryVolumeRaidsRequest struct {
	DeviceTypeID string `json:"deviceTypeId" validate:"required"`
	VolumeType   string `json:"volumeType" validate:"required"`
}

func (req *QueryVolumeRaidsRequest) Validate(logger *log.Logger) error {

	return validate(req, logger)
}
