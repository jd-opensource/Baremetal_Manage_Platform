package projectLogic

import (
	"time"

	"coding.jd.com/aidc-bmp/bmp-openapi/constant"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/instanceDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/projectDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/sharingProjectDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/userDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/baseLogic"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/instanceLogic"

	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
	responseTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/response"
	"coding.jd.com/aidc-bmp/bmp-openapi/util"
	log "coding.jd.com/aidc-bmp/bmp_log"
	namespacePrefix "git.jd.com/cps-golang/ironic-common/ironic/common/NamespacePrefixs"
	commonUtil "git.jd.com/cps-golang/ironic-common/ironic/util"
)

func ProjectEntity2Project(logger *log.Logger, o *projectDao.Project, shareproject []*sharingProjectDao.SharingProject) *responseTypes.Project {
	tz := logger.GetPoint("timezone").(string)
	query := map[string]interface{}{
		"is_del":     0,
		"project_id": o.ProjectID,
	}
	cnt, err := instanceLogic.GetInstanceCount(logger, query)
	if err != nil {
		logger.Warnf("instanceLogic.GetInstanceCount error:%s", err.Error())
	}

	shares := []responseTypes.ShareProject{}

	if len(shareproject) > 0 {
		for _, v := range shareproject {
			shares = append(shares, responseTypes.ShareProject{
				ProjectName:    v.ProjectName,
				ProjectID:      v.ProjectID,
				OwnerUserID:    v.OwnerUserID,
				OwnerUserName:  v.OwnerUserName,
				SharedUserID:   v.SharedUserID,
				SharedUserName: v.SharedUserName,
				CreatedTime:    util.TimestampToString(int64(v.CreatedTime), tz),
			})
		}

	}

	return &responseTypes.Project{
		ID:            o.ID,
		ProjectName:   o.ProjectName,
		ProjectID:     o.ProjectID,
		InstanceCount: cnt,
		ShareProjects: shares,
		CreatedBy:     o.CreatedBy,
		CreatedTime:   util.TimestampToString(int64(o.CreatedTime), tz),
		UpdatedBy:     o.UpdatedBy,
		UpdatedTime:   util.TimestampToString(int64(o.UpdatedTime), tz),
	}
}

//只带owned标志位,不关心sharing具体内容
func ProjectListEntity2Project(logger *log.Logger, o *projectDao.AllProject) *responseTypes.Project {
	tz := logger.GetPoint("timezone").(string)
	query := map[string]interface{}{
		"is_del":     0,
		"project_id": o.ProjectID,
	}
	cnt, err := instanceLogic.GetInstanceCount(logger, query)
	if err != nil {
		logger.Warnf("instanceLogic.GetInstanceCount error:%s", err.Error())
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

	list, _ := projectDao.GetAllProject(logger, map[string]interface{}{
		"project_name": param.ProjectName,
		"user_id":      userId,
		"is_del":       0,
	})
	if len(list) != 0 {
		logger.Warn("projectName exist:", param.ProjectName)
		panic(constant.BuildInvalidArgumentWithArgs("项目名称已存在", "projectName exist"))
	}
	if param.IsDefault == 1 {
		list, _ = projectDao.GetAllProject(logger, map[string]interface{}{
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

	list, _ := projectDao.GetAllProject(logger, map[string]interface{}{
		"project_name": param.ProjectName,
		"user_id":      userId,
		"is_del":       0,
	})
	if len(list) != 0 {
		logger.Warn("projectName exist:", param.ProjectName)
		panic(constant.BuildInvalidArgumentWithArgs("项目名称已存在", "projectName exist"))
	}

	p.ProjectName = param.ProjectName
	p.UpdatedBy = logger.GetPoint("username").(string)
	p.UpdatedTime = int(time.Now().Unix())
	if err := projectDao.UpdateProjectById(logger, p); err != nil {
		logger.Warn("ModifyProject UpdateProjectById sql error:", ProjectId, err.Error())
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
	count, err := projectDao.GetProjectAndSharedCount(logger, userId, param.Owned)
	if err != nil {
		logger.Warnf("GetProjectCount dao error, query:%s, error:%s", util.ObjToJson(query), err.Error())
		return nil, 0, err
	}
	var entityList []*projectDao.AllProject
	if param.IsAll == baseLogic.IS_ALL {
		entityList, err = projectDao.GetMultiProjectIncludeShare(logger, userId, 0, nil, nil, 0, 0)
	} else {
		entityList, err = projectDao.GetMultiProjectIncludeShare(logger, userId, 0, nil, nil, offset, limit)
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

func CreateShareProject(logger *log.Logger, projectID string, param *requestTypes.ShareProjectRequest) error {
	project, err := projectDao.GetProjectById(logger, projectID)
	if project != nil {
		logger.Warnf("CreateShareProject.GetProjectById error, project_id:%s, error:%s", projectID, err.Error())
		panic(constant.BuildCANCELLEDWithArgs("项目不存在", "project not found"))
	}
	owner, err := userDao.GetUserById(logger, param.OwnerID)
	if owner != nil {
		logger.Warnf("CreateShareProject.GetUserById error, owner_id:%s, error:%s", param.OwnerID, err.Error())
		panic(constant.BuildCANCELLEDWithArgs("用户不存在", "user not found"))
	}
	sharer, err := userDao.GetUserById(logger, param.SharerID)
	if sharer != nil {
		logger.Warnf("CreateShareProject.GetUserById error, sharer_id:%s, error:%s", param.OwnerID, err.Error())
		panic(constant.BuildCANCELLEDWithArgs("用户不存在", "user not found"))
	}

	if project.UserID == param.SharerID {
		panic(constant.BuildCANCELLEDWithArgs("项目不能分享给自己", "project can not shared to the project owner"))
	}

	sp, _ := sharingProjectDao.GetOneSharingProject(logger, projectID, param.OwnerID, param.SharerID)
	if sp != nil {
		panic(constant.BuildCANCELLEDWithArgs("项目已被分享给相同用户", "project already shared to the user"))
	}
	sp = &sharingProjectDao.SharingProject{
		ProjectID:      projectID,
		OwnerUserID:    param.OwnerID,
		SharedUserID:   param.SharerID,
		OwnerUserName:  owner.UserName,
		SharedUserName: sharer.UserName,
	}
	_, err = sharingProjectDao.AddSharingProject(logger, sp)
	if err != nil {
		logger.Warnf("CreateShareProject.AddSharingProject error, project_id:%s, owner_id:%s, sharer_id:%s, error:%s", projectID, param.OwnerID, param.SharerID, err.Error())
		panic(constant.INTERNAL_ERROR)
	}

	return nil
}

func CancelShareProject(logger *log.Logger, projectID string, param *requestTypes.ShareProjectRequest) (err error) {

	sp, _ := sharingProjectDao.GetOneSharingProject(logger, projectID, param.OwnerID, param.SharerID)
	if sp == nil {
		panic(constant.BuildCANCELLEDWithArgs("共享项目不存在", "shared project not found"))
	}
	err = sharingProjectDao.DeleteOneSharingProject(logger, sp)
	if err != nil {
		logger.Warnf("CreateShareProject.DeleteSharingProject error, project_id:%s, owner_id:%s, sharer_id:%s, error:%s", projectID, param.OwnerID, param.SharerID, err.Error())
		panic(constant.INTERNAL_ERROR)
	}
	return nil
}
