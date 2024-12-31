package messageDao

import (
	"fmt"
	"strings"

	"coding.jd.com/aidc-bmp/bmp-openapi/dao"
	log "coding.jd.com/aidc-bmp/bmp_log"
)

// WebMessage message
type WebMessage struct {
	//以下为消息通用字段
	ID int `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	// message uuid
	MessageID string `gorm:"column:message_id;type:varchar(255);not null" json:"message_id"`
	// user_id
	UserID string `gorm:"index:i_message_userid;column:user_id;type:varchar(255);not null" json:"user_id"`
	// user_name
	UserName string `gorm:"column:user_name;type:varchar(255);not null" json:"user_name"`
	// 资源类型 实例/设备
	ResourceType string `gorm:"column:resource_type;type:varchar(255);not null" json:"resource_type"`
	// 操作结果 fail/succ
	Result string `gorm:"column:result;type:varchar(255);not null" json:"result"`
	// 消息时间戳
	FinishTime int `gorm:"column:finish_time;type:int(11);not null" json:"finish_time"`
	// 0未读, 1已读
	HasRead int8 `gorm:"index:i_message_read;column:has_read;type:tinyint(4);not null" json:"has_read"`
	// 是否删除0未删除 1已删除
	IsDel int8 `gorm:"column:is_del;type:tinyint(4);not null" json:"is_del"`
	// 消息类型
	MessageType string `gorm:"column:message_type;type:varchar(255);not null" json:"message_type"`
	// 消息子类型
	MessageSubType string `gorm:"column:message_sub_type;type:varchar(255);not null" json:"message_sub_type"`
	//实例id
	InstanceID string `gorm:"column:instance_id;type:varchar(255);not null" json:"instance_id"`
	// 实例name
	InstanceName string `gorm:"column:instance_name;type:varchar(255);not null" json:"instance_name"`
	//设备sn
	SN string `gorm:"column:sn;type:varchar(255);not null" json:"sn"`
	// 详情, 带外监控日志的详情请用content
	Detail string `gorm:"column:detail;type:varchar(4096);not null" json:"detail"`

	//故障类型, 带外故障消息专属字段
	FaultType string `gorm:"column:fault_type;type:varchar(10);not null" json:"fault_type"`
	//故障内容，带外故障消息专属字段
	Content string `gorm:"column:content;type:varchar(4096);not null" json:"content"`
	//系统原始日志id, 带外故障消息专属字段
	LogID string `gorm:"column:logid;type:varchar(256);not null" json:"logid"`
	//故障告警时间，带外故障消息专属字段
	AlertTime int `gorm:"column:alert_time;type:int(11);not null" json:"alert_time"`
	//告警次数,始终为1，带外故障消息专属字段
	AlertCount int `gorm:"column:alert_count;type:int(11);not null" json:"alert_count"`

	//版本名称, 许可证到期消息专属字段
	LicenseName string `gorm:"column:license_name;type:varchar(64);not null" json:"license_name"`
	//版本号, 许可证到期消息专属字段
	LicenseNumber string `gorm:"column:license_number;type:varchar(64);not null" json:"license_number"`
	//开始时间, 许可证到期消息专属字段
	LienseStartTime int `gorm:"column:license_start_time;type:int(11);not null" json:"license_start_time"`
	//结束时间, 许可证到期消息专属字段
	LienseEndTime int `gorm:"column:license_end_time;type:int(11);not null" json:"license_end_time"`
	//inbond rule uuid
	RuleID string `gorm:"column:rule_id;type:varchar(256);not null" json:"rule_id"`
	//inbond rule name
	RuleName  string `gorm:"column:rule_name;type:varchar(256);not null" json:"rule_name"`
	IsRecover int8   `gorm:"column:is_recover;type:tinyint(4);not null" json:"is_recover"` // 是否为恢复通知，0不是，1是
}

func (t *WebMessage) TableName() string {
	return "web_message"
}

// AddWebMessage insert a new Os into database and returns
// last inserted Id on success.
func AddWebMessage(logger *log.Logger, m *WebMessage) (id int64, err error) {

	//m.CreateTime = time.Now()
	//m.UpdateTime = time.Now()
	return dao.CreateAndGetId(logger, dao.IronicWdb, m)
}

// GetOsById retrieves Os by Id. Returns error if
// Id doesn't exist
func GetLatestUnreadWebMessageByUserId(logger *log.Logger, userId string) (v *WebMessage, err error) {
	v = &WebMessage{}
	err = dao.Where(logger, dao.IronicRdb, "user_id = ? and has_read = 0 and is_del = 0", userId).Order("id desc").Limit(1).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

// func GetLatestUnreadWebMessageByUserId(logger *log.Logger, userId string) (v *WebMessage, err error) {
// 	v = &WebMessage{}
// 	err = dao.Where(logger, dao.IronicRdb, "user_id = ? and  is_del = 0", userId).Take(v).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return v, nil
// }

// GetAllOs retrieves all os matches certain condition. Returns empty list if
// no records exist
func GetAllWebMessagesByUserId(logger *log.Logger, userId string) (ml []*WebMessage, err error) {

	query := map[string]interface{}{
		"user_id": userId,
		"is_del":  0,
	}
	var db = dao.Model(logger, dao.IronicRdb, WebMessage{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return

}

func GetWebMessagesByUserIdCount(logger *log.Logger, query map[string]interface{}) (n int64, err error) {
	var db = dao.Model(logger, dao.IronicRdb, WebMessage{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return 0, err
	}
	err = db.Count(&n).Error
	return
}

// GetMultiOs retrieves all Os matches certain condition. Returns empty list if
// no records exist
func GetMultiWebMessagesByUserId(logger *log.Logger, query map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []*WebMessage, err error) {

	var db = dao.Model(logger, dao.IronicRdb, WebMessage{})
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

func GetAllWebMessages(logger *log.Logger, query map[string]interface{}) (ml []*WebMessage, err error) {

	var db = dao.Model(logger, dao.IronicRdb, WebMessage{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}

// UpdateOs updates Os by Id and returns error if
// the record to be updated doesn't exist
func UpdateWebMessageId(logger *log.Logger, m *WebMessage) (err error) {

	//m.UpdateTime = time.Now()
	return dao.Model(logger, dao.IronicWdb, WebMessage{}).Where("id = ?", m.ID).Updates(m).Error
}

func UpdateWebMessages(logger *log.Logger, query map[string]interface{}, updates map[string]interface{}) (err error) {

	// updates["updated_time"] = time.Now().Unix()
	var db = dao.Model(logger, dao.IronicWdb, WebMessage{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return err
	}
	err = db.Updates(updates).Error
	return

}

// GetOsById retrieves Os by Id. Returns error if
// Id doesn't exist
func GetOneWebMessage(logger *log.Logger, query map[string]interface{}, idOrder string) (v *WebMessage, err error) {
	v = &WebMessage{}
	var db = dao.Model(logger, dao.IronicWdb, WebMessage{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	if idOrder == "desc" {
		err = db.Order("id desc").Limit(1).Take(v).Error
	} else {
		err = db.Order("id asc").Limit(1).Take(v).Error
	}

	if err != nil {
		return nil, err
	}
	return v, nil
}
