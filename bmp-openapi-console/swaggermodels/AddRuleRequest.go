package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters addRule
type AddRuleRequest struct {
	WriteRequestHeader

	// in:body
	Body requestTypes.AddRuleRequest
}
