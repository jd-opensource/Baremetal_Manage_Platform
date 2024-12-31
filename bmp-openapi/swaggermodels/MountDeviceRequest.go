package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters mountDevice
type MountDeviceRequest struct {
	ReadRequestHeader

	// in: body
	Body requestTypes.MountDevicesRequest
}
