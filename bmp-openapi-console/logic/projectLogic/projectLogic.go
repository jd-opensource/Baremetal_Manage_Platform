package projectLogic

import (
	"strings"
	"time"

	"coding.jd.com/aidc-bmp/bmp-openapi/constant"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/instanceDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/projectDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/sharingProjectDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/userDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/baseLogic"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/instanceLogic"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/monitorRuleLogic"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/userLogic"

	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
	responseTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/response"
	"coding.jd.com/aidc-bmp/bmp-openapi/util"
	log "coding.jd.com/aidc-bmp/bmp_log"
	namespacePrefix "git.jd.com/cps-golang/ironic-common/ironic/common/NamespacePrefixs"
	commonUtil "git.jd.com/cps-golang/ironic-common/ironic/util"
)

func ProjectEntity2Project(logger *log.Logger, o *projectDao.Project, shareproject []*sharingProjectDao.SharingProject) *responseTypes.Project {

	u, err := userLogic.GetUserById(logger, o.UserID)
	if err != nil {
		logger.Warnf("Project.UserID get invalid, project_id:%s, user_id:%s, err:%s", o.ProjectID, o.UserID, err.Error())
		return nil
	}

	tz := logger.GetPoint("timezone").(string)
	shares := []responseTypes.ShareProject{}
	owned := 2

	var cnt int64
	if logger.GetPoint("userId").(string) == o.UserID { //拥有者身份
		//拥有者
		owned = 1
		query := map[string]interface{}{
			"is_del":     0,
			"project_id": o.ProjectID,
		}
		cnt, err = instanceLogic.GetInstanceCount(logger, query)
		if err != nil {
			logger.Warnf("instanceLogic.GetInstanceCount error:%s", err.Error())
		}

		for _, v := range shareproject {
			shares = append(shares, responseTypes.ShareProject{
				ProjectName:       v.ProjectName,
				ProjectID:         v.ProjectID,
				OwnerUserID:       v.OwnerUserID,
				OwnerUserName:     v.OwnerUserName,
				SharedUserID:      v.SharedUserID,
				SharedUserName:    v.SharedUserName,
				CreatedTime:       util.TimestampToString(int64(v.CreatedTime), tz),
				SharedInstanceIDs: v.SharedInstanceIDs,
			})
		}
	} else { //协作者身份
		for _, v := range shareproject {
			if v.SharedUserID == logger.GetPoint("userId").(string) {
				if v.SharedInstanceIDs != "" {
					cnt = int64(len(strings.Split(v.SharedInstanceIDs, ",")))
				}
			}
		}
	}

	return &responseTypes.Project{
		ID:            o.ID,
		ProjectName:   o.ProjectName,
		Description:   o.Description,
		ProjectID:     o.ProjectID,
		InstanceCount: cnt,
		ShareProjects: shares,
		Owned:         owned,
		CreatedBy:     o.CreatedBy,
		OwnedBy:       u.UserName,
		CreatedTime:   util.TimestampToString(int64(o.CreatedTime), tz),
		UpdatedBy:     o.UpdatedBy,
		UpdatedTime:   util.TimestampToString(int64(o.UpdatedTime), tz),
	}
}

//只带owned标志位,不关心sharing具体内容
func ProjectListEntity2Project(logger *log.Logger, o *projectDao.AllProject) *responseTypes.Project {

	userId := logger.GetPoint("userId").(string)
	tz := logger.GetPoint("timezone").(string)
	query := map[string]interface{}{
		"is_del":     0,
		"project_id": o.ProjectID,
	}
	var cnt int64
	var err error
	var shareAll bool
	if o.Owned == 2 { //共享，考虑部分共享的情况
		s, err := sharingProjectDao.GetOneSharingProject(logger, o.ProjectID, "", userId)
		if err != nil {
			logger.Warnf("GetOneSharingProject error:%s", err.Error())
		}
		if s.SharedInstanceIDs == "" {
			cnt = 0
		} else {
			instanceids := strings.Split(s.SharedInstanceIDs, ",")
			cnt = int64(len(instanceids))
		}
	}

	if o.Owned == 1 || shareAll {
		cnt, err = instanceLogic.GetInstanceCount(logger, query)
		if err != nil {
			logger.Warnf("instanceLogic.GetInstanceCount error:%s", err.Error())
		}
	}

	shares := []responseTypes.ShareProject{}

	return &responseTypes.Project{
		ID:            o.ID,
		ProjectName:   o.ProjectName,
		ProjectID:     o.ProjectID,
		InstanceCount: cnt,
		ShareProjects: shares,
		Owned:         o.Owned,
		CreatedBy:     o.CreatedBy,
		CreatedTime:   util.TimestampToString(int64(o.CreatedTime), tz),
		UpdatedBy:     o.UpdatedBy,
		UpdatedTime:   util.TimestampToString(int64(o.UpdatedTime), tz),
	}
}

func GetProjectById(logger *log.Logger, ProjectId string) (*responseTypes.Project, error) {
	entity, err := projectDao.GetProjectById(logger, ProjectId)
	if err != nil {
		logger.Warn("GetProjectByUuid sql error:", ProjectId, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}
	shareEntity, err := sharingProjectDao.GetSharingsByProjectId(logger, ProjectId)
	if err != nil {
		logger.Warnf("GetProjectById.GetSharingsByProjectId error, project_id:%s, error:%s", ProjectId, err.Error())
	}
	return ProjectEntity2Project(logger, entity, shareEntity), nil
}

func CreateProject(logger *log.Logger, param *requestTypes.CreateProjectRequest, extraUserId string) (string, error) {

	var userId string
	if extraUserId == "" {
		userId = logger.GetPoint("userId").(string)
	} else {
		userId = extraUserId //兼容创建用户时创建默认项目，因为此时项目所属的userid不是当前登录用户的userid
	}
	// project_name不做唯一性限制
	// list, _ := projectDao.GetAllProject(logger, map[string]interface{}{
	// 	"project_name": param.ProjectName,
	// 	"user_id":      userId,
	// 	"is_del":       0,
	// })
	// if len(list) != 0 {
	// 	logger.Warn("projectName exist:", param.ProjectName)
	// 	panic(constant.BuildInvalidArgumentWithArgs("项目名称已存在", "projectName exist"))
	// }
	if param.IsDefault == 1 {
		list, _ := projectDao.GetAllProject(logger, map[string]interface{}{
			"is_default": 1,
			"user_id":    userId,
			"is_del":     0,
		})
		if len(list) != 0 {
			logger.Warn("默认项目 exist:", param.ProjectName)
			panic(constant.BuildInvalidArgumentWithArgs("默认项目只能有一个", "defaultProject only exist one"))
		}
	}
	ProjectEntity := &projectDao.Project{
		UserID:      userId,
		ProjectName: param.ProjectName,
		IsDefault:   param.IsDefault,
		IsSystem:    param.IsSystem,
		CreatedBy:   logger.GetPoint("username").(string),
		CreatedTime: int(time.Now().Unix()),
		UpdatedBy:   logger.GetPoint("username").(string),
		UpdatedTime: int(time.Now().Unix()),
	}
	ProjectEntity.ProjectID = commonUtil.GetRandString("project-", namespacePrefix.INSTANCE_ID_LENGTH, false, true, true)

	if _, err := projectDao.AddProject(logger, ProjectEntity); err != nil {
		logger.Warnf("CreateProject AddProject sql error, entity:%s, error:%s", util.ObjToJson(ProjectEntity), err.Error())
		return "", err
	}
	return ProjectEntity.ProjectID, nil
}

func ModifyProject(logger *log.Logger, param *requestTypes.ModifyProjectRequest, ProjectId string) error {
	userId := logger.GetPoint("userId").(string)
	p, err := projectDao.GetProjectById(logger, ProjectId)
	if err != nil {
		logger.Warn("GetProjectByUuid sql error:", ProjectId, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}
	if p.UserID != userId {
		panic(constant.PermissionDenyForObject)
	}

	// list, _ := projectDao.GetAllProject(logger, map[string]interface{}{
	// 	"project_name": param.ProjectName,
	// 	"user_id":      userId,
	// 	"is_del":       0,
	// })
	// if len(list) != 0 {
	// 	logger.Warn("projectName exist:", param.ProjectName)
	// 	panic(constant.BuildInvalidArgumentWithArgs("项目名称已存在", "projectName exist"))
	// }

	p.ProjectName = param.ProjectName
	p.UpdatedBy = logger.GetPoint("username").(string)
	p.UpdatedTime = int(time.Now().Unix())
	if err := projectDao.UpdateProjectById(logger, p); err != nil {
		logger.Warn("ModifyProject UpdateProjectById sql error:", ProjectId, err.Error())
		return err
	}

	q := map[string]interface{}{
		"is_del":     0,
		"project_id": p.ProjectID,
	}
	u := map[string]interface{}{
		"project_name": param.ProjectName,
	}
	if err := sharingProjectDao.UpdateMultiShareProjects(logger, q, u); err != nil {
		logger.Warn("ModifyProject UpdateMultiShareProjects sql error:", ProjectId, err.Error())
		return err
	}

	return nil
}

func ModifyProjectDescription(logger *log.Logger, param *requestTypes.ModifyProjectDescriptionRequest, ProjectId string) error {
	userId := logger.GetPoint("userId").(string)
	p, err := projectDao.GetProjectById(logger, ProjectId)
	if err != nil {
		logger.Warn("GetProjectByUuid sql error:", ProjectId, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}
	if p.UserID != userId {
		panic(constant.PermissionDenyForObject)
	}

	p.Description = param.Description
	p.UpdatedBy = logger.GetPoint("username").(string)
	p.UpdatedTime = int(time.Now().Unix())
	if err := projectDao.UpdateProjectById(logger, p); err != nil {
		logger.Warn("ModifyProjectDescription UpdateProjectById sql error:", ProjectId, err.Error())
		return err
	}
	return nil
}

/*
	删除项目时，有如下要求：
	1,默认项目也可以删除。如果删除的是默认项目，则自动选一个最早的项目作为默认项目
	2,每个用户至少保留一个项目。
	3,项目下如果有资源则不可删除
*/

func DeleteProject(logger *log.Logger, ProjectId string) error {
	userId := logger.GetPoint("userId").(string)
	query := map[string]interface{}{
		"user_id": userId,
		"is_del":  0,
	}
	entities, err := projectDao.GetAllProject(logger, query)
	if err != nil {
		panic(constant.INTERNAL_ERROR)
	}
	idx := -1
	for i, entity := range entities {
		if entity.ProjectID == ProjectId {
			idx = i
		}
	}
	if idx == -1 {
		panic(constant.PermissionDenyForObject)
	}
	if len(entities) < 2 {
		panic(constant.BuildCANCELLEDWithArgs("每个用户至少保留一个项目", "at least one project remained for user"))
	}

	instanceQuery := map[string]interface{}{
		"project_id": ProjectId,
		"is_del":     0,
	}
	instanceCnt, err := instanceDao.GetInstanceCount(logger, instanceQuery)
	if err != nil {
		panic(constant.INTERNAL_ERROR)
	}
	if instanceCnt > 0 {
		panic(constant.BuildCANCELLEDWithArgs("项目下有资源不可删除", "project contains resource, please release resource first"))
	}

	deleteEntity := entities[idx]
	defaultEntity := entities[0]
	if idx == 0 {
		defaultEntity = entities[1]
	}

	deleteEntity.IsDel = 1
	deleteEntity.UpdatedTime = int(time.Now().Unix())
	deleteEntity.DeletedTime = int(time.Now().Unix())
	if err := projectDao.UpdateProjectById(logger, deleteEntity); err != nil {
		logger.Warn("DeleteProject1 UpdateProjectById sql error:", ProjectId, err.Error())
		return err
	}
	defaultEntity.IsDefault = 1
	defaultEntity.UpdatedTime = int(time.Now().Unix())
	defaultEntity.DeletedTime = int(time.Now().Unix())
	if err := projectDao.UpdateProjectById(logger, defaultEntity); err != nil {
		logger.Warn("DeleteProject2 UpdateProjectById sql error:", ProjectId, err.Error())
		return err
	}

	//删除项目时，此项目的共享关系都删除
	if err := sharingProjectDao.DeleteSharingProjectByProjectId(logger, ProjectId); err != nil {
		logger.Warn("DeleteSharingProjectByProjectId sql error:", ProjectId, err.Error())
		return err
	}

	userEntity, err := userDao.GetUserById(logger, userId)
	if err != nil {
		logger.Warn("DeleteProject GetUserById error:", userId, err.Error())
		return err
	}
	userEntity.DefaultProjectID = defaultEntity.ProjectID
	if err := userDao.UpdateUserById(logger, userEntity); err != nil {
		logger.Warn("DeleteProject UpdateUser's defaultProjectById error:", userId, ProjectId, err.Error())
		return err
	}

	return nil
}

func GetProjectList(logger *log.Logger, param requestTypes.QueryProjectsRequest, p util.Pageable) ([]*responseTypes.Project, int64, error) {
	userId := logger.GetPoint("userId").(string)
	offset, limit := p.Pageable2offset()
	query := map[string]interface{}{
		"user_id": userId,
		"is_del":  0,
	}
	if param.ProjectName != "" {
		query["project_name"] = param.ProjectName
	}
	//新增按用户名查找拥有的项目列表
	if param.OwnerName != "" {
		ownerUser, _ := userLogic.GetUserByName(logger, &requestTypes.GetUserByNameRequest{
			UserName: param.OwnerName,
		})
		if ownerUser != nil {
			userId = ownerUser.UserID
			param.Owned = 1
		}
	}
	//新增按被转移用户的用户名查找它拥有的项目列表
	if param.SharerName != "" {
		sharerUser, _ := userLogic.GetUserByName(logger, &requestTypes.GetUserByNameRequest{
			UserName: param.SharerName,
		})
		if sharerUser != nil {
			userId = sharerUser.UserID
			param.Owned = 1
		}
	}

	count, err := projectDao.GetProjectAndSharedCount(logger, userId, param.Owned)
	if err != nil {
		logger.Warnf("GetProjectCount dao error, query:%s, error:%s", util.ObjToJson(query), err.Error())
		return nil, 0, err
	}
	var entityList []*projectDao.AllProject

	var sortby []string
	var order []string
	if param.OrderByCreatetime == "desc" {
		sortby = []string{"c.created_time"}
		order = []string{"desc"}
	} else {
		sortby = []string{"c.created_time"}
		order = []string{"asc"}
	}

	if param.IsAll == baseLogic.IS_ALL {
		entityList, err = projectDao.GetMultiProjectIncludeShare(logger, userId, param.Owned, sortby, order, 0, 0)
	} else {
		entityList, err = projectDao.GetMultiProjectIncludeShare(logger, userId, param.Owned, sortby, order, offset, limit)
	}

	if err != nil {
		logger.Warnf("GetMultiProjectIncludeShare sql error， user_id:%s, error:%s", userId, err.Error())
		return nil, 0, err
	}
	Projects := []*responseTypes.Project{}
	var defaultProject *responseTypes.Project
	for _, entity := range entityList {
		// shareEntity, err := sharingProjectDao.GetSharingsByProjectId(logger, entity.ProjectID)
		if err != nil {
			logger.Warnf("GetProjectList.GetSharingsByProjectId error, project_id:%s, error:%s", entity.ProjectID, err.Error())
		}
		if entity.IsDefault == 1 {
			defaultProject = ProjectListEntity2Project(logger, entity)
		} else {
			Projects = append(Projects, ProjectListEntity2Project(logger, entity))
		}
	}
	if defaultProject == nil {
		return Projects, count, nil
	}

	return append([]*responseTypes.Project{defaultProject}, Projects...), count, nil
}

func AddShareProject(logger *log.Logger, projectID string, param *requestTypes.ShareProjectRequest) error {
	project, err := projectDao.GetProjectById(logger, projectID)
	if err != nil {
		logger.Warnf("addShareProject.GetProjectById error, project_id:%s, error:%s", projectID, err.Error())
		panic(constant.BuildCANCELLEDWithArgs("项目不存在", "project not found"))
	}
	owner, err := userDao.GetUserById(logger, param.OwnerID)
	if err != nil {
		logger.Warnf("addShareProject.GetUserById error, owner_id:%s, error:%s", param.OwnerID, err.Error())
		panic(constant.BuildCANCELLEDWithArgs("用户不存在", "user not found"))
	}
	sharer, err := userDao.GetUserById(logger, param.SharerID)
	if err != nil {
		logger.Warnf("addShareProject.GetUserById error, sharer_id:%s, error:%s", param.OwnerID, err.Error())
		panic(constant.BuildCANCELLEDWithArgs("用户不存在", "user not found"))
	}

	if project.UserID == param.SharerID {
		panic(constant.BuildCANCELLEDWithArgs("项目不能分享给自己", "project can not shared to the project owner"))
	}

	//只有分享给控制台用户
	if sharer.RoleID == baseLogic.ROLE_OPERATOR_UUID {
		panic(constant.BuildCANCELLEDWithArgs("目标用户没有控制台访问权限", "target user access denied for console"))
	}

	q1 := map[string]interface{}{
		"is_del":     0,
		"project_id": projectID,
	}
	instanceEntities, err := instanceDao.GetAllInstance(logger, q1)
	if err != nil {
		logger.Warnf("AddShareProject.GetAllInstance error, query.project_id:%s, error:%s", projectID, err.Error())
	}

	nowInstanceIds := make([]string, len(instanceEntities))
	for idx, instanceEntity := range instanceEntities {
		nowInstanceIds[idx] = instanceEntity.InstanceID
	}

	if param.InstanceIDs == "all" {

		param.InstanceIDs = strings.Join(nowInstanceIds, ",")
	} else if param.InstanceIDs != "" { //校验instanceid合法性
		instanceIds := strings.Split(param.InstanceIDs, ",")

		if len(instanceIds) > len(nowInstanceIds) {
			logger.Warnf("实例id参数校验错误, instanceIds:%v", instanceIds)
			panic(constant.BuildCANCELLEDWithArgs("实例id参数校验错误", "instanceId param error"))
		}
		for _, v := range instanceIds {
			if !util.InArray(v, nowInstanceIds) {
				logger.Warnf("实例id参数校验错误, instanceId:%s", v)
				panic(constant.BuildCANCELLEDWithArgs("实例id参数校验错误", "instanceId param error"))
			}
		}
	}

	spEntity, err := sharingProjectDao.GetOneSharingProject(logger, projectID, param.OwnerID, param.SharerID)

	if spEntity == nil { //新增
		sp := &sharingProjectDao.SharingProject{
			ProjectID:      projectID,
			OwnerUserID:    param.OwnerID,
			SharedUserID:   param.SharerID,
			OwnerUserName:  owner.UserName,
			SharedUserName: sharer.UserName,
			// Premission:     param.Premisson, //本版本只支持write
			Premission:        "write", //本版本只支持write
			ProjectName:       project.ProjectName,
			CreatedBy:         logger.GetPoint("username").(string),
			CreatedTime:       int(time.Now().Unix()),
			UpdatedBy:         logger.GetPoint("username").(string),
			UpdatedTime:       int(time.Now().Unix()),
			SharedInstanceIDs: param.InstanceIDs,
		}
		_, err = sharingProjectDao.AddSharingProject(logger, sp)
		if err != nil {
			logger.Warnf("addShareProject.AddSharingProject error, project_id:%s, owner_id:%s, sharer_id:%s, error:%s", projectID, param.OwnerID, param.SharerID, err.Error())
			panic(constant.INTERNAL_ERROR)
		}
	} else { //修改
		q := map[string]interface{}{
			"is_del":         0,
			"project_id":     projectID,
			"owner_user_id":  param.OwnerID,
			"shared_user_id": param.SharerID,
		}
		u := map[string]interface{}{
			"shared_instance_ids": param.InstanceIDs,
		}
		err = sharingProjectDao.UpdateShareProjectByMap(logger, q, u)
		if err != nil {
			logger.Warnf("addShareProject.UpdateShareProjectByMap error, project_id:%s, owner_id:%s, sharer_id:%s, error:%s", projectID, param.OwnerID, param.SharerID, err.Error())
			panic(constant.INTERNAL_ERROR)
		}
	}

	return nil
}

func CancelShareProject(logger *log.Logger, projectID string, param *requestTypes.CalcelShareProjectRequest) (err error) {

	sp, err := sharingProjectDao.GetOneSharingProject(logger, projectID, param.OwnerID, param.SharerID)
	if err != nil {
		logger.Warnf("shared project not found, projectID:%s, OwnerID:%s, SharerID:%s, error:%s", projectID, param.OwnerID, param.SharerID, err.Error())
		panic(constant.BuildCANCELLEDWithArgs("共享项目不存在", "shared project not found"))
	}
	err = sharingProjectDao.DeleteOneSharingProject(logger, sp)
	if err != nil {
		logger.Warnf("CreateShareProject.DeleteSharingProject error, project_id:%s, owner_id:%s, sharer_id:%s, error:%s", projectID, param.OwnerID, param.SharerID, err.Error())
		panic(constant.INTERNAL_ERROR)
	}
	return nil
}

func DescribeShareProject(logger *log.Logger, projectID string) ([]*sharingProjectDao.SharingProject, error) {
	q := map[string]interface{}{
		"is_del":     0,
		"project_id": projectID,
	}

	return sharingProjectDao.GetAllShareProjects(logger, q)

}

//项目转移时，项目下面的规则全部失效
func MoveProject(logger *log.Logger, projectID string, param *requestTypes.MoveProjectRequest) error {

	project, err := projectDao.GetProjectById(logger, projectID)
	if err != nil {
		logger.Warnf("CreateShareProject.GetProjectById error, project_id:%s, error:%s", projectID, err.Error())
		panic(constant.BuildCANCELLEDWithArgs("项目不存在", "project not found"))
	}
	owner, err := userDao.GetUserById(logger, param.OwnerID)
	if err != nil {
		logger.Warnf("CreateShareProject.GetUserById error, owner_id:%s, error:%s", param.OwnerID, err.Error())
		panic(constant.BuildCANCELLEDWithArgs("用户不存在", "user not found"))
	}
	mover, err := userDao.GetUserById(logger, param.MoverID)
	if err != nil {
		logger.Warnf("CreateShareProject.GetUserById error, sharer_id:%s, error:%s", param.OwnerID, err.Error())
		panic(constant.BuildCANCELLEDWithArgs("用户不存在", "user not found"))
	}

	if project.UserID == param.MoverID {
		panic(constant.BuildCANCELLEDWithArgs("项目不能转移给自己", "project can not shared to the project owner"))
	}

	//只有转移给控制台用户
	if mover.RoleID == baseLogic.ROLE_OPERATOR_UUID {
		panic(constant.BuildCANCELLEDWithArgs("目标用户没有控制台访问权限", "target user access denied for console"))
	}

	//如果只有一个项目，禁止转移
	projects, _ := projectDao.GetAllProjectByUserId(logger, param.OwnerID)
	if len(projects) == 0 {
		panic(constant.BuildCANCELLEDWithArgs("目标用户下无项目", "no project for current user"))
	}
	if len(projects) == 1 {
		panic(constant.BuildCANCELLEDWithArgs("目标用户只此一个项目，不支持转移", "operation denied, only one project for current user"))
	}

	sp, _ := sharingProjectDao.GetOneSharingProject(logger, projectID, param.OwnerID, param.MoverID)
	if sp != nil {
		//TODO 需要删除共享信息
		sp.IsDel = 1
		err = sharingProjectDao.DeleteOneSharingProject(logger, sp)
		if err != nil {
			logger.Warnf("CreateShareProject.DeleteOneSharingProject error, project_id:%s, owner_id:%s, sharer_id:%s, error:%s", projectID, param.OwnerID, param.MoverID, err.Error())
			panic(constant.INTERNAL_ERROR)
		}
	}

	// 更改该project的其他已有共享信息
	q := map[string]interface{}{
		"is_del":     0,
		"project_id": projectID,
	}
	u := map[string]interface{}{
		"owner_user_id":   param.MoverID,
		"owner_user_name": mover.UserName,
	}
	if err := sharingProjectDao.UpdateMultiShareProjects(logger, q, u); err != nil {
		logger.Warnf("CreateShareProject.UpdateMultiShareProjects error, project_id:%s, owner_id:%s, sharer_id:%s, error:%s", projectID, param.OwnerID, param.MoverID, err.Error())
		panic(constant.INTERNAL_ERROR)
	}

	//project下实例的user_id更新
	iq := map[string]interface{}{
		"is_del":     0,
		"project_id": projectID,
	}
	iu := map[string]interface{}{
		"user_id": param.MoverID,
	}
	if err := instanceDao.UpdateInstances(logger, iq, iu); err != nil {
		logger.Warnf("moveProject.UpdateInstances error, project_id:%s, error:%s", projectID, err.Error())
	}

	//项目改成非默认项目
	u = map[string]interface{}{
		"user_id":    param.MoverID,
		"is_default": 0,
	}
	if err := projectDao.UpdateProjectByMap(logger, q, u); err != nil {
		logger.Warnf("moveProject.UpdateProjectByMap error, project_id:%s, owner_id:%s, mover_id:%s, error:%s", projectID, param.OwnerID, param.MoverID, err.Error())
		panic(constant.INTERNAL_ERROR)
	}
	if project.IsDefault == 1 { //如果是默认项目被转移的话，需要选一个新的默认项目
		newDefaultProject := projects[0]
		if newDefaultProject.ProjectID == projectID {
			newDefaultProject = projects[1]
		}
		newDefaultProject.IsDefault = 1
		if err := projectDao.UpdateProjectById(logger, newDefaultProject); err != nil {
			logger.Warnf("moveProject.UpdatedefaultProject error, project_id:%s, owner_id:%s, mover_id:%s, error:%s, new_default_project_id:%s", projectID, param.OwnerID, param.MoverID, err.Error(), newDefaultProject.ProjectID)
			// panic(constant.INTERNAL_ERROR)
		}
		owner.DefaultProjectID = newDefaultProject.ProjectID
		if err := userDao.UpdateUserById(logger, owner); err != nil {
			logger.Warnf("moveProject.UpdatedefaultProjectfor user error, project_id:%s, owner_id:%s, mover_id:%s, error:%s, new_default_project_id:%s", projectID, param.OwnerID, param.MoverID, err.Error(), newDefaultProject.ProjectID)
			// panic(constant.INTERNAL_ERROR)
		}
	}

	if err := monitorRuleLogic.DelMonitorRulesByProjectID(logger, projectID); err != nil {
		logger.Warnf("moveProject.DelMonitorRulesByProjectID error, project_id:%s, owner_id:%s, mover_id:%s, error:%s", projectID, param.OwnerID, param.MoverID, err.Error())
	}

	return nil
}

//项目部分转移时，规则下面的被转移实例被摘除
func MoveInstances(logger *log.Logger, param *requestTypes.MoveInstancesRequest) error {

	ownerProject, err := projectDao.GetProjectById(logger, param.OwnerProjectID)
	if err != nil {
		logger.Warnf("MoveInstances.GetProjectById error, project_id:%s, error:%s", param.OwnerProjectID, err.Error())
		panic(constant.BuildCANCELLEDWithArgs("项目不存在", "project not found"))
	}
	_, err = userDao.GetUserById(logger, param.OwnerID)
	if err != nil {
		logger.Warnf("MoveInstances.GetUserById error, owner_id:%s, error:%s", param.OwnerID, err.Error())
		panic(constant.BuildCANCELLEDWithArgs("用户不存在", "user not found"))
	}
	if ownerProject.UserID != param.OwnerID {
		logger.Warnf("MoveInstances ower_id not match with owner_project_id, owner_id:%s, owner_project_id:%s", param.OwnerID, param.OwnerProjectID)
		panic(constant.BuildCANCELLEDWithArgs("owner用户和项目不匹配", "owner user not match with project"))
	}

	moverProject, err := projectDao.GetProjectById(logger, param.MoverProjectID)
	if err != nil {
		logger.Warnf("MoveInstances.GetProjectById error, project_id:%s, error:%s", param.OwnerProjectID, err.Error())
		panic(constant.BuildCANCELLEDWithArgs("项目不存在", "project not found"))
	}
	mover, err := userDao.GetUserById(logger, param.MoverID)
	if err != nil {
		logger.Warnf("CreateShareProject.GetUserById error, sharer_id:%s, error:%s", param.OwnerID, err.Error())
		panic(constant.BuildCANCELLEDWithArgs("用户不存在", "user not found"))
	}
	if moverProject.UserID != param.MoverID {
		logger.Warnf("MoveInstances mover not match with mover_project_id, mover_id:%s, mover_project_id:%s", param.MoverID, param.MoverProjectID)
		panic(constant.BuildCANCELLEDWithArgs("mover用户和项目不匹配", "move user not match with project"))
	}

	// if moverProject.UserID == param.MoverID {
	// 	panic(constant.BuildCANCELLEDWithArgs("项目不能转移给自己", "project can not shared to the project owner"))
	// }

	//只有转移给控制台用户
	if mover.RoleID == baseLogic.ROLE_OPERATOR_UUID {
		panic(constant.BuildCANCELLEDWithArgs("目标用户没有控制台访问权限", "target user access denied for console"))
	}

	moveInstanceIds := []string{}
	if param.InstanceIDs != "" {
		moveInstanceIds = strings.Split(param.InstanceIDs, ",")
	}
	//直接将instance改成新的project_id
	q := map[string]interface{}{
		"is_del":         0,
		"instance_id.in": moveInstanceIds,
	}
	updates := map[string]interface{}{
		"project_id": param.MoverProjectID,
		"user_id":    param.MoverID,
	}
	if err := instanceDao.UpdateInstances(logger, q, updates); err != nil {
		logger.Warnf("moveInstance.UpdateInstances error, instance_ids:%s, project_id:%s, user_id:%s, error:%s", param.InstanceIDs, param.MoverProjectID, param.MoverID, err.Error())
		panic(constant.INTERNAL_ERROR)
	}

	//param.InstanceIDs 之前如果被共享给其他用户，则共享关系也要改变，实例的原共享关系废除，新共享关系建立
	q2 := map[string]interface{}{
		"project_id":    param.OwnerProjectID,
		"owner_user_id": param.OwnerID,
	}
	shareProjectEntities, err := sharingProjectDao.GetAllShareProjects(logger, q2)
	if err != nil {
		logger.Warnf("MoveInstances.GetAllShareProjects error, project_id:%s, owner_user_id:%s, error:%s", param.OwnerProjectID, param.OwnerID, err.Error())
	}
	for _, spEntity := range shareProjectEntities {
		shareInstanceIds := []string{}

		if spEntity.SharedInstanceIDs != "" {
			shareInstanceIds = strings.Split(spEntity.SharedInstanceIDs, ",")
		}

		innerInstanceIds := util.IntersectArray(shareInstanceIds, moveInstanceIds) //交集
		//innerInstanceIds这部分从老的shareproject移除
		remainShareInstancesIds := util.DifferenceArray(shareInstanceIds, innerInstanceIds)
		spEntity.SharedInstanceIDs = strings.Join(remainShareInstancesIds, ",")
		//spEntity.SharedInstanceIDs可能为空，所以用UpdateShareProjectByMap，不用UpdateShareProjectById

		tq := map[string]interface{}{
			"project_id":     spEntity.ProjectID,
			"shared_user_id": spEntity.SharedUserID,
			"is_del":         0,
		}
		tu := map[string]interface{}{
			"shared_instance_ids": spEntity.SharedInstanceIDs,
		}
		if err := sharingProjectDao.UpdateShareProjectByMap(logger, tq, tu); err != nil {
			logger.Warnf("moveInstance.UpdateShareProjectByMap.ShareInstancesIds error, project_id:%s, owner_user_id:%s, shareInstanceIds:%v, error:%s", param.OwnerProjectID, param.OwnerID, spEntity.SharedInstanceIDs, err.Error())
		}

		//innerInstanceIds这部分的ownerid要变更，需更新/插入shareproject
		if param.MoverID == spEntity.SharedUserID { //之前的共享人和现在的拥有者是同一个人的话，不用再添加共享
			continue
		}
		sp1, _ := sharingProjectDao.GetOneSharingProject(logger, param.MoverProjectID, param.MoverID, spEntity.SharedUserID)
		if sp1 != nil { //更新
			shareInstanceIds := util.UnionArray(innerInstanceIds, strings.Split(sp1.SharedInstanceIDs, ","))
			sp1.SharedInstanceIDs = strings.TrimSpace(strings.Join(shareInstanceIds, ","))
			tq := map[string]interface{}{
				"project_id":     sp1.ProjectID,
				"shared_user_id": sp1.SharedUserID,
				"is_del":         0,
			}
			tu := map[string]interface{}{
				"shared_instance_ids": sp1.SharedInstanceIDs,
			}
			if err := sharingProjectDao.UpdateShareProjectByMap(logger, tq, tu); err != nil {
				logger.Warnf("moveInstance.UpdateShareProject.UpdateShareProjectByMap2 error, project_id:%s, owner_user_id:%s, shareInstanceIds:%v, error:%s", param.MoverProjectID, param.MoverID, sp1.SharedInstanceIDs, err.Error())
			}

		} else { //插入
			sp := &sharingProjectDao.SharingProject{
				ProjectID:      param.MoverProjectID,
				OwnerUserID:    param.MoverID,
				SharedUserID:   spEntity.SharedUserID,
				OwnerUserName:  mover.UserName,
				SharedUserName: spEntity.SharedUserName,
				// Premission:     param.Premisson, //本版本只支持write
				Premission:        "write", //本版本只支持write
				ProjectName:       spEntity.ProjectName,
				CreatedBy:         logger.GetPoint("username").(string),
				CreatedTime:       int(time.Now().Unix()),
				UpdatedBy:         logger.GetPoint("username").(string),
				UpdatedTime:       int(time.Now().Unix()),
				SharedInstanceIDs: strings.TrimSpace(strings.Join(innerInstanceIds, ",")),
			}
			_, err = sharingProjectDao.AddSharingProject(logger, sp)
			if err != nil {
				logger.Warnf("moveInstance.AddSharingProject error, project_id:%s, owner_id:%s, sharer_id:%s, shareInstanceids:%v, error:%s", param.MoverProjectID, param.MoverID, spEntity.SharedUserID, innerInstanceIds, err.Error())
			}
		}
	}

	if err := monitorRuleLogic.DelRMonitorRuleInstances(logger, param.OwnerProjectID, param.InstanceIDs); err != nil {
		logger.Warnf("moveInstance.DelRMonitorRuleInstances error, project_id:%s, owner_id:%s, mover_id:%s, instanceids:%v, error:%s", param.MoverProjectID, param.OwnerID, param.MoverID, param.InstanceIDs, err.Error())
	}

	return nil
}
