package swaggermodels

import (
	"coding.jd.com/aidc-bmp/oob-log-alert/dao"
)

// swagger:response ruleList
type RuleListResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result dao.RuleListResponse `json:"result"`
	}
}
