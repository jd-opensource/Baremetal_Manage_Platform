package mailDao

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/dao"
	log "coding.jd.com/aidc-bmp/bmp_log"
)

// MailMessage message
type MailMessage struct {
	//以下为消息通用字段
	ID int `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	// 邮件服务器地址
	ServerAddr string `gorm:"column:server_addr;type:varchar(255);not null" json:"server_addr"`
	// 邮件服务器端口
	ServerPort string `gorm:"column:server_port;type:varchar(255);not null" json:"server_port"`
	// 管理员邮箱
	AdminAddr string `gorm:"column:admin_addr;type:varchar(255);not null" json:"admin_addr"`
	// 管理员邮箱密码
	AdminPass string `gorm:"column:admin_pass;type:varchar(255);not null" json:"admin_pass"`
	//收件人邮箱
	To string `gorm:"column:to;type:varchar(255);not null" json:"to"`
	// 是否推送邮件告警0不推送，1推送
	IsPush string `gorm:"column:is_push;type:type:varchar(8);not null" json:"is_push"`
	// 是否删除0未删除 1已删除
	IsDel int8 `gorm:"column:is_del;type:tinyint(4);not null" json:"is_del"`
}

func (t *MailMessage) TableName() string {
	return "mail_message"
}

// AddMailMessage insert a new Os into database and returns
// last inserted Id on success.
func AddMailMessage(logger *log.Logger, m *MailMessage) (id int64, err error) {

	//m.CreateTime = time.Now()
	//m.UpdateTime = time.Now()
	return dao.CreateAndGetId(logger, dao.IronicWdb, m)
}

// GetMailMessage retrieves Os by Id. Returns error if
// Id doesn't exist
func GetMailMessage(logger *log.Logger) (v *MailMessage, err error) {
	v = &MailMessage{}
	err = dao.Where(logger, dao.IronicRdb, "is_del = 0").Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

func UpdateMailMessage(logger *log.Logger, query map[string]interface{}, updates map[string]interface{}) (err error) {

	// updates["updated_time"] = time.Now().Unix()
	var db = dao.Model(logger, dao.IronicWdb, MailMessage{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return err
	}
	err = db.Updates(updates).Error
	return
}

func UpdateMailAnywhere(logger *log.Logger, m *MailMessage) (err error) {

	//m.UpdateTime = time.Now()
	return dao.Model(logger, dao.IronicWdb, MailMessage{}).Save(m).Error
}
