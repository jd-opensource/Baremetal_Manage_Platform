package swaggermodels

import (
	"coding.jd.com/aidc-bmp/oob-log-alert/dao"
)

// swagger:response changeUse
type ChangeUseResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result dao.CommonResponse `json:"result"`
	}
}
