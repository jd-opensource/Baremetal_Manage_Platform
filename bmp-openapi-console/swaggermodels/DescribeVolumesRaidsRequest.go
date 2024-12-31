package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters describeVolumesRaids
type DescribeVolumesRaidsRequest struct {
	ReadRequestHeader

	// in: query
	requestTypes.QueryVolumesRaidsRequest
}
