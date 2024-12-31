package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters associatedImage
type AssociatedImageRequest struct {
	WriteRequestHeader

	// in: body
	Body requestTypes.AssociateImageRequest
}
