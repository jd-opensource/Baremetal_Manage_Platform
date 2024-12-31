package commandDao

import (
	"fmt"
	"strings"
	"time"

	"coding.jd.com/aidc-bmp/bmp-scheduler/dao"
	log "git.jd.com/cps-golang/log"
)

// Command 指令
type Command struct {
	ID              int       `gorm:"primaryKey;column:id" json:"-"`                   // ID
	RequestID       string    `gorm:"column:request_id" json:"requestId"`              // 请求ID
	Sn              string    `gorm:"column:sn" json:"sn"`                             // 设备SN
	InstanceID      string    `gorm:"column:instance_id" json:"instanceId"`            // 实例Id
	Action          string    `gorm:"column:action" json:"action"`                     // 操作
	Type            string    `gorm:"column:type" json:"type"`                         // 操作类型：agent, driver, network
	Status          string    `gorm:"column:status" json:"status"`                     // 状态: wait,running,finish,error
	ParentCommandID int64     `gorm:"column:parent_command_id" json:"parentCommandId"` // 父指令Id
	ExecuteCount    int       `gorm:"column:execute_count" json:"executeCount"`        // 执行次数
	TimeoutTime     time.Time `gorm:"column:timeout_time" json:"timeoutTime"`          // timeout time
	TimeoutPolicy   string    `gorm:"column:timeout_policy" json:"timeoutPolicy"`      // timeout policy
	Task            string    `gorm:"column:task" json:"task"`                         // 任务名称
	CreatedBy       string    `gorm:"column:created_by" json:"createdBy"`              // 创建者
	UpdatedBy       string    `gorm:"column:updated_by" json:"updatedBy"`              // 更新者
	CreatedTime     int       `gorm:"column:created_time" json:"createdTime"`          // 创建时间戳
	UpdatedTime     int       `gorm:"column:updated_time" json:"updatedTime"`          // 更新时间戳
	DeletedTime     int       `gorm:"column:deleted_time" json:"deletedTime"`          // 删除时间戳
	IsDel           int8      `gorm:"column:is_del" json:"isDel"`                      // 是否删除0未删除 1已删除
}

//var db *gorm.DB = dao.IronicRdb

func (t *Command) TableName() string {
	return "command"
}

// AddCommand insert a new Command into database and returns
// last inserted Id on success.
func AddCommand(logger *log.Logger, m *Command) (id int64, err error) {
	m.CreatedTime = int(time.Now().Unix())
	m.UpdatedTime = int(time.Now().Unix())
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

func GetCommandCount(logger *log.Logger, query map[string]interface{}) (n int64, err error) {

	db, err := dao.WhereBuild(dao.Model(logger, dao.IronicRdb, Command{}), query)
	if err != nil {
		return n, err
	}
	err = db.Count(&n).Error
	return
}

func GetAllCommand(logger *log.Logger, query map[string]interface{}) (ml []*Command, err error) {
	g := dao.Model(logger, dao.IronicRdb, Command{})
	g, err = dao.WhereBuild(g, query)
	if err != nil {
		return nil, err
	}
	err = g.Find(&ml).Error
	return

}

// GetMultiCommand retrieves all Command matches certain condition. Returns empty list if
// no records exist
func GetMultiCommand(logger *log.Logger, query map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []*Command, err error) {

	var db = dao.Model(logger, dao.IronicRdb, Command{})
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

func GetOneCommand(logger *log.Logger, query map[string]interface{}) (l *Command, err error) {
	l = &Command{}
	var db = dao.Model(logger, dao.IronicRdb, Command{})
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
func UpdateCommandById(logger *log.Logger, m *Command) (err error) {

	m.UpdatedTime = int(time.Now().Unix())
	err = dao.Model(logger, dao.IronicWdb, Command{}).Where("id = ?", m.ID).Updates(m).Error
	return
}

func CountBySnAndStatus(logger *log.Logger, sn, status string) (int64, error) {
	param := map[string]interface{}{
		"sn":     sn,
		"status": status,
		"is_del": 0,
	}
	return GetCommandCount(logger, param)
}

func GetFirstCommand(logger *log.Logger, query map[string]interface{}) (l *Command, err error) {
	l = &Command{}
	var db = dao.Model(logger, dao.IronicRdb, Command{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Order("id asc").Limit(1).Offset(0).Take(l).Error
	if err != nil {
		return nil, err
	}
	return l, nil

}

func QueryBySnAndStatus(logger *log.Logger, sn, status string) ([]*Command, error) {
	param := map[string]interface{}{
		"sn":     sn,
		"status": status,
		"is_del": 0,
	}
	return GetMultiCommand(logger, param, nil, []string{"id"}, []string{"asc"}, 0, 0)
}

func UpdateMultiCommands(logger *log.Logger, query map[string]interface{}, updates map[string]interface{}) (err error) {

	updates["updated_time"] = time.Now().Unix()
	var db = dao.Model(logger, dao.IronicWdb, Command{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return err
	}
	err = db.Updates(updates).Error
	return
}
