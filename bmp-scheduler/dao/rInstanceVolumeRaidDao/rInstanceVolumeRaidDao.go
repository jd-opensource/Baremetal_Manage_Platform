package rInstanceVolumeRaidDao

import (

	//"coding.jd.com/aidc-bmp/bmp-scheduler/logic/baseLogic"

	"coding.jd.com/aidc-bmp/bmp-scheduler/dao"
	log "git.jd.com/cps-golang/log"
)

// RInstanceVolumeRaid 实例每个数据卷选择的raid
type RInstanceVolumeRaid struct {
	ID          int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`                        // 主键
	InstanceID  string `gorm:"index:i_instance_id;column:instance_id;type:varchar(255);not null" json:"instance_id"` // instance uuid
	VolumeID    string `gorm:"column:volume_id;type:varchar(255);not null" json:"volume_id"`                         // volume uuid
	VolumeType  string `gorm:"index:i_volume_type;column:volume_type;type:varchar(255);not null" json:"volume_type"` // data|system
	RaidCan     string `gorm:"column:raid_can;type:varchar(16);not null" json:"raid_can"`                            // RAID配置： (RAID,NO RAID)
	RaidID      string `gorm:"column:raid_id;type:varchar(255);not null" json:"raid_id"`                             // RAID模式：raidid,一个一条
	RaidName    string `gorm:"column:raid_name;type:varchar(255);not null" json:"raid_name"`                         // RAID名称
	CreatedBy   string `gorm:"column:created_by;type:varchar(255);not null" json:"created_by"`                       // 创建者
	UpdatedBy   string `gorm:"column:updated_by;type:varchar(255);not null" json:"updated_by"`                       // 更新者
	CreatedTime int    `gorm:"column:created_time;type:int(255);not null" json:"created_time"`                       // 创建时间
	UpdatedTime int    `gorm:"column:updated_time;type:int(11);not null" json:"updated_time"`                        // 更新时间
	DeletedTime int    `gorm:"column:deleted_time;type:int(11);not null" json:"deleted_time"`                        // 删除时间
	IsDel       int8   `gorm:"column:is_del;type:tinyint(4);not null" json:"is_del"`                                 // 是否删除0未删除 1已删除
}

func (t *RInstanceVolumeRaid) TableName() string {
	return "r_instance_volume_raid"
}

// AddVolume insert a new Volume into database and returns
// last inserted Id on success.
func AddRInstanceVolumeRaid(logger *log.Logger, m *RInstanceVolumeRaid) (id int64, err error) {

	//m.CreateTime = time.Now()
	//m.UpdateTime = time.Now()
	return dao.CreateAndGetId(logger, dao.IronicWdb, m)
}

// GetVolumeById retrieves Volume by Id. Returns error if
// Id doesn't exist
func GetRInstanceVolumeRaids(logger *log.Logger, instanceId string) (ml []*RInstanceVolumeRaid, err error) {
	ml = []*RInstanceVolumeRaid{}
	err = dao.Where(logger, dao.IronicRdb, "instance_id = ? and is_del = 0", instanceId).Find(&ml).Error
	if err != nil {
		return nil, err
	}
	return ml, nil
}

// UpdateVolume updates Volume by Id and returns error if
// the record to be updated doesn't exist
func UpdateRInstanceVolumeRaidById(logger *log.Logger, m *RInstanceVolumeRaid) (err error) {

	//m.UpdateTime = time.Now()
	return dao.Model(logger, dao.IronicWdb, RInstanceVolumeRaid{}).Where("id = ?", m.ID).Save(m).Error
}
