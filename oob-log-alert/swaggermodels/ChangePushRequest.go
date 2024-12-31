package swaggermodels

import (
	"coding.jd.com/aidc-bmp/oob-log-alert/dao"
)

// swagger:parameters changePush
type ChangePushRequest struct {
	WriteRequestHeader

	// in:body
	Body dao.ChangePushRequest
}
