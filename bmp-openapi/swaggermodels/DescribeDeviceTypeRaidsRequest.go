package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters describeDeviceTypeRaids
type DescribeDeviceTypeRaidsRequest struct {
	ReadRequestHeader

	// in: query
	requestTypes.QueryDeviceTypeRaidRequest
}
