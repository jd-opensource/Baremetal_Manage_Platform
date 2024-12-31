package models

import (
	"encoding/json"

	"coding.jd.com/bmp/agent-proxy-jdstack/util"
	log "git.jd.com/cps-golang/log"
)

// GetDeviceTagsFromRedis read CPS list from redis, according IDC ID in config file
func GetDeviceTagsFromRedis(logger *log.Logger, key string, filed string) ([]string, error) {

	tags := []string{}
	data, err := util.HGetObjectFromRedis(key, filed)
	if err != nil {
		logger.Warnf("get device tags by sn error, key: %s, field: %s, error: %s", key, filed, err.Error())
		return []string{}, err
	}
	err = json.Unmarshal([]byte(data), &tags)
	if err != nil {
		logger.Warnf("get device tags json unmarshal error: %s, origin msg: %s", err.Error(), data)
		return []string{}, err

	}
	return tags, nil
}

// PutDeviceTags2Redis ...
func PutDeviceTags2Redis(logger *log.Logger, instanceID string, reqObj ProxyPut) error {
	devicestags := []string{}
	mountpointTags := []string{}
	nicTags := []string{}
	for _, dataPointX := range reqObj.DataPoints {
		if dataPointX.Metric == "bmp.disk.bytes.read" {
			if dataPointX.Tags["device"] != "" {
				devicestags = append(devicestags, dataPointX.Tags["device"])
			}
		} else if dataPointX.Metric == "bmp.disk.partition.used" {
			if dataPointX.Tags["device"] != "" {
				mountpointTags = append(mountpointTags, dataPointX.Tags["device"])
			}

		} else if dataPointX.Metric == "bmp.network.bytes.ingress" {
			if dataPointX.Tags["device"] != "" {
				nicTags = append(nicTags, dataPointX.Tags["device"])
			}
		}
	}

	datas := map[string]interface{}{}
	if len(devicestags) > 0 {
		tagByte, err := json.Marshal(devicestags)
		if err != nil {
			logger.Warnf("putDeviceTags2Redis marshal error:%s", err.Error())
		} else {
			datas["disk"] = string(tagByte)
		}
	}
	if len(mountpointTags) > 0 {
		tagByte, err := json.Marshal(mountpointTags)
		if err != nil {
			logger.Warnf("putMountpointTags2Redis marshal error:%s", err.Error())
		} else {
			datas["mountpoint"] = string(tagByte)
		}
	}
	if len(nicTags) > 0 {
		tagByte, err := json.Marshal(nicTags)
		if err != nil {
			logger.Warnf("putNicTags2Redis marshal error:%s", err.Error())
		} else {
			datas["nic"] = string(tagByte)
		}
	}

	if len(datas) > 0 {
		if err := util.HMSetObjectToRedis(instanceID, datas); err != nil {
			logger.Warnf("HMSetObjectToRedis error:", err.Error())
			return err
		}
	}

	return nil
}
