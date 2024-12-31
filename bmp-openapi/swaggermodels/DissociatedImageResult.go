package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A DissociatedImageResult is an response struct that is used to dissociated image result.
// swagger:response dissociatedImage
type DissociatedImageResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.CommonResponse `json:"result"`
	}
}
