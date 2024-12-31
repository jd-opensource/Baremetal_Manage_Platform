package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A UnLockProjectInstanceResult is an response struct that is used to describe unlock instance result.
// swagger:response unLockProjectInstance
type UnLockProjectInstanceResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.CommonResponse `json:"result"`
	}
}
