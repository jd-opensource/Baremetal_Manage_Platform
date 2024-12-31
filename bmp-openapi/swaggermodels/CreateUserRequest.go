package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters createUser
type CreateUserRequest struct {
	WriteRequestHeader

	// in:body
	Body requestTypes.CreateUserRequest
}
