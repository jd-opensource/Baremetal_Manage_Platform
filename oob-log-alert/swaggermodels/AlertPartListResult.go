package swaggermodels

// swagger:response alertPartList
type AlertPartListResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result []string `json:"result"`
	}
}
