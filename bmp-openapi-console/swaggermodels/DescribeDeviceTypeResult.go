package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// swagger:response describeDeviceType
type DescribeDeviceTypeResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.DeviceTypeInfo `json:"result"`
	}
}
