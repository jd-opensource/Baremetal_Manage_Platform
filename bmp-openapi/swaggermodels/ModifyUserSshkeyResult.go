package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A ModifyUserSshkeyResult is an response struct that is used to modify sshkey result.
// swagger:response modifyUserSshkey
type ModifyUserSshkeyResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.CommonResponse `json:"result"`
	}
}
