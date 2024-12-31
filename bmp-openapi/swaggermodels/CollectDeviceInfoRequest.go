package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters collectDeviceInfo
type CollectDeviceInfoRequest struct {
	WriteRequestHeader

	// in:body
	Body requestTypes.CollectDeviceInfoRequest
}
