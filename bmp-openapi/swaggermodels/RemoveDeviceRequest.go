package swaggermodels

// swagger:parameters removeDevice
type RemoveDeviceRequest struct {
	WriteRequestHeader

	// in: path
	DeviceID string `json:"device_id"`
}
