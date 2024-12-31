package monitorProxyLogic

import (
	"encoding/json"
	"fmt"
	"strings"

	"coding.jd.com/aidc-bmp/bmp-openapi/logic/baseLogic"
	"coding.jd.com/aidc-bmp/bmp-openapi/types/request"
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
	"coding.jd.com/aidc-bmp/bmp-openapi/util"
	log "coding.jd.com/aidc-bmp/bmp_log"
	beego "github.com/beego/beego/v2/server/web"
)

const (
	agentStatusUrl string = "/api/v1/describeAgentStatus"
	deviceTagsUrl  string = "/api/v1/tags"
)

var AgentStatusMapZh map[int]string = map[int]string{
	0: "未安装",
	1: "正常",
	2: "异常",
}
var AgentStatusMapEn map[int]string = map[int]string{
	0: "Not Installed",
	1: "Normal",
	2: "Abnormal",
}

var monitorProxyServer string

func Init() error {
	bmp_monitor_proxy_host, err := beego.AppConfig.String("bmp_monitor_proxy_host")
	if err != nil {
		return err
	}
	bmp_monitor_proxy_port, err := beego.AppConfig.String("bmp_monitor_proxy_port")
	if err != nil {
		return err
	}
	monitorProxyServer = fmt.Sprintf("http://%s:%s", bmp_monitor_proxy_host, bmp_monitor_proxy_port)
	return nil
}

type DescribeAgentStatusResponse struct {
	// 操作失败结果。成功时有此结构
	Result []response.AgentStatusItem `json:"result"`
	// 操作失败结果。失败时有此结构
	Error interface{} `json:"error"`
	// 请求traceId
	// required: true
	RequestId string `json:"requestId"`
}

func DescribeAgentStatus(logger *log.Logger, param *request.DesrcibeAgentStatusRequest) (*response.AgentStatusResponse, error) {
	url := monitorProxyServer + agentStatusUrl
	instanceIds := strings.Split(param.InstanceID, ",")
	body := map[string]interface{}{
		"instanceIds": instanceIds,
	}
	header := map[string]string{
		"Traceid": logger.GetPoint("logid").(string),
	}
	logger.Infof("DescribeAgentStatus.request:%s", util.InterfaceToJson(body))
	resp, err := util.DoHttpPost(logger, url, header, body)
	if err != nil {
		logger.Warnf("DescribeAgentStatus post error, instanceIds:%s, error:%s", param.InstanceID, err.Error())
		return nil, err
	}
	logger.Infof("DescribeAgentStatus.resp:%s", string(resp))

	res := DescribeAgentStatusResponse{}
	if err := json.Unmarshal(resp, &res); err != nil {
		logger.Warnf("DescribeAgentStatus post body parse error, instanceIds:%s, error:%s", param.InstanceID, err.Error())
		return nil, err
	}

	for idx, item := range res.Result {
		if logger.GetPoint("language").(string) == baseLogic.EN_US {
			item.StatusName = AgentStatusMapEn[item.Status]
		} else {
			item.StatusName = AgentStatusMapZh[item.Status]
		}
		res.Result[idx] = item
	}

	return &response.AgentStatusResponse{
		AgentStatus: res.Result,
	}, nil

}

type DescribeDeviceTagsResponse struct {
	// 操作失败结果。成功时有此结构
	Result response.TagsResponse `json:"result"`
	// 操作失败结果。失败时有此结构
	Error interface{} `json:"error"`
	// 请求traceId
	// required: true
	RequestId string `json:"requestId"`
}

func DescribeDeviceTags(logger *log.Logger, param *request.DesrcibeTagsRequest) (*response.TagsResponse, error) {
	url := monitorProxyServer + deviceTagsUrl

	body := map[string]interface{}{
		"instanceId": param.InstanceID,
		"metric":     param.TagName,
	}
	header := map[string]string{
		"Traceid": logger.GetPoint("logid").(string),
	}
	resp, err := util.DoHttpPost(logger, url, header, body)
	if err != nil {
		logger.Warnf("DescribeDeviceTags post error, instanceIds:%s, error:%s", param.InstanceID, err.Error())
		return nil, err
	}

	res := DescribeDeviceTagsResponse{}
	if err := json.Unmarshal(resp, &res); err != nil {
		logger.Warnf("DescribeDeviceTags post body parse error, instanceIds:%s, error:%s", param.InstanceID, err.Error())
		return nil, err
	}
	return &response.TagsResponse{
		Tags: res.Result.Tags,
	}, nil
}
