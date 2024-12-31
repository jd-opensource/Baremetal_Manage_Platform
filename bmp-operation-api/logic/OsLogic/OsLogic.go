package OsLogic

import (
	"errors"

	sdkParameters "coding.jd.com/aidc-bmp/bmp-openapi-sdk/client/os"
	sdk "coding.jd.com/aidc-bmp/bmp-operation-api/services/openApi"
	requestTypes "coding.jd.com/aidc-bmp/bmp-operation-api/types/request"
	"coding.jd.com/aidc-bmp/bmp-operation-api/types/response"
	"coding.jd.com/aidc-bmp/bmp-operation-api/util"
	log "git.jd.com/cps-golang/log"
)

// 【raid列表】
func DescribeOSs(logger *log.Logger, param *requestTypes.QueryOssRequest) (*response.OsList, error) {
	userId := logger.GetPoint("userId").(string)
	logid := logger.GetPoint("logid").(string)
	req := sdkParameters.NewDescribeOSsParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	req.WithBmpUserID(&userId)
	req.WithOsType(&param.OsType)
	req.WithOsName(&param.OsName)
	req.WithOsVersion(&param.OsVersion)
	//req.WithPlatform(&param.Platform)
	req.WithIsAll(&param.IsAll)
	logger.Info("DescribeOSs openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.Os.DescribeOSs(req, nil)
	logger.Info("DescribeOSs openapi sdk resp:", util.InterfaceToJson(client))

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
	resp := response.OsList{}

	for _, v := range client.Payload.Result.Oss {
		obj := response.Os{}
		obj.ID = v.ID
		obj.OsID = v.OsID
		obj.OsName = v.OsName
		obj.OsType = v.OsType
		obj.OsVersion = v.OsVersion
		obj.Architecture = v.Architecture
		obj.Bits = int(v.Bits)
		obj.SysUser = v.SysUser
		obj.CreatedTime = util.UtcToTime(v.CreatedTime)
		obj.UpdatedTime = util.UtcToTime(v.UpdatedTime)
		obj.CreatedBy = v.CreatedBy
		obj.UpdatedBy = v.UpdatedBy
		resp.Oss = append(resp.Oss, &obj)
	}
	return &resp, nil
}
