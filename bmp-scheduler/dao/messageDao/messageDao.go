package messageDao

import (
	"fmt"
	"strings"

	"coding.jd.com/aidc-bmp/bmp-scheduler/dao"
	log "git.jd.com/cps-golang/log"
)

// WebMessage message
type WebMessage struct {
	//以下为消息通用字段
	ID             int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	MessageID      string `gorm:"column:message_id;type:varchar(255);not null" json:"message_id"`                  // message_id
	UserID         string `gorm:"index:i_message_userid;column:user_id;type:varchar(255);not null" json:"user_id"` // user_id
	UserName       string `gorm:"column:user_name;type:varchar(255);not null" json:"user_name"`                    // user_id
	ResourceType   string `gorm:"column:resource_type;type:varchar(255);not null" json:"resource_type"`            // 资源类型
	Result         string `gorm:"column:result;type:varchar(255);not null" json:"result"`                          // 操作结果
	FinishTime     int    `gorm:"column:finish_time;type:int(11);not null" json:"finish_time"`                     // 完成时间戳
	HasRead        int8   `gorm:"index:i_message_read;column:has_read;type:tinyint(4);not null" json:"has_read"`   // 0未读, 1已读
	IsDel          int8   `gorm:"column:is_del;type:tinyint(4);not null" json:"is_del"`                            // 是否删除0未删除 1已删除
	MessageType    string `gorm:"column:message_type;type:varchar(255);not null" json:"message_type"`              // 消息类型
	MessageSubType string `gorm:"column:message_sub_type;type:varchar(255);not null" json:"message_sub_type"`      // 消息子类型
	InstanceID     string `gorm:"column:instance_id;type:varchar(255);not null" json:"instance_id"`                //实例id
	// 实例name
	InstanceName string `gorm:"column:instance_name;type:varchar(255);not null" json:"instance_name"`
	SN           string `gorm:"column:sn;type:varchar(255);not null" json:"sn"`          //设备sn
	Detail       string `gorm:"column:detail;type:varchar(4096);not null" json:"detail"` //详情，带外监控的消息比较特殊

	//以下为带外故障消息专属字段
	FaultType  string `gorm:"column:fault_type;type:varchar(10);not null" json:"fault_type"` // 故障类型
	Content    string `gorm:"column:content;type:varchar(4096);not null" json:"content"`
	LogID      string `gorm:"column:logid;type:varchar(256);not null" json:"logid"`        //带外原始日志的id
	AlertTime  int    `gorm:"column:alert_time;type:int(11);not null" json:"alert_time"`   //故障告警时间
	AlertCount int    `gorm:"column:alert_count;type:int(11);not null" json:"alert_count"` //告警次数,始终为1

	//以下为许可证到期消息专属字段
	LicenseName     string `gorm:"column:license_name;type:varchar(64);not null" json:"license_name"`         // 版本名称
	LicenseNumber   string `gorm:"column:license_number;type:varchar(64);not null" json:"license_number"`     // 版本号
	LienseStartTime int    `gorm:"column:license_start_time;type:int(11);not null" json:"license_start_time"` //开始时间
	LienseEndTime   int    `gorm:"column:license_end_time;type:int(11);not null" json:"license_end_time"`     //结束时间

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
func GetWebMessageById(logger *log.Logger, id int64) (v *WebMessage, err error) {
	v = &WebMessage{}
	err = dao.Where(logger, dao.IronicRdb, "id = ? and is_del = 0", id).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

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
	err = db.Offset(offset).Limit(limit).Find(&ml).Error
	return
}

// UpdateOs updates Os by Id and returns error if
// the record to be updated doesn't exist
func UpdateOsById(logger *log.Logger, m *WebMessage) (err error) {

	//m.UpdateTime = time.Now()
	return dao.Model(logger, dao.IronicWdb, WebMessage{}).Where("id = ?", m.ID).Updates(m).Error
}
