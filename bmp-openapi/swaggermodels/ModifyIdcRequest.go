package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters modifyIdc
type ModifyIdcRequest struct {
	WriteRequestHeader

	// in: path
	IdcID string `json:"idc_id"`

	// in: body
	Body requestTypes.ModifyIdcRequest
}
