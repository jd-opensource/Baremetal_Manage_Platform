package swaggermodels

import (
	"coding.jd.com/aidc-bmp/oob-log-alert/dao"
)

// swagger:response getLogCollectionsBySn
type GetLogCollectionsBySnResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result dao.GetLogCollectionsResponse `json:"result"`
	}
}
