package swaggermodels

// swagger:parameters deleteUserProject
type DeleteUserProjectRequest struct {
	WriteRequestHeader

	// in: path
	ProjectID string `json:"project_id"`
}
