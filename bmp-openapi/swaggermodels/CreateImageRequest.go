package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters createImage
type CreateImageRequest struct {
	WriteRequestHeader

	// in:body
	Body requestTypes.CreateImageRequest
}
