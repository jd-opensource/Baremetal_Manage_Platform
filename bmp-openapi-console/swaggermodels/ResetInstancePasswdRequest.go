package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters resetProjectInstancePasswd
type ResetInstancePasswdRequest struct {
	ReadRequestHeader

	// in: path
	InstanceID string `json:"instance_id"`
	// in: body
	Body requestTypes.ResetInstancePasswdRequest
}
