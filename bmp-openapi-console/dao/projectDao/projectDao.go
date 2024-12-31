package projectDao

import (
	"fmt"
	"strings"

	"coding.jd.com/aidc-bmp/bmp-openapi/dao"
	log "coding.jd.com/aidc-bmp/bmp_log"
)

// Project Project
type Project struct {
	ID          int    `gorm:"primaryKey;column:id" json:"-"`
	ProjectID   string `gorm:"column:project_id" json:"projectId"`     // 项目uuid
	ProjectName string `gorm:"column:project_name" json:"projectName"` // 项目名称
	Description string `gorm:"column:description" json:"description"`  // 项目描述
	IsDefault   int8   `gorm:"column:is_default" json:"isDefault"`     // 是否默认项目0否 1是
	IsSystem    int8   `gorm:"column:is_system" json:"isSystem"`       // 是否系统项目 0否 1是
	UserID      string `gorm:"column:user_id" json:"userId"`           // 项目拥有者用户id
	CreatedBy   string `gorm:"column:created_by" json:"createdBy"`     // 创建者
	UpdatedBy   string `gorm:"column:updated_by" json:"updatedBy"`     // 更新者
	CreatedTime int    `gorm:"column:created_time" json:"createdTime"` // 创建时间戳
	UpdatedTime int    `gorm:"column:updated_time" json:"updatedTime"` // 更新时间戳
	DeletedTime int    `gorm:"column:deleted_time" json:"deletedTime"` // 删除时间戳
	IsDel       int8   `gorm:"column:is_del" json:"isDel"`             // 是否删除0未删除 1已删除
}

func (t *Project) TableName() string {
	return "project"
}

// AddProject insert a new Project into database and returns
// last inserted Id on success.
func AddProject(logger *log.Logger, m *Project) (id int64, err error) {

	//m.CreateTime = time.Now()
	//m.UpdateTime = time.Now()
	return dao.CreateAndGetId(logger, dao.IronicWdb, m)
}

// GetProjectById retrieves Project by Id. Returns error if
// Id doesn't exist
func GetProjectById(logger *log.Logger, ProjectId string) (v *Project, err error) {
	v = &Project{}
	err = dao.Where(logger, dao.IronicRdb, "project_id = ? and is_del = 0", ProjectId).Take(v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

// GetProjectByUuid retrieves Project by Uuid. Returns error if
// Id doesn't exist
func GetProjectCount(logger *log.Logger, query map[string]interface{}) (n int64, err error) {
	var db = dao.Model(logger, dao.IronicRdb, Project{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return 0, err
	}
	err = db.Count(&n).Error
	return
}

type AllProject struct {
	Project
	//1表示拥有，2表示共享
	Owned int `gorm:"column:owned" json:"owned"`
}

// GetProjectAndSharedCount retrieves Project by Uuid. Returns error if
// Id doesn't exist
func GetProjectAndSharedCount(logger *log.Logger, user_id string, owned int) (n int64, err error) {
	leftsql := fmt.Sprintf(`select a.project_id, a.project_name, a.user_id, a.created_time, 1 as owned from project a where a.is_del = 0 and a.user_id = '%s'`, user_id)
	rightsql := fmt.Sprintf(`select b.project_id, b.project_name, b.shared_user_id as user_id, b.created_time, 2 as owned from sharing_project b where b.is_del = 0 and b.shared_user_id = '%s'`, user_id)
	var sql string
	if owned == 1 { //只查有拥有权的项目
		sql = fmt.Sprintf(`select count(*) from (%s) c`, leftsql)
	} else if owned == 2 { //只查有共享权的项目
		sql = fmt.Sprintf(`select count(*) from (%s) c`, rightsql)
	} else { //查有拥有权和共享权的项目
		sql = fmt.Sprintf(`select count(*) from (%s union %s) c`, leftsql, rightsql)
	}
	err = dao.Raw(logger, dao.IronicRdb, sql).Count(&n).Error
	return
}

// GetAllProject retrieves all Project matches certain condition. Returns empty list if
// no records exist
func GetAllProject(logger *log.Logger, query map[string]interface{}) (ml []*Project, err error) {

	var db = dao.Model(logger, dao.IronicRdb, Project{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Order("id ASC").Find(&ml).Error
	return
}

func GetAllProjectByUserId(logger *log.Logger, userId string) (ml []*Project, err error) {

	query := map[string]interface{}{
		"user_id": userId,
		"is_del":  0,
	}
	var db = dao.Model(logger, dao.IronicRdb, Project{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return nil, err
	}
	err = db.Order("id ASC").Find(&ml).Error
	return
}

func GetAllProjectIncludeShare(logger *log.Logger, user_id string, owned int) (ml []*AllProject, err error) {

	leftsql := fmt.Sprintf(`select a.is_default, a.project_id, a.project_name, a.user_id, a.created_time, a.updated_time, 1 as owned from project a where a.is_del = 0 and a.user_id = '%s'`, user_id)
	rightsql := fmt.Sprintf(`select 0 as is_default, b.project_id, b.project_name, b.shared_user_id as user_id, m.created_time as created_time, m.updated_time as updated_time, 2 as owned from sharing_project b join project m on b.project_id = m.project_id where b.is_del = 0 and b.shared_user_id = '%s'`, user_id)
	var sql string
	if owned == 1 { //只查有拥有权的项目
		sql = fmt.Sprintf(`select c.* from (%s) c`, leftsql)
	} else if owned == 2 { //只查有共享权的项目,暂时还用不到
		sql = fmt.Sprintf(`select c.* from (%s) c`, rightsql)
	} else { //查有拥有权和共享权的项目
		sql = fmt.Sprintf(`select c.* from (%s union %s) c order by c.created_time desc`, leftsql, rightsql)
	}
	err = dao.Raw(logger, dao.IronicRdb, sql).Scan(&ml).Error
	return
}

// GetMultiProject retrieves all Project matches certain condition. Returns empty list if
// no records exist
func GetMultiProject(logger *log.Logger, query map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []*Project, err error) {

	var db = dao.Model(logger, dao.IronicRdb, Project{})
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

func GetMultiProjectIncludeShare(logger *log.Logger, user_id string, owned int, sortby []string, order []string,
	offset int64, limit int64) (ml []*AllProject, err error) {

	leftsql := fmt.Sprintf(`select a.is_default, a.project_id, a.project_name, a.user_id, a.created_time, a.updated_time, 1 as owned from project a where a.is_del = 0 and a.user_id = '%s'`, user_id)
	rightsql := fmt.Sprintf(`select 0 as is_default, b.project_id, b.project_name, b.shared_user_id as user_id, m.created_time as created_time, m.updated_time as updated_time, 2 as owned from sharing_project b join project m on b.project_id = m.project_id where b.is_del = 0 and b.shared_user_id = '%s'`, user_id)
	var sql string
	if owned == 1 { //只查有拥有权的项目
		sql = fmt.Sprintf(`select c.* from (%s) c`, leftsql)
	} else if owned == 2 { //只查有共享权的项目,暂时还用不到
		sql = fmt.Sprintf(`select c.* from (%s) c`, rightsql)
	} else { //查有拥有权和共享权的项目
		sql = fmt.Sprintf(`select c.* from (%s union %s) c`, leftsql, rightsql)
	}

	db := dao.Raw(logger, dao.IronicRdb, sql)

	// if len(sortby) == 0 {
	// 	db = db.Order("c.created_time desc")
	// } else {
	orderConditions := []string{}
	for i := 0; i < len(sortby); i++ {
		orderConditions = append(orderConditions, fmt.Sprintf("%s %s", sortby[i], order[i]))
	}
	db = db.Order(strings.Join(orderConditions, ","))
	// }
	if offset == 0 && limit == 0 {
		err = db.Find(&ml).Error
		return
	}
	err = db.Offset(offset).Limit(limit).Find(&ml).Error
	return
}

// UpdateProject updates Project by Id and returns error if
// the record to be updated doesn't exist
func UpdateProjectById(logger *log.Logger, m *Project) (err error) {

	//m.UpdateTime = time.Now()
	return dao.Model(logger, dao.IronicWdb, Project{}).Where("project_id = ?", m.ProjectID).Updates(m).Error
}

func UpdateProjectByMap(logger *log.Logger, query map[string]interface{}, updates map[string]interface{}) (err error) {

	var db = dao.Model(logger, dao.IronicWdb, Project{})
	db, err = dao.WhereBuild(db, query)
	if err != nil {
		return err
	}
	err = db.Updates(updates).Error
	return
}

func UpdateUserDefaultProject(logger *log.Logger, userId, projectId string) error {
	err := dao.IronicWdb.Exec("update project set is_default = 0 where is_del = 0 and user_id = ?", userId).Error
	if err != nil {
		return err
	}

	err = dao.IronicWdb.Exec("update project set is_default = 1 where project_id = ?", projectId).Error
	if err != nil {
		return err
	}
	return nil
}
