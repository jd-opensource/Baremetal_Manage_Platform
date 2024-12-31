package ApikeyLogic

import (
	"errors"

	sdkModels "coding.jd.com/aidc-bmp/bmp-openapi-console-sdk/models"

	"coding.jd.com/aidc-bmp/bmp-console-api/thirdpart/openApi"
	requestTypes "coding.jd.com/aidc-bmp/bmp-console-api/types/request"
	"coding.jd.com/aidc-bmp/bmp-console-api/util"
	sdkApikeyParameters "coding.jd.com/aidc-bmp/bmp-openapi-console-sdk/client/api_key"
	log "coding.jd.com/aidc-bmp/bmp_log"
)

func CreateApikey(logger *log.Logger, param *sdkModels.CreateApikeyRequest) (string, error) {

	logid := logger.GetPoint("logid").(string)

	req := sdkApikeyParameters.NewCreateUserApikeyParams()
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	req.WithBody(param)
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	logger.Info("CreateApikey openapi request:", util.InterfaceToJson(req))
	responseApi, err := openApi.SdkClient.APIKey.CreateUserApikey(req, nil)
	logger.Info("CreateApikey openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("CreateApikey openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return "", err
		}
		return "", errors.New(respErr.Message)
	}
	return *responseApi.Payload.Result.ApikeyID, nil

}

func DeleteApikey(logger *log.Logger, apikeyId string) error {
	logid := logger.GetPoint("logid").(string)
	req := sdkApikeyParameters.NewDeleteUserApikeyParams()
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	req.WithApikeyID(apikeyId)
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	logger.Info("DeleteApikey openapi request:", util.InterfaceToJson(req))
	responseApi, err := openApi.SdkClient.APIKey.DeleteUserApikey(req, nil)
	logger.Info("DeleteApikey openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("DeleteApikey openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return err
		}
		return errors.New(respErr.Message)
	}
	return nil
}

func GetApikeyList(logger *log.Logger, param requestTypes.GetApikeyListRequest) (*sdkModels.ApikeyList, error) {

	logid := logger.GetPoint("logid").(string)
	req := sdkApikeyParameters.NewDescribeUserAPIKeysParams()
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	req.WithPageNumber(&param.PageNumber)
	req.WithPageSize(&param.PageSize)
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	logger.Info("GetApikeyList openapi request:", util.InterfaceToJson(req))
	responseApi, err := openApi.SdkClient.APIKey.DescribeUserAPIKeys(req, nil)
	logger.Info("GetApikeyList openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("GetApikeyList openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}
	return responseApi.Payload.Result, nil

}
