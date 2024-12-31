package models

import (
	"encoding/json"
)

type ProxyPut struct {
	SN         string       `json:"sn"`
	DataPoints []DataPointA `json:"dataPoints"`
}

type DataPointA struct {

	/* 监控指标名称，长度不超过255字节，只允许英文、数字、下划线_、点.,  [0-9][a-z] [A-Z] [. _ ]， 其它会返回err  */
	Metric string `json:"metric"`

	/* 数据维度，数据类型为map类型，最多五个标签，尽量不传或少传。总长度不大于255字节，只允许英文、数字、下划线_、点., [0-9][a-z] [A-Z] [. _ ]，  其它会返回err (Optional) */
	Tags map[string]string `json:"tags"`

	/* 秒级时间戳，早于当前时间30天的不能写入；建议的上报时间戳：上报时间间隔的整数倍，如上报间隔为5s，则建议上报的时间戳为 timestamp = current timestamp - (current timestamp % time interval) = 1487647187 - （1487647187 % 5） = 1487647187 -2 = 1487647185  */
	Timestamp int64 `json:"timestamp"`

	/* 监控的值。number or string。最大值为long.MAX_VALUE=9223372036854775807=263-1。累计类型的指标，累计到最大值后需要翻转为0，重新开始计数。翻转后不影响速率的计算。  */
	Value float64 `json:"value"`
}

func (o ProxyPut) JSON() ([]byte, error) {

	if json, err := json.Marshal(o); err != nil {
		return nil, err
	} else {
		return json, nil
	}
}

func (o *ProxyPut) FromJSON(str []byte) error {

	if err := json.Unmarshal(str, o); err != nil {
		return err
	} else {
		return nil
	}
}
