package request

import (
	"git.jd.com/cps-golang/ironic-common/exception"
	validator "github.com/go-playground/validator/v10"
)

type QueryDeviceStockRequest struct {
	IdcID        string `json:"idc_id"`
	DeviceTypeId string `json:"deviceTypeId"`
	Count        int64  `json:"count"`
}

func (c *QueryDeviceStockRequest) Validate() {
	if err := validator.New().Struct(c); err != nil {
		panic(exception.CommonParamValid)
	}
}
