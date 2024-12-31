package auditLogsDao

import (
	"fmt"
	"strings"
	"time"

	"coding.jd.com/aidc-bmp/bmp-openapi/dao"
	log "coding.jd.com/aidc-bmp/bmp_log"
)

// 操作日志
type AuditLogs struct {
	ID              int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Logid           string `gorm:"column:logid;type:varchar(255);not null" json:"logid"` // log uuid
	Sn              string `gorm:"index:i_sn;column:sn;type:varchar(255);not null" json:"sn"`
	DeviceID        string `gorm:"index:i_device_id;column:device_id;type:varchar(255);not null" json:"device_id"`
	InstanceID      string `gorm:"index:i_instance_id;column:instance_id;type:varchar(255);not null" json:"instance_id"`
	Operation       string `gorm:"index:i_operation;column:operation;type:varchar(255);not null" json:"operation"` // action
	OperateUserID   string `gorm:"column:operate_user_id;type:varchar(255);not null" json:"operate_user_id"`       // user_id
	OperateUserName string `gorm:"column:operate_user_name;type:varchar(255);not null" json:"operate_user_name"`   // user_name
	Result          string `gorm:"column:result;type:varchar(255);not null" json:"result"`                         // success/fail
	FailReason      string `gorm:"column:fail_reason;type:varchar(255);not null" json:"fail_reason"`               // reason for fail                    // 完成时间戳
	CreatedTime     int    `gorm:"column:created_time" json:"createdTime"`                                         // 创建时间戳
	UpdatedTime     int    `gorm:"column:updated_time" json:"updatedTime"`                                         // 更新时间戳
	DeletedTime     int    `gorm:"column:deleted_time" json:"deletedTime"`                                         // 删除时间戳
	IsDel           int8   `gorm:"column:is_del;type:tinyint(4);not null" json:"is_del"`                           // 是否删除0未删除 1已删除
}

//var db *gorm.DB = dao.IronicRdb

func (t *AuditLogs) TableName() string {
	return "audit_logs"
}

// AddCommand insert a new Command into database and returns
// last inserted Id on success.
func AddAuditLogs(logger *log.Logger, m *AuditLogs) (id int64, err error) {
	m.CreatedTime = int(time.Now().Unix())
	m.UpdatedTime = int(time.Now().Unix())
	return dao.CreateAndGetId(logger, dao.IronicWdb, m)
}

// GetCommandById retrieves Command by Id. Returns error if
// Id doesn't exist
func GetAuditLogsById(logger *log.Logger, id int64) (v *AuditLogs, err error) {
	v = &AuditLogs{}
	err = dao.Where(logger, dao.IronicRdb, "id = ? and is_del = 0", id).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

// GetCommandById retrieves Command by parentId. Returns error if
// Id doesn't exist
// GetCommandById retrieves Command by parentId. Returns error if
// Id doesn't exist
func GetAuditLogsByUUId(logger *log.Logger, uuid string) (v *AuditLogs, err error) {
	v = &AuditLogs{}
	err = dao.Where(logger, dao.IronicRdb, "logid = ? and is_del = 0", uuid).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

func GetAuditLogsCount(logger *log.Logger, query map[string]interface{}) (n int64, err error) {

	db, err := dao.WhereBuild(dao.Model(logger, dao.IronicRdb, AuditLogs{}), query)
	if err != nil {
		return n, err
	}
	err = db.Count(&n).Error
	return
}

func GetAllAuditLogs(logger *log.Logger, query map[string]interface{}) (ml []*AuditLogs, err error) {
	g := dao.Model(logger, dao.IronicRdb, AuditLogs{})
	g, err = dao.WhereBuild(g, query)
	if err != nil {
		return nil, err
	}
	err = g.Find(&ml).Error
	return

}

// GetMultiCommand retrieves all Command matches certain condition. Returns empty list if
// no records exist
func GetMultiAuditLogs(logger *log.Logger, query map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []*AuditLogs, err error) {

	var db = dao.Model(logger, dao.IronicRdb, AuditLogs{})
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
	err = db.Offset(offset).Limit(limit).Find(&ml).Error
	return
}

func QueryAuditLogs(logger *log.Logger, query map[string]interface{}, offset, limit int64) (ml []*AuditLogs, err error) {
	return GetMultiAuditLogs(logger, query, nil, []string{"created_time"}, []string{"desc"}, offset, limit)
	/*var db = dao.Model(logger, dao.IronicRdb, DeviceType{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return*/
}

func GetOneAuditLogs(logger *log.Logger, query map[string]interface{}) (l *AuditLogs, err error) {
	l = &AuditLogs{}
	var db = dao.Model(logger, dao.IronicRdb, AuditLogs{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Take(l).Error
	if err != nil {
		return nil, err
	}
	return l, nil
}

// UpdateCommand updates Command by Id and returns error if
// the record to be updated doesn't exist
func UpdateAuditLogsById(logger *log.Logger, m *AuditLogs) (err error) {

	m.UpdatedTime = int(time.Now().Unix())
	err = dao.Model(logger, dao.IronicWdb, AuditLogs{}).Where("id = ?", m.ID).Updates(m).Error
	return
}

func UpdateMultiAuditLogs(logger *log.Logger, query map[string]interface{}, updates map[string]interface{}) (err error) {

	updates["updated_time"] = time.Now().Unix()
	var db = dao.Model(logger, dao.IronicWdb, AuditLogs{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return err
	}
	err = db.Updates(updates).Error
	return
}
