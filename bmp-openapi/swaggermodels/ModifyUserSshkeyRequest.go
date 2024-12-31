package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters modifyUserSshkey
type ModifyUserSshkeyRequest struct {
	WriteRequestHeader

	// in: path
	SshkeyID string `json:"sshkey_id"`

	// in: body
	Body requestTypes.ModifySshkeyRequest
}
