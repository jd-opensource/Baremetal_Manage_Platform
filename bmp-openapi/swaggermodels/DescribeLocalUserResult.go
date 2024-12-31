package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/controllers"
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A DescribeLocalUserResult is an response struct that is used to describe getuser.
// swagger:response describeLocalUser
type DescribeLocalUserResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.UserInfo          `json:"result"`
		Err    *controllers.ErrorResponse `json:"error"`
	}
}
