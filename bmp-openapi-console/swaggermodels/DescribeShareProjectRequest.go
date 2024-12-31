package swaggermodels

// swagger:parameters describeShareProject
type DescribeShareProjectRequest struct {
	ReadRequestHeader
	// in: path
	ProjectID string `json:"project_id"`
}
