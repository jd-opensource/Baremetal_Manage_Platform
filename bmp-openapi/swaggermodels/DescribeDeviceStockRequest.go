package swaggermodels

// swagger:parameters describeDeviceStock
type DescribeDeviceStockRequest struct {
	ReadRequestHeader

	// in: query
	DeviceTypeID string `json:"deviceTypeId"`
}
