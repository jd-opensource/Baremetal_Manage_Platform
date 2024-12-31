package swaggermodels

import requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"

// swagger:parameters deleteDeviceTypeRaid
type DeleteDeviceTypeRaid struct {
	WriteRequestHeader

	// in:body
	Body requestTypes.DeleteDeviceTypeRaidRequest
}
