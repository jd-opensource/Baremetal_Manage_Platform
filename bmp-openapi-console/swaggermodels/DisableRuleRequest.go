package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters disableRule
type DisableRuleRequest struct {
	WriteRequestHeader

	// in:body
	Body requestTypes.DisableRuleRequest
}
