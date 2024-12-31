package response

import (
	sdkModels "coding.jd.com/aidc-bmp/bmp-openapi-sdk/models"
)

type InstancesResponse struct {
	PagingResponse
	Instances []Instance `json:"instances"`
}

//详情使用
type InstanceInfo struct {
	Instance Instance `json:"instance"`
}

type CreateInstanceResponse struct {
	InstanceIDs []string `json:"instanceIds"`
}

//列表
type Instance struct {
	sdkModels.Instance
	CpuInfo    string `json:"cpuInfo"`
	MemInfo    string `json:"memInfo"`
	SystemInfo string `json:"systemInfo"`
	DataInfo   string `json:"dataInfo"`
	NicInfo    string `json:"nicInfo"`
	GpuInfo    string `json:"gpuInfo"`
}
