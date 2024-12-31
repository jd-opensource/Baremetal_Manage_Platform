package CpsLogContentCollectionDao

import (
	"time"

	"coding.jd.com/aidc-bmp/bmp-openapi/dao"
	log "coding.jd.com/aidc-bmp/bmp_log"
)

// CpsLogContentCollection [...]
type CpsLogContentCollection struct {
	ID          int       `gorm:"primary_key;column:id;type:int(32) unsigned;not null" json:"-"` // ID
	Content     string    `gorm:"column:content;type:varchar(256);not null" json:"content"`
	FaultConfID int       `gorm:"index:rule_id_idx;column:fault_conf_id;type:int(32) unsigned;not null" json:"fault_conf_id"`
	IsDel       bool      `gorm:"index:idx_isdel;column:is_del;type:tinyint(1);not null" json:"is_del"` // : 0,; 1,
	Sn          string    `gorm:"index:sn_idx;column:sn;type:varchar(32);not null" json:"sn"`           // sn
	Count       uint32    `gorm:"column:count;type:int(32) unsigned;not null" json:"count"`             // count
	Level       string    `gorm:"index:level_idx;column:level;type:varchar(16);not null" json:"level"`
	EventTime   time.Time `gorm:"column:event_time;type:timestamp" json:"event_time"`
	CollectTime time.Time `gorm:"column:collect_time;type:timestamp" json:"collect_time"`
	UpdateTime  time.Time `gorm:"column:update_time;type:timestamp" json:"update_time"`
}

func (t *CpsLogContentCollection) TableName() string {
	return "cps_log_content_collection"
}

func UpdateCpsLogContentCollections(logger *log.Logger, query map[string]interface{}, updates map[string]interface{}) (err error) {

	// updates["updated_time"] = time.Now().Unix()
	var db = dao.Model(logger, dao.IronicWdb, CpsLogContentCollection{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return err
	}
	err = db.Updates(updates).Error
	return

}
