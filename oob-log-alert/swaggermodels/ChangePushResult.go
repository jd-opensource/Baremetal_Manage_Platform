package swaggermodels

import (
	"coding.jd.com/aidc-bmp/oob-log-alert/dao"
)

// swagger:response changePush
type ChangePushResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result dao.CommonResponse `json:"result"`
	}
}
