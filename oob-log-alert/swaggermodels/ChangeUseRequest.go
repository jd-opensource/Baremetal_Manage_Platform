package swaggermodels

import (
	"coding.jd.com/aidc-bmp/oob-log-alert/dao"
)

// swagger:parameters changeUse
type ChangeUseRequest struct {
	WriteRequestHeader

	// in:body
	Body dao.ChangeUseRequest
}
