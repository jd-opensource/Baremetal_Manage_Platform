package dao

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	log "coding.jd.com/aidc-bmp/bmp_log"
	"github.com/jinzhu/gorm"
)

type GetLogCollectionsRequest struct {
	IdcId string `json:"idc_id"`
	Sn    string `json:"sn"`
	// 故障等级
	Level string `json:"level"`
	// 故障配件
	Spec       string `json:"spec"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
	DealStatus string `json:"deal_status"` //0是未处理，1是已处理
	// ExportType 非空要导出
	ExportType string `json:"exportType"`
	PageRequest
}

type GetLogCollectionsBySnRequest struct {
	Sn string `json:"sn"`
	// ExportType 非空要导出
	ExportType string `json:"exportType"`
	PageRequest
}

type GetLogCollectionsResponse struct {
	Detail []*LogRespItem `json:"detail"`
	PageResult
}

type LogRespItem struct {
	LogId    int    `json:"logid"`
	Sn       string `json:"sn"`
	DeviceId string `json:"device_id"`
	IdcId    string `json:"idc_id"`
	IdcName  string `json:"idc_name"`
	IloIp    string `json:"ilo_ip"`
	// 故障名称
	AlertName string `json:"alert_name"`
	// 故障描述
	AlertDesc string `json:"alert_desc"`
	// 故障类型
	AlertType string `json:"alert_type"`
	// 故障配件
	AlertPart string `json:"alert_part"`
	// 故障等级
	AlertLevel string `json:"alert_level"`
	// 故障原始日志
	AlertContent string `json:"alert_content"`
	// 带外故障发现时间
	CollectTime time.Time `json:"collect_time"`
	// 首次故障报警时间
	EventTime time.Time `json:"event_time"`
	// 更新时间
	UpdateTime time.Time `json:"update_time"`
	// 告警次数
	Count int `json:"count"`
	// 状态：0为未处理，1为已恢复
	Status int8 `json:"status"`
	// 设备状态
	ManageStatus string `json:"manage_status"`
}

func (d LogRespItem) MarshalJSON() ([]byte, error) {
	type Alias LogRespItem
	return json.Marshal(struct {
		Alias
		CollectTime string `json:"collect_time"`
		EventTime   string `json:"event_time"`
		UpdateTime  string `json:"update_time"`
	}{
		Alias:       Alias(d),
		CollectTime: d.CollectTime.Format("2006-01-02 15:04:05"),
		EventTime:   d.EventTime.Format("2006-01-02 15:04:05"),
		UpdateTime:  d.UpdateTime.Format("2006-01-02 15:04:05"),
	})
}

type DealLogCollectionRequest struct {
	Cid int `json:"logid"`
}

// CpsLogContentCollection [...]
type CpsLogContentCollection struct {
	ID          int       `gorm:"primary_key;column:id;type:int(32) unsigned;not null" json:"-"` // ID
	Content     string    `gorm:"column:content;type:varchar(256);not null" json:"content"`
	FaultConfID uint32    `gorm:"index:rule_id_idx;column:fault_conf_id;type:int(32) unsigned;not null" json:"fault_conf_id"`
	IsDel       bool      `gorm:"index:idx_isdel;column:is_del;type:tinyint(1);not null" json:"is_del"` // : 0,; 1,
	Sn          string    `gorm:"index:sn_idx;column:sn;type:varchar(32);not null" json:"sn"`           // sn
	Count       uint32    `gorm:"column:count;type:int(32) unsigned;not null" json:"count"`             // count
	Level       string    `gorm:"index:level_idx;column:level;type:varchar(16);not null" json:"level"`
	EventTime   time.Time `gorm:"column:event_time;type:timestamp" json:"event_time"`
	CollectTime time.Time `gorm:"column:collect_time;type:timestamp" json:"collect_time"`
	UpdateTime  time.Time `gorm:"column:update_time;type:timestamp" json:"update_time"`
}

func (u *CpsLogContentCollection) TableName() string {
	return "cps_log_content_collection"
}

func GetAllCpsLogContentCollection(logger *log.Logger) (ml []*CpsLogContentCollection, err error) {

	query := map[string]interface{}{
		"is_del": 0,
	}

	var db = Model(logger, IronicRdb, CpsLogContentCollection{})
	db, err = WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}

func InsertAndGetCollectionID(logger *log.Logger, sn, content string, ruleID int, collectTime, eventTime time.Time, level string) (n int64, err error) {

	var db = Model(logger, IronicRdb, CpsLogContentCollection{})
	c := &CpsLogContentCollection{}
	if e := db.Where("sn = ? and content = ? and is_del = 0", sn, content).Take(c).Error; e != nil {
		if e == gorm.ErrRecordNotFound {
			c.Content = content
			c.FaultConfID = uint32(ruleID)
			c.Count = 1
			c.Sn = sn
			c.Level = level
			c.EventTime = eventTime
			c.CollectTime = collectTime
			c.UpdateTime = collectTime
			id, err := InsertLogCollection(logger, c)
			if err != nil {
				return 0, fmt.Errorf("InsertLogCollection Error, content : %s", content)
			}
			logger.Infof("InsertLogCollection success, content:%s, id:%d", content, id)
			return id, nil
		} else {
			logger.Warnf("InsertAndGetCollectionID.Unexpected, err:%+v, content:%s", e, content)
			return 0, err
		}
	}
	c.Count += 1
	c.CollectTime = collectTime
	c.UpdateTime = collectTime
	if err := UpdateCpsLogContentCollectionById(logger, c); err != nil {
		return 0, err
	}

	return int64(c.ID), nil
}

func InsertLogCollection(logger *log.Logger, lc *CpsLogContentCollection) (n int64, err error) {

	n, err = CreateAndGetId(logger, IronicWdb, lc)
	if err != nil {
		return 0, fmt.Errorf("InsertAndGetCollectionID.Instert error, sn:%s, content : %s, err:%s", lc.Sn, lc.Content, err.Error())
	}
	return

}

func UpdateCpsLogContentCollectionById(logger *log.Logger, m *CpsLogContentCollection) (err error) {

	return Model(logger, IronicWdb, CpsLogContentCollection{}).Where("id = ?", m.ID).Updates(m).Error

}

type OperationLogsItem struct {
	CpsLogContentCollection
	// cps_ops_cps
	IdcId        string `json:"idc_id"`
	Brand        string `json:"brand"`
	Model        string `json:"model"`
	DeviceId     string `json:"device_id"`
	IloIp        string `json:"ilo_ip"`
	Cabinet      string `json:"cabinet"`
	ManageStatus string `json:"manage_status"`
	// Clazz      string `json:"clazz"`
	// cps_fault_rules
	Name      string `json:"name"`
	Parts     string `json:"parts"`
	FaultType string `json:"fault_type"`
	Condition string `json:"condition"`
	Threshold string `json:"threshold"`
	Desc      string `json:"desc"`
	IsPush    int8   `json:"is_push"`
	IsUse     int8   `json:"is_use"`
}

func GetLogCollectionsforOperation(logger *log.Logger, sn, level, spec, idcId, startTime, endTime string, dealStatus string, limit, offset int64) (ml []*OperationLogsItem, n int64, err error) {

	var baseSql string = ` from cps_log_content_collection a join device b join cps_fault_rules c on a.sn = b.sn and a.fault_conf_id = c.id where c.is_del = 0 `

	if idcId != "" {
		baseSql += ` and b.idc_id = '` + idcId + `'`
	}
	if startTime != "" {
		baseSql += ` and a.event_time >= '` + startTime + `'`
	}
	if endTime != "" {
		baseSql += ` and a.event_time <= '` + endTime + `'`
	}
	if sn != "" {
		baseSql += ` and a.sn = '` + sn + `'`
	}
	if level != "" {
		baseSql += ` and c.level = '` + level + `'`
	}
	if spec != "" {
		baseSql += ` and c.parts = '` + spec + `'`
	}
	if dealStatus != "" {
		if dealStatus == "0" { //未处理
			baseSql += ` and a.count > 0`
		} else { //已处理
			baseSql += ` and a.count = 0`
		}

	}

	countSql := `select count(*) ` + baseSql
	err = Raw(logger, IronicRdb, countSql).Count(&n).Error
	if err != nil {
		return nil, 0, err
	}

	logger.Info("GetLogCollectionsforOperation.countSql is:", countSql)

	var selectSql string
	selectedColumn := `a.*, b.idc_id, b.brand, b.model, b.ilo_ip, b.cabinet, b.manage_status, b.device_id, c.name, c.parts, c.fault_type, c.condition, c.threshold, c.desc, c.is_push, c.is_use`
	if limit == 0 {
		selectSql = `select ` + selectedColumn + baseSql + " order by a.collect_time desc"
	} else {
		selectSql = `select ` + selectedColumn + baseSql + " order by a.collect_time desc" + fmt.Sprintf(" limit %d offset %d", limit, offset)
	}
	logger.Info("GetLogCollectionsforOperation.selectSql is:", selectSql)
	err = Raw(logger, IronicRdb, selectSql).Scan(&ml).Error
	return
}

// console是push的
func GetLogCollectionsforConsole(logger *log.Logger, sn string, limit, offset int64) (ml []*CpsLogContentCollection, n int64, err error) {

	err = Raw(logger, IronicRdb, "select count(*) from cps_log_content_collection a join cps_fault_rules b on a.fault_conf_id = b.id where sn = ? and b.is_push = 1", sn).Count(&n).Error
	if err != nil {
		return nil, 0, err
	}
	if limit == 0 {
		err = Raw(logger, IronicRdb, "select * from cps_log_content_collection a join cps_fault_rules b on a.fault_conf_id = b.id where sn = ? and b.is_push = 1 order by a.id desc", sn).Scan(&ml).Error
		return
	} else {
		err = Raw(logger, IronicRdb, "select * from cps_log_content_collection a join cps_fault_rules b on a.fault_conf_id = b.id where sn = ? and b.is_push = 1 order by a.id desc limit ? offset ?", sn, limit, offset).Scan(&ml).Error
		return
	}
}

func DealCollectionByID(logger *log.Logger, cid int) (err error) {

	updates := map[string]interface{}{
		"count": 0,
	}
	err = Model(logger, IronicWdb, CpsLogContentCollection{}).Where("id = ?", cid).Updates(updates).Error
	if err != nil {
		logger.Warnf("DealCollectionByID.CpsLogContentCollection.update err, id:%d, err:%s", cid, err.Error())
		return
	}
	updates = map[string]interface{}{
		"is_dealed": 1,
	}
	err = Model(logger, IronicWdb, CpsLogItems{}).Where("collection_id = ?", cid).Updates(updates).Error
	if err != nil {
		logger.Warnf("DealCollectionByID.CpsLogItems.update err, id:%d, err:%s", cid, err.Error())
		return
	}
	return
}

type OrmOwnedStatusItem struct {
	Sn           string `json:"sn"`
	CollectionId int    `json:"collection_id"`
	RuleId       int    `json:"rule_id"`
	Count        int    `json:"count"`
	Owned        string `json:"owned"`
}

func GetOwnedStatusBySns(logger *log.Logger, sns []string) (ml []*OrmOwnedStatusItem, err error) {
	snsc := "'" + strings.Join(sns, "','") + "'"
	sql := fmt.Sprintf("select a.sn, b.id as collection_id, b.fault_conf_id as rule_id, b.count, c.owned from device a join cps_log_content_collection b join cps_fault_rules c on a.sn = b.sn and b.fault_conf_id = c.id where a.sn in (%s)", snsc)
	err = Raw(logger, IronicRdb, sql).Scan(&ml).Error
	return
}
