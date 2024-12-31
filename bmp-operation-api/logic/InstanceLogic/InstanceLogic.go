package InstanceLogic

import (
	// "errors"

	"errors"
	"fmt"
	"strconv"
	"strings"

	sdkInstanceParameters "coding.jd.com/aidc-bmp/bmp-openapi-sdk/client/instance"
	"coding.jd.com/aidc-bmp/bmp-operation-api/logic/BaseLogic"
	"coding.jd.com/aidc-bmp/bmp-operation-api/services/openApi"
	sdk "coding.jd.com/aidc-bmp/bmp-operation-api/services/openApi"
	requestTypes "coding.jd.com/aidc-bmp/bmp-operation-api/types/request"
	"coding.jd.com/aidc-bmp/bmp-operation-api/types/response"

	sdkModels "coding.jd.com/aidc-bmp/bmp-openapi-sdk/models"
	"coding.jd.com/aidc-bmp/bmp-operation-api/util"
	log "git.jd.com/cps-golang/log"
)

// 【设备管理】【开机】
func StartInstance(logger *log.Logger, param *requestTypes.StartInstanceRequest) (*response.CommonResponse, error) {
	logid := logger.GetPoint("logid").(string)
	req := sdkInstanceParameters.NewStartProjectInstanceParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	req.WithInstanceID(param.InstanceID)
	logger.Info("StartProjectInstance openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.Instance.StartProjectInstance(req, nil)
	logger.Info("StartProjectInstance openapi sdk resp:", util.InterfaceToJson(client))

	if err != nil || client.Payload == nil {
		logger.Info("openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}

	res := &response.CommonResponse{}
	if client.Payload.Result.Success {
		res.Success = true
	} else {
		res.Success = false
	}
	return res, nil
}

// 【设备管理】【重启】
func RestartInstance(logger *log.Logger, param *requestTypes.RestartInstanceRequest) (*response.CommonResponse, error) {
	logid := logger.GetPoint("logid").(string)
	req := sdkInstanceParameters.NewRestartProjectInstanceParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	req.WithInstanceID(param.InstanceID)
	logger.Info("RestartProjectInstance openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.Instance.RestartProjectInstance(req, nil)
	logger.Info("RestartProjectInstance openapi sdk resp:", util.InterfaceToJson(client))

	if err != nil || client.Payload == nil {
		logger.Info("openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}

	res := &response.CommonResponse{}
	if client.Payload.Result.Success {
		res.Success = true
	} else {
		res.Success = false
	}
	return res, nil
}

// 【设备管理】【关机】
func StopInstance(logger *log.Logger, param *requestTypes.StopInstanceRequest) (*response.CommonResponse, error) {
	logid := logger.GetPoint("logid").(string)
	req := sdkInstanceParameters.NewStopProjectInstanceParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	req.WithInstanceID(param.InstanceID)
	logger.Info("StopProjectInstance openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.Instance.StopProjectInstance(req, nil)
	logger.Info("StopProjectInstance openapi sdk resp:", util.InterfaceToJson(client))

	if err != nil || client.Payload == nil {
		logger.Info("openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}

	res := &response.CommonResponse{}
	if client.Payload.Result.Success {
		res.Success = true
	} else {
		res.Success = false
	}
	return res, nil
}

// 【设备管理
func ResetInstanceStatus(logger *log.Logger, instanceId string) (*response.CommonResponse, error) {
	logid := logger.GetPoint("logid").(string)
	req := sdkInstanceParameters.NewResetInstanceStatusParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	req.WithInstanceID(instanceId)
	logger.Info("ResetInstanceStatus openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.Instance.ResetInstanceStatus(req, nil)
	logger.Info("ResetInstanceStatus openapi sdk resp:", util.InterfaceToJson(client))

	if err != nil || client.Payload == nil {
		logger.Info("openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}

	res := &response.CommonResponse{}
	if client.Payload.Result.Success {
		res.Success = true
	} else {
		res.Success = false
	}
	return res, nil
}

// 【设备管理】【关机】
func DeleteInstance(logger *log.Logger, instanceId string) (*response.CommonResponse, error) {
	logid := logger.GetPoint("logid").(string)
	req := sdkInstanceParameters.NewDeleteProjectInstanceParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	req.WithInstanceID(instanceId)
	logger.Info("DeleteProjectInstance openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.Instance.DeleteProjectInstance(req, nil)
	logger.Info("DeleteProjectInstance openapi sdk resp:", util.InterfaceToJson(client))

	if err != nil || client.Payload == nil {
		logger.Info("openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}

	res := &response.CommonResponse{}
	if client.Payload.Result.Success {
		res.Success = true
	} else {
		res.Success = false
	}
	return res, nil
}

func LockInstance(logger *log.Logger, instanceId string) (bool, error) {
	logid := logger.GetPoint("logid").(string)
	request := sdkInstanceParameters.NewLockProjectInstanceParams()
	request.WithInstanceID(instanceId)
	request.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	request.WithBmpUserID(&userId)
	language := logger.GetPoint("language").(string)
	request.WithBmpLanguage(&language)
	logger.Info("lockinstance openapi request:", util.InterfaceToJson(request))
	responseApi, err := openApi.SdkClient.Instance.LockProjectInstance(request, nil)
	logger.Info("lockinstance openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("lockinstance openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return false, err
		}
		return false, errors.New(respErr.Message)
	}
	return responseApi.Payload.Result.Success, nil
}

func UnLockInstance(logger *log.Logger, instanceId string) (bool, error) {
	logid := logger.GetPoint("logid").(string)
	request := sdkInstanceParameters.NewUnLockProjectInstanceParams()
	request.WithInstanceID(instanceId)
	request.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	request.WithBmpUserID(&userId)
	language := logger.GetPoint("language").(string)
	request.WithBmpLanguage(&language)
	logger.Info("unlockinstance openapi request:", util.InterfaceToJson(request))
	responseApi, err := openApi.SdkClient.Instance.UnLockProjectInstance(request, nil)
	logger.Info("unlockinstance openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("lockinstance openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return false, err
		}
		return false, errors.New(respErr.Message)
	}
	return responseApi.Payload.Result.Success, nil
}

// 【设备管理】【批量开机】
func StartInstances(logger *log.Logger, instanceIDs []string) (bool, error) {
	logid := logger.GetPoint("logid").(string)
	req := sdkInstanceParameters.NewStartInstancesParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)

	body := sdkModels.StartInstancesRequest{
		InstanceIds: instanceIDs,
	}
	req.WithBody(&body)

	logger.Info("StartInstances StartInstances openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.Instance.StartInstances(req, nil)
	logger.Info("StartInstances StartInstances openapi sdk resp:", util.InterfaceToJson(client))

	if err != nil || client.Payload == nil {
		logger.Info("openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return false, err
		}
		return false, errors.New(respErr.Message)
	}

	res := &response.CommonResponse{}
	if client.Payload.Result.Success {
		res.Success = true
	} else {
		res.Success = false
	}
	return true, nil
}

// 【设备管理】【批量关机】
func StopInstances(logger *log.Logger, instanceIDs []string) (bool, error) {
	logid := logger.GetPoint("logid").(string)
	req := sdkInstanceParameters.NewStopInstancesParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)

	body := sdkModels.StopInstancesRequest{
		InstanceIds: instanceIDs,
	}
	req.WithBody(&body)

	logger.Info("StopInstances openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.Instance.StopInstances(req, nil)
	logger.Info("StopInstances openapi sdk resp:", util.InterfaceToJson(client))

	if err != nil || client.Payload == nil {
		logger.Info("openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return false, err
		}
		return false, errors.New(respErr.Message)
	}

	res := &response.CommonResponse{}
	if client.Payload.Result.Success {
		res.Success = true
	} else {
		res.Success = false
	}
	return true, nil
}

// 【设备管理】【批量重启】
func RestartInstances(logger *log.Logger, instanceIDs []string) (bool, error) {
	logid := logger.GetPoint("logid").(string)
	req := sdkInstanceParameters.NewRestartInstancesParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)

	body := sdkModels.RestartInstancesRequest{
		InstanceIds: instanceIDs,
	}
	req.WithBody(&body)

	logger.Info("RestartInstances openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.Instance.RestartInstances(req, nil)
	logger.Info("RestartInstances openapi sdk resp:", util.InterfaceToJson(client))

	if err != nil || client.Payload == nil {
		logger.Info("openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return false, err
		}
		return false, errors.New(respErr.Message)
	}

	res := &response.CommonResponse{}
	if client.Payload.Result.Success {
		res.Success = true
	} else {
		res.Success = false
	}
	return true, nil
}

// 【设备管理】【批量编辑实例名称】
func ModifyInstances(logger *log.Logger, instanceIDs []string, instanceName string) (bool, error) {
	logid := logger.GetPoint("logid").(string)
	req := sdkInstanceParameters.NewModifyInstancesParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)

	body := sdkModels.ModifyInstancesRequest{
		InstanceIDs:  instanceIDs,
		InstanceName: instanceName,
	}
	req.WithBody(&body)

	logger.Info("ModifyInstances openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.Instance.ModifyInstances(req, nil)
	logger.Info("ModifyInstances openapi sdk resp:", util.InterfaceToJson(client))

	if err != nil || client.Payload == nil {
		logger.Info("openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return false, err
		}
		return false, errors.New(respErr.Message)
	}

	res := &response.CommonResponse{}
	if client.Payload.Result.Success {
		res.Success = true
	} else {
		res.Success = false
	}
	return true, nil
}

// 【设备管理】【批量删除】
func DeleteInstances(logger *log.Logger, instanceIDs []string) (bool, error) {
	logid := logger.GetPoint("logid").(string)
	req := sdkInstanceParameters.NewDeleteInstancesParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)

	body := sdkModels.DeleteInstancesRequest{
		InstanceIDs: instanceIDs,
	}
	req.WithBody(&body)

	logger.Info("DeleteInstances openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.Instance.DeleteInstances(req, nil)
	logger.Info("DeleteInstances openapi sdk resp:", util.InterfaceToJson(client))

	if err != nil || client.Payload == nil {
		logger.Info("openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return false, err
		}
		return false, errors.New(respErr.Message)
	}

	res := &response.CommonResponse{}
	if client.Payload.Result.Success {
		res.Success = true
	} else {
		res.Success = false
	}
	return true, nil
}

func ResetInstancePassword(logger *log.Logger, instanceId string, password string) (bool, error) {

	logid := logger.GetPoint("logid").(string)
	request := sdkInstanceParameters.NewResetProjectInstancePasswdParams()
	request.WithInstanceID(instanceId)
	request.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	request.WithBmpUserID(&userId)
	language := logger.GetPoint("language").(string)
	request.WithBmpLanguage(&language)

	body := sdkModels.ResetInstancePasswdRequest{
		Password: password,
	}
	request.WithBody(&body)

	logger.Info("ResetProjectInstancePasswd openapi request:", util.InterfaceToJson(request))
	responseApi, err := openApi.SdkClient.Instance.ResetProjectInstancePasswd(request, nil)
	logger.Info("ResetProjectInstancePasswd openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("ResetProjectInstancePasswd openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return false, err
		}
		return false, errors.New(respErr.Message)
	}
	return responseApi.Payload.Result.Success, nil
}

func ResetInstancesPassword(logger *log.Logger, instanceIds []string, password string) (bool, error) {

	logid := logger.GetPoint("logid").(string)
	request := sdkInstanceParameters.NewResetProjectInstancesPasswdParams()
	request.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	request.WithBmpUserID(&userId)
	language := logger.GetPoint("language").(string)
	request.WithBmpLanguage(&language)

	body := sdkModels.ResetInstancesPasswdRequest{
		Password:    password,
		InstanceIDs: instanceIds,
	}
	request.WithBody(&body)

	logger.Info("ResetProjectInstancesPasswd openapi request:", util.InterfaceToJson(request))
	responseApi, err := openApi.SdkClient.Instance.ResetProjectInstancesPasswd(request, nil)
	logger.Info("ResetProjectInstancesPasswd openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("ResetProjectInstancesPasswd openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return false, err
		}
		return false, errors.New(respErr.Message)
	}
	return responseApi.Payload.Result.Success, nil
}

//ReinstallInstance sdk重装
func ReinstallInstance(logger *log.Logger, param requestTypes.ReinstallInstanceRequest) (bool, error) {

	logid := logger.GetPoint("logid").(string)

	req := sdkInstanceParameters.NewReinstallProjectInstanceParams()
	req.WithTraceID(logid)
	req.WithBody(&param.ReinstallInstanceRequest)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithInstanceID(param.InstanceId)
	logger.Info("ReinstallProjectInstance openapi request:", util.InterfaceToJson(req))
	responseApi, err := openApi.SdkClient.Instance.ReinstallProjectInstance(req, nil)
	logger.Info("ReinstallProjectInstance openapi response:", util.InterfaceToJson(responseApi))
	fmt.Println("ReinstallProjectInstance openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("ReinstallProjectInstance openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return false, err
		}
		return false, errors.New(respErr.Message)
	}

	return responseApi.Payload.Result.Success, nil
}

func DescribeInstance(logger *log.Logger, param requestTypes.QueryInstanceRequest) (*response.InstanceInfo, error) {

	logid := logger.GetPoint("logid").(string)
	request := sdkInstanceParameters.NewDescribeProjectInstanceParams()
	userId := logger.GetPoint("userId").(string)
	request.WithBmpUserID(&userId)
	request.WithTraceID(logid)
	request.WithInstanceID(param.InstanceId)
	language := logger.GetPoint("language").(string)
	request.WithBmpLanguage(&language)
	logger.Info("DescribeProjectInstance openapi request:", util.InterfaceToJson(request))
	responseApi, err := openApi.SdkClient.Instance.DescribeProjectInstance(request, nil)
	logger.Info("DescribeProjectInstance openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("DescribeProjectInstance openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}

	obj := response.Instance{}
	v := responseApi.Payload.Result.Instance

	obj = SdkInstance2RespInstance(logger, *v)

	return &response.InstanceInfo{
		Instance: obj,
	}, nil
}

func SdkInstance2RespInstance(logger *log.Logger, v sdkModels.Instance) response.Instance {
	/*
		CPU：Intel E5-2650V4（2*12物理核，2.2GHz）
		内存：256GB DDR4
		系统盘：300GB (2*300GB SAS, RAID1)
		数据盘：72TB (12*6TB SATA, NO RAID)
		网卡：2*25GE
		GPU：P40*4
	*/
	//cputpl := "%s %s (%d*%d物理核, %sGHz)"
	//memtpl := "%dGB %s %dGHZ"
	//systemtpl := "%dGB(%dGB %s %s*%d,%s)"
	//datatpl := "%dGB(%dGB %s %s*%d,%s)"
	//nictpl := "%d*%dGE"
	//gputpl := "%s %s*%d"
	language := logger.GetPoint("language").(string)
	//if language == constants.LANGUAGE_EN {
	//	cputpl = "%s %s (%d*%dcores, %sGHz)"
	//}

	//raidMap := make(map[string]response.DeviceTypeRaid) //	请求一次，避免循环请求

	// raid, _ := raidLogic.QueryRaid(logger, v.SystemVolumeRaidID)
	// raidStr := ""
	// if raid != nil {
	// 	raidStr = raid.Name
	// }

	cpu_ := strconv.Itoa(int(v.CPUAmount)) + "*" + strconv.Itoa(int(v.CPUCores)) + "物理核," + v.CPUFrequency + "GHz"
	if language == BaseLogic.EN_US {
		cpu_ = strconv.Itoa(int(v.CPUAmount)) + "*" + strconv.Itoa(int(v.CPUCores)) + "cores," + v.CPUFrequency + "GHz"
	}
	mem_ := strconv.Itoa(int(v.MemAmount)*int(v.MemSize)) + "GB(" + strconv.Itoa(int(v.MemSize)) + "GB*" + strconv.Itoa(int(v.MemAmount)) + ")" + v.MemType + " " + strconv.Itoa(int(v.MemFrequency)) + "MHz"
	//系统盘：显示总容量 （单块容量 接口类型 类型*数量 RAID 模式）
	//240GB 240GB SATA SSD*2 RAID1
	_svcap := int(v.SystemVolumeAmount) * int(v.SystemVolumeSize)
	if strings.ToUpper(v.SystemVolumeRaidName) == "RAID1" {
		_svcap = _svcap / 2
	}
	sv_, systemVolumeType := GetSystemDiskDetailByRaid(logger, v.InstanceVolumeRaids)
	// 显示品牌 显卡型号 个数，例 NVIDIA P40 * 4
	gpu_ := v.GpuManufacturer + " " + v.GpuModel + "*" + strconv.Itoa(int(v.GpuAmount))
	if v.GpuAmount == 0 {
		gpu_ = ""
	}
	nicInfo := strconv.Itoa(int(v.NicAmount)) + "*" + strconv.Itoa(int(v.NicRate)) + "GE"
	v.SystemVolumeType = systemVolumeType
	return response.Instance{
		Instance:   v,
		CpuInfo:    cpu_,
		MemInfo:    mem_,
		SystemInfo: sv_,
		NicInfo:    nicInfo,
		GpuInfo:    gpu_,
	}
}

func Trim(str string) string {
	s := strings.TrimRight(str, "0")
	s = strings.TrimRight(s, ".")
	return s
}

func GetSystemDiskDetailByRaid(logger *log.Logger, ivr []*sdkModels.InstanceVolumeRaid) (string, string) {

	volumeDetail := ""

	systemVolumeType := ""

	for _, vr := range ivr {
		if vr.Volume.VolumeType == "system" {

			v := vr.Volume
			systemVolumeType = v.DiskType

			// avaliableValue := .0
			all := 0.0
			volumeSize, _ := strconv.ParseFloat(v.VolumeSize, 64)
			size := volumeSize
			all = volumeSize * float64(v.VolumeAmount)
			if vr.Raid != nil {
				if util.InArray(vr.Raid.Name, []string{"RAID1", "RAID10"}) {
					all = all / 2
				} else if util.InArray(vr.Raid.Name, []string{"RAID51H"}) {
					all = volumeSize * float64(v.VolumeAmount-2)
				} else if util.InArray(vr.Raid.Name, []string{"RAID5"}) {
					all = volumeSize * float64(v.VolumeAmount-1)
				}

			}
			var allStr, sizeStr string
			if v.VolumeUnit == "TB" { //保留两位小数
				allStr = strconv.FormatFloat(all, 'f', 2, 64)
				sizeStr = strconv.FormatFloat(size, 'f', 2, 64)
			} else { //不保留小数
				allStr = strconv.Itoa(int(all))
				sizeStr = strconv.Itoa(int(size))

			}

			volumeDetail = allStr + v.VolumeUnit + "(" +
				(strconv.Itoa(int(v.VolumeAmount)) + "*" + sizeStr + v.VolumeUnit) + " " + v.InterfaceType + ")"

		} else {
			continue
		}
	}

	//NVME另外算法 TODO
	return volumeDetail, systemVolumeType

}
