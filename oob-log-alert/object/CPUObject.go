package object

import (
	"encoding/json"
	"fmt"
)

// CPURecord CPU record
type CPURecord struct {
	FQDD                   string // CPU.Socket.1
	DeviceType             string // CPU
	NumberOfProcessorCores string // 6
	Model                  string // Intel(R) Xeon(R) CPU E5-2620 0 @ 2.00GHz
	CurrentClockSpeed      string // 2000 MHz
	CPUStatus              string // Disk 0 in Backplane 1 of Integrated RAID Controller 1
	PrimaryStatus          string // OK

}

// CPUObject , the type of CPU information from Out-of-Band
type CPUObject struct {
	Status bool        `json:"status,omitempty"`
	Count  int         `json:"count,omitempty"`
	CPUs   []CPURecord `json:"detail,omitempty"`
}

//JSON  return JSON string
func (c CPUObject) JSON() (string, error) {

	json, err := json.Marshal(c.CPUs)
	if err != nil {
		fmt.Println("CPUObject to JSON error:" + err.Error())
		return "", err
	}
	return string(json), nil
}
