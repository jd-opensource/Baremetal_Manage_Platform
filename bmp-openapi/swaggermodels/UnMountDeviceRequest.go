package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters unMountDevice
type UnMountDeviceRequest struct {
	ReadRequestHeader

	// in: body
	Body requestTypes.UnMountDevicesRequest
}
