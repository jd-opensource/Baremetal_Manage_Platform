package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters modifyLocalUserPassword
type ModifyLocalUserPasswordRequest struct {
	ReadRequestHeader

	// in: body
	Body requestTypes.ModifyUserPasswordRequest
}
