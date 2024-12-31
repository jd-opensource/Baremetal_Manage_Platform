package models

import (
	"encoding/json"
	"fmt"

	sdkmodels "git.jd.com/jcloud-api-gateway/jcloud-sdk-go/services/monitor/models"
)

type ProxyPut struct {
	SN         string      `json:"sn"`
	DataPoints []sdkmodels.DataPointX `json:"dataPoints"`
}

func (o ProxyPut) JSON() ([]byte, error) {

	if json, err := json.Marshal(o); err != nil {
		fmt.Println("ProxyPut to JSON error:" + err.Error())
		return nil, err
	} else {
		return json, nil
	}
}

func (o *ProxyPut) FromJSON(str []byte) error {

	if err := json.Unmarshal(str, o); err != nil {
		fmt.Println("From JSON to ProxyPut error:" + err.Error())
		return err
	} else {
		return nil
	}
}
