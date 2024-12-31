package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-operation-api/types/request"
	responseTypes "coding.jd.com/aidc-bmp/bmp-operation-api/types/response"
)

// swagger:parameters uploadImage
type UploadImageRequest struct {
	WriteHeader
	// in:body
	Body requestTypes.UploadImageRequest
}

// swagger:response uploadImage
type UploadImageResponse struct {
	// in: body
	Body struct {
		Result responseTypes.ImageUpload `json:"result"`
		CommonRespBody
	}
}

// swagger:parameters createImage
type CreateImageRequest struct {
	WriteHeader
	// in:body
	Body requestTypes.CreateImageRequest
}

// CreateImageResponse is an response struct
// swagger:response createImage
type CreateImageResponse struct {
	// in: body
	Body struct {
		Result responseTypes.ImageId `json:"result"`
		CommonRespBody
	}
}

// swagger:parameters describeImages
type DescribeImagesRequest struct {
	ReadHeader
	// in: query
	requestTypes.QueryImagesRequest
}

// DescribeImagesResponse is an response struct
// swagger:response describeImages
type DescribeImagesResponse struct {
	// in: body
	Body struct {
		Result responseTypes.ImageList `json:"result"`
		CommonRespBody
	}
}

// swagger:parameters deleteImage
type DeleteImageRequest struct {
	WriteHeader

	// in: path
	ImageID string `json:"image_id"`
}

// DeleteImageResponse is an response struct
// swagger:response deleteImage
type DeleteImageResponse struct {
	// in: body
	Body struct {
		Result responseTypes.CommonResponse `json:"result"`
		CommonRespBody
	}
}

// swagger:parameters describeImageDeviceTypes
type DescribeImageDeviceTypesRequest struct {
	ReadHeader
	// in: query
	requestTypes.QueryImageDeviceTypesRequest
}

// DescribeImageDeviceTypesResponse is an response struct
// swagger:response describeImageDeviceTypes
type DescribeImageDeviceTypesResponse struct {
	// in: body
	Body struct {
		Result responseTypes.DeviceTypePage `json:"result"`
		CommonRespBody
	}
}
