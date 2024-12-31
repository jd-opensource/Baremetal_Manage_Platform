package controllers

import (
	"encoding/json"

	"coding.jd.com/bmp/agent-proxy-jdstack/api/models"
	requestTypes "coding.jd.com/bmp/agent-proxy-jdstack/types/request"
	errorCode "git.jd.com/cps-golang/ironic-common/ironic/common/ErrorCode"
	httpStatus "git.jd.com/cps-golang/ironic-common/ironic/enums/HttpStatus"
)

// DeviceTagController ...
type DeviceTagController struct {
	BaseController
}

type GetDeviceTagsRequest struct {
	InstanceID string `json:"instanceId"`
	Metric     string `json:"metric"`
}

// GetDeviceTags ...
func (d *DeviceTagController) GetDeviceTags() {

	req := &requestTypes.GetDeviceTagsRequest{}

	if err := json.Unmarshal(d.Ctx.Input.RequestBody, req); err != nil {
		d.logPoints.Warn("DescribeAgentStatus parse body error:", err.Error())
		d.SetErrorResponse(httpStatus.BAD_REQUEST, err.Error(), errorCode.INVALID_ARGUMENT)
		return
	}

	req.Validate(d.logPoints)

	var tags []string
	var err error

	tags, err = models.GetDeviceTagsFromRedis(d.logPoints, req.InstanceID, req.Metric)
	if err != nil {
		d.logPoints.Warnf("GetDeviceTagsFromRedis error, instanceID :%s, metric: %s, error: %s", req.InstanceID, req.Metric, err.Error())
		d.Res.Result = map[string][]string{
			"tags": tags,
		}
		return
	}

	//根据原型要求调整默认顺序
	//挂载点默认/，其他
	//网卡默认bond0,没有的话默认eth0,其他
	switch req.Metric {
	case "mountpoint":
		for k, v := range tags {
			if v == "/" {
				var part1 []string
				var part2 []string
				part1 = tags[:k]
				if k < len(tags)-1 {
					part2 = tags[k+1:]
				}
				tags = append(append([]string{v}, part1...), part2...)
				break
			}
		}
	case "nic":
		flag := false
		for _, v := range tags {
			if v == "bond0" {
				tags = []string{v}
				flag = true
				break
			}
		}
		if !flag {
			for k, v := range tags {
				if v == "eth0" {
					var part1 []string
					var part2 []string
					part1 = tags[:k]
					if k < len(tags)-1 {
						part2 = tags[k+1:]
					}
					tags = append(append([]string{v}, part1...), part2...)
					break
				}
			}
		}
	}

	d.Res.Result = map[string][]string{
		"tags": tags,
	}
	return
}
