package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters createDeviceTypeRaid
type CreateDeviceTypeRaidRequest struct {
	WriteRequestHeader

	// in:body
	Body requestTypes.CreateDeviceTypeRaidRequest
}
