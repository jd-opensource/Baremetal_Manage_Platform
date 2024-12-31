package swaggermodels

import (
	"coding.jd.com/aidc-bmp/oob-log-alert/dao"
)

// swagger:response getLogCollections
type GetLogCollectionsResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result dao.GetLogCollectionsResponse `json:"result"`
	}
}
