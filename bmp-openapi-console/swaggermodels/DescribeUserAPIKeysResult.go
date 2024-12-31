package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A DescribeUserAPIKeysResult is an response struct that is used to describe getapikeylist.
// swagger:response describeUserAPIKeys
type DescribeUserAPIKeysResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.ApikeyList `json:"result"`
	}
}
