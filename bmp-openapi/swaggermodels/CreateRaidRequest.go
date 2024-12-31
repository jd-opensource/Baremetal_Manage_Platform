package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters createRaid
type CreateRaidRequest struct {
	WriteRequestHeader

	// in:body
	Body requestTypes.CreateRaidRequest
}
