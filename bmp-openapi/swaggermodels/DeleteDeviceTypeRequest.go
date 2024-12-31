package swaggermodels

// swagger:parameters deleteDeviceType
type DeleteDeviceTypeRequest struct {
	WriteRequestHeader

	// in: path
	DeviceTypeID string `json:"device_type_id"`
}
