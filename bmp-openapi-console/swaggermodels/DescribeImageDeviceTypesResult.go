package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// swagger:response describeImageDeviceTypes
type DescribeImageDeviceTypesResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.DeviceTypeList `json:"result"`
	}
}
