package rVolumeRaidDao

import (
	"fmt"
	"strings"

	"coding.jd.com/aidc-bmp/bmp-openapi/dao"
	log "coding.jd.com/aidc-bmp/bmp_log"
)

// Volume RAID 关联关系表
type RVolumeRaid struct {
	ID           int    `gorm:"primaryKey;column:id" json:"-"`             // 主键
	VolumeID     string `gorm:"column:volume_id" json:"volumeId"`          // 卷uuid
	DeviceTypeID string `gorm:"column:device_type_id" json:"deviceTypeId"` // 设备类型uuid

	RaidCan     string `gorm:"column:raid_can" json:"raidCan"`         // RAID配置： (RAID,NO RAID)
	RaidID      string `gorm:"column:raid_id" json:"raidId"`           // RAID模式：RAID1,RIAD10等
	RaidName    string `gorm:"column:raid_name" json:"raidName"`       // RAID名称
	CreatedBy   string `gorm:"column:created_by" json:"createdBy"`     // 创建者
	UpdatedBy   string `gorm:"column:updated_by" json:"updatedBy"`     // 更新者
	CreatedTime int    `gorm:"column:created_time" json:"createdTime"` // 创建时间
	UpdatedTime int    `gorm:"column:updated_time" json:"updatedTime"` // 更新时间
	DeletedTime int    `gorm:"column:deleted_time" json:"deletedTime"` // 删除时间
	IsDel       int8   `gorm:"column:is_del" json:"isDel"`             // 是否删除0未删除 1已删除
}

func (t *RVolumeRaid) TableName() string {
	return "r_volume_raid"
}

// AddVolume insert a new Volume into database and returns
// last inserted Id on success.
func AddRVolumeRaid(logger *log.Logger, m *RVolumeRaid) (id int64, err error) {

	//m.CreateTime = time.Now()
	//m.UpdateTime = time.Now()
	return dao.CreateAndGetId(logger, dao.IronicWdb, m)
}

// GetVolumeById retrieves Volume by Id. Returns error if
// Id doesn't exist
func GetRVolumeRaidsByVId(logger *log.Logger, VolumeId string) (ml []*RVolumeRaid, err error) {
	ml = []*RVolumeRaid{}
	err = dao.Where(logger, dao.IronicRdb, "volume_id = ? and is_del = 0", VolumeId).Find(&ml).Error
	if err != nil {
		return nil, err
	}
	return ml, nil
}

// GetVolumeByUuid retrieves Volume by Uuid. Returns error if
// Id doesn't exist
func GetRVolumeRaidCount(logger *log.Logger, query map[string]interface{}) (n int64, err error) {
	var db = dao.Model(logger, dao.IronicRdb, RVolumeRaid{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return 0, err
	}
	err = db.Count(&n).Error
	return
}

// GetAllVolume retrieves all Volume matches certain condition. Returns empty list if
// no records exist
func GetAllRVolumeRaid(logger *log.Logger, query map[string]interface{}) (ml []*RVolumeRaid, err error) {

	var db = dao.Model(logger, dao.IronicRdb, RVolumeRaid{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}

// GetMultiVolume retrieves all Volume matches certain condition. Returns empty list if
// no records exist
func GetMultiRVolumeRaid(logger *log.Logger, query map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []*RVolumeRaid, err error) {

	var db = dao.Model(logger, dao.IronicRdb, RVolumeRaid{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	if len(fields) > 0 {
		db = db.Select(fields)
	}
	orderConditions := []string{}
	for i := 0; i < len(sortby); i++ {
		orderConditions = append(orderConditions, fmt.Sprintf("%s %s", sortby[i], order[i]))
	}
	db = db.Order(strings.Join(orderConditions, ","))
	if offset == 0 && limit == 0 {
		err = db.Find(&ml).Error
		return
	}
	err = db.Offset(offset).Limit(limit).Find(&ml).Error
	return
}

// UpdateVolume updates Volume by Id and returns error if
// the record to be updated doesn't exist
func UpdateRVolumeRaidById(logger *log.Logger, m *RVolumeRaid) (err error) {

	//m.UpdateTime = time.Now()
	return dao.Model(logger, dao.IronicWdb, RVolumeRaid{}).Where("id = ?", m.ID).Save(m).Error
}

func UpdateRVolumeRaidByWhere(logger *log.Logger, query map[string]interface{}, updates *RVolumeRaid) (err error) {
	var db = dao.Model(logger, dao.IronicWdb, RVolumeRaid{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return err
	}
	err = db.Updates(updates).Error
	return
}

func GetSystemVolumeRaid(logger *log.Logger, deviceTypeId string) (ml []*RVolumeRaid, err error) {
	ml = []*RVolumeRaid{}

	sql := `select a.* from r_volume_raid a inner join volume b on a.volume_id = b.volume_id and b.volume_type = 'system' and a.is_del = 0 and b.is_del = 0`

	err = dao.Raw(logger, dao.IronicRdb, sql).Scan(&ml).Error
	if err != nil {
		return nil, err
	}
	return ml, nil
}
