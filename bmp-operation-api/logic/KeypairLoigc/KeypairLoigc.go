package KeypairLoigc

import (
	"errors"

	"coding.jd.com/aidc-bmp/bmp-operation-api/services/openApi"
	requestTypes "coding.jd.com/aidc-bmp/bmp-operation-api/types/request"
	"coding.jd.com/aidc-bmp/bmp-operation-api/types/response"

	sdkSshkeyParameters "coding.jd.com/aidc-bmp/bmp-openapi-sdk/client/sshkey"
	"coding.jd.com/aidc-bmp/bmp-operation-api/util"
	log "git.jd.com/cps-golang/log"
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
