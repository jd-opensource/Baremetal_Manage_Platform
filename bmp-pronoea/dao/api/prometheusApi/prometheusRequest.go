package prometheusApi

import (
	"bytes"
	"encoding/json"
)

type PrometheusRangeQueryRequest struct {
	Query         string          `json:"query"`
	StartUTCTime  string          `json:"start"`
	EndUTCTime    string		  `json:"end"`
	Step          string		  `json:"step"`
}

func (p *PrometheusRangeQueryRequest) Decode() (map[string]interface{}, error) {
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

