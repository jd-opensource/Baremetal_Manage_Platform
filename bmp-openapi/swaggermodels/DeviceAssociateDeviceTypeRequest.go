package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters deviceAssociateDeviceType
type DeviceAssociateDeviceTypeRequest struct {
	WriteRequestHeader

	// in: body
	Body requestTypes.DeviceAssociateDeviceTypeRequest
}
