package swaggermodels

// swagger:response getMessageTypes
type GetMessageTypesResult struct {
	// in: body
	Body struct {
		Result    map[string][]string `json:"result"`
		RequestId string              `json:"requestId"`
	}
}
