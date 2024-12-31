package swaggermodels

// swagger:parameters unLockProjectInstance
type UnLockProjectInstanceResultRequest struct {
	WriteRequestHeader

	// in: path
	InstanceID string `json:"instance_id"`
}
