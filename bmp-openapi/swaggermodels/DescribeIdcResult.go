package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A DescribeIdcResult is an response struct that is used to describe idc.
// swagger:response describeIdc
type DescribeIdcResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.IdcInfo `json:"result"`
	}
}
