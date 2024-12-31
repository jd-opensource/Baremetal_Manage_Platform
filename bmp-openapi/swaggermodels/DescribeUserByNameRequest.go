package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters describeUserByName
type DescribeUserByNameRequest struct {
	ReadRequestHeader

	// in: query
	requestTypes.GetUserByNameRequest
}
