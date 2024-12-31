package query

import (
	"encoding/json"
	"fmt"
	"git.jd.com/bmp-pronoea/common"
	"git.jd.com/bmp-pronoea/common/util"
	"git.jd.com/bmp-pronoea/models"
	"git.jd.com/bmp-pronoea/models/rules/bmpRules"
	"strings"
)

const (
	PROMEQL_OPERATOR_REGULAR = "=~"
	PROMEQL_OPERATOR_EQ = "="
)

var MetricLabelsOperatorRelation = map[string]map[string]string{
	models.METRIC_NAME_BMP_GAUGE : map[string]string {
		"metric_name" : PROMEQL_OPERATOR_REGULAR,
	},
	models.METRIC_NAME_BMP_COUNTER : map[string]string {
		"metric_name" : PROMEQL_OPERATOR_REGULAR,
	},
}

func CreateQueryParam(metricName, sampleMethod string, step int, labels map[string]string) (string, error) {
	if util.IsEmpty(metricName) {
		return "", fmt.Errorf("metric name empty! ")
	}
	queryParam := metricName
	if labels == nil || len(labels) <= 0 {
		return queryParam, nil
	}

	queryLabels := make([]string, 0)
	rateTemplate := ""
	for labelName, labelValue := range labels {
		if util.IsEmpty(labelName) {
			continue
		}

		operatorTmp := PROMEQL_OPERATOR_EQ
		if _, exist := MetricLabelsOperatorRelation[metricName]; exist {
			if tmp, existT := MetricLabelsOperatorRelation[metricName][labelName]; existT {
				operatorTmp = tmp
			}
		}

		queryLabel := fmt.Sprintf("%v%v'%v'", labelName, operatorTmp, labelValue)
		queryLabels = append(queryLabels, queryLabel)

		if labelName == "metric_name" {
			if value, exist := bmpRules.SampleDefaultMethodMap[labelValue]; exist {
				rateTemplate = value
			}
		}

		//是否需要添加rate查询
		/*
		if util.IsEmpty(rateTemplate) && labelName == "metric_name" {
			// 有这样的 (bmp.network.packets.ingress)|(bmp.network.packets.egress)，取第一个
			labelFirstValue := labelValue
			labelValueList := strings.Split(labelValue, common.DELIMITER_VERTICAL)
			if labelValueList != nil && len(labelValueList) >= 1 {
				labelFirstValue = labelValueList[0]
				labelFirstValue = strings.Replace(labelFirstValue, common.CHARACTER_LEFT_PARENTHESIS, "", -1)
				labelFirstValue = strings.Replace(labelFirstValue, common.CHARACTER_RIGTH_PARENTHESIS, "", -1)
			}

			logs.Info("metric_name query %v, %v", labelValue, labelFirstValue)

			if value, exist := bmpRules.SampleDefaultMethodMap[labelFirstValue]; exist {
				rateTemplate = value
			}
		}
		 */
	}

	if queryLabels != nil && len(queryLabels) > 0 {
		queryParam = fmt.Sprintf("%v{%v}", queryParam, strings.Join(queryLabels, util.SEPARATOR_COMMA))
	}

	if util.IsNotEmpty(rateTemplate) {
		//rate(bmp_monitor_counter{metric_name="bmp.disk.bytes.read"}[2m])
		queryParam = fmt.Sprintf(rateTemplate, queryParam, fmt.Sprintf("%vs", step+60))
	}

	if tmpTpl, exist := bmpRules.AlertExprMethodMap[sampleMethod]; exist {
		queryParam = fmt.Sprintf(tmpTpl, queryParam)
	}

	return queryParam, nil
}

func CreateQueryParamList(metricName, sampleMethod string, step int, labels map[string]string) (map[string]string, error) {
	metricNameValue := ""
	_, exist := labels["metric_name"]
	if exist {
		metricNameValue = labels["metric_name"]
	}
	metricNameValueList := make([]string, 0)
	if util.IsNotEmpty(metricNameValue) {
		metricNameValueList = strings.Split(metricNameValue, common.DELIMITER_VERTICAL)
	}

	queryParamList := make(map[string]string, 0)

	if len(metricNameValueList) <= 0 {
		queryParam, err := CreateQueryParam(metricName, sampleMethod, step, labels)
		if err != nil {
			return nil, err
		}
		labelsByte, err := json.Marshal(labels)
		if err != nil {
			return nil, fmt.Errorf("json.Marshal error:%v, %v", err, labels)
		}
		queryParamList[string(labelsByte)] = queryParam
		return queryParamList, nil
	}

	//len(metricNameValueList) > 0
	for _, metricNameTmpValue := range metricNameValueList {
		metricNameTmpValue = strings.Replace(metricNameTmpValue, common.CHARACTER_LEFT_PARENTHESIS, "", -1)
		metricNameTmpValue = strings.Replace(metricNameTmpValue, common.CHARACTER_RIGTH_PARENTHESIS, "", -1)

		labels["metric_name"] = metricNameTmpValue
		queryParam, err := CreateQueryParam(metricName, sampleMethod, step, labels)
		if err != nil {
			return nil, err
		}
		labelsByte, err := json.Marshal(labels)
		if err != nil {
			return nil, fmt.Errorf("json.Marshal error:%v, %v", err, labels)
		}
		queryParamList[string(labelsByte)] = queryParam
	}
	return queryParamList, nil
}
