package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A DeleteUserApikeyResult is an response struct that is used to describe delete Apikey result.
// swagger:response deleteUserApikey
type DeleteUserApikeyResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.CommonResponse `json:"result"`
	}
}
