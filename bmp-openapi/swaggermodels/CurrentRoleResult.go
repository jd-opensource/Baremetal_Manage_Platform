package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A CurrentRoleResult is an response struct that is used to describe role.
// swagger:response currentRole
type CurrentRoleResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.RoleInfo `json:"result"`
	}
}
