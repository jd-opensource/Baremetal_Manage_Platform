package swaggermodels

// swagger:parameters describeIdc
type DescribeIdcRequest struct {
	ReadRequestHeader

	// in: path
	IdcID string `json:"idc_id"`
}
