package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters enableRule
type EnableRuleRequest struct {
	WriteRequestHeader

	// in:body
	Body requestTypes.EnableRuleRequest
}
