package swaggermodels

// swagger:parameters describeDevice
type DescribeDeviceRequest struct {
	ReadRequestHeader

	// in: path
	DeviceID string `json:"device_id"`
}
