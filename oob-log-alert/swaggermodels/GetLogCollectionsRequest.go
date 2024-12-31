package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/oob-log-alert/dao"
)

// swagger:parameters getLogCollections
type GetLogCollectionsRequest struct {
	ReadRequestHeader
	// in: query
	requestTypes.GetLogCollectionsRequest
}
