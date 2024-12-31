package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters dissociatedDeviceType
type DissociatedDeviceTypeRequest struct {
	WriteRequestHeader

	// in: body
	Body requestTypes.DissociatedImageRequest
}
