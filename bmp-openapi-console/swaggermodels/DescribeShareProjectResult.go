package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A DescribeShareProjectResult is an response struct that is used to describe getapikey.
// swagger:response describeShareProject
type DescribeShareProjectResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.ShareProjectInfo `json:"result"`
	}
}
