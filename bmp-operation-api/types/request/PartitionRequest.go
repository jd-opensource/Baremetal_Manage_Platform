package request

type QueryDefaultSystemPartitionRequest struct {
	ImageID      string `json:"imageId" validate:"required"`
	DeviceTypeID string `json:"deviceTypeId" validate:"required"`
}
