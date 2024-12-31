package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters describeProjectInstances
type DescribeProjectInstancesRequest struct {
	ReadRequestHeader

	// in: query
	requestTypes.QueryInstancesRequest
}
