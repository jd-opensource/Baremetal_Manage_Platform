package KeypairLoigc

import (
	"errors"

	"coding.jd.com/aidc-bmp/bmp-console-api/thirdpart/openApi"
	requestTypes "coding.jd.com/aidc-bmp/bmp-console-api/types/request"
	"coding.jd.com/aidc-bmp/bmp-console-api/types/response"

	"coding.jd.com/aidc-bmp/bmp-console-api/util"
	sdkSshkeyParameters "coding.jd.com/aidc-bmp/bmp-openapi-console-sdk/client/sshkey"
	sdkModels "coding.jd.com/aidc-bmp/bmp-openapi-console-sdk/models"
	log "coding.jd.com/aidc-bmp/bmp_log"
)

func DescribeKeypairs(logger *log.Logger, param requestTypes.QueryKeypairsRequest) (*response.KeyPairsResponse, error) {

	logid := logger.GetPoint("logid").(string)
	req := sdkSshkeyParameters.NewDescribeUserSSHKeysParams()
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	if param.Name != "" {
		req.WithName(&param.Name)
	}
	pageNumber := param.PageNumber
	pageSize := param.PageSize
	if pageNumber != 0 {
		req.WithPageNumber(&pageNumber)
	}
	if pageSize != 0 {
		req.WithPageSize(&pageSize)
	}
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	logger.Info("describekeypairs openapi request:", util.InterfaceToJson(req))
	responseApi, err := openApi.SdkClient.Sshkey.DescribeUserSSHKeys(req, nil)
	logger.Info("describekeypairs openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("describekeypairs openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}

	resp := response.KeyPairsResponse{}
	resp.PagingResponse.PageNumber = responseApi.Payload.Result.PageNumber
	resp.PagingResponse.PageSize = responseApi.Payload.Result.PageSize
	resp.PagingResponse.TotalCount = responseApi.Payload.Result.TotalCount

	obj := response.KeyPair{}
	for _, v := range responseApi.Payload.Result.Sshkeys {

		obj.Name = v.Name
		obj.KeypairId = v.SshkeyID

		req := sdkSshkeyParameters.NewGetInstancesBySshkeyParams()
		userId := logger.GetPoint("userId").(string)
		req.WithBmpUserID(&userId)
		req.WithSshkeyID(v.SshkeyID)
		req.WithBmpLanguage(&language)
		logger.Info("getInstancesBySshkey openapi request:", util.InterfaceToJson(req))
		respInstance, _ := openApi.SdkClient.Sshkey.GetInstancesBySshkey(req, nil)
		logger.Info("getInstancesBySshkey openapi response:", util.InterfaceToJson(respInstance))
		if err != nil {
			respErr, b := util.GetOpenapiError(err)
			if !b {
				logger.Warn("getInstancesBySshkey openapi error:", err.Error())
			}
			logger.Warn("getInstancesBySshkey openapi error:", respErr.Message)
		}
		obj.IsCheckDelete = "yes"
		if len(respInstance.Payload.Result.InstanceIds) != 0 {
			// for _, value := range respInstance.Result.Instances {
			// 	if value.Status == "creating" || value.Status == "reinstalling" {
			// 		logger.Info("list keypair "+v.KeypairId+" is using by instance:", value.InstanceId)
			// 		obj.IsCheckDelete = "no"
			// 		break
			// 	}
			// }
			obj.IsCheckDelete = "no"
		}
		obj.FingerPrint = v.FingerPrint
		obj.CreatedTime = v.CreatedTime
		obj.UpdatedTime = v.UpdatedTime
		resp.Keypairs = append(resp.Keypairs, obj)

	}

	return &resp, nil
}

func DescribeKeypair(logger *log.Logger, keypairID string) (*response.KeyPairInfo, error) {

	logid := logger.GetPoint("logid").(string)
	req := sdkSshkeyParameters.NewDescribeUserSSHKeyParams()
	req.WithSshkeyID(keypairID)
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	logger.Info("describekeypair openapi request:", util.InterfaceToJson(req))
	responseApi, err := openApi.SdkClient.Sshkey.DescribeUserSSHKey(req, nil)
	logger.Info("describekeypair openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("describekeypair openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}

	obj := response.KeyPairInfo{}
	v := responseApi.Payload.Result.Sshkey
	obj.Keypair.Name = v.Name
	obj.Keypair.KeypairId = v.SshkeyID

	insReq := sdkSshkeyParameters.NewGetInstancesBySshkeyParams()
	req.WithSshkeyID(v.SshkeyID)
	req.WithBmpUserID(&userId)
	req.WithBmpLanguage(&language)
	logger.Info("getInstancesBySshkey openapi request:", util.InterfaceToJson(insReq))
	respInstance, err := openApi.SdkClient.Sshkey.GetInstancesBySshkey(insReq, nil)
	logger.Info("getInstancesBySshkey openapi response:", util.InterfaceToJson(respInstance))
	if err != nil {
		respErr, b := util.GetOpenapiError(err)
		if !b {
			logger.Warn("getInstancesBySshkey openapi error:", err.Error())
		}
		logger.Warn("getInstancesBySshkey openapi error:", respErr.Message)
	}
	obj.Keypair.IsCheckDelete = "yes"
	if len(respInstance.Payload.Result.InstanceIds) != 0 {
		// for _, value := range respInstance.Result.Instances {
		// 	if value.Status == "creating" || value.Status == "reinstalling" {
		// 		logger.Info("list keypair "+v.KeypairId+" is using by instance:", value.InstanceId)
		// 		obj.IsCheckDelete = "no"
		// 		break
		// 	}
		// }
		obj.Keypair.IsCheckDelete = "no"
	}

	//obj.Keypair.IsCheckDelete = "yes"
	obj.Keypair.FingerPrint = v.FingerPrint
	obj.Keypair.CreatedTime = v.CreatedTime
	obj.Keypair.UpdatedTime = v.UpdatedTime

	return &obj, nil
}

func CreateKeypairs(logger *log.Logger, param requestTypes.CreateKeypairRequest) (*response.KeyPair, error) {

	logid := logger.GetPoint("logid").(string)
	req := sdkSshkeyParameters.NewCreateUserSshkeyParams()
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	body := &sdkModels.CreateSshkeyRequest{
		Name: &param.Name,
		Key:  &param.Key,
	}
	req.WithBody(body)
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	logger.Info("createkeypairs openapi request:", util.InterfaceToJson(req))
	responseApi, err := openApi.SdkClient.Sshkey.CreateUserSshkey(req, nil)
	logger.Info("createkeypairs openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("createkeypairs openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}
	obj := response.KeyPair{}
	// v := responseApi.Payload.Result
	obj.Name = param.Name
	// obj.KeypairId = v.KeypairId

	return &obj, nil
}

func ModifyKeypair(logger *log.Logger, keypairId, name string) (bool, error) {

	logid := logger.GetPoint("logid").(string)
	req := sdkSshkeyParameters.NewModifyUserSshkeyParams()
	req.WithTraceID(logid)
	req.WithSshkeyID(keypairId)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	body := &sdkModels.ModifySshkeyRequest{
		Name: &name,
	}
	req.WithBody(body)
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	logger.Info("ModifyUserSshkey openapi request:", util.InterfaceToJson(req))
	responseApi, err := openApi.SdkClient.Sshkey.ModifyUserSshkey(req, nil)
	if err != nil {
		logger.Info("ModifyUserSshkey openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return false, err
		}
		return false, errors.New(respErr.Message)
	}
	logger.Info("ModifyUserSshkey openapi response:", util.InterfaceToJson(responseApi))
	return responseApi.Payload.Result.Success, nil
}

func DeleteKeypairs(logger *log.Logger, keypairID string) (bool, error) {

	logid := logger.GetPoint("logid").(string)
	req := sdkSshkeyParameters.NewDeleteUserSshkeyParams()
	req.WithTraceID(logid)
	req.WithSshkeyID(keypairID)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	logger.Info("deletekeypairs openapi request:", util.InterfaceToJson(req))
	responseApi, err := openApi.SdkClient.Sshkey.DeleteUserSshkey(req, nil)
	if err != nil {
		logger.Info("deletekeypairs openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return false, err
		}
		return false, errors.New(respErr.Message)
	}
	logger.Info("deletekeypairs openapi response:", util.InterfaceToJson(responseApi))
	return responseApi.Payload.Result.Success, nil
}
