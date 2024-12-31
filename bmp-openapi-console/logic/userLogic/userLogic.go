package userLogic

import (
	"fmt"
	"strings"
	"time"

	"coding.jd.com/aidc-bmp/bmp-openapi/constant"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/projectDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/roleDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/baseLogic"

	userDao "coding.jd.com/aidc-bmp/bmp-openapi/dao/userDao"
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
	responseTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/response"
	"coding.jd.com/aidc-bmp/bmp-openapi/util"
	log "coding.jd.com/aidc-bmp/bmp_log"
)

func userEntity2user(logger *log.Logger, o *userDao.User) *responseTypes.User {

	tz := logger.GetPoint("timezone").(string)
	if tz == "" {
		tz = o.Timezone
	}
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
	entityList := []*userDao.User{}
	fmt.Println(param.IsAll)
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
	for _, entity := range entityList {
		users = append(users, userEntity2user(logger, entity))
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
