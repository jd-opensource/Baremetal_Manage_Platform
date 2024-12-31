package swaggermodels

import "coding.jd.com/aidc-bmp/bmp-openapi/types/response"

// A CreateUserSshkeyResult is an response struct that is used to describe create sshkey result.
// swagger:response createUserSshkey
type CreateUserSshkeyResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.SshkeyId `json:"result"`
	}
}
