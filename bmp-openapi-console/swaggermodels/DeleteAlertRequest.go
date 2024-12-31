package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters deleteAlert
type DeleteAlertRequest struct {
	WriteRequestHeader

	// in:body
	Body requestTypes.DeleteAlertRequest
}
