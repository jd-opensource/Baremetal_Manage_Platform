package swaggermodels

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// A AssociatedImageResult is an response struct that is used to associated image result.
// swagger:response associatedImage
type AssociatedImageResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.CommonResponse `json:"result"`
	}
}
