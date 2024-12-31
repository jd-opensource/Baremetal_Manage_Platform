package swaggermodels

import (
	"coding.jd.com/aidc-bmp/oob-log-alert/dao"
)

// swagger:parameters ruleResetPushAndUse
type RuleResetPushAndUseRequest struct {
	WriteRequestHeader

	// in:body
	Body dao.RuleResetPushAndUseRequest
}
