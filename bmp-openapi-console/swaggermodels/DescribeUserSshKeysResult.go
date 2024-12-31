package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A DescribeUserSshKeysResult is an response struct that is used to describe get sshkey list.
// swagger:response describeUserSshKeys
type DescribeUserSshKeysResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.SshkeyList `json:"result"`
	}
}
