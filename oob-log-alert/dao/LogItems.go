package dao

import (
	"time"

	log "coding.jd.com/aidc-bmp/bmp_log"
	"coding.jd.com/aidc-bmp/oob-log-alert/util"
)

// CpsLogItems [...]
type CpsLogItems struct {
	ID           int       `gorm:"primary_key;column:id;type:int(32) unsigned;not null" json:"-"`                                    // ID
	CollectionID int       `gorm:"index:collection_id_idx;column:collection_id;type:int(32) unsigned;not null" json:"collection_id"` // id
	CollectTime  time.Time `gorm:"column:collect_time;type:timestamp" json:"collect_time"`
	EventTime    time.Time `gorm:"column:event_time;type:timestamp" json:"event_time"`
	IsDealed     bool      `gorm:"index:is_dealed_idx;column:is_dealed;type:tinyint(1);not null" json:"is_dealed"` // : 0,, 1,
	IsDel        bool      `gorm:"index:idx_isdel;column:is_del;type:tinyint(1);not null" json:"is_del"`           // : 0,; 1,
	Sn           string    `gorm:"index:sn_idx;column:sn;type:varchar(32);not null" json:"sn"`                     // sn
}

func (u *CpsLogItems) TableName() string {
	return "cps_log_items"
}

func GetAllCpsLogItems(logger *log.Logger) (ml []*CpsLogItems, err error) {

	query := map[string]interface{}{
		"is_del": 0,
	}

	var db = Model(logger, IronicRdb, CpsLogItems{})
	db, err = WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}

func GetAllToDealedCpsLogItems(logger *log.Logger) (ml []*CpsLogItems, err error) {

	query := map[string]interface{}{
		"is_del":    0,
		"is_dealed": 0,
	}

	var db = Model(logger, IronicRdb, CpsLogItems{})
	db, err = WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}

func GetToDealedLogItemsBySn(logger *log.Logger, sn string) (ml []*CpsLogItems, err error) {

	query := map[string]interface{}{
		"is_del":    0,
		"is_dealed": 0,
		"sn":        sn,
	}

	var db = Model(logger, IronicRdb, CpsLogItems{})
	db, err = WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}

func BatchInsertLogItems(logger *log.Logger, items []*CpsLogItems) (err error) {

	tx := GetGormTx(logger)
	tx.Begin()
	for _, item := range items {
		err = tx.Create(item).Error
		if err != nil {
			logger.Warnf("BatchInsertLogItems.create error, content:%s", util.ToString(item))
			tx.Rollback()
			return
		}
	}
	tx.Commit()
	return
}
