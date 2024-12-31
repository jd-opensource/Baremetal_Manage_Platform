package swaggermodels

import "coding.jd.com/aidc-bmp/bmp-openapi/types/response"

// A CreateRoleResult is an response struct that is used to describe create role result.
// swagger:response createRole
type CreateRoleResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.RoleId `json:"result"`
	}
}
