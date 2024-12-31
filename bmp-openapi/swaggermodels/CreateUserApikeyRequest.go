package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters createUserApikey
type CreateUserApikeyRequest struct {
	WriteRequestHeader

	// in:body
	Body requestTypes.CreateApikeyRequest
}
