package agent

import (
	"encoding/json"

	domainTypes "coding.jd.com/aidc-bmp/bmp-scheduler/types/domain"
)

type MakeRaid struct {
	Action        string         `json:"action"`
	Sn            string         `json:"sn"`
	MakeRaidDatas []MakeRaidData `json:"raid_datas"`
}

type MakeRaidData struct {
	RaidDriver string   `json:"raid_driver"`
	AdapterID  int      `json:"adapter_id"`
	Volumes    []Volume `json:"volumes"`
}

type Volume struct {
	VolumeID      string         `json:"volume_id"`
	IsRootDevice  bool           `json:"is_root_device"`
	IsDataDevice  bool           `json:"is_data_device"`
	RaidLevel     string         `json:"raid_level"`
	PhysicalDisks []PhysicalDisk `json:"physical_disks"`
}

type PhysicalDisk struct {
	Enclosure int `json:"enclosure"`
	Slot      int `json:"slot"`
}

//解决struct不能设置默认值的问题
func (d MakeRaid) MarshalJSON() ([]byte, error) {
	type Alias MakeRaid
	if d.Action == "" {
		d.Action = "MakeRaid"
	}
	return json.Marshal((Alias)(d))
}

type MakeRaidAgentRespData struct {
	MakeRaidDatas []MakeRaidAgentRespRaidData `json:"raid_datas"`
}

type MakeRaidAgentRespRaidData struct {
	AdapterID int                       `json:"adapter_id"`
	Volumes   []MakeRaidAgentRespVolume `json:"volumes"`
}

type MakeRaidAgentRespVolume struct {
	VolumeId        string                      `json:"volume_id"`
	RootDeviceHints domainTypes.RootDeviceHints `json:"disk_hints"`
}
