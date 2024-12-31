package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters createUserSshkey
type CreateUserSshkeyRequest struct {
	WriteRequestHeader

	// in:body
	Body requestTypes.CreateSshkeyRequest
}
