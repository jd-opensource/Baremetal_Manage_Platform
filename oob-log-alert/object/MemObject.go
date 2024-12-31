package object

import (
	"encoding/json"
	"fmt"
)

// MemRecord Memory record
type MemRecord struct {
	FQDD                  string // DIMM.Socket.A1
	DeviceType            string // Memory
	MemoryType            string // DDR-3
	SerialNumber          string // 134799E0
	Manufacturer          string // Samsung
	Speed                 string // 1333 MHz
	CurrentOperatingSpeed string // 1333 MHz
	Rank                  string // Double Rank
	Size                  string // 32768 MB
	PrimaryStatus         string // OK

}

// MemObject , the type of Memory information from Out-of-Band
type MemObject struct {
	Status bool        `json:"status,omitempty"`
	Count  int         `json:"count,omitempty"`
	Mems   []MemRecord `json:"detail,omitempty"`
}

//JSON  return JSON string
func (c MemObject) JSON() (string, error) {

	json, err := json.Marshal(c.Mems)
	if err != nil {
		fmt.Println("MemObject to JSON error:" + err.Error())
		return "", err
	}
	return string(json), nil
}
