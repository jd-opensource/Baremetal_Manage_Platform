package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-operation-api/types/request"
	responseTypes "coding.jd.com/aidc-bmp/bmp-operation-api/types/response"
)

// swagger:parameters createDeviceType
type CreateDeviceTypeRequest struct {
	WriteHeader
	// in:body
	Body requestTypes.CreateDeviceTypeRequest
}

// CreateDeviceTypeResponse is an response struct
// swagger:response createDeviceType
type CreateDeviceTypeResponse struct {
	// in: body
	Body struct {
		Result responseTypes.CreateDeviceTypeResult `json:"result"`
		CommonRespBody
	}
}

// swagger:parameters queryDeviceTypes
type QueryDeviceTypesRequest struct {
	WriteHeader
	// in: query
	requestTypes.QueryDeviceTypesRequest
}

// QueryDeviceTypesResponse is an response struct
// swagger:response queryDeviceTypes
type QueryDeviceTypesResponse struct {
	// in: body
	Body struct {
		Result responseTypes.DeviceTypePage `json:"result"`
		CommonRespBody
	}
}

// swagger:parameters queryDeviceType
type QueryDeviceTypeRequest struct {
	ReadHeader
	// in: path
	DeviceTypeID string `json:"device_type_id"`
}

// QueryDeviceTypeResponse is an response struct
// swagger:response queryDeviceType
type QueryDeviceTypeResponse struct {
	// in: body
	Body struct {
		Result responseTypes.DeviceType `json:"result"`
		CommonRespBody
	}
}

// swagger:parameters queryDeviceTypeImages
type QueryDeviceTypeImagesRequest struct {
	ReadHeader
	// in: query
	requestTypes.QueryDeviceTypeImagePageRequest
}

// QueryDeviceTypeImagesResponse is an response struct
// swagger:response queryDeviceTypeImages
type QueryDeviceTypeImagesResponse struct {
	// in: body
	Body struct {
		Result responseTypes.DeviceTypeImageList `json:"result"`
		CommonRespBody
	}
}

// swagger:parameters associateDeviceTypeImage
type AssociateDeviceTypeImageRequest struct {
	WriteHeader
	// in: body
	Body requestTypes.AssociateDeviceTypeImageRequest
}

// AssociateDeviceTypeImageResponse is an response struct
// swagger:response associateDeviceTypeImage
type AssociateDeviceTypeImageResponse struct {
	// in: body
	Body struct {
		Result responseTypes.CommonResponse `json:"result"`
		CommonRespBody
	}
}

// swagger:parameters disassociateDeviceTypeImage
type DisassociateDeviceTypeImageRequest struct {
	WriteHeader
	// in: body
	Body requestTypes.DisassociateDeviceTypeImageRequest
}

// DisassociateDeviceTypeImageResponse is an response struct
// swagger:response disassociateDeviceTypeImage
type DisassociateDeviceTypeImageResponse struct {
	// in: body
	Body struct {
		Result responseTypes.CommonResponse `json:"result"`
		CommonRespBody
	}
}

// swagger:parameters modifyDeviceType
type ModifyDeviceTypeRequest struct {
	WriteHeader
	// in: path
	DeviceTypeID string `json:"device_type_id"`
	// in: body
	Body requestTypes.ModifyDeviceTypeRequest
}

// ModifyDeviceTypeResponse is an response struct
// swagger:response modifyDeviceType
type ModifyDeviceTypeResponse struct {
	// in: body
	Body struct {
		Result responseTypes.CommonResponse `json:"result"`
		CommonRespBody
	}
}

// swagger:parameters deleteDeviceType
type DeleteDeviceTypeRequest struct {
	WriteHeader
	// in: path
	DeviceTypeID string `json:"device_type_id"`
	// in: body
	Body requestTypes.ModifyDeviceTypeRequest
}

// DeleteDeviceTypeResponse is an response struct
// swagger:response deleteDeviceType
type DeleteDeviceTypeResponse struct {
	// in: body
	Body struct {
		Result responseTypes.CommonResponse `json:"result"`
		CommonRespBody
	}
}
