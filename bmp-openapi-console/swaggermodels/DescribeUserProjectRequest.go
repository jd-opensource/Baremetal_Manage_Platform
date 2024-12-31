package swaggermodels

// swagger:parameters describeUserProject
type DescribeUserProjectRequest struct {
	ReadRequestHeader

	// in: path
	ProjectID string `json:"project_id"`
}
