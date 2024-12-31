package swaggermodels

// swagger:response alertLevelList
type AlertLevelListResult struct {
	ResponseHeader
	// in: body
	Body struct {
		Result []string `json:"result"`
	}
}
