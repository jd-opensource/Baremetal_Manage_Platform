package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A CreateImageResult is an response struct that is used to describe create image result.
// swagger:response createImage
type CreateImageResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.ImageId `json:"result"`
	}
}
