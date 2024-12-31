package bmpApi

import (
	"bytes"
	"encoding/json"
)

type BmpAlertRequest struct {
	Alerts []*BmpAlertDetail `json:"alerts"`
}

type BmpAlertDetail struct {
	RuleName       string `json:"ruleName" validate:"required"` //规则名称
	RuleID         string `json:"ruleId" validate:"required"`
	InstanceID     string `json:"instanceId" validate:"required"`
	Trigger        string `json:"trigger" validate:"required"` //触发条件
	CalculationUnit    string `json:"calculationUnit" validate:"required"` //样本单位
	AlertValue     string `json:"alertValue" validate:"required"` //告警值
	AlertTimestamp int64  `json:"alertTimestamp" validate:"required"`//告警时间戳
	AlertPeriod    int    `json:"alertPeriod" validate:"required"` //告警持续时间
}

func (p *BmpAlertRequest) Decode() (map[string]interface{}, error) {
	var err error
	data, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	mapData := make(map[string]interface{})
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()
	if err = decoder.Decode(&mapData); err != nil {
		return nil, err
	}

	return mapData, nil
}
