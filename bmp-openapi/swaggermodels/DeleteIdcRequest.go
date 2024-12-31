package swaggermodels

// swagger:parameters deleteIdc
type DeleteIdcRequest struct {
	WriteRequestHeader

	// in: path
	IdcID string `json:"idc_id"`
}
