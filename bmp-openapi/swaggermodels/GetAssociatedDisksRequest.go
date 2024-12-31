package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters getAssociatedDisks
type GetAssociatedDisksRequest struct {
	ReadRequestHeader

	// in: query
	requestTypes.GetAssociatedDisksRequest
}
