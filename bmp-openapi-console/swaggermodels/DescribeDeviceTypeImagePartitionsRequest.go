package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters describeDeviceTypeImagePartitions
type DescribeDeviceTypeImagePartitionsRequest struct {
	ReadRequestHeader

	// in: query
	requestTypes.QueryDeviceTypeImagePartitionRequest
}
