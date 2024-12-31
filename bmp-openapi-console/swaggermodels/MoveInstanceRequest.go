package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters moveUserInstances
type MoveInstanceRequest struct {
	WriteRequestHeader
	// in: body
	Body requestTypes.MoveInstancesRequest
}
