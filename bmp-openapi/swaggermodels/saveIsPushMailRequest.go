package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters saveIsPushMail
type SaveIsPushMailRequest struct {
	WriteRequestHeader

	// in: body
	Body requestTypes.SaveIsPushMailRequest
}
