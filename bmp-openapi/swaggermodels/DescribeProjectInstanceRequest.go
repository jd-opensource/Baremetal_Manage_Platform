package swaggermodels

// swagger:parameters describeProjectInstance
type DescribeProjectInstanceRequest struct {
	ReadRequestHeader
	// 实例uuid
	// in: path
	InstanceID string `json:"instance_id"`
}
