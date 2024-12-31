package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters modifyImage
type ModifyImageRequest struct {
	ReadRequestHeader

	// in: path
	ImageID string `json:"image_id"`

	// in: body
	Body requestTypes.ModifyImageRequest
}
