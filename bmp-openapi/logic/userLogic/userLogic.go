package userLogic

import (
	"strings"
	"time"

	"coding.jd.com/aidc-bmp/bmp-openapi/constant"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/instanceDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/projectDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/roleDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/sharingProjectDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/baseLogic"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/projectLogic"

	userDao "coding.jd.com/aidc-bmp/bmp-openapi/dao/userDao"
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
	responseTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/response"
	"coding.jd.com/aidc-bmp/bmp-openapi/util"
	log "coding.jd.com/aidc-bmp/bmp_log"
	namespacePrefix "git.jd.com/cps-golang/ironic-common/ironic/common/NamespacePrefixs"
	commonUtil "git.jd.com/cps-golang/ironic-common/ironic/util"
)

func userEntity2user(logger *log.Logger, o *userDao.User) *responseTypes.User {

	tz := logger.GetPoint("timezone").(string)
	role, _ := roleDao.GetRoleById(logger, o.RoleID)
	if role == nil {
		role = &roleDao.Role{}
	}
	project, _ := projectDao.GetProjectById(logger, o.DefaultProjectID)
	if project == nil {
		project = &projectDao.Project{}
	}

	roleName := role.RoleNameCn
	if logger.GetPoint("language") == baseLogic.EN_US {
		roleName = role.RoleNameEn
	}
	return &responseTypes.User{
		ID:                 o.ID,
		UserName:           o.UserName,
		Password:           o.Password,
		UserID:             o.UserID,
		RoleID:             o.RoleID,
		RoleName:           roleName,
		PhonePrefix:        o.PhonePrefix,
		PhoneNumber:        o.PhoneNumber,
		DefaultProjectID:   o.DefaultProjectID,
		DefaultProjectName: project.ProjectName,
		Timezone:           o.Timezone,
		Language:           o.Language,
		Email:              o.Email,
		Description:        o.Description,
		CreatedBy:          o.CreatedBy,
		CreatedTime:        util.TimestampToString(int64(o.CreatedTime), tz),
		UpdatedBy:          o.UpdatedBy,
		UpdatedTime:        util.TimestampToString(int64(o.UpdatedTime), tz),
	}
}

func GetUserById(logger *log.Logger, userId string) (*responseTypes.User, error) {
	entity, err := userDao.GetUserById(logger, userId)
	if err != nil {
		logger.Warn("GetuserByUuid sql error:", userId, err.Error())
		return nil, err
	}
	logger.Infof("GetUserById userid:%s, entity:%s", userId, util.InterfaceToJson(entity))
	return userEntity2user(logger, entity), nil
}

func GetUserByName(logger *log.Logger, param *requestTypes.GetUserByNameRequest) (*responseTypes.User, error) {
	entity, err := userDao.GetUserByName(logger, param.UserName)
	if err != nil {
		logger.Warn("GetuserByName sql error:", param.UserName, err.Error())
		return nil, err
	}
	return userEntity2user(logger, entity), nil
}

func CreateUser(logger *log.Logger, param *requestTypes.CreateUserRequest) (string, error) {
	list, _ := userDao.GetAllUser(logger, map[string]interface{}{
		"user_name": param.UserName,
		"is_del":    baseLogic.IS_NOT_DEL,
	})
	if len(list) != 0 {
		logger.Warn("user name exist:", param.UserName)
		panic(constant.BuildInvalidArgumentWithArgs("用户名 "+param.UserName+" 已存在", "username "+param.UserName+" exist"))
	}
	if param.RoleID == baseLogic.ROLE_ADMIN_UUID { //如果是超级管理角色
		listRole, _ := userDao.GetAllUser(logger, map[string]interface{}{
			"role_id": baseLogic.ROLE_ADMIN_UUID,
			"is_del":  baseLogic.IS_NOT_DEL,
		})
		if len(listRole) != 0 {
			logger.Warn("roleID  exist:", param.UserName)
			panic(constant.BuildInvalidArgumentWithArgs("平台拥有者已存在,只能存在一个平台拥有着", "Owner exist"))
		}
	}
	if _, err := roleDao.GetRoleById(logger, param.RoleID); err != nil {
		chmsg := "角色Id不存在"
		enmsg := "roleId not exist"
		panic(constant.BuildNotFoundWithArgs(chmsg, enmsg))
	}
	user_entity := &userDao.User{
		UserName:    param.UserName,
		RoleID:      param.RoleID,
		Email:       param.Email,
		PhonePrefix: param.PhonePrefix,
		PhoneNumber: param.PhoneNumber,
		Password:    util.BcryptEncode(param.Password),
		Description: param.Description,
		Timezone:    param.Timezone,
		Language:    param.Language,
		CreatedBy:   logger.GetPoint("username").(string),
		CreatedTime: int(time.Now().Unix()),
		UpdatedBy:   logger.GetPoint("username").(string),
		UpdatedTime: int(time.Now().Unix()),
	}
	user_entity.UserID = commonUtil.GetRandString("user-", namespacePrefix.INSTANCE_ID_LENGTH, false, true, true)

	pn := baseLogic.DEFAULTPROJECTNAME
	if logger.GetPoint("language").(string) == baseLogic.EN_US {
		pn = baseLogic.DEFAULTPROJECTNAME_EN
	}

	projectId, err := projectLogic.CreateProject(logger, &requestTypes.CreateProjectRequest{
		ProjectName: pn,
		IsDefault:   1,
		IsSystem:    1,
	}, user_entity.UserID)
	if err != nil {
		logger.Warnf("Createuser Add defaultProject sql error, entity:%s, error:%s", util.ObjToJson(user_entity), err.Error())
		return "", err
	}
	user_entity.DefaultProjectID = projectId
	if _, err := userDao.Adduser(logger, user_entity); err != nil {
		logger.Warnf("Createuser Adduser sql error, entity:%s, error:%s", util.ObjToJson(user_entity), err.Error())
		return "", err
	}
	return user_entity.UserID, nil
}

func ModifyUser(logger *log.Logger, param *requestTypes.ModifyUserRequest, userId string) error {

	entity, err := userDao.GetUserById(logger, userId)
	if err != nil {
		logger.Warn("GetuserByUuid sql error:", userId, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}
	if param.Description != nil {
		entity.Description = *param.Description
	}
	if param.PhonePrefix != nil {
		entity.PhonePrefix = *param.PhonePrefix
	}
	if param.PhoneNumber != nil {
		entity.PhoneNumber = *param.PhoneNumber
	}
	if param.Email != nil {
		entity.Email = *param.Email
	}
	entity.UpdatedBy = logger.GetPoint("username").(string)
	entity.UpdatedTime = int(time.Now().Unix())
	if err := userDao.UpdateUserById(logger, entity); err != nil {
		logger.Warn("Modifyuser UpdateuserById sql error:", userId, err.Error())
		return err
	}
	return nil
}

func ModifyLocalUser(logger *log.Logger, param *requestTypes.ModifyLocalUserRequest) error {

	userId := logger.GetPoint("userId").(string)
	u, err := userDao.GetUserById(logger, userId)
	if err != nil {
		logger.Warn("GetuserByUuid sql error:", userId, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}
	if param.DefaultProjectID != nil {
		u.DefaultProjectID = *param.DefaultProjectID
	}
	if param.Email != nil {
		u.Email = *param.Email
	}
	if param.PhonePrefix != nil {
		u.PhonePrefix = *param.PhonePrefix
	}
	if param.PhoneNumber != nil {
		u.PhoneNumber = *param.PhoneNumber
	}
	if param.Language != nil {
		u.Language = *param.Language
	}
	if param.Timezone != nil {
		u.Timezone = *param.Timezone
	}

	u.UpdatedBy = logger.GetPoint("username").(string)
	u.UpdatedTime = int(time.Now().Unix())
	if err := userDao.UpdateUserById(logger, u); err != nil {
		logger.Warn("Modifyuser UpdateuserById sql error:", userId, err.Error())
		return err
	}
	if param.DefaultProjectID != nil {
		if err := projectDao.UpdateUserDefaultProject(logger, userId, *param.DefaultProjectID); err != nil {
			logger.Warn("Modifyuser UpdateUserDefaultProject sql error:", userId, err.Error())
			return err
		}
	}
	return nil
}

func DeleteUser(logger *log.Logger, userId string) error {
	entity, err := userDao.GetUserById(logger, userId)
	if err != nil {
		logger.Warn("GetuserByUuid sql error:", userId, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}
	if entity.RoleID == baseLogic.ROLE_ADMIN_UUID {
		panic(constant.BuildInvalidArgumentWithArgs("平台拥有着不能被删除", "Can't delete Owner"))
	}

	q := map[string]interface{}{
		"is_del":  0,
		"user_id": userId,
	}
	instanceC, _ := instanceDao.GetInstanceCount(logger, q)
	if instanceC > 0 {
		logger.Warnf("user have instance, user_id:%s, count:%d", userId, instanceC)
		panic(constant.BuildInvalidArgumentWithArgs("此用户内还有实例资源未被释放，请释放或转移给其它用户后才能删除用户", "There are resources in this user that have not been released. Please release or transfer to other users before deleting the user"))
	}

	//删除项目和共享项目
	deleteUpdate := map[string]interface{}{
		"is_del": 1,
	}
	if err := projectDao.UpdateProjects(logger, q, deleteUpdate); err != nil {
		logger.Warnf("delete project for user drop error, user_id:%s, err:%s", userId, err.Error())
	}

	q1 := map[string]interface{}{
		"is_del":        0,
		"owner_user_id": userId,
	}
	if err := sharingProjectDao.UpdateShareProjects(logger, q1, deleteUpdate); err != nil {
		logger.Warnf("delete share project for user drop error, user_id:%s, err:%s", userId, err.Error())
	}

	q2 := map[string]interface{}{
		"is_del":         0,
		"shared_user_id": userId,
	}
	if err := sharingProjectDao.UpdateShareProjects(logger, q2, deleteUpdate); err != nil {
		logger.Warnf("delete share project for user drop error, user_id:%s, err:%s", userId, err.Error())
	}

	// instanceList, err := instanceDao.GetAllInstance(logger, map[string]interface{}{
	// 	"user_id": userId,
	// 	"is_del":  baseLogic.IS_NOT_DEL,
	// })
	// if len(instanceList) != 0 {
	// 	logger.Warn("user have instances:", userId)
	// 	panic(constant.BuildInvalidArgumentWithArgs("此用户内还有资源未被释放，请释放或转移给其它用户后才能删除用户", "There are resources in this user that have not been released. Please release or transfer to other users before deleting the user"))
	// }
	// deviceList, err := deviceDao.GetAllDevice(logger, map[string]interface{}{
	// 	"user_id": userId,
	// 	"is_del":  baseLogic.IS_NOT_DEL,
	// })
	// if len(deviceList) != 0 {
	// 	logger.Warn("user have devices:", userId)
	// 	panic(constant.BuildInvalidArgumentWithArgs("此用户内还有资源未被释放，请释放或转移给其它用户后才能删除用户", "There are resources in this user that have not been released. Please release or transfer to other users before deleting the user"))
	// }
	entity.IsDel = 1
	entity.UpdatedBy = logger.GetPoint("username").(string)
	entity.UpdatedTime = int(time.Now().Unix())
	entity.DeletedTime = int(time.Now().Unix())
	if err := userDao.UpdateUserById(logger, entity); err != nil {
		logger.Warn("Deleteuser UpdateuserById sql error:", userId, err.Error())
		return err
	}

	//删除项目和共享项目

	return nil
}

func GetUserList(logger *log.Logger, param *requestTypes.QueryUsersRequest, p util.Pageable) ([]*responseTypes.User, int64, error) {
	offset, limit := p.Pageable2offset()
	query := map[string]interface{}{
		"user_name.like":     "%" + param.UserName + "%",
		"default_project_id": param.DefaultProjectID,
		"is_del":             0,
	}
	if param.RoleID != "" {
		query["role_id.in"] = strings.Split(param.RoleID, ",")
	}
	count, err := userDao.GetUserCount(logger, query)
	if err != nil {
		logger.Warnf("GetUserCount dao error, query:%s, error:%s", util.ObjToJson(query), err.Error())
		return nil, 0, err
	}
	var entityList []*userDao.User
	if param.IsAll == baseLogic.IS_ALL {
		entityList, err = userDao.GetMultiUser(logger, query, []string{}, []string{"id"}, []string{"desc"}, 0, 0)
	} else {
		entityList, err = userDao.GetMultiUser(logger, query, []string{}, []string{"id"}, []string{"desc"}, offset, limit)
	}
	if err != nil {
		logger.Warn("QueryByuserIds GetAlluser sql error:", err.Error())
		return nil, 0, err
	}
	users := []*responseTypes.User{}
	pcu, err := projectDao.GetProjectCountGroupByUser(logger)
	if err != nil {
		logger.Warnf("GetProjectCountGroupByUser err:%s", err.Error())
	}

	spcu, err := sharingProjectDao.GetShareProjectCountGroupByUser(logger)
	if err != nil {
		logger.Warnf("GetShareProjectCountGroupByUser err:%s", err.Error())
	}

	icu, err := instanceDao.GetInstanceCountGroupByUser(logger)
	if err != nil {
		logger.Warnf("GetInstanceCountGroupByUser err:%s", err.Error())
	}
	for _, entity := range entityList {
		u := userEntity2user(logger, entity)
		for _, vv := range pcu {
			if vv.UserID == u.UserID {
				u.ProjectCount = int(vv.Count)
				break
			}
		}
		for _, vv := range spcu {
			if vv.UserID == u.UserID {
				u.ProjectCount += int(vv.Count)
				break
			}
		}

		for _, vv := range icu {
			if vv.UserID == u.UserID {
				u.InstanceCount = int(vv.Count)
				break
			}
		}
		users = append(users, u)
	}

	return users, count, nil
}

func VerifyUser(logger *log.Logger, req *requestTypes.VerifyUserRequest) bool {
	entity, _ := userDao.GetAllUser(logger, map[string]interface{}{
		"user_name": req.UserName,
		"is_del":    baseLogic.IS_NOT_DEL,
	})
	if len(entity) == 0 { //说明用户名不存在
		panic(constant.BuildInvalidArgumentWithArgs("用户名不存在", "userName not exist"))
	}

	permit := false
	for _, v := range entity {
		if v.RoleID == req.RoleID || v.RoleID == "role-admin-uuid-01" {
			permit = true
			break
		}
	}

	if !permit { //无权限
		panic(constant.BuildInvalidArgumentWithArgs("用户权限不足", "permission denied"))
	}
	logger.Info("请求信息", util.ObjToJson(req), "数据库hash值", entity[0].Password)
	res := util.BcryptVerify(req.Password, entity[0].Password)
	if !res {
		panic(constant.BuildInvalidArgumentWithArgs("密码错误", "password error"))
	} else {
		return true
	}
}

func ModifyUserPassword(logger *log.Logger, param *requestTypes.ModifyUserPasswordRequest) error {
	userId := logger.GetPoint("userId").(string)
	u, err := userDao.GetUserById(logger, userId)
	if err != nil {
		logger.Warn("GetuserByUuid sql error:", userId, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}
	if !util.BcryptVerify(param.OldPassword, u.Password) {
		panic(constant.BuildInvalidArgumentWithArgs("旧密码输入错误", "oldPassword error"))
	}
	u.Password = util.BcryptEncode(param.Password)
	if err := userDao.UpdateUserById(logger, u); err != nil {
		logger.Warn("ModifyUserPassword UpdateuserById sql error:", userId, err.Error())
		return err
	}
	return nil
}
