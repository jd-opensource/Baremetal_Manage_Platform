package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters createOS
type CreateOsRequest struct {
	WriteRequestHeader

	// in:body
	Body requestTypes.CreateOSRequest
}
