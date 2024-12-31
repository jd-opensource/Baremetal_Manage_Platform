package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters dialMail
type DialMailRequest struct {
	WriteRequestHeader

	// in: body
	Body requestTypes.MailDialRequest
}
