package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A DescribeOsResult is an response struct that is used to describe os.
// swagger:response describeOS
type DescribeOsResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.OsInfo `json:"result"`
	}
}
