package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters restartInstances
type RestartInstancesRequest struct {
	WriteRequestHeader

	// in:body
	Body requestTypes.RestartInstancesRequest
}
