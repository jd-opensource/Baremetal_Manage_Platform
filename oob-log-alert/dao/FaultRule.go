package dao

import (
	"fmt"
	"strings"
	"time"

	log "coding.jd.com/aidc-bmp/bmp_log"
)

type RuleListRequest struct {
	AlertName  string `json:"alert_name"`
	AlertSpec  string `json:"alert_spec"`
	AlertLevel string `json:"alert_level"`
	// ExportType 非空要导出
	ExportType string `json:"exportType"`
	PageRequest
}

type RuleListResponse struct {
	Detail []*RuleListResponseItem `json:"detail"`
	PageResult
}

type RuleListResponseItem struct {
	RuleId     int    `json:"rule_id"`
	AlertName  string `json:"alert_name"`
	AlertSpec  string `json:"alert_spec"`
	AlertLevel string `json:"alert_level"`
	AlertDesc  string `json:"alert_desc"`
	AlertType  string `json:"alert_type"`
	// 判定条件
	AlertCondition string `json:"alert_condition"`
	// 判定阈值
	AlertThreshold string `json:"alert_threshold"`
	// 是否推送，0表示关闭，1表示开启
	PushStatus int `json:"push_status"`
	// 是否启用，0表示关闭，1表示开启
	UseStatus int `json:"use_status"`
}

type ChangePushRequest struct {
	RuleID     int `json:"rule_id"`
	PushStatus int `json:"push_status"`
}

type ChangeUseRequest struct {
	RuleID    int `json:"rule_id"`
	UseStatus int `json:"use_status"`
}

type RuleResetPushAndUseRequest struct {
	RuleIDs string `json:"rule_ids"`
}

// CpsFaultRules [...]
type CpsFaultRules struct {
	ID            int       `gorm:"primary_key;column:id;type:int(32) unsigned;not null" json:"-"`            // ID
	Name          string    `gorm:"column:name;type:varchar(52);not null" json:"name"`                        // 故障名称
	Parts         string    `gorm:"index:parts_idx;column:parts;type:varchar(12);not null" json:"parts"`      // 故障配件
	FaultType     string    `gorm:"column:fault_type;type:varchar(10);not null" json:"fault_type"`            // 故障类型
	Condition     string    `gorm:"column:condition;type:varchar(20);not null" json:"condition"`              // 判定条件
	Threshold     string    `gorm:"column:threshold;type:varchar(52);not null" json:"threshold"`              // 判定阀值
	Level         string    `gorm:"column:level;type:varchar(16);not null" json:"level"`                      // 故障等级
	Desc          string    `gorm:"column:desc;type:varchar(152);not null" json:"desc"`                       // 故障描述
	IsPush        bool      `gorm:"index:is_push_idx;column:is_push;type:tinyint(1);not null" json:"is_push"` // 故障是否推送 0-否 1-是
	IsUse         bool      `gorm:"index:is_use_idx;column:is_use;type:tinyint(1);not null" json:"is_use"`    // 规则是否启动 0-否 1-是
	IsDefaultPush bool      `gorm:"column:is_default_push;type:tinyint(1);not null" json:"is_default_push"`   // 故障是否默认推送 0-否 1-是
	IsDefaultUse  bool      `gorm:"column:is_default_use;type:tinyint(1);not null" json:"is_default_use"`     // 规则是否默认启动 0-否 1-是
	IsDel         bool      `gorm:"column:is_del;type:tinyint(1);not null" json:"is_del"`                     // 是否删除 0-否 1-是
	CreateTime    time.Time `gorm:"column:create_time;type:timestamp" json:"create_time"`                     // 创建时间
	Owned         string    `gorm:"index:owned_idx;column:owned;type:varchar(12);not null" json:"owned"`
}

func (u *CpsFaultRules) TableName() string {
	return "cps_fault_rules"
}

func GetFaultRules(logger *log.Logger) (ml []*CpsFaultRules, err error) {

	query := map[string]interface{}{
		"is_use": 1,
		"is_del": 0,
	}

	var db = Model(logger, IronicRdb, CpsFaultRules{})
	db, err = WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}

func GetFaultPushedRules(logger *log.Logger) (ml []*CpsFaultRules, err error) {

	query := map[string]interface{}{
		"is_use":  1,
		"is_del":  0,
		"is_push": 1,
	}

	var db = Model(logger, IronicRdb, CpsFaultRules{})
	db, err = WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}

func GetFaultUsedRules(logger *log.Logger) (ml []*CpsFaultRules, err error) {

	query := map[string]interface{}{
		"is_use": 1,
		"is_del": 0,
	}

	var db = Model(logger, IronicRdb, CpsFaultRules{})
	db, err = WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}

func GetFaultRulesByRids(logger *log.Logger, rids []int) (ml []*CpsFaultRules, err error) {

	query := map[string]interface{}{
		// "is_use": 1,//规则后面会变，所以不能加此过滤条件
		"is_del": 0,
		"id.in":  rids,
	}

	var db = Model(logger, IronicRdb, CpsFaultRules{})
	db, err = WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return

}

func GetRuleList(logger *log.Logger, name, part, level string, limit, offset int64) (ml []*CpsFaultRules, n int64, err error) {

	query := map[string]interface{}{
		"is_del": 0,
	}
	if name != "" {
		query["name.like"] = "%" + name + "%"
	}
	if part != "" {
		query["parts"] = part
	}
	if level != "" {
		query["level"] = level
	}
	var db = Model(logger, IronicRdb, CpsFaultRules{})
	db, err = WhereBuild(db, query)
	if err != nil {
		return nil, 0, err
	}
	err = db.Count(&n).Error
	if err != nil {
		return nil, 0, err
	}
	err = db.Offset(offset).Limit(limit).Find(&ml).Error
	return
}

func ChangePushStatusByRuleId(logger *log.Logger, ruleId int, pushStatus int8) (err error) {

	query := map[string]interface{}{
		"is_del": 0,
		"id":     ruleId,
	}
	updates := map[string]interface{}{
		"is_push": pushStatus,
	}
	var db = Model(logger, IronicWdb, CpsFaultRules{})
	db, err = WhereBuild(db, query)
	if err != nil {
		return err
	}
	err = db.Updates(updates).Error
	return
}

func ChangeUseStatusByRuleId(logger *log.Logger, ruleId int, useStatus int8) (err error) {

	query := map[string]interface{}{
		"is_del": 0,
		"id":     ruleId,
	}
	updates := map[string]interface{}{
		"is_use": useStatus,
	}
	var db = Model(logger, IronicWdb, CpsFaultRules{})
	db, err = WhereBuild(db, query)
	if err != nil {
		return err
	}
	err = db.Updates(updates).Error
	return
}

func RuleResetPushAndUse(logger *log.Logger, ruleIds []int) (err error) {

	var ris string
	for _, v := range ruleIds {
		ris = ris + fmt.Sprintf(",%d", v)
	}
	ris = strings.TrimLeft(ris, ",")
	sql := fmt.Sprintf("update cps_fault_rules set is_push = is_default_push, is_use = is_default_use where id in (%s)", ris)
	err = IronicWdb.Exec(sql).Error
	return

}
