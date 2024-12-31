package rMonitorRulesInstanceDao

import (
	"fmt"
	"strings"

	"coding.jd.com/aidc-bmp/bmp-openapi/dao"
	"coding.jd.com/aidc-bmp/bmp-openapi/util"
	log "coding.jd.com/aidc-bmp/bmp_log"
)

// RMonitorRulesInstance 带内监控规则-实例关联表
type RMonitorRulesInstance struct {
	ID           int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	RuleID       string `gorm:"index:i_rule_id;column:rule_id;type:varchar(255);not null" json:"rule_id"`             // rule uuid
	RuleName     string `gorm:"column:rule_name;type:varchar(255);not null" json:"rule_name"`                         // 规则名称
	InstanceID   string `gorm:"index:i_instance_id;column:instance_id;type:varchar(255);not null" json:"instance_id"` // instance uuid
	InstanceName string `gorm:"column:instance_name;type:varchar(255);not null" json:"instance_name"`                 // instance name
	Tags         string `gorm:"column:tags;type:varchar(255);not null" json:"tags"`                                   // disk/mountpoint/nic tag
	CreatedTime  int    `gorm:"column:created_time;type:int(11);not null" json:"created_time"`                        // 创建时间戳
	UpdatedTime  int    `gorm:"column:updated_time;type:int(11);not null" json:"updated_time"`                        // 更新时间戳
	DeletedTime  int    `gorm:"column:deleted_time;type:int(11);not null" json:"deleted_time"`                        // 更新时间戳
	IsDel        int8   `gorm:"column:is_del;type:tinyint(4);not null" json:"is_del"`                                 // 是否删除0未删除 1已删除
}

func (t *RMonitorRulesInstance) TableName() string {
	return "r_monitor_rules_instance"
}

// AddWebMessage insert a new Os into database and returns
// last inserted Id on success.
func AddRMonitorRulesInstance(logger *log.Logger, m *RMonitorRulesInstance) (id int64, err error) {

	//m.CreateTime = time.Now()
	//m.UpdateTime = time.Now()
	return dao.CreateAndGetId(logger, dao.IronicWdb, m)
}

func AddMultiRMonitorRulesInstance(logger *log.Logger, ml []*RMonitorRulesInstance) (err error) {
	for _, v := range ml {
		if _, err := dao.CreateAndGetId(logger, dao.IronicWdb, v); err != nil {
			logger.Warnf("AddMultiRMonitorRulesInstance error, v:%s, error:%s", util.InterfaceToJson(v), err.Error())
		}
	}
	//m.CreateTime = time.Now()
	//m.UpdateTime = time.Now()
	return nil
}

func GetRMonitorRulesInstanceCount(logger *log.Logger, query map[string]interface{}) (n int64, err error) {
	var db = dao.Model(logger, dao.IronicRdb, RMonitorRulesInstance{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return 0, err
	}
	err = db.Count(&n).Error
	return
}

// GetMultiOs retrieves all Os matches certain condition. Returns empty list if
// no records exist
func GetMultiRMonitorRulesInstance(logger *log.Logger, query map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []*RMonitorRulesInstance, err error) {

	var db = dao.Model(logger, dao.IronicRdb, RMonitorRulesInstance{})
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
	if offset == 0 && limit == 0 {
		err = db.Find(&ml).Error
		return
	}
	err = db.Offset(offset).Limit(limit).Find(&ml).Error
	return
}

func GetAllRMonitorRulesInstance(logger *log.Logger, query map[string]interface{}) (ml []*RMonitorRulesInstance, err error) {

	var db = dao.Model(logger, dao.IronicRdb, RMonitorRulesInstance{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}

// UpdateOs updates Os by Id and returns error if
// the record to be updated doesn't exist
func UpdateRMonitorRulesInstance(logger *log.Logger, m *RMonitorRulesInstance) (err error) {

	//m.UpdateTime = time.Now()
	return dao.Model(logger, dao.IronicWdb, RMonitorRulesInstance{}).Where("id = ?", m.ID).Updates(m).Error
}

func UpdateMultiRMonitorRulesInstance(logger *log.Logger, query map[string]interface{}, updates map[string]interface{}) (err error) {

	// updates["updated_time"] = time.Now().Unix()
	var db = dao.Model(logger, dao.IronicWdb, RMonitorRulesInstance{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return err
	}
	err = db.Updates(updates).Error
	return

}
