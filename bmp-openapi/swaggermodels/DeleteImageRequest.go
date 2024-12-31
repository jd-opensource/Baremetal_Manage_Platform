package swaggermodels

// swagger:parameters deleteImage
type DeleteImageRequest struct {
	WriteRequestHeader

	// in: path
	ImageID string `json:"image_id"`
}
