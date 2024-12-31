package swaggermodels

import "coding.jd.com/aidc-bmp/bmp-openapi/types/response"

// A CreateInstancesResult is an response struct that is used to describe create instance result.
// swagger:response createUser
type CreateUserResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.UserId `json:"result"`
	}
}
