package instanceUserDataDao

import (
	"time"

	log "coding.jd.com/aidc-bmp/bmp_log"

	"coding.jd.com/aidc-bmp/bmp-openapi/dao"
)

// InstanceUserData 用户自定义脚本表
type InstanceUserData struct {
	ID         uint64    `gorm:"primary_key;column:id;type:bigint(20) unsigned;not null" json:"-"`       // ID
	InstanceID string    `gorm:"unique;column:instance_id;type:varchar(36);not null" json:"instance_id"` // 实例ID
	UserData   string    `gorm:"column:user_data;type:mediumtext" json:"user_data"`                      // 用户自定义脚本
	CreateTime time.Time `gorm:"column:create_time;type:datetime;not null" json:"create_time"`           // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;type:datetime" json:"update_time"`                    // 更新时间
	IsDel      int8      `gorm:"column:is_del;type:tinyint(3) unsigned;not null" json:"is_del"`          // 是否删除(0-未删, 1-已删)
}

func (t *InstanceUserData) TableName() string {
	return "instance_user_data"
}

// AddInstanceUserData insert a new InstanceUserData into database and returns
// last inserted Id on success.
func AddInstanceUserData(logger *log.Logger, m *InstanceUserData) (id int64, err error) {
	m.CreateTime = time.Now()
	m.UpdateTime = time.Now()
	return dao.CreateAndGetId(logger, dao.IronicWdb, m)
}

// //TODO 有时间试试gorm v2支持的批量插入
func AddMultiInstanceUserData(logger *log.Logger, m []*InstanceUserData) (id int64, err error) {

	tx := dao.GetGormTx(logger)
	tx.Begin()
	for _, instanceUserData := range m {
		instanceUserData.CreateTime = time.Now()
		instanceUserData.UpdateTime = time.Now()
		if err := tx.Create(instanceUserData).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
	}
	tx.Commit()
	return int64(len(m)), nil
}

// GetInstanceUserDataById retrieves InstanceUserData by Id. Returns error if
// Id doesn't exist
func GetInstanceUserDataById(logger *log.Logger, id int64) (v *InstanceUserData, err error) {
	v = &InstanceUserData{}
	err = dao.Where(logger, dao.IronicRdb, "id = ? and is_del = 0", id).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

// GetAllInstanceUserData retrieves all InstanceUserData matches certain condition. Returns empty list if
// no records exist
func GetAllInstanceUserData(logger *log.Logger, query map[string]interface{}) (ml []*InstanceUserData, err error) {

	var db = dao.Model(logger, dao.IronicRdb, InstanceUserData{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}

func GetOneInstanceUserData(logger *log.Logger, query map[string]interface{}) (l *InstanceUserData, err error) {
	l = &InstanceUserData{}
	var db = dao.Model(logger, dao.IronicRdb, InstanceUserData{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Take(l).Error
	if err != nil {
		return nil, err
	}
	return l, nil
}

// UpdateInstanceUserDataById updates InstanceUserData by Id and returns error if
// the record to be updated doesn't exist
func UpdateInstanceUserDataById(logger *log.Logger, m *InstanceUserData) (err error) {

	m.UpdateTime = time.Now()
	return dao.Model(logger, dao.IronicWdb, InstanceUserData{}).Where(logger, dao.IronicRdb, "id = ?", m.ID).Updates(m).Error
}

func GetByInstanceId(logger *log.Logger, instanceId string) (l *InstanceUserData, err error) {
	l = &InstanceUserData{}
	err = dao.Where(logger, dao.IronicRdb, "instance_id = ? and is_del = 0", instanceId).Take(l).Error
	if err != nil {
		return nil, err
	}
	return l, nil

}
