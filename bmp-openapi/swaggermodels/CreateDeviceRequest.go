package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters createDevice
type CreateDeviceRequest struct {
	WriteRequestHeader

	// in:body
	Body requestTypes.CreateDevicesRequest
}
