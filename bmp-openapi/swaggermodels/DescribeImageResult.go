package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A DescribeImageResult is an response struct that is used to describe image.
// swagger:response describeImage
type DescribeImageResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.ImageInfo `json:"result"`
	}
}
