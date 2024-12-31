package monitorDataLogic

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"coding.jd.com/aidc-bmp/bmp-openapi/constant"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/instanceDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/userDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/monitorProxyLogic"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/monitorRuleLogic"
	"coding.jd.com/aidc-bmp/bmp-openapi/types/request"
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
	"coding.jd.com/aidc-bmp/bmp-openapi/util"
	log "coding.jd.com/aidc-bmp/bmp_log"
	beego "github.com/beego/beego/v2/server/web"
)

var pronoeaUrl string

func Init() error {
	bmp_pronoea_host, err := beego.AppConfig.String("bmp_pronoea_host")
	if err != nil {
		return err
	}
	bmp_pronoea_port, _ := beego.AppConfig.String("bmp_pronoea_port")
	if err != nil {
		return err
	}
	pronoeaUrl = fmt.Sprintf("http://%s:%s/api/query", bmp_pronoea_host, bmp_pronoea_port)
	return nil
}

type PronoeaMonitorDataResponse struct {
	Code      int                         `json:"Code"`
	RequestID string                      `json:"RequestId"`
	Message   string                      `json:"Message"`
	Result    []*response.DataEveryMetric `json:"Result"`
}

// func getSampleMethod(logger *log.Logger, metric string) string {
// 	rateMetrics := []string{
// 		"bmp.disk.counts.read",
// 		"bmp.disk.counts.write",
// 		"bmp.network.packets.ingress",
// 		"bmp.network.packets.egress",
// 		"bmp.disk.bytes.read",
// 		"bmp.disk.bytes.write",
// 		"bmp.network.bytes.ingress",
// 		"bmp.network.bytes.egress",
// 	}
// 	increaseMetrics := []string{}
// 	if util.InArrayString(metric, rateMetrics) {
// 		return "rate"
// 	}
// 	if util.InArrayString(metric, increaseMetrics) {
// 		return "increase"
// 	}
// 	return ""

// }

func getTagNameByMetric(metric string) string {
	mountpointMetrics := []string{

		"bmp.disk.partition.used",
		"bmp.disk.partition.util",
	}
	diskMetrics := []string{
		"bmp.disk.used",
		"bmp.disk.util",
		"bmp.disk.bytes.read",
		"bmp.disk.bytes.write",
		"bmp.disk.counts.read",
		"bmp.disk.counts.write",
	}
	nicMetrics := []string{
		"bmp.network.bytes.ingress",
		"bmp.network.bytes.egress",
		"bmp.network.packets.ingress",
		"bmp.network.packets.egress",
	}
	if util.InArrayString(metric, mountpointMetrics) {
		return "mountpoint"
	}
	if util.InArrayString(metric, diskMetrics) {
		return "disk"
	}
	if util.InArrayString(metric, nicMetrics) {
		return "nic"
	}
	return ""
}

func GetMonitorData(logger *log.Logger, param request.GetMonitorDataRequest) ([]*response.DataEveryMetric, error) {
	metrics := strings.Split(param.MetricName, ",")
	if len(metrics) == 0 {
		panic(constant.BuildCANCELLEDWithArgs("param metricName错误", "param metricName invalid"))
	}
	instanceEntity, err := instanceDao.GetInstanceById(logger, param.InstanceID)
	if err != nil {
		logger.Warnf("GetMonitorData.GetInstanceById error, instanceId:%s, error:%s", param.InstanceID, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs("instance不存在", "instanceId not exist"))

	}
	if param.IdcID != "" && instanceEntity.IDcID != param.IdcID {
		panic(constant.BuildInvalidArgumentWithArgs("instance和idcId不匹配", "instanceId not match idcId"))
	}
	userEntity, err := userDao.GetUserById(logger, instanceEntity.UserID)
	if err != nil {
		logger.Warnf("GetMonitorData.GetUserById error, userId:%s, error:%s", instanceEntity.UserID, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs("instance所属user不存在", "instance.user not exist"))
	}
	if param.UserName != "" && param.UserName != userEntity.UserName {
		panic(constant.BuildInvalidArgumentWithArgs("instance和username不匹配", "instanceId not match username"))
	}

	url := pronoeaUrl
	tableName := monitorRuleLogic.MapMetric2Table[metrics[0]]
	if tableName == "" {
		tableName = "bmp_monitor_gauge"
	}
	req := response.MonitorDataQuery{
		TableName:    tableName,
		SampleMethod: param.DownSampleType,
		StartTime:    param.StartTime,
		EndTime:      param.EndTime,
		Step:         param.TimeInterval,
	}
	if strings.Contains(param.MetricName, "bmp.memory.used") || strings.Contains(param.MetricName, "bmp.disk.used") || strings.Contains(param.MetricName, "bmp.disk.partition.used") {
		req.Func = map[string]int{
			"/": 1000000000,
		}
	}
	if strings.Contains(param.MetricName, "bmp.network.bytes.ingress") || strings.Contains(param.MetricName, "bmp.network.bytes.egress") {
		req.Func = map[string]int{
			"*": 8,
		}
	}

	//支持多个，传参格式 metric_name = "(cps.disk.partition.used)|(cps.disk.partition.util)"
	metricNames := strings.Split(param.MetricName, ",")
	if len(metricNames) > 1 {
		kw := []string{}
		for _, mn := range metricNames {
			kw = append(kw, fmt.Sprintf("(%s)", mn))
		}
		param.MetricName = strings.Join(kw, "|")
	}
	labels := map[string]string{
		"metric_name": param.MetricName,
		"instance_id": param.InstanceID,
	}

	if param.Device != "" {
		labels["device"] = param.Device
	} else { //磁盘和网口相关的指标，当没有给tag时(监控图首页),给一个默认tag，否则前端展示有问题
		tagName := getTagNameByMetric(metricNames[0])
		if tagName != "" {
			tagValues, err := monitorProxyLogic.DescribeDeviceTags(logger, &request.DesrcibeTagsRequest{
				InstanceID: param.InstanceID,
				TagName:    tagName,
			})
			if err != nil || tagValues == nil || len(tagValues.Tags) == 0 {
				logger.Warnf("GetMonitorData.DescribeDeviceTags error, instanceId:%s", param.InstanceID)
			} else {
				labels["device"] = tagValues.Tags[0]
			}
		}

	}
	req.Labels = labels
	if param.LastManyTime != 0 {
		req.EndTime = time.Now().Unix()
		req.StartTime = time.Now().Add(time.Duration(0-param.LastManyTime) * time.Hour).Unix()
	}
	logger.Infof("GetMonitorData.DoHttpPost url:%s, request:%s", url, util.InterfaceToJson(req))

	header := map[string]string{
		"Traceid": logger.GetPoint("logid").(string),
	}
	resp, err := util.DoHttpPost(logger, url, header, req)
	if err != nil {
		logger.Warnf("DescribeDeviceTags post error, instanceIds:%s, error:%s", param.InstanceID, err.Error())
		return nil, err
	}
	logger.Infof("GetMonitorData.DoHttpPost response:%s", string(resp))

	res := PronoeaMonitorDataResponse{}
	if err := json.Unmarshal(resp, &res); err != nil {
		logger.Warnf("GetMonitorData post body parse error, instanceIds:%s, error:%s", param.InstanceID, err.Error())
		return nil, err
	}
	return res.Result, nil
}
