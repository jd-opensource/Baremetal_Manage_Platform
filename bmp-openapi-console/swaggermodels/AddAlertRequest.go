package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters addAlert
type AddAlertRequest struct {
	WriteRequestHeader

	// in:body
	Body requestTypes.AddAlertRequest
}
