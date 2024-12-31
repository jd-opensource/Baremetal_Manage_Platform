package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters modifyUserApikey
type ModifyUserApikeyRequest struct {
	ReadRequestHeader

	// in: path
	ApikeyID string `json:"apikey_id"`

	// in: body
	Body requestTypes.ModifyApikeyRequest
}
