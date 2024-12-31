package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A ModifyLocalUserPasswordResult is an response struct that is used to describe getuser.
// swagger:response modifyLocalUserPassword
type ModifyLocalUserPasswordResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.CommonResponse `json:"result"`
	}
}
