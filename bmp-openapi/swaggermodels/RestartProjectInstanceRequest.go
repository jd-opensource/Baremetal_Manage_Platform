package swaggermodels

// swagger:parameters restartProjectInstance
type RestartProjectInstanceRequest struct {
	WriteRequestHeader

	// in: path
	InstanceID string `json:"instance_id"`
}
