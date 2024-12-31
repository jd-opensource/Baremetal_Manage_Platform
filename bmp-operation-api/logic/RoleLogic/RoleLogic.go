package RoleLogic

import (
	"errors"

	sdkParameters "coding.jd.com/aidc-bmp/bmp-openapi-sdk/client/role"
	sdk "coding.jd.com/aidc-bmp/bmp-operation-api/services/openApi"
	requestTypes "coding.jd.com/aidc-bmp/bmp-operation-api/types/request"
	"coding.jd.com/aidc-bmp/bmp-operation-api/types/response"
	"coding.jd.com/aidc-bmp/bmp-operation-api/util"
	log "git.jd.com/cps-golang/log"
)

// 【raid列表】
func DescribeRoles(logger *log.Logger, param *requestTypes.DescribeRolesRequest) (*response.RoleList, error) {
	userId := logger.GetPoint("userId").(string)
	logid := logger.GetPoint("logid").(string)
	req := sdkParameters.NewDescribeRolesParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	req.WithBmpUserID(&userId)
	req.WithIsAll(&param.IsAll)
	logger.Info("DescribeRoles openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.Role.DescribeRoles(req, nil)
	logger.Info("DescribeRoles openapi sdk resp:", util.InterfaceToJson(client))

	if err != nil || client.Payload == nil {
		logger.Info("openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}
	//if len(result.Payload.Result.Oss) == 0 {
	//	return nil, errors.New("Oss not found")
	//}
	resp := response.RoleList{}

	for _, v := range client.Payload.Result.Roles {
		obj := response.Role{}

		obj.RoleID = v.RoleID
		obj.DescriptionEn = v.DescriptionEn
		obj.DescriptionCn = v.DescriptionCn
		obj.RoleNameEn = v.RoleNameEn
		obj.RoleNameCn = v.RoleNameCn
		obj.RoleName = v.RoleName
		obj.Permission = v.Permission
		obj.UserCount = int(v.UserCount)
		obj.CreatedTime = util.UtcToTime(v.CreatedTime)
		obj.UpdatedTime = util.UtcToTime(v.UpdatedTime)
		obj.CreatedBy = v.CreatedBy
		obj.UpdatedBy = v.UpdatedBy
		resp.Roles = append(resp.Roles, &obj)
	}
	resp.PageNumber = client.Payload.Result.PageNumber
	resp.PageSize = client.Payload.Result.PageSize
	resp.TotalCount = client.Payload.Result.TotalCount
	return &resp, nil
}

func CurrentRole(logger *log.Logger) (*response.Role, error) {
	userId := logger.GetPoint("userId").(string)
	logid := logger.GetPoint("logid").(string)
	req := sdkParameters.NewCurrentRoleParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	req.WithBmpUserID(&userId)
	logger.Info("CurrentRole openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.Role.CurrentRole(req, nil)
	logger.Info("CurrentRole openapi sdk resp:", util.InterfaceToJson(client))

	if err != nil || client.Payload == nil {
		return nil, err
	}
	o := client.Payload.Result.Role

	r := &response.Role{
		RoleName:      o.RoleName,
		Permission:    o.Permission,
		RoleNameEn:    o.RoleNameEn,
		RoleNameCn:    o.RoleNameCn,
		RoleID:        o.RoleID,
		DescriptionEn: o.DescriptionEn,
		DescriptionCn: o.DescriptionCn,
		CreatedBy:     o.CreatedBy,
		CreatedTime:   o.CreatedTime,
		UpdatedBy:     o.UpdatedBy,
		UpdatedTime:   o.UpdatedTime,
	}

	return r, nil
}
