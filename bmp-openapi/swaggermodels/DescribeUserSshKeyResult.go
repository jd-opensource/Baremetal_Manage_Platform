package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A DescribeUserSshKeyResult is an response struct that is used to describe sshkey.
// swagger:response describeUserSshKey
type DescribeUserSshKeyResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.SshkeyInfo `json:"result"`
	}
}
