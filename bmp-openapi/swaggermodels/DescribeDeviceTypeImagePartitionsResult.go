package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// swagger:response describeDeviceTypeImagePartitions
type DescribeDeviceTypeImagePartitionsResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.QueryDeviceTypeImagePartitionResponse `json:"result"`
	}
}
