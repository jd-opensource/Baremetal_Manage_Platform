package response

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/rVolumeRaidDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/raidDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/volumeDao"
)

type VolumeRaids struct {
	volumeDao.Volume
	Raids []*rVolumeRaidDao.RVolumeRaid `json:"raids"`
}

type InstanceVolumeRaid struct {
	Volume *volumeDao.Volume `json:"volume"`
	Raid   *raidDao.Raid     `json:"raid"`
}
