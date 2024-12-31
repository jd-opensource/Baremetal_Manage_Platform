package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A GetRoleListResult is an response struct that is used to describe get role list.
// swagger:response describeRoles
type DescribeRolesResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.RoleList `json:"result"`
	}
}
