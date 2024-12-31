package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A ModifyUserApikeyResult is an response struct that is used to describe modifyApikey result.
// swagger:response modifyUserApikey
type ModifyUserApikeyResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.CommonResponse `json:"result"`
	}
}
