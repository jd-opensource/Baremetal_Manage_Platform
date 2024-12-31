package swaggermodels

import "coding.jd.com/aidc-bmp/bmp-openapi/types/response"

// A CreateOsResult is an response struct that is used to describe create os result.
// swagger:response createOS
type CreateOsResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.OsId `json:"result"`
	}
}
