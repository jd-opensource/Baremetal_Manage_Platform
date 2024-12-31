package commandDao

import (
	"fmt"
	"strings"
	"time"

	"coding.jd.com/aidc-bmp/bmp-openapi/dao"
	log "coding.jd.com/aidc-bmp/bmp_log"
)

// Command 指令
type Command struct {
	ID              int64     `gorm:"primary_key;column:id;type:bigint(20) unsigned;not null" json:"id"`                                    // ID
	RequestID       string    `gorm:"index:i_command_request_id;column:request_id;type:varchar(64);not null" json:"requestId"`              // 请求ID
	Sn              string    `gorm:"index:i_command_sn;column:sn;type:varchar(64);not null" json:"sn"`                                     // 设备SN
	InstanceID      string    `gorm:"index:i_command_instance_id;column:instance_id;type:varchar(36)" json:"instance_id"`                   // 实例Id
	Action          string    `gorm:"column:action;type:varchar(32);not null" json:"action"`                                                // 操作
	Type            string    `gorm:"column:type;type:varchar(16);not null" json:"type"`                                                    // 操作类型：agent, driver, network
	Status          string    `gorm:"index:i_command_status;column:status;type:varchar(16)" json:"status"`                                  // 状态: wait,running,finish,error
	ParentCommandID int64     `gorm:"index:i_parent_command_id;column:parent_command_id;type:bigint(20) unsigned" json:"parent_command_id"` // 父指令Id
	ExecuteCount    int       `gorm:"column:execute_count;type:int(10) unsigned;not null" json:"execute_count"`                             // 执行次数
	TimeoutTime     time.Time `gorm:"column:timeout_time;type:datetime" json:"timeout_time"`                                                // timeout time
	TimeoutPolicy   string    `gorm:"column:timeout_policy;type:varchar(16)" json:"timeout_policy"`                                         // timeout policy
	CreatedTime     int       `gorm:"column:created_time" json:"createdTime"`                                                               // 创建时间戳
	UpdatedTime     int       `gorm:"column:updated_time" json:"updatedTime"`                                                               // 更新时间戳
	DeletedTime     int       `gorm:"column:deleted_time" json:"deletedTime"`                                                               // 删除时间戳
	IsDel           int8      `gorm:"column:is_del" json:"isDel"`                                                                           // 是否删除0未删除 1已删除
}

//var db *gorm.DB = dao.IronicRdb

func (t *Command) TableName() string {
	return "command"
}

// AddCommand insert a new Command into database and returns
// last inserted Id on success.
func AddCommand(logger *log.Logger, m *Command) (id int64, err error) {
	//m.CreateTime = time.Now()
	//m.UpdateTime = time.Now()
	return dao.CreateAndGetId(logger, dao.IronicWdb, m)
}

// GetCommandById retrieves Command by Id. Returns error if
// Id doesn't exist
func GetCommandById(logger *log.Logger, id int64) (v *Command, err error) {
	v = &Command{}
	err = dao.Where(logger, dao.IronicRdb, "id = ? and is_del = 0", id).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

// GetCommandById retrieves Command by parentId. Returns error if
// Id doesn't exist
func GetCommandByParentId(logger *log.Logger, parentId int64) (v *Command, err error) {
	v = &Command{}
	err = dao.Where(logger, dao.IronicRdb, "parent_command_id = ? and is_del = 0", parentId).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

//TODO map[string]string gorm为什么不支持
func GetCommandCount(logger *log.Logger, query map[string]interface{}) (n int64, err error) {

	var db = dao.Model(logger, dao.IronicWdb, Command{})
	db, err = dao.WhereBuild(db, query)
	err = db.Count(&n).Error
	return
}

// GetAllCommand retrieves all Command matches certain condition. Returns empty list if
// no records exist
func GetAllCommand(logger *log.Logger, query map[string]interface{}) (ml []*Command, err error) {

	var db = dao.Model(logger, dao.IronicWdb, Command{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Find(&ml).Error
	return
}

// GetMultiCommand retrieves all Command matches certain condition. Returns empty list if
// no records exist
func GetMultiCommand(logger *log.Logger, query map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []*Command, err error) {

	var db = dao.Model(logger, dao.IronicWdb, Command{})
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

// UpdateCommand updates Command by Id and returns error if
// the record to be updated doesn't exist
func UpdateCommandById(logger *log.Logger, m *Command) (err error) {
	//m.UpdateTime = time.Now()
	err = dao.Model(logger, dao.IronicWdb, Command{}).Where("id = ?", m.ID).Updates(m).Error
	return
}

func DeleteCommandBySn(logger *log.Logger, sn string) (err error) {
	//m.UpdateTime = time.Now()
	err = dao.Model(logger, dao.IronicWdb, Command{}).Where("sn = ? and status <> ?", sn, "finish").Update("is_del", 1).Error
	return
}

func DeleteCommandByInstanceId(logger *log.Logger, instanceId string) (err error) {
	//m.UpdateTime = time.Now()
	err = dao.Model(logger, dao.IronicWdb, Command{}).Where("instance_id = ? and status <> ?", instanceId, "finish").Update("is_del", 1).Error
	return
}

func DeleteCommandByInstanceIdAndTask(logger *log.Logger, instanceId string, task string) (err error) {
	//m.UpdateTime = time.Now()
	err = dao.Model(logger, dao.IronicWdb, Command{}).Where("instance_id = ? and task = ? and is_del = ? and status <> ?", instanceId, task, 0, "finish").Update("is_del", 1).Error
	return
}
