package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A DescribeRuleResult is an response struct that is used to describe get rule.
// swagger:response describeRules
type DescribeRulesResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.RuleList `json:"result"`
	}
}