package swaggermodels

import "coding.jd.com/aidc-bmp/bmp-openapi/types/response"

// A VerifyUserResult is an response struct that is used to describe verify user result.
// swagger:response verifyUser
type VerifyUserResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.CommonResponse `json:"result"`
	}
}
