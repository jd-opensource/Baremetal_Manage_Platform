package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A DescribeImagesResult is an response struct that is used to describe images.
// swagger:response describeImages
type DescribeImagesResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.ImageList `json:"result"`
	}
}
