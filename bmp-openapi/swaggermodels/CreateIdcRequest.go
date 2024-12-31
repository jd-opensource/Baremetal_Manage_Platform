package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters createIdc
type CreateIdcRequest struct {
	WriteRequestHeader

	// in:body
	Body requestTypes.CreateIdcRequest
}
