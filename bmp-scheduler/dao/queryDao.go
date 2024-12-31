package dao

import (
	clog "git.jd.com/cps-golang/log"
	"github.com/jinzhu/gorm"
)

// SELECT * FROM `foods` LIMIT 1
//eg.  db.Take(logger, rdb, &food)
func Take(logger *clog.Logger, gdb *gorm.DB, out interface{}, where ...interface{}) *gorm.DB {
	gl := &gLogger{traceid: logger.GetPoint("logid").(string)}
	db := gdb.New()
	db.SetLogger(gl)
	return db.Take(out, where...)
}

// SELECT * FROM `foods`   ORDER BY `foods`.`id` ASC LIMIT 1
//eg.  db.First(logger, rdb, &food)
func First(logger *clog.Logger, gdb *gorm.DB, out interface{}, where ...interface{}) *gorm.DB {
	gl := &gLogger{traceid: logger.GetPoint("logid").(string)}
	db := gdb.New()
	db.SetLogger(gl)
	return db.First(out, where...)
}

// SELECT * FROM `foods`   ORDER BY `foods`.`id` DESC LIMIT 1
//eg.  db.Last(logger, rdb, &food)
func Last(logger *clog.Logger, gdb *gorm.DB, out interface{}, where ...interface{}) *gorm.DB {
	gl := &gLogger{traceid: logger.GetPoint("logid").(string)}
	db := gdb.New()
	db.SetLogger(gl)
	return db.Last(out, where...)
}

// SELECT * FROM `foods` 可单可复
//eg.  db.Find(logger, rdb, &foods)
func Find(logger *clog.Logger, gdb *gorm.DB, out interface{}, where ...interface{}) *gorm.DB {
	gl := &gLogger{traceid: logger.GetPoint("logid").(string)}
	db := gdb.New()
	db.SetLogger(gl)
	return db.Find(out, where...)
}

// // 查询一列值
//eg. Model(Food{}).Pluck("title", &titles)
// func Pluck(logger *clog.Logger, model interface{}, column string, target interface{}) error {
// 	//等价于：SELECT title FROM `foods` 提取了title字段，保存到titles变量
// 	//这里Model函数是为了绑定一个模型实例，可以从里面提取表名。
// 	gl := &gLogger{traceid: logger.GetPoint("logid").(string)}
// 	db := gdb.New()
// 	db.SetLogger(gl)
// 	return gdb.Model(model).Pluck("title", target).Error
// }

// 设置where条件
//eg. dao.Where(logger, rdb, "id = ?", 10).Take(&food)
//eg. dao.Where(logger, rdb, "id in (?)", []int{1,2,5,6}).Take(&food)
//eg. dao.Where(logger, rdb, "create_time >= ? and create_time <= ?", "2018-11-06 00:00:00", "2018-11-06 23:59:59").Find(&foods)
//eg. dao.Where(logger, rdb, "title like ?", "%可乐%").Find(&foods)
func Where(logger *clog.Logger, gdb *gorm.DB, query interface{}, args ...interface{}) *gorm.DB {
	gl := &gLogger{traceid: logger.GetPoint("logid").(string)}
	db := gdb.New()
	db.SetLogger(gl)
	return db.Where(query, args...)
}

// 指定返回的字段
//eg. dao.Select(logger, rdb, "id,title").Where("id = ?", 1).Take(&food)
//eg. dao.Model(logger, rdb, Food{}).Select("count(*) as total").Pluck("total", &total)
func Select(logger *clog.Logger, gdb *gorm.DB, query interface{}, args ...interface{}) *gorm.DB {
	gl := &gLogger{traceid: logger.GetPoint("logid").(string)}
	db := gdb.New()
	db.SetLogger(gl)
	return db.Select(query, args...)
}

// 指定绑定的模型,一般和Pluck一起使用
//eg. dao.Model(logger, rdb, Food{}).Select("count(*) as total").Pluck("total", &total)
func Model(logger *clog.Logger, gdb *gorm.DB, value interface{}) *gorm.DB {
	gl := &gLogger{traceid: logger.GetPoint("logid").(string)}
	db := gdb.New()
	db.SetLogger(gl)
	return db.Model(value)
}

// 指定order
//eg. dao.Order(logger, rdb, "create_time desc").Limit(10).Offset(0).Find(&foods)
//eg. dao.Where(logger, rdb, "create_time >= ?", "2018-11-06 00:00:00").Order("create_time desc").Find(&foods)
func Order(logger *clog.Logger, gdb *gorm.DB, value interface{}, reorder ...bool) *gorm.DB {
	gl := &gLogger{traceid: logger.GetPoint("logid").(string)}
	db := gdb.New()
	db.SetLogger(gl)
	return db.Order(value, reorder...)
}

// limit
//eg. dao.Order(logger, rdb, "create_time desc").Limit(10).Offset(0).Find(&foods)
//eg. dao.Limit(logger, rdb, 10).Offset(0).Find(&foods)
func Limit(logger *clog.Logger, gdb *gorm.DB, limit interface{}) *gorm.DB {
	gl := &gLogger{traceid: logger.GetPoint("logid").(string)}
	db := gdb.New()
	db.SetLogger(gl)
	return db.Limit(limit)
}

// Offset
//eg. dao.Order(logger, rdb, "create_time desc").Limit(10).Offset(0).Find(&foods)
//eg. dao.Offset(logger, rdb, 10).Find(&foods)
func Offset(logger *clog.Logger, gdb *gorm.DB, limit interface{}) *gorm.DB {
	gl := &gLogger{traceid: logger.GetPoint("logid").(string)}
	db := gdb.New()
	db.SetLogger(gl)
	return db.Offset(limit)
}

//eg. dao.Model(logger, rdb, Food{}).Count(&total)
// COUNT

//eg. dao.Model(logger, rdb, Food{}).Select("type, count(*) as  total").Group("type").Having("total > 0").Scan(&results)
//group by,必须要搭配select一起使用

//Raw
//eg. dao.Raw(logger, rdb, sql, "2018-11-06 00:00:00").Scan(&results)
func Raw(logger *clog.Logger, gdb *gorm.DB, sql string, values ...interface{}) *gorm.DB {
	gl := &gLogger{traceid: logger.GetPoint("logid").(string)}
	db := gdb.New()
	db.SetLogger(gl)
	return db.Raw(sql, values...)
}
