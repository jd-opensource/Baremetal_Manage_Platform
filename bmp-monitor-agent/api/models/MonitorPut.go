package models

import (
	"encoding/json"
	"fmt"

	sdkmodels "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/monitor/models"
)

type MonitorPut struct {
	Source     string                 `json:"source"`
	Region     string                 `json:"region,omitempty"`
	InstanceId string                 `json:"instance_id,omitempty"`
	SN         string                 `json:"sn,omitempty"`
	DataPoints []sdkmodels.DataPointX `json:"data_points"`
}

type MonitorPutResponse struct {
	RequestId  string `json:"request_id,omitempty"`
	Data       string `json:"data,omitempty"`
	Message    string `json:"message,omitempty"`
	StatusCode int    `json:"status_code"`
}

type MonitoCloudPutResponse struct {
	RequestId string `json:"requestId,omitempty"`
	Result    struct {
		Failed  int `json:"failed,omitempty"`
		Success int `json:"success,omitempty"`
	} `json:"result,omitempty"`
}

func (o MonitorPut) JSON() ([]byte, error) {

	if json, err := json.Marshal(o); err != nil {
		fmt.Println("MonitorPut to JSON error:" + err.Error())
		return nil, err
	} else {
		return json, nil
	}
}

func (o *MonitorPut) FromJSON(str []byte) error {

	if err := json.Unmarshal(str, o); err != nil {
		fmt.Println("From JSON to MonitorPut error:" + err.Error())
		return err
	} else {
		return nil
	}
}
