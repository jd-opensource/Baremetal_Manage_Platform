package models

import (
	"encoding/json"
	"fmt"
)

const (
	AGENT_UNKNOWN = iota
	AGENT_OK
	AGENT_STOP
)

type AgentStatus struct {
	InstanceID   string `json:"instance_id"`
	SN           string `json:"sn"`
	AgentVersion string `json:"agent_version"`
	Timestamp    int64  `json:"timestamp"`
	Status       int    `json:"status"`
}

func (o AgentStatus) JSON() ([]byte, error) {

	if json, err := json.Marshal(o); err != nil {
		fmt.Println("AgentStatus to JSON error:" + err.Error())
		return nil, err
	} else {
		return json, nil
	}
}

func (o *AgentStatus) FromJSON(str []byte) error {

	if err := json.Unmarshal(str, o); err != nil {
		fmt.Println("From JSON to AgentStatus error:" + err.Error())
		return err
	} else {
		return nil
	}
}