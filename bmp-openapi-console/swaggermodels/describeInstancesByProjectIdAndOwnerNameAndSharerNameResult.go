package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A DescribeInstancesByProjectIdAndOwnerNameAndSharerNameResult is an response struct that is used to get instance list.
// swagger:response describeInstancesByProjectIdAndOwnerNameAndSharerName
type DescribeInstancesByProjectIdAndOwnerNameAndSharerNameResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.InstanceForShareList `json:"result"`
	}
}
