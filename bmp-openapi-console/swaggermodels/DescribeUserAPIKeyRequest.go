package swaggermodels

// swagger:parameters describeUserAPIKey
type DescribeUserAPIKeyRequest struct {
	ReadRequestHeader
	// required: true
	// in: path
	ApikeyID string `json:"apikey_id"`
}
