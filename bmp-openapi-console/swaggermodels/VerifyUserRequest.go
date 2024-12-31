package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters verifyUser
type VerifyUserRequest struct {
	WriteRequestHeader

	// in:body
	Body requestTypes.VerifyUserRequest
}
