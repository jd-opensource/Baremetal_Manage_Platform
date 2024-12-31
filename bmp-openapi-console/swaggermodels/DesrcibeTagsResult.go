package swaggermodels

import (
	responseTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/response"
)

// swagger:response desrcibeTags
type DesrcibeTagsResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result responseTypes.TagsResponse `json:"result"`
	}
}
