package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A ModifyLocalUserResult is an response struct that is used to describe getuser.
// swagger:response modifyLocalUser
type ModifyLocalUserResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.CommonResponse `json:"result"`
	}
}
