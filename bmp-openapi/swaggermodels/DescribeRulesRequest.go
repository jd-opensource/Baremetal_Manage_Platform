package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters describeRules
type DescribeRulesRequest struct {
	ReadRequestHeader

	// in: query
	requestTypes.DescribeRulesRequest
}
