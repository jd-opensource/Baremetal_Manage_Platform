package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters startInstances
type StartInstancesRequest struct {
	WriteRequestHeader

	// in:body
	Body requestTypes.StartInstancesRequest
}
