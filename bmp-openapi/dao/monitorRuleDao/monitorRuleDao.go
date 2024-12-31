package monitorRuleDao

import (
	"fmt"
	"strings"

	"coding.jd.com/aidc-bmp/bmp-openapi/dao"
	log "coding.jd.com/aidc-bmp/bmp_log"
)

// MonitorRules 带内监控规则表
type MonitorRules struct {
	ID            int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	RuleID        string `gorm:"index:i_rule_id;column:rule_id;type:varchar(255);not null" json:"rule_id"` // rule uuid
	Status        int8   `gorm:"column:status;type:tinyint(4);not null" json:"status"`                     // 1表示启用 2表示不启用 3表示已告警
	RuleName      string `gorm:"column:rule_name;type:varchar(255);not null" json:"rule_name"`             // 规则名称
	Dimension     string `gorm:"column:dimension;type:varchar(255);not null" json:"dimension"`             // 维度 [instance disk mountpoint nic]
	Resource      string `gorm:"column:resource;type:varchar(255);not null" json:"resource"`               // 资源类型 [只支持instance一种]
	TriggerOption string `gorm:"column:trigger_option;type:varchar(255);not null" json:"trigger_option"`   // 触发条件, json
	NoticeOption  string `gorm:"column:notice_option;type:varchar(255);not null" json:"notice_option"`     // 通知策略, json
	CreatedTime   int    `gorm:"column:created_time;type:int(11);not null" json:"created_time"`            // 创建时间戳
	UpdatedTime   int    `gorm:"column:updated_time;type:int(11);not null" json:"updated_time"`            // 更新时间戳
	DeletedTime   int    `gorm:"column:deleted_time;type:int(11);not null" json:"deleted_time"`            // 更新时间戳
	IsDel         int8   `gorm:"column:is_del;type:tinyint(4);not null" json:"is_del"`                     // 是否删除0未删除 1已删除
	UserID        string `gorm:"column:user_id;type:varchar(255);not null" json:"user_id"`                 // user uuid
	UserName      string `gorm:"column:user_name;type:varchar(255);not null" json:"user_name"`             // user name

}

func (t *MonitorRules) TableName() string {
	return "monitor_rules"
}

// AddWebMessage insert a new Os into database and returns
// last inserted Id on success.
func AddMonitorRules(logger *log.Logger, m *MonitorRules) (id int64, err error) {

	//m.CreateTime = time.Now()
	//m.UpdateTime = time.Now()
	return dao.CreateAndGetId(logger, dao.IronicWdb, m)
}

func GetMonitorRuleById(logger *log.Logger, ruleId string) (v *MonitorRules, err error) {
	v = &MonitorRules{}
	err = dao.Where(logger, dao.IronicRdb, "rule_id = ? and is_del = 0", ruleId).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

// GetAllOs retrieves all os matches certain condition. Returns empty list if
// no records exist
func GetAllMonitorRulesByUserId(logger *log.Logger, userId string) (ml []*MonitorRules, err error) {

	query := map[string]interface{}{
		"user_id": userId,
		"is_del":  0,
	}
	var db = dao.Model(logger, dao.IronicRdb, MonitorRules{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return

}

func GetMonitorRulesCount(logger *log.Logger, query map[string]interface{}) (n int64, err error) {
	var db = dao.Model(logger, dao.IronicRdb, MonitorRules{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return 0, err
	}
	err = db.Count(&n).Error
	return
}

// GetMultiOs retrieves all Os matches certain condition. Returns empty list if
// no records exist
func GetMultiMonitorRules(logger *log.Logger, query map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []*MonitorRules, err error) {

	var db = dao.Model(logger, dao.IronicRdb, MonitorRules{})
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

func GetAllMonitorRules(logger *log.Logger, query map[string]interface{}) (ml []*MonitorRules, err error) {

	var db = dao.Model(logger, dao.IronicRdb, MonitorRules{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}

// UpdateOs updates Os by Id and returns error if
// the record to be updated doesn't exist
func UpdateMonitorRules(logger *log.Logger, m *MonitorRules) (err error) {

	//m.UpdateTime = time.Now()
	return dao.Model(logger, dao.IronicWdb, MonitorRules{}).Where("id = ?", m.ID).Updates(m).Error
}

func UpdateMultiMonitorRules(logger *log.Logger, query map[string]interface{}, updates map[string]interface{}) (err error) {

	// updates["updated_time"] = time.Now().Unix()
	var db = dao.Model(logger, dao.IronicWdb, MonitorRules{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return err
	}
	err = db.Updates(updates).Error
	return

}
