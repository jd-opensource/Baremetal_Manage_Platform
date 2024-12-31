package request

import (
	"git.jd.com/cps-golang/ironic-common/exception"
	validator "github.com/go-playground/validator/v10"
)

type ResourceTagRequest struct {
	Region string `json:"region"`
	Params struct {
		ResourceVo struct {
			ResourceIds  []string `json:"resourceIds"`
			ServiceCodes []string `json:"serviceCodes"`
		} `json:"resourceVo"`
	} `json:"params"`
}

func (c *ResourceTagRequest) Validate() {
	if err := validator.New().Struct(c); err != nil {
		panic(exception.CommonParamValid)
	}
}
