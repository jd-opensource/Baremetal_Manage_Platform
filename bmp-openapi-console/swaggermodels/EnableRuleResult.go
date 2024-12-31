package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A EnableRuleResult is an response struct that is used to describe create rule result.
// swagger:response enableRule
type EnableRuleResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.CommonResponse `json:"result"`
	}
}
