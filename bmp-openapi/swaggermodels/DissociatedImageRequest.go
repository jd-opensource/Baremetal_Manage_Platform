package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters dissociatedImage
type DissociatedImageRequest struct {
	WriteRequestHeader

	// in: body
	Body requestTypes.DissociatedImageRequest
}
