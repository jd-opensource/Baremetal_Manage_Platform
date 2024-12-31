package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters deleteInstances
type DeleteInstancesRequest struct {
	ReadRequestHeader

	// in: body
	Body requestTypes.DeleteInstancesRequest
}
