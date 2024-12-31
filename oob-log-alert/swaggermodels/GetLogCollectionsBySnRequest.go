package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/oob-log-alert/dao"
)

// swagger:parameters getLogCollectionsBySn
type GetLogCollectionsBySnRequest struct {
	ReadRequestHeader
	// in: query
	requestTypes.GetLogCollectionsBySnRequest
}
