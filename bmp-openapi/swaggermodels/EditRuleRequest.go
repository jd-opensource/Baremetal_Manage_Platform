package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters editRule
type EditRuleRequest struct {
	WriteRequestHeader

	// in:body
	Body requestTypes.EditRuleRequest
}
