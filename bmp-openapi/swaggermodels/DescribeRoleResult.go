package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A GetRoleResult is an response struct that is used to describe role.
// swagger:response describeRole
type DescribeRoleResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.RoleInfo `json:"result"`
	}
}
