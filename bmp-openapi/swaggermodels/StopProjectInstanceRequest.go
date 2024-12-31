package swaggermodels

// swagger:parameters stopProjectInstance
type StopProjectInstanceRequest struct {
	WriteRequestHeader

	// in: path
	InstanceID string `json:"instance_id"`
}
