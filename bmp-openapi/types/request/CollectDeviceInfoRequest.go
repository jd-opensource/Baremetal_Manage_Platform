package request

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/constant"

	log "coding.jd.com/aidc-bmp/bmp_log"
	validator "github.com/go-playground/validator/v10"
)

type CollectDeviceInfoRequest struct {
	// 待采集项
	// required: true
	Collects []CollectItem `json:"collects" validate:"required"`
}

type CollectItem struct {
	// sn
	// required: true
	Sn string `json:"sn"`
	// 可选参数。 传参:True, False 。True表示传入的raid_driver值将覆盖已适配机器的raid_driver
	// Extensions:
	// x-nullable: true
	AllowOverride bool `json:"allowOverride"`
}

func (c *CollectDeviceInfoRequest) Validate(logger *log.Logger) {
	if err := validator.New().Struct(c); err != nil {
		logger.Warn("CollectDeviceInfoRequest Validate error:", err.Error())
		panic(constant.INVALID_ARGUMENT)
	}

}
