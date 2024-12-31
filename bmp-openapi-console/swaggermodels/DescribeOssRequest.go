package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters describeOSs
type DescribeOssRequest struct {
	ReadRequestHeader

	// in: query
	requestTypes.QueryOssRequest
}
