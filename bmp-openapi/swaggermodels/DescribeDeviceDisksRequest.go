package swaggermodels

// swagger:parameters describeDeviceDisks
type DescribeDeviceDisksRequest struct {
	ReadRequestHeader

	// in: path
	DeviceID string `json:"device_id"`
}
