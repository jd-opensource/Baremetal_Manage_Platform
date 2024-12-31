package swaggermodels

// swagger:parameters describeImage
type DescribeImageRequest struct {
	ReadRequestHeader

	// in: path
	ImageID string `json:"image_id"`
}
