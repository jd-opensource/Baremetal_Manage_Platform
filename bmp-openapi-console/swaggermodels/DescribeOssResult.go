package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A DescribeOssResult is an response struct that is used to describe oss.
// swagger:response describeOSs
type DescribeOssResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.OsList `json:"result"`
	}
}
