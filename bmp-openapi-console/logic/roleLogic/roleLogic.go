package roleLogic

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/constant"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/userDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/baseLogic"

	"coding.jd.com/aidc-bmp/bmp-openapi/dao/roleDao"
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
	responseTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/response"
	"coding.jd.com/aidc-bmp/bmp-openapi/util"
	log "coding.jd.com/aidc-bmp/bmp_log"
)

func RoleEntity2Role(logger *log.Logger, o *roleDao.Role) *responseTypes.Role {
	tz := logger.GetPoint("timezone").(string)
	userCount, _ := userDao.GetUserCount(logger, map[string]interface{}{
		"role_id": o.RoleID,
	})
	roleName := o.RoleNameCn
	if logger.GetPoint("language").(string) == baseLogic.EN_US {
		roleName = o.RoleNameEn
	}
	permission := o.DescriptionCn
	if logger.GetPoint("language").(string) == baseLogic.EN_US {
		permission = o.DescriptionEn
	}
	return &responseTypes.Role{
		ID:            o.ID,
		RoleName:      roleName,
		Permission:    permission,
		RoleNameEn:    o.RoleNameEn,
		RoleNameCn:    o.RoleNameCn,
		RoleID:        o.RoleID,
		UserCount:     int(userCount),
		DescriptionEn: o.DescriptionEn,
		DescriptionCn: o.DescriptionCn,
		CreatedBy:     o.CreatedBy,
		CreatedTime:   util.TimestampToString(int64(o.CreatedTime), tz),
		UpdatedBy:     o.UpdatedBy,
		UpdatedTime:   util.TimestampToString(int64(o.UpdatedTime), tz),
	}
}

func GetRoleById(logger *log.Logger, RoleId string) (*responseTypes.Role, error) {
	entity, err := roleDao.GetRoleById(logger, RoleId)
	if err != nil {
		logger.Warn("GetRoleByUuid sql error:", RoleId, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}
	return RoleEntity2Role(logger, entity), nil
}

func CurrentRole(logger *log.Logger) (*responseTypes.Role, error) {

	userId := logger.GetPoint("userId").(string)
	u, err := userDao.GetUserById(logger, userId)
	if err != nil {
		logger.Warn("GurrentRole.GetUserById sql error:", userId, err.Error())
		panic(constant.INVALID_ARGUMENT)
	}

	entity, err := roleDao.GetRoleById(logger, u.RoleID)
	if err != nil {
		logger.Warn("GurrentRole.GetRoleByUuid sql error:", u.RoleID, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}
	return RoleEntity2Role(logger, entity), nil
}

func GetRoleList(logger *log.Logger, param requestTypes.QueryRolesRequest, p util.Pageable) ([]*responseTypes.Role, int64, error) {
	offset, limit := p.Pageable2offset()
	query := map[string]interface{}{
		//"role_name": param.RoleName,
		//"role_id":            param.RoleID,
		//"default_project_id": param.DefaultProjectID,
		"is_del": 0,
	}
	count, err := roleDao.GetRoleCount(logger, query)
	if err != nil {
		logger.Warnf("GetRoleCount dao error, query:%s, error:%s", util.ObjToJson(query), err.Error())
		return nil, 0, err
	}
	entityList := []*roleDao.Role{}
	if param.IsAll == baseLogic.IS_ALL {
		entityList, err = roleDao.GetMultiRole(logger, query, []string{}, []string{"id"}, []string{"desc"}, 0, 0)
	} else {
		entityList, err = roleDao.GetMultiRole(logger, query, []string{}, []string{"id"}, []string{"desc"}, offset, limit)
	}
	if err != nil {
		logger.Warn("QueryByRoleIds GetAllRole sql error:", err.Error())
		return nil, 0, err
	}
	Roles := []*responseTypes.Role{}
	for _, entity := range entityList {
		Roles = append(Roles, RoleEntity2Role(logger, entity))
	}
	return Roles, count, nil
}
