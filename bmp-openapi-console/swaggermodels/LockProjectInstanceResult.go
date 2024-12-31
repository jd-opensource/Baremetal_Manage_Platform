package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A LockProjectInstanceResult is an response struct that is used to describe lock instance result.
// swagger:response lockProjectInstance
type LockProjectInstanceResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.CommonResponse `json:"result"`
	}
}
