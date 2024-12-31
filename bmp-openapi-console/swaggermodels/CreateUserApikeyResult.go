package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A CreateUserApikeyResult is an response struct that is used to describe create apikey result.
// swagger:response createUserApikey
type CreateUserApikeyResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.ApikeyId `json:"result"`
	}
}
