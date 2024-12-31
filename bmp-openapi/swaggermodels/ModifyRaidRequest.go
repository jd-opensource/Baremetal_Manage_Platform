package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters modifyRaid
type ModifyRaidRequest struct {
	WriteRequestHeader

	// in: path
	RaidID string `json:"raid_id"`

	// in: body
	Body requestTypes.ModifyRaidRequest
}
