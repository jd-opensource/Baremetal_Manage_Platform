package swaggermodels

// swagger:parameters describeDeviceType
type DescribeDeviceTypeRequest struct {
	ReadRequestHeader

	// in: path
	DeviceTypeID string `json:"device_type_id"`
}
