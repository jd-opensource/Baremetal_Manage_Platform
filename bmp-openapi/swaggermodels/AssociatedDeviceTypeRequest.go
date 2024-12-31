package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters associatedDeviceType
type AssociatedDeviceTypeRequest struct {
	WriteRequestHeader

	// in: body
	Body requestTypes.AssociateDeviceTypeRequest
}
