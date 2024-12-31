package models

import (
	"encoding/json"
)

type Instance struct {
	RequestID string         `json:"requestId"`
	Result    InstanceResult `json:"result,omitempty"`
	Error     struct {
		Code    int    `json:"code,omitempty"`
		Status  string `json:"status,omitempty"`
		Message string `json:"message,omitempty"`
	} `json:"error,omitempty"`
}

type InstanceResult struct {
	Instance InstanceDetail `json:"instance"`
}

type InstanceDetail struct {
	// 实例uuid
	InstanceID string `json:"instanceId"` // 实例ID（uuid）
	// 机房uuid
	IDcID string `json:"idcId"`
	// 机房名称
	IdcName string `json:"idcName"`
	// 实例所属项目UUID
	ProjectID string `json:"projectId"`
	// 实例所属用户UUID
	UserID string `json:"userId"`
	// 实例名称
	InstanceName string `json:"instanceName"`
	// 设备uuid
	DeviceID string `json:"deviceId"`
	// 设备SN
	Sn string `json:"sn"`
	// 带外管理IP
	IloIP string `json:"iloIp"`
	// 机型uuid
	DeviceTypeID string `json:"deviceTypeId"`
}

func (o Instance) JSON() ([]byte, error) {

	if json, err := json.Marshal(o); err != nil {
		return nil, err
	} else {
		return json, nil
	}
}

func (o *Instance) FromJSON(str []byte) error {

	if err := json.Unmarshal(str, o); err != nil {
		return err
	} else {
		return nil
	}
}
