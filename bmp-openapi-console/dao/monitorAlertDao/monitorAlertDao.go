package monitorAlertDao

import (
	"fmt"
	"strings"

	"coding.jd.com/aidc-bmp/bmp-openapi/dao"
	log "coding.jd.com/aidc-bmp/bmp_log"
)

// MonitorAlerts 告警历史表
type MonitorAlerts struct {
	ID              int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	AlertID         string `gorm:"index:i_rule_id;column:alert_id;type:varchar(255);not null" json:"alert_id"` // alert uuid
	AlertTime       int    `gorm:"column:alert_time;type:int(11);not null" json:"alert_time"`                  // 告警时间戳
	RuleID          string `gorm:"index:i_rule_id;column:rule_id;type:varchar(255);not null" json:"rule_id"`   // rule uuid
	RuleName        string `gorm:"column:rule_name;type:varchar(255);not null" json:"rule_name"`               // rule name
	Resource        string `gorm:"column:resource;type:varchar(255);not null" json:"resource"`                 // 资源类型 [只支持instance一种]
	ResourceID      string `gorm:"column:resource_id;type:varchar(255);not null" json:"resource_id"`           // 资源id,目前就是实例id
	ResourceName    string `gorm:"column:resource_name;type:varchar(255);not null" json:"resource_name"`       // 资源名称,目前就是实例名称
	Trigger         string `gorm:"column:trigger;type:varchar(255);not null" json:"trigger"`                   // 触发条件,接口需要翻译
	AlertValue      string `gorm:"column:alert_value;type:varchar(255);not null" json:"alert_value"`           // 报警值
	CalculationUnit string `gorm:"column:calculation_unit;type:varchar(255);not null" json:"calculation_unit"` // 单位，%，GB，bps等，个没有单位
	AlertLevel      int8   `gorm:"column:alert_level;type:tinyint(4);not null" json:"alert_level"`             // 1表示一般，2表示严重，3表示紧急
	AlertPeriod     int    `gorm:"column:alert_period;type:int(11);not null" json:"alert_period"`              // 告警持续时间(分钟为单位)
	UserID          string `gorm:"column:user_id;type:varchar(255);not null" json:"user_id"`                   // 通知对象 userid
	UserName        string `gorm:"column:user_name;type:varchar(255);not null" json:"user_name"`               // 通知对象 用户名
	CreatedTime     int    `gorm:"column:created_time;type:int(11);not null" json:"created_time"`              // 创建时间戳
	IsDel           int8   `gorm:"column:is_del;type:tinyint(4);not null" json:"is_del"`                       // 是否删除0未删除 1已删除
	IsRecover       int8   `gorm:"column:is_recover;type:tinyint(4);not null" json:"is_recover"`               // 是否为恢复通知，0不是，1是
	ProjectID       string `gorm:"column:project_id" json:"project_id"`                                        // 项目uuid
}

func (t *MonitorAlerts) TableName() string {
	return "monitor_alerts"
}

// AddWebMessage insert a new Os into database and returns
// last inserted Id on success.
func AddMonitorAlerts(logger *log.Logger, m *MonitorAlerts) (id int64, err error) {

	//m.CreateTime = time.Now()
	//m.UpdateTime = time.Now()
	return dao.CreateAndGetId(logger, dao.IronicWdb, m)
}

func GetMonitorAlertsById(logger *log.Logger, alertId string) (v *MonitorAlerts, err error) {
	v = &MonitorAlerts{}
	err = dao.Where(logger, dao.IronicRdb, "alert_id = ? and is_del = 0", alertId).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

// GetAllOs retrieves all os matches certain condition. Returns empty list if
// no records exist
func GetAllMonitorAlertsByUserId(logger *log.Logger, userId string) (ml []*MonitorAlerts, err error) {

	query := map[string]interface{}{
		"user_id": userId,
		"is_del":  0,
	}
	var db = dao.Model(logger, dao.IronicRdb, MonitorAlerts{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return

}

func GetAllMonitorAlertsByRuleId(logger *log.Logger, RuleId string) (ml []*MonitorAlerts, err error) {

	query := map[string]interface{}{
		"rule_id": RuleId,
		"is_del":  0,
	}
	var db = dao.Model(logger, dao.IronicRdb, MonitorAlerts{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return

}

func GetMonitorAlertsCount(logger *log.Logger, query map[string]interface{}) (n int64, err error) {
	var db = dao.Model(logger, dao.IronicRdb, MonitorAlerts{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return 0, err
	}
	err = db.Count(&n).Error
	return
}

// GetMultiOs retrieves all Os matches certain condition. Returns empty list if
// no records exist
func GetMultiMonitorAlerts(logger *log.Logger, query map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []*MonitorAlerts, err error) {

	var db = dao.Model(logger, dao.IronicRdb, MonitorAlerts{})
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

func GetAllMonitorAlerts(logger *log.Logger, query map[string]interface{}) (ml []*MonitorAlerts, err error) {

	var db = dao.Model(logger, dao.IronicRdb, MonitorAlerts{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}

// UpdateOs updates Os by Id and returns error if
// the record to be updated doesn't exist
func UpdateMonitorAlerts(logger *log.Logger, m *MonitorAlerts) (err error) {

	//m.UpdateTime = time.Now()
	return dao.Model(logger, dao.IronicWdb, MonitorAlerts{}).Where("id = ?", m.ID).Updates(m).Error
}

func UpdateMultiMonitorAlerts(logger *log.Logger, query map[string]interface{}, updates map[string]interface{}) (err error) {

	// updates["updated_time"] = time.Now().Unix()
	var db = dao.Model(logger, dao.IronicWdb, MonitorAlerts{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return err
	}
	err = db.Updates(updates).Error
	return

}
