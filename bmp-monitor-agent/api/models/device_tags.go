package models

import (
	"encoding/json"
	"fmt"

	"coding.jd.com/bmp/cps-agent/util"

	"github.com/astaxie/beego/logs"
)

// DeviceTagsResponse api返回结构
type DeviceTagsResponse struct {
	Code    int
	Message string
	Data    interface{}
}

var TagsMap = map[string]string{
	"disk":       "CPS_TAGS_DISK",
	"mountpoint": "CPS_TAGS_MOUNTPOINT",
	"nic":        "CPS_TAGS_NIC",
}

// GetDeviceTagsFromRedis read CPS list from redis, according IDC ID in config file
func GetDeviceTagsFromRedis(key string, filed string) ([]string, error) {

	tags := []string{}
	data, err := util.HGetObjectFromRedis(key, filed)
	if err != nil {
		logs.Warn("get device tags by sn error. %s", err.Error())
		return []string{}, err
	}
	err = json.Unmarshal([]byte(data), &tags)
	if err != nil {
		logs.Warn("get device tags json unmarshal error. %s", err.Error())
		return []string{}, err
	}
	return tags, nil
}

// PutDeviceTags2Redis ...
func PutDeviceTags2Redis(instanceID string, reqObj ProxyPut) error {
	devicestags := []string{}
	mountpointTags := []string{}
	nicTags := []string{}
	for _, dataPointX := range reqObj.DataPoints {
		if dataPointX.Metric == "cps.disk.bytes.read" {
			if tagsMap, ok := (*dataPointX.Tags).(map[string]interface{}); ok {
				if devicestag, ok := tagsMap["device"]; ok {
					devicestags = append(devicestags, devicestag.(string))
				}
			}
		} else if dataPointX.Metric == "cps.disk.partition.used" {
			if tagsMap, ok := (*dataPointX.Tags).(map[string]interface{}); ok {
				if mountpointTag, ok := tagsMap["device"]; ok {
					mountpointTags = append(mountpointTags, mountpointTag.(string))
				}
			}
		} else if dataPointX.Metric == "cps.network.bytes.ingress" {
			if tagsMap, ok := (*dataPointX.Tags).(map[string]interface{}); ok {
				if nicTag, ok := tagsMap["device"]; ok {
					nicTags = append(nicTags, nicTag.(string))
				}
			}
		} else {
			continue
		}
	}

	datas := map[string]interface{}{}
	if len(devicestags) > 0 {
		tagByte, err := json.Marshal(devicestags)
		if err != nil {
			logs.Warn("putDeviceTags2Redis marshal error:", err.Error())
		} else {
			datas["disk"] = string(tagByte)
		}
	}
	if len(mountpointTags) > 0 {
		tagByte, err := json.Marshal(mountpointTags)
		if err != nil {
			logs.Warn("putMountpointTags2Redis marshal error:", err.Error())
		} else {
			datas["mountpoint"] = string(tagByte)
		}
	}
	if len(nicTags) > 0 {
		tagByte, err := json.Marshal(nicTags)
		if err != nil {
			logs.Warn("putNicTags2Redis marshal error:", err.Error())
		} else {
			datas["nic"] = string(tagByte)
		}
	}

	vv, _ := json.Marshal(datas)
	fmt.Println("final tags is:", string(vv))

	if len(datas) > 0 {
		if err := util.HMSetObjectToRedis(instanceID, datas); err != nil {
			logs.Warn("HMSetObjectToRedis error:", err.Error())
			return err
		}
	}

	return nil
}
