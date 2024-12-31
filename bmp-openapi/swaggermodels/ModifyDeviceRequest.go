package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters modifyDevice
type ModifyDeviceRequest struct {
	ReadRequestHeader

	// in: path
	DeviceID string `json:"device_id"`

	// in: body
	Body requestTypes.ModifyDevicesRequest
}
