package request

import (
	sdkModels "coding.jd.com/aidc-bmp/bmp-openapi-sdk/models"
)

// type InstanceRequest struct {
// 	PageNumber string `json:"pageNumber"`
// 	PageSize   string `json:"pageSize"`
// 	CabinetNo  string `json:"cabinetNo"`
// 	Erp        string `json:"erp"`
// 	Host       string `json:"host"`
// }

type StartInstanceRequest struct {
	InstanceID string `json:"instanceId"`
}

type RestartInstanceRequest struct {
	InstanceID string `json:"instanceId"`
}

type StopInstanceRequest struct {
	InstanceID string `json:"instanceId"`
}

// func (c *InstanceRequest) Validate() {
// 	if err := validator.New().Struct(c); err != nil {
// 		panic(exception.CommonParamValid)
// 	}
// }

type StartInstancesRequest struct {
	InstanceIDs []string `json:"instanceIds"`
}

type ModifyInstancesRequest struct {
	InstanceIDs  []string `json:"instanceIds"`
	InstanceName string   `json:"instanceName"`
}

type ResetPasswdRequest struct {
	HeaderRequest
	InstanceId string `json:"instanceId" valid:"Required"`
	Password   string `json:"password"`
}

type ResetInstancesPasswdRequest struct {
	HeaderRequest
	InstanceIds []string `json:"instanceIds" valid:"Required"`
	Password    string   `json:"password"`
}

type ReinstallInstanceRequest struct {
	HeaderRequest
	InstanceId string `json:"instanceId" valid:"Required"`
	sdkModels.ReinstallInstanceRequest
}

type QueryInstanceRequest struct {
	HeaderRequest
	InstanceId string `json:"instanceId" valid:"Required"`
}
