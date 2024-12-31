package swaggermodels

// swagger:parameters resetInstanceStatus
type ResetInstanceStatusRequest struct {
	WriteRequestHeader

	// in: path
	InstanceID string `json:"instance_id"`
}
