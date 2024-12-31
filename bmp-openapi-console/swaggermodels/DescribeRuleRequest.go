package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters describeRule
type DescribeRuleRequest struct {
	ReadRequestHeader

	// in: query
	requestTypes.DescribeRuleRequest
}
