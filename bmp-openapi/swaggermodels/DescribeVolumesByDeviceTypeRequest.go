package swaggermodels

// swagger:parameters describeVolumesByDeviceType
type DescribeVolumesByDeviceType struct {
	ReadRequestHeader

	// in: path
	DeviceTypeID string `json:"device_type_id"`
}
