package swaggermodels

import (
	"coding.jd.com/aidc-bmp/oob-log-alert/dao"
)

// swagger:parameters dealLogCollection
type DealLogCollectionRequest struct {
	WriteRequestHeader

	// in:body
	Body dao.DealLogCollectionRequest
}