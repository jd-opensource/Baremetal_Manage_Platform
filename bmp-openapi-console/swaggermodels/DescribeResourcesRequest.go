package swaggermodels

import requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"

// swagger:parameters describeResources
type DescribeResourcesRequest struct {
	ReadRequestHeader

	// in: path
	requestTypes.QueryResourcesRequest
}
