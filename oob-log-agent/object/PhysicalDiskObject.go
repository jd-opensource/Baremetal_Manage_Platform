package object

import (
	"encoding/json"
	"fmt"
)

// PhysicalDiskRecord Physical Disk record
type PhysicalDiskRecord struct {
	FQDD              string // Disk.Bay.0:Enclosure.Internal.0-1:RAID.Integrated.1-1
	DeviceType        string // PhysicalDisk
	SerialNumber      string // WXU1E13TLDN2
	ManufacturingYear string // 2013
	ManufacturingWeek string // 23
	ManufacturingDay  string // 7
	Manufacturer      string // WD
	Model             string // WD3001BKHG
	Slot              string // 0
	DeviceDescription string // Disk 0 in Backplane 1 of Integrated RAID Controller 1
	UsedSizeInBytes   string // 299439751168 Bytes
	FreeSizeInBytes   string // 0 Bytes
	SizeInBytes       string // 299439751168 Bytes
	PrimaryStatus     string // OK
	RollupStatus      string // OK
	RaidStatus        string // Online
	HotSpareStatus    string // No
	MediaType         string // HDD
	MaxCapableSpeed   string // 6Gbs
	DriveFormFactor   string // 2.5 inch
}

// PhysicalDiskObject , the type of Physical Disk information from Out-of-Band
type PhysicalDiskObject struct {
	Status bool                 `json:"status,omitempty"`
	Count  int                  `json:"count,omitempty"`
	PDisks []PhysicalDiskRecord `json:"detail,omitempty"`
}

//JSON  return JSON string
func (o PhysicalDiskObject) JSON() (string, error) {

	json, err := json.Marshal(o.PDisks)
	if err != nil {
		fmt.Println("PhysicalDiskObject to JSON error:" + err.Error())
		return "", err
	}
	return string(json), nil
}
