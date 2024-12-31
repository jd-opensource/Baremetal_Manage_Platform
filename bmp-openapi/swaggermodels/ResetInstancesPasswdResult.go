package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A ResetProjectInstancesPasswdResult is an response struct that is used to describe reset instance password.
// swagger:response resetProjectInstancesPasswd
type ResetProjectInstancesPasswdResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.CommonResponse `json:"result"`
	}
}
