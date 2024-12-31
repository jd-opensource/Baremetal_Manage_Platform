package exec

import (
	"encoding/json"
	"git.jd.com/bmp-pronoea/common/util"
	"git.jd.com/bmp-pronoea/dao/api"
	"git.jd.com/bmp-pronoea/dao/api/prometheusApi"
)

func GetMetricRangeDataApi(requestId string, queryReq *prometheusApi.PrometheusRangeQueryRequest) (*prometheusApi.MetricDataListResponse, error) {
	api.Locker.Lock()
	defer api.Locker.Unlock()

	a, err := api.GetApi(api.PROMETHEUS_API)
	if err != nil {
		return nil, err
	}
	base := a.GetBaseApi()

	currentTime := util.GetCurrentTimestamp()
	headers := a.GetRequestHeader(currentTime)
	base.Options = &api.Options{
		Headers: headers,
		Data: map[string]api.IBaseApiRequest{
			api.PARAMS_TYPE_QUERY: queryReq,
		},
	}

	resp, err := api.Send(requestId, a, prometheusApi.METRIC_RANGE_QUERY)
	if err != nil {
		return nil, err
	}
	getMetricRangeDataResp := new(prometheusApi.MetricDataListResponse)
	if err = json.Unmarshal(resp, getMetricRangeDataResp); err != nil {
		return nil, err
	}

	return getMetricRangeDataResp, nil
}