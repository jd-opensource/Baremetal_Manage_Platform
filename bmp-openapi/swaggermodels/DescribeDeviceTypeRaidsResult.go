package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// swagger:response describeDeviceTypeRaids
type DescribeDeviceTypeRaidsResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result []*response.RDeviceTypeRaid `json:"result"`
	}
}
