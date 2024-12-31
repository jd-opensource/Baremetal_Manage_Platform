package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/controllers"
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// swagger:response describeUser
type DescribeUserResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.UserInfo          `json:"result"`
		Err    *controllers.ErrorResponse `json:"error"`
	}
}
