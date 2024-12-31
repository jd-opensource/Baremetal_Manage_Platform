package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A DescribeProjectInstancesResult is an response struct that is used to get instance list.
// swagger:response describeProjectInstances
type DescribeProjectInstancesResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.InstanceList `json:"result"`
	}
}
