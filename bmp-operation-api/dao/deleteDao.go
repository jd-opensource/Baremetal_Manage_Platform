package dao

import (
	clog "git.jd.com/cps-golang/log"
	"github.com/jinzhu/gorm"
)

//删除记录
//eg1.根据模型删除
/*
food := Food{}
dao.Where(logger, rdb, "id = ?", 2).Take(&food).Delete(&food)
*/
//eg2.根据Where条件删除数据
// dao.Where(logger, rdb, "type = ?", 5).Delete(&Food{})//注意这里Delete函数需要传递一个空的模型变量指针
func Delete(logger *clog.Logger, gdb *gorm.DB, value interface{}, where ...interface{}) *gorm.DB {
	gl := &gLogger{traceid: logger.GetPoint("logid").(string)}
	db := gdb.New()
	db.SetLogger(gl)
	return db.Delete(value, where...)
}
