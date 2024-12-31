package swaggermodels

// swagger:parameters describeOS
type DescribeOsRequest struct {
	ReadRequestHeader

	// in: path
	OsID string `json:"os_id"`
}
