package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters modifyUser
type ModifyUserRequest struct {
	ReadRequestHeader

	// in: path
	UserID string `json:"user_id"`
	// in: body
	Body requestTypes.ModifyUserRequest
}
