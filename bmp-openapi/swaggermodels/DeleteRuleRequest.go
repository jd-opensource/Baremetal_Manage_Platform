package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters deleteRule
type DeleteRuleRequest struct {
	WriteRequestHeader

	// in:body
	Body requestTypes.DeleteRuleRequest
}
