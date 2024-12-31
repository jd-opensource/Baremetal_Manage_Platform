package swaggermodels

// swagger:parameters deleteDevice
type DeleteDeviceRequest struct {
	WriteRequestHeader

	// in: path
	DeviceID string `json:"device_id"`
}
