package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters describeUserAPIKeys
type DescribeUserAPIKeysRequest struct {
	ReadRequestHeader

	// in: query
	requestTypes.QueryApikeysRequest
}
