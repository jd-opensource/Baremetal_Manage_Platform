package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters modifyLocalUser
type ModifyLocalUserRequest struct {
	ReadRequestHeader
	// in: body
	Body requestTypes.ModifyLocalUserRequest
}
