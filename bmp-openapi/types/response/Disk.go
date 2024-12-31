package response

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/diskDao"
)

type Disk struct {
	diskDao.Disk
	//是否选中
	Selected bool `json:"selected"`
}
