package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A ResetProjectInstancePasswdResult is an response struct that is used to describe reset instance password.
// swagger:response resetProjectInstancePasswd
type ResetProjectInstancePasswdResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.CommonResponse `json:"result"`
	}
}
