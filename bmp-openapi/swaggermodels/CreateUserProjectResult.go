package swaggermodels

import "coding.jd.com/aidc-bmp/bmp-openapi/types/response"

// swagger:response createUserProject
type CreateUserProjectResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result response.ProjectId `json:"result"`
	}
}
