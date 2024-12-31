package swaggermodels

// swagger:parameters startProjectInstance
type StartProjectInstanceRequest struct {
	WriteRequestHeader

	// in: path
	InstanceID string `json:"instance_id"`
}
