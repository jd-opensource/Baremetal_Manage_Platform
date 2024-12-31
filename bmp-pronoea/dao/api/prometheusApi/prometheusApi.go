package prometheusApi

import (
	"fmt"
	"git.jd.com/bmp-pronoea/dao/api"
	"os"
)

type PrometheusApi struct {
	api.BaseApi
}

// 方法名
const (
	METRIC_RANGE_QUERY = "metricQueryRange"      // 查询样本数据
	PROMETHEUS_RELOAD  = "prometheusReload"      // 重启prometheus
)


var apis = map[string]*api.UrlPath{
	METRIC_RANGE_QUERY: {
		Method: "GET",
		Path:   "/api/v1/query_range",
	},
	PROMETHEUS_RELOAD: {
		Method: "POST",
		Path:   "/-/reload",
	},
}

func (a *PrometheusApi) ServiceConfig(configHeader string) error {
	networkConfig, err := a.ParseConfigByHeader(configHeader)
	if err != nil {
		return err
	}

	a.Host = networkConfig.Host
	a.Domain = networkConfig.Domain
	a.TimeOut = networkConfig.Timeout
	a.ConnectTimeOut = networkConfig.ConnectTimeout
	a.MaxIdleConnsPerHost = networkConfig.MaxIdleConnsPerHost

	return nil
}

func (a *PrometheusApi) GetBaseApi() *api.BaseApi {
	return &a.BaseApi
}

func (a *PrometheusApi) GetRequestHeader(timeStamp int64) map[string]string {
	headers := make(map[string]string, 0)
	headers["Timestamp"] = fmt.Sprintf("%v", timeStamp)
	/*
	headers["Authentication"] = a.Config["auth"]
	md5Base := fmt.Sprintf("%v%v%v", a.Config["ak"], a.Config["sk"], timeStamp)
	sign := util.GetMd5(md5Base)
	headers["Signature"] = sign
	 */

	return headers
}

func NewPrometheusApi() (*PrometheusApi, error) {
	a := &PrometheusApi{
		api.BaseApi{
			Apis:   apis,
			Config: make(map[string]string),
		},
	}
	if err := a.ServiceConfig("prometheus"); err != nil {
		return nil, err
	}
	return a, nil
}

func init() {
	a, err := NewPrometheusApi()
	if err != nil {
		fmt.Printf("init prometheus api error:%s", err.Error())
		os.Exit(1)
	}
	api.ApiMap[api.PROMETHEUS_API] = a
}