package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters associateDeviceDisks
type AssociateDeviceDisksRequest struct {
	WriteRequestHeader

	// in: body
	Body requestTypes.AssociateDeviceDisksRequest
}
