package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A DeleteUserSshkeyResult is an response struct that is used to describe delete sshkey result.
// swagger:response deleteUserSshkey
type DeleteUserSshkeyResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.CommonResponse `json:"result"`
	}
}
