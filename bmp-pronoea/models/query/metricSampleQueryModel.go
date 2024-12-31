package query

import (
	"encoding/json"
	"fmt"
	"git.jd.com/bmp-pronoea/common/util"
	"git.jd.com/bmp-pronoea/dao/api/prometheusApi"
	"git.jd.com/bmp-pronoea/dao/api/prometheusApi/exec"
	"git.jd.com/bmp-pronoea/types/request"
	"git.jd.com/bmp-pronoea/types/response"
	"github.com/astaxie/beego/logs"
	"strconv"
)

//查询prometheus 样本数据
func SearchPrometheusSamples (
	requestId string,
	req *request.MetricRangeQueryRequest) ([]*response.MetricRangeQueryResponse, error) {

	if err := req.Validate(req); err != nil {
		return nil, err
	}

	queryParamMap, err := CreateQueryParamList(req.TableName, req.SampleMethod, req.Step, req.Labels)
	if err != nil {
		return nil, fmt.Errorf("search prometheus sample fail %v", err)
	}
	/*
	queryParam, err := CreateQueryParam(req.TableName, req.SampleMethod, req.Step, req.Labels)
	if err != nil {
		return nil, fmt.Errorf("search prometheus sample fail %v", err)
	}
	 */

	startUtcTime, err := util.GetUtcTimeByTimeStamp(req.StartTime)
	if err != nil {
		return nil, fmt.Errorf("search prometheus sample fail %v", err)
	}
	endUtcTime, err := util.GetUtcTimeByTimeStamp(req.EndTime)
	if err != nil {
		return nil, fmt.Errorf("search prometheus sample fail %v", err)
	}

	resultList := make([]*response.MetricRangeQueryResponse, 0)
	for labelsTmp, queryParam := range queryParamMap {
		queryReq := &prometheusApi.PrometheusRangeQueryRequest{
			Query : queryParam,
			StartUTCTime : startUtcTime,
			EndUTCTime : endUtcTime,
			Step    : fmt.Sprintf("%vs", req.Step),
		}

		queryResp, err := exec.GetMetricRangeDataApi(requestId, queryReq)
		if err != nil {
			return nil, fmt.Errorf("search prometheus sample fail %v", err)
		}
		if queryResp.Status == "error" {
			return nil, fmt.Errorf("search prometheus sample fail %v", queryResp.ErrorMsg)
		}

		results := make([]*response.MetricRangeQueryResponse, 0)
		if queryResp.Data == nil || queryResp.Data.Result == nil || len(queryResp.Data.Result) <= 0 {
			return results, nil
		}

		results, err = turnPrometheusDataToResponse(requestId, labelsTmp, queryResp.Data.Result, req)
		if err != nil {
			return nil, fmt.Errorf("search prometheus sample fail %v", err)
		}
		resultList = append(resultList, results...)
	}

	return resultList, nil
}

func turnPrometheusDataToResponse(requestId, queryLables string, result []*prometheusApi.MetricDataDetail, req *request.MetricRangeQueryRequest) ([]*response.MetricRangeQueryResponse, error) {
	returnResList := make([]*response.MetricRangeQueryResponse, 0)

	funcMap := make(map[string]int, 0)
	if req.Func != nil && len(req.Func) > 0 {
		for operator, value := range req.Func {
			if value == 0 || util.IsEmpty(operator){
				continue
			}
			funcMap[operator] = value
		}
	}

	for _, tmp := range result {
		if tmp.Values == nil || len(tmp.Values) <= 0 {
			continue
		}

		dataList := make([]*response.MetricRangeQueryData, 0)
		for _, value := range tmp.Values {
			if value == nil || len(value) != 2 {
				continue
			}
			tmpDetail := &response.MetricRangeQueryData{}
			tmpDetail.Timestamp = util.TurnInterfaceToInt64(value[0])
			tmpDetail.Value = util.TurnInterfaceToString(value[1])
			if len(funcMap) > 0 {
				valueFloat64, err := util.TurnStringToFloat64(tmpDetail.Value, 2)
				if err == nil {
					for operator, operatorVal := range funcMap {
						if operator == "/" {
							tmpDetail.Value = strconv.FormatFloat(
								valueFloat64 / float64(operatorVal),
								'f', 2, 64)
						}
						if operator == "*" {
							tmpDetail.Value = strconv.FormatFloat(
								valueFloat64 * float64(operatorVal),
								'f', 2, 64)
						}
					}
				} else {
					logs.Info(requestId, fmt.Sprintf("字符串转换Float64失败 %v, %v ", err, tmpDetail.Value))
				}
			}

			dataList = append(dataList, tmpDetail)
		}

		labelsTmp := make(map[string]string, 0)
		err := json.Unmarshal([]byte(queryLables), &labelsTmp)
		if err != nil {
			return nil, fmt.Errorf("json.Unmarshal error %v, %v", err, queryLables)
		}
		for key, value := range labelsTmp {
			delete(labelsTmp, key)
			keyTmp := util.UnderLineToHumpCamel(key)
			labelsTmp[keyTmp] = value
		}

		tmpQuery := &request.MetricRangeQueryRequest{
			TableName:  req.TableName,
			StartTime:  req.StartTime,
			EndTime:    req.EndTime,
			Step:       req.Step,
			Labels:     labelsTmp,
			Func:       req.Func,
			SampleMethod: req.SampleMethod,
		}
		returnRes := &response.MetricRangeQueryResponse{
			Data : dataList,
			Query: tmpQuery,
		}

		returnResList = append(returnResList, returnRes)
	}
	return returnResList, nil

}
