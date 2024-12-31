package swaggermodels

// swagger:parameters deleteProjectInstance
type DeleteProjectInstanceRequest struct {
	WriteRequestHeader

	// in: path
	InstanceID string `json:"instance_id"`
}
