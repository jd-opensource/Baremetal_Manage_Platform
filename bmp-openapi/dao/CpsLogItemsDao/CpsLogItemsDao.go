package CpsLogItemsDao

import (
	"time"

	"coding.jd.com/aidc-bmp/bmp-openapi/dao"
	log "coding.jd.com/aidc-bmp/bmp_log"
)

// CpsLogItems [...]
type CpsLogItems struct {
	ID           uint32    `gorm:"primary_key;column:id;type:int(32) unsigned;not null" json:"-"`                                    // ID
	CollectionID uint32    `gorm:"index:collection_id_idx;column:collection_id;type:int(32) unsigned;not null" json:"collection_id"` // id
	CollectTime  time.Time `gorm:"column:collect_time;type:timestamp" json:"collect_time"`
	EventTime    time.Time `gorm:"column:event_time;type:timestamp" json:"event_time"`
	IsDealed     bool      `gorm:"index:is_dealed_idx;column:is_dealed;type:tinyint(1);not null" json:"is_dealed"` // : 0,, 1,
	IsDel        bool      `gorm:"index:idx_isdel;column:is_del;type:tinyint(1);not null" json:"is_del"`           // : 0,; 1,
	Sn           string    `gorm:"index:sn_idx;column:sn;type:varchar(32);not null" json:"sn"`                     // sn
}

func (t *CpsLogItems) TableName() string {
	return "cps_log_items"
}

func UpdateCpsLogItemss(logger *log.Logger, query map[string]interface{}, updates map[string]interface{}) (err error) {

	// updates["updated_time"] = time.Now().Unix()
	var db = dao.Model(logger, dao.IronicWdb, CpsLogItems{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return err
	}
	err = db.Updates(updates).Error
	return

}
