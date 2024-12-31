package request

type QueryRaidsRequest struct {
	DeviceTypeID string `json:"deviceTypeId" validate:"required"`
	VolumeType   string `json:"volumeType" validate:"required"`
}
