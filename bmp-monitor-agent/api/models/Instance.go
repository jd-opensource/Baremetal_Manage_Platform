package models

import (
	"encoding/json"
	"fmt"
)

type Instance struct {
	RequestID string `json:"requestId"`
	Result    struct {
		Az        string `json:"az,omitempty"`
		Bandwidth int    `json:"bandwidth,omitempty"`
		Charge    struct {
			ChargeStatus string `json:"chargeStatus,omitempty"`
		} `json:"charge,omitempty"`
		DataRaidType   string `json:"dataRaidType,omitempty"`
		DataRaidTypeID string `json:"dataRaidTypeId,omitempty"`
		Description    string `json:"description,omitempty"`
		DeviceType     string `json:"deviceType,omitempty"`
		EnableInternet string `json:"enableInternet,omitempty"`
		EnableIpv6     string `json:"enableIpv6,omitempty"`
		ImageType      string `json:"imageType,omitempty"`
		InstanceID     string `json:"instanceId,omitempty"`
		LineType       string `json:"lineType,omitempty"`
		Name           string `json:"name,omitempty"`
		NetworkType    string `json:"networkType,omitempty"`
		OsName         string `json:"osName,omitempty"`
		OsType         string `json:"osType,omitempty"`
		OsTypeID       string `json:"osTypeId,omitempty"`
		OsVersion      string `json:"osVersion,omitempty"`
		PrivateIP      string `json:"privateIp,omitempty"`
		PublicIP       string `json:"publicIp,omitempty"`
		Region         string `json:"region,omitempty"`
		Status         string `json:"status,omitempty"`
		SubnetID       string `json:"subnetId,omitempty"`
		SysRaidType    string `json:"sysRaidType,omitempty"`
		SysRaidTypeID  string `json:"sysRaidTypeId,omitempty"`
	} `json:"result,omitempty"`
	Error struct {
		Code    int    `json:"code,omitempty"`
		Status  string `json:"status,omitempty"`
		Message string `json:"message,omitempty"`
	} `json:"error,omitempty"`
}

func (o Instance) JSON() ([]byte, error) {

	if json, err := json.Marshal(o); err != nil {
		fmt.Println("Instance to JSON error:" + err.Error())
		return nil, err
	} else {
		return json, nil
	}
}

func (o *Instance) FromJSON(str []byte) error {

	if err := json.Unmarshal(str, o); err != nil {
		fmt.Println("From JSON to Instance error:" + err.Error())
		return err
	} else {
		return nil
	}
}
