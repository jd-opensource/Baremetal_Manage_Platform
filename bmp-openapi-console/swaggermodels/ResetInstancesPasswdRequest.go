package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters resetProjectInstancesPasswd
type ResetProjectInstancesPasswdRequest struct {
	WriteRequestHeader

	// in: body
	Body requestTypes.ResetInstancesPasswdRequest
}
