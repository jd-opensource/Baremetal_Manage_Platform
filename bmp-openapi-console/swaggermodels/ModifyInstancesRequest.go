package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters modifyInstances
type ModifyInstancesRequest struct {
	ReadRequestHeader

	// in: body
	Body requestTypes.ModifyInstancesRequest
}
