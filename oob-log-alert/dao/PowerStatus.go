package dao

import (
	"time"
)

type PowerStatus struct {
	Id           int       `json:"id"`
	SN           string    `json:"sn"`
	InstanceID   string    `json:"instance_id"`
	InstanceName string    `json:"instance_name"`
	IdcId        string    `json:"idc_id"`
	Pin          string    `json:"pin"`
	IloIp        string    `json:"ilo_ip"`
	Status       string    `json:"status"`
	IloStatus    string    `json:"ilo_status"`
	Ipv4         string    `json:"ipv4"`
	DeviceType   string    `json:"device_type"`
	UpdateTime   time.Time `json:"update_time"`
}
