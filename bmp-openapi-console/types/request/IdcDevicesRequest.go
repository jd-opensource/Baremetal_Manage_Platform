package request

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/constant"
	log "coding.jd.com/aidc-bmp/bmp_log"
	validator "github.com/go-playground/validator/v10"
)

type CreateIdcDevicesRequest struct {
	Sns        []string `json:"sns" validate:"required"`
	DeviceType string   `json:"device_type"`
}

func (c *CreateIdcDevicesRequest) Validate(logger *log.Logger) {
	if err := validator.New().Struct(c); err != nil {
		logger.Warn("CreateIdcDevicesRequest.Validate error:", err.Error())
		panic(constant.INVALID_ARGUMENT)
	}
}
