package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters desrcibeTags
type DesrcibeTagsRequest struct {
	ReadRequestHeader

	// in: query
	requestTypes.DesrcibeTagsRequest
}
