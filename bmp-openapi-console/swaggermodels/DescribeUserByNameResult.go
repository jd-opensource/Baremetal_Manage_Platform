package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A GetUserByNameResult is an response struct that is used to describe user.
// swagger:response describeUserByName
type DescribeUserByNameResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.UserInfo `json:"result"`
	}
}
