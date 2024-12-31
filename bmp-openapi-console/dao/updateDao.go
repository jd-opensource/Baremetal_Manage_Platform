package dao

import (
	clog "coding.jd.com/aidc-bmp/bmp_log"
	"github.com/jinzhu/gorm"
)

//更新 根据主键id，更新所有模型字段值。查询->set->save
//eg. dao.Save(logger, rdb, &food)
func Save(logger *clog.Logger, gdb *gorm.DB, value interface{}) *gorm.DB {
	gl := &gLogger{traceid: logger.GetPoint("logid").(string)}
	db := gdb.New()
	db.SetLogger(gl)
	return db.Save(value)
}

//更新单个字段值
//eg1.更新一条记录(根据主键) dao.Model(logger, rdb, &food).Update("price", 25)
//eg2.更新所有记录 dao.Model(logger, rdb, Food{}).Update("price", 25)
//eg3.根据自定义条件更新记录 dao.Model(logger, rdb, Food{}).Where("create_time > ?", "2018-11-06 20:00:00").Update("price", 25)
func Update(logger *clog.Logger, gdb *gorm.DB, attrs ...interface{}) *gorm.DB {
	gl := &gLogger{traceid: logger.GetPoint("logid").(string)}
	db := gdb.New()
	db.SetLogger(gl)
	return db.Update(attrs...)
}

//更新多个字段值
/*
updataFood := Food{
	Stock:120,
	Title:"柠檬雪碧",
	//Age:0
}
*/
//eg1. 通过结构体变量设置更新字段 dao.Model(logger, rdb, &food).Updates(&updataFood) 问题是:这样只会更新updataFood中字段为非默认值的字段,比如updataFood定义时就算加上了Age:0，这个字段也不会被更新成目标0
//eg2. 根据自定义条件更新记录 dao.Model(logger, rdb, Food{}).Where("price > ?", 10).Updates(&updataFood) 问题同上
//eg3. 想更新所有字段值，包括零值，就是不想忽略掉空值字段时,使用map类型替代上面的结构体变量
/*
data := make(map[string]interface{})
data["stock"] = 0 //零值字段
data["price"] = 35
dao.Model(logger, rdb, Food{}).Where("id = ?", 2).Updates(data)
*/
func Updates(logger *clog.Logger, gdb *gorm.DB, values interface{}, ignoreProtectedAttrs ...bool) *gorm.DB {
	gl := &gLogger{traceid: logger.GetPoint("logid").(string)}
	db := gdb.New()
	db.SetLogger(gl)
	return db.Updates(values, ignoreProtectedAttrs...)
}

//更新表达式
//eg. db.Model(&food).Update("stock", gorm.Expr("stock + 1")) 等价于UPDATE `foods` SET `stock` = stock + 1  WHERE `foods`.`id` = '2'
