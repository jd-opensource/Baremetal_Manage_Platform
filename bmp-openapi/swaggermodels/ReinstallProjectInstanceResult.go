package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A ReinstallProjectInstanceResult is an response struct that is used to describe reinstall instance.
// swagger:response reinstallProjectInstance
type ReinstallProjectInstanceResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.CommonResponse `json:"result"`
	}
}
