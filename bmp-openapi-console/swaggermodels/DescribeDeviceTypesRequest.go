package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters describeDeviceTypes
type DescribeDeviceTypesRequest struct {
	ReadRequestHeader

	// in: query
	requestTypes.QueryDeviceTypesRequest
}
