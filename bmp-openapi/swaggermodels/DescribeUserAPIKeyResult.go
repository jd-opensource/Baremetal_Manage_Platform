package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A DescribeUserAPIKeyResult is an response struct that is used to describe getapikey.
// swagger:response describeUserAPIKey
type DescribeUserAPIKeyResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.ApikeyInfo `json:"result"`
	}
}
