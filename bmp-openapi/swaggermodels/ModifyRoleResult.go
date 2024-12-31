package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A ModifyRoleResult is an response struct that is used to describe modify role.
// swagger:response modifyRole
type ModifyRoleResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.CommonResponse `json:"result"`
	}
}
