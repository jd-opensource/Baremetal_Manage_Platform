package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/oob-log-alert/dao"
)

// swagger:parameters ruleList
type RuleListRequest struct {
	ReadRequestHeader
	// in: query
	requestTypes.RuleListRequest
}
