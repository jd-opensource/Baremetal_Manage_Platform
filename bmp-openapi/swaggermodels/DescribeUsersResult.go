package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// swagger:response describeUsers
type DescribeUsersResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.UserList `json:"result"`
	}
}
