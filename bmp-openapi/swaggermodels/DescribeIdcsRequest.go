package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters describeIdcs
type DescribIdcsRequest struct {
	ReadRequestHeader

	// in: query
	requestTypes.QueryIdcsRequest
}
