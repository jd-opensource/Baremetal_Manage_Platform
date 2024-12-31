package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters describeInstancesByProjectIdAndOwnerNameAndSharerName
type DescribeInstancesByProjectIdAndOwnerNameAndSharerNameRequest struct {
	ReadRequestHeader

	// in: query
	requestTypes.DescribeInstancesByProjectIdAndOwnerNameAndSharerNameRequest
}
