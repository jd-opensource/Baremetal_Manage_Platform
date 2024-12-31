package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// swagger:response describeDeviceTypeImages
type DescribeDeviceTypeImagesResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.ImageList `json:"result"`
	}
}
