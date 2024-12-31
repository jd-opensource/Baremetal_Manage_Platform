package object

import (
	"encoding/json"
	"fmt"
)

// NicRecord NIC record
type NicRecord struct {
	FQDD                string // NIC.Integrated.1-1-1
	DeviceType          string // NIC
	PermanentMACAddress string // 24:6E:96:B0:FA:8E
	CurrentMACAddress   string // 24:6E:96:B0:FA:8E
	ProductName         string // Intel(R) Ethernet 10G 4P X520/I350 rNDC - 24:6E:96:03:43:E0
	DeviceDescription   string // Integrated NIC 1 Port 1 Partition 1
	LinkSpeed           string // 10 Gbps
	LinkDuplex          string // Full Duplex
	LinkStatus          string // OK

}

// NicObject , the type of NIC information from Out-of-Band
type NicObject struct {
	Status bool        `json:"status,omitempty"`
	Count  int         `json:"count,omitempty"`
	Nics   []NicRecord `json:"detail,omitempty"`
}

//JSON  return JSON string
func (c NicObject) JSON() (string, error) {

	json, err := json.Marshal(c.Nics)
	if err != nil {
		fmt.Println("NicObject to JSON error:" + err.Error())
		return "", err
	}
	return string(json), nil
}
