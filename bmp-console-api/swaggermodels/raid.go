package swaggermodels

import (
	requestTypes "coding.jd.com/aidc-bmp/bmp-console-api/types/request"
	"coding.jd.com/aidc-bmp/bmp-console-api/types/response"
)

// swagger:parameters queryVolumeRaids
type QueryRaidsByDeviceTypeIDAndVolumeTypeRequest struct {
	ReadHeader

	// in: query
	requestTypes.QueryVolumeRaidsRequest
}

// QueryRaidsByDeviceTypeIDAndVolumeTypeResponse is an response struct
// swagger:response queryVolumeRaids
type QueryRaidsByDeviceTypeIDAndVolumeTypeResponse struct {
	// in: body
	Body struct {
		Result    []*response.VolumeRaid `json:"result"`
		RequestId string                 `json:"requestId"`
	}
}
