package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters stopInstances
type StopInstancesRequest struct {
	WriteRequestHeader

	// in:body
	Body requestTypes.StopInstancesRequest
}
