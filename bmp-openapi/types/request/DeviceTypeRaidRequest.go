package request

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/constant"
	log "coding.jd.com/aidc-bmp/bmp_log"
	"github.com/go-playground/validator/v10"
)

type CreateDeviceTypeRaidRequest struct {
	RaidID               string `json:"raidId"`
	DeviceType           string `json:"deviceType"`
	DeviceTypeID         string `json:"DeviceTypeId"`
	VolumeType           string `json:"volumeType"`
	VolumeDetail         string `json:"volumeDetail"`
	AvailableValue       int    `json:"availableValue"`
	SystemPartitionCount int    `json:"systemPartitionCount"`
	DiskType             string `json:"diskType"`
}

func (c *CreateDeviceTypeRaidRequest) Validate(logger *log.Logger) {
	if err := validator.New().Struct(c); err != nil {
		logger.Warn("CreateDeviceTypeRaidRequest.Validate error:", err.Error())
		panic(constant.INVALID_ARGUMENT)
	}
}

type DeleteDeviceTypeRaidRequest struct {
	RaidId       string `json:"raidId"`
	DeviceType   string `json:"deviceType"`
	VolumeType   string `json:"volumeType"`
	VolumeDetail string `json:"volumeDetail"`
}

func (c *DeleteDeviceTypeRaidRequest) Validate(logger *log.Logger) {
	if err := validator.New().Struct(c); err != nil {
		logger.Warn("DeleteDeviceTypeRaidRequest.Validate error:", err.Error())
		panic(constant.INVALID_ARGUMENT)
	}
}
