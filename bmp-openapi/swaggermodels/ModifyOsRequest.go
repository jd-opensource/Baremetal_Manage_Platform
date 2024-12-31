package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters modifyOS
type ModifyOsRequest struct {
	ReadRequestHeader

	// in: path
	OsID string `json:"os_id"`

	// in: body
	Body requestTypes.ModifyOSRequest
}
