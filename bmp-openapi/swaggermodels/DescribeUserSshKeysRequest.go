package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
)

// swagger:parameters describeUserSshKeys
type DescribeUserSshKeysRequest struct {
	ReadRequestHeader

	// in: query
	requestTypes.QuerySshkeysRequest
}
