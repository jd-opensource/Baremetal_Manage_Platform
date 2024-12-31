package response

import (
	sdkModels "coding.jd.com/aidc-bmp/bmp-openapi-console-sdk/models"
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
	CpuInfo    string                          `json:"cpuInfo"`
	MemInfo    string                          `json:"memInfo"`
	SystemInfo string                          `json:"systemInfo"`
	DataInfo   string                          `json:"dataInfo"`
	NicInfo    string                          `json:"nicInfo"`
	GpuInfo    string                          `json:"gpuInfo"`
	Volumes    []*sdkModels.InstanceVolumeRaid `json:"volumes"`
}

type ModifyInstanceResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

//监控返回
type InstanceMonitorDiskResponse struct {
	InstanceId string   `json:"instanceId"`
	Disk       []string `json:"disk"`
}

type InstanceMonitorMountpointsResponse struct {
	InstanceId  string   `json:"instanceId"`
	Mountpoints []string `json:"mountpoints"`
}

type InstanceMonitorNicsResponse struct {
	InstanceId string   `json:"instanceId"`
	Nics       []string `json:"nics"`
}
