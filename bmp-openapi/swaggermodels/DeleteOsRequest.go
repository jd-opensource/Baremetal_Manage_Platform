package swaggermodels

// swagger:parameters deleteOS
type DeleteOsRequest struct {
	WriteRequestHeader

	// in: path
	OsID string `json:"os_id"`
}
