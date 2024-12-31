package swaggermodels

// swagger:parameters lockProjectInstance
type LockProjectInstanceRequest struct {
	WriteRequestHeader

	// in: path
	InstanceID string `json:"instance_id"`
}
