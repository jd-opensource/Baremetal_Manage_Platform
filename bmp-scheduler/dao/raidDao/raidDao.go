package raidDao

import (
	"fmt"
	"strings"
	"time"

	"coding.jd.com/aidc-bmp/bmp-scheduler/dao"
	log "git.jd.com/cps-golang/log"
)

// Raid raid
type Raid struct {
	ID            uint64 `gorm:"primaryKey;column:id" json:"-"`              // ID
	RaidID        string `gorm:"column:raid_id" json:"raidId"`               // raid uuid
	Name          string `gorm:"column:name" json:"name"`                    // raid类型 noraid，raid0，raid1，raid10
	DescriptionEn string `gorm:"column:description_en" json:"descriptionEn"` // RAID英文描述
	DescriptionZh string `gorm:"column:description_zh" json:"descriptionZh"` // RAID中文描述
	CreatedBy     string `gorm:"column:created_by" json:"createdBy"`         // 创建者
	UpdatedBy     string `gorm:"column:updated_by" json:"updatedBy"`         // 更新者
	CreatedTime   int    `gorm:"column:created_time" json:"createdTime"`     // 创建时间戳
	UpdatedTime   int    `gorm:"column:updated_time" json:"updatedTime"`     // 更新时间戳
	DeletedTime   int    `gorm:"column:deleted_time" json:"deletedTime"`     // 删除时间戳
	IsDel         int8   `gorm:"column:is_del" json:"isDel"`                 // 是否删除0未删除 1已删除
}

func (t *Raid) TableName() string {
	return "raid"
}

// AddRaid insert a new Raid into database and returns
// last inserted Id on success.
func AddRaid(logger *log.Logger, m *Raid) (id int64, err error) {

	m.CreatedTime = int(time.Now().Unix())
	m.UpdatedTime = int(time.Now().Unix())
	return dao.CreateAndGetId(logger, dao.IronicWdb, m)
}

// GetRaidById retrieves Raid by Id. Returns error if
// Id doesn't exist
func GetRaidById(logger *log.Logger, id int64) (v *Raid, err error) {
	v = &Raid{}
	err = dao.Where(logger, dao.IronicRdb, "id = ? and is_del = 0", id).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

// GetRaidByUuid retrieves Raid by Id. Returns error if
// Id doesn't exist
func GetRaidByUuid(logger *log.Logger, uuid string) (v *Raid, err error) {
	v = &Raid{}
	err = dao.Where(logger, dao.IronicRdb, "raid_id = ? and is_del = 0", uuid).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

// GetAllRaid retrieves all raid matches certain condition. Returns empty list if
// no records exist
func GetAllRaid(logger *log.Logger, query map[string]interface{}) (ml []*Raid, err error) {

	var db = dao.Model(logger, dao.IronicRdb, Raid{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}

// GetMultiRaid retrieves all Raid matches certain condition. Returns empty list if
// no records exist
func GetMultiRaid(logger *log.Logger, query map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []*Raid, err error) {

	var db = dao.Model(logger, dao.IronicRdb, Raid{})
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
	err = db.Offset(offset).Limit(limit).Find(&ml).Error
	return
}

// UpdateRaid updates Raid by Id and returns error if
// the record to be updated doesn't exist
func UpdateRaidById(logger *log.Logger, m *Raid) (err error) {

	m.UpdatedTime = int(time.Now().Unix())
	return dao.Model(logger, dao.IronicWdb, Raid{}).Where("id = ?", m.ID).Updates(m).Error
}

type DRaid struct {
	RaidId        string `json:"raid_id"`
	NameEn        string `json:"name_en"`
	NameZh        string `json:"name_zh"`
	DeviceType    string `json:"device_type"`
	VolumeType    string `json:"volume_type"`
	VolumeDetail  string `json:"volume_detail"`
	DescriptionZh string `json:"description_zh"`
	DescriptionEn string `json:"description_en"`
}

func GetAllRraidDevice(logger *log.Logger, raid_id, device_type, volume_type string) (ml []*DRaid, err error) {

	sql := `select t1.uuid as raid_id, t1.name_en, t1.name_zh, t1.description_zh, t1.description_en, t2.device_type, t2.volume_type, t2.volume_detail from raid t1 left join r_device_type_raid t2 on t1.uuid = t2.raid_id where t1.is_del = 0 and t2.is_del = 0`
	if raid_id != "" {
		sql = sql + fmt.Sprintf(" and t1.uuid = '%s'", raid_id)
	}
	if device_type != "" {
		sql = sql + fmt.Sprintf(" and t2.device_type = '%s'", device_type)
	}
	if volume_type != "" {
		sql = sql + fmt.Sprintf(" and t2.volume_type = '%s'", volume_type)
	}

	err = dao.Raw(logger, dao.IronicRdb, sql).Scan(&ml).Error
	return

}
