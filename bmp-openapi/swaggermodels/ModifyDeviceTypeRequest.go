package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters modifyDeviceType
type ModifyDeviceTypeRequest struct {
	ReadRequestHeader

	// in: path
	DeviceTypeID string `json:"device_type_id"`

	// in: body
	Body requestTypes.ModifyDeviceTypeRequest
}
