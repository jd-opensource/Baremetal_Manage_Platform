package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A StopProjectInstanceResult is an response struct that is used to describe stop instance result.
// swagger:response stopProjectInstance
type StopProjectInstanceResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.CommonResponse `json:"result"`
	}
}
