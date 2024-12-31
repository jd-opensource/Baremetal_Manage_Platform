package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters createDeviceType
type CreateDeviceTypeRequest struct {
	WriteRequestHeader

	// in:body
	Body requestTypes.CreateDeviceTypeRequest
}
