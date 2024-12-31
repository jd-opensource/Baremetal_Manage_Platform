package dao

import (
	clog "git.jd.com/cps-golang/log"
	"github.com/jinzhu/gorm"
)

// 插入单条记录，v2版本才支持插入多条记录？
//eg. db.Create(logger, rdb, u)
func Create(logger *clog.Logger, gdb *gorm.DB, out interface{}) *gorm.DB {
	gl := &gLogger{traceid: logger.GetPoint("logid").(string)}
	db := gdb.New()
	db.SetLogger(gl)
	return db.Create(out)
}

//插入并获取刚插入数据的自增Id
//eg. db.CreateAndGetId(logger, rdb, u)
func CreateAndGetId(logger *clog.Logger, gdb *gorm.DB, out interface{}) (int64, error) {
	// 事务约束，否则并发下可能获取到的是错误的id
	gl := &gLogger{traceid: logger.GetPoint("logid").(string)}
	db := gdb.New()
	db.SetLogger(gl)
	tx := db.Begin()
	if err := tx.Create(out).Error; err != nil {
		tx.Rollback()
		return 0, err
	}
	var id []int64
	if err := tx.Raw("select LAST_INSERT_ID() as id").Pluck("id", &id).Error; err != nil {
		tx.Rollback()
		return 0, err
	}
	tx.Commit()
	return id[0], nil
}
