package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters reinstallProjectInstance
type ReinstallProjectInstanceRequest struct {
	ReadRequestHeader

	// in: path
	InstanceID string `json:"instance_id"`
	// in: body
	Body requestTypes.ReinstallInstanceRequest
}
