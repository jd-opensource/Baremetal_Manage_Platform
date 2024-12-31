package ImageLoigc

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"mime/multipart"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"coding.jd.com/aidc-bmp/bmp-operation-api/constant"

	sdkDeviceTypeParameters "coding.jd.com/aidc-bmp/bmp-openapi-sdk/client/device_type"
	sdkParameters "coding.jd.com/aidc-bmp/bmp-openapi-sdk/client/image"
	sdkModels "coding.jd.com/aidc-bmp/bmp-openapi-sdk/models"
	"coding.jd.com/aidc-bmp/bmp-operation-api/logic/BaseLogic"
	"coding.jd.com/aidc-bmp/bmp-operation-api/services/excel"
	sdk "coding.jd.com/aidc-bmp/bmp-operation-api/services/openApi"
	requestTypes "coding.jd.com/aidc-bmp/bmp-operation-api/types/request"
	"coding.jd.com/aidc-bmp/bmp-operation-api/types/response"
	"coding.jd.com/aidc-bmp/bmp-operation-api/util"
	log "git.jd.com/cps-golang/log"
	beego "github.com/beego/beego/v2/server/web"
)

// 【用户管理】【用户列表】
func DescribeImages(logger *log.Logger, param *requestTypes.QueryImagesRequest) (*response.ImageList, error) {

	logid := logger.GetPoint("logid").(string)
	req := sdkParameters.NewDescribeImagesParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	req.WithImageName(&param.ImageName)
	req.WithDeviceTypeID(&param.DeviceTypeID)
	req.WithIsAll(&param.IsAll)
	req.WithPageSize(&param.PageSize)
	req.WithPageNumber(&param.PageNumber)
	req.WithArchitecture(&param.Architecture)
	req.WithOsType(&param.OsType)
	req.WithSource(&param.Source)
	req.WithImageID(&param.ImageID)
	logger.Info("DescribeImages openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.Image.DescribeImages(req, nil)
	logger.Info("DescribeImages openapi sdk resp:", util.InterfaceToJson(client))

	if err != nil || client.Payload == nil {
		logger.Info("openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}
	resp := &response.ImageList{}
	if len(client.Payload.Result.Images) == 0 {
		resp.Images = make([]response.Image, 0)
		return resp, nil
	}
	for _, v := range client.Payload.Result.Images {
		obj := response.Image{} //不能放在外层！！！
		obj.ID = v.ID
		obj.ImageID = v.ImageID
		obj.ImageName = v.ImageName
		obj.OsID = v.OsID
		obj.Format = v.Format
		obj.BootMode = v.BootMode
		obj.Filename = v.Filename
		obj.URL = v.URL
		obj.Hash = v.Hash
		obj.Description = v.Description
		obj.Source = v.Source
		obj.SystemPartition = v.SystemPartition
		obj.OsVersion = v.OsVersion
		obj.OsType = v.OsType
		obj.CreatedBy = v.CreatedBy
		obj.CreatedTime = util.UtcToTime(v.CreatedTime)
		obj.UpdatedBy = v.UpdatedBy
		obj.UpdatedTime = util.UtcToTime(v.UpdatedTime)
		obj.Architecture = v.Architecture
		obj.SourceName = v.SourceName
		obj.IsBind = v.IsBind
		obj.DeviceTypeNum = int(v.DeviceTypeNum)
		resp.Images = append(resp.Images, obj)
	}
	resp.PageNumber = client.Payload.Result.PageNumber
	resp.PageSize = client.Payload.Result.PageSize
	resp.TotalCount = client.Payload.Result.TotalCount
	return resp, nil
}

// 【用户管理】【用户列表】
func DescribeImage(logger *log.Logger, imageId string) (*response.Image, error) {

	logid := logger.GetPoint("logid").(string)
	req := sdkParameters.NewDescribeImageParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	req.WithImageID(imageId)
	logger.Info("DescribeImage openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.Image.DescribeImage(req, nil)
	logger.Info("DescribeImage openapi sdk resp:", util.InterfaceToJson(client))
	fmt.Println("sss", err)
	if err != nil || client.Payload == nil {
		logger.Info("openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}
	v := client.Payload.Result.Image
	obj := &response.Image{} //不能放在外层！！！
	obj.ID = v.ID
	obj.ImageID = v.ImageID
	obj.ImageName = v.ImageName
	obj.OsID = v.OsID
	obj.Format = v.Format
	obj.BootMode = v.BootMode
	obj.Filename = v.Filename
	obj.URL = v.URL
	obj.Hash = v.Hash
	obj.Description = v.Description
	obj.Source = v.Source
	obj.SystemPartition = v.SystemPartition
	obj.OsVersion = v.OsVersion
	obj.OsType = v.OsType
	obj.CreatedBy = v.CreatedBy
	obj.CreatedTime = util.UtcToTime(v.CreatedTime)
	obj.UpdatedBy = v.UpdatedBy
	obj.UpdatedTime = util.UtcToTime(v.UpdatedTime)
	obj.Architecture = v.Architecture
	obj.SourceName = v.SourceName
	obj.IsBind = v.IsBind
	obj.DeviceTypeNum = int(v.DeviceTypeNum)

	return obj, nil
}
func ExportImage(logger *log.Logger, result []response.Image) (string, string, error) {
	language := logger.GetPoint("language").(string)
	e := excel.NewExcelInstance()
	index := e.SetNewSheet("Sheet1")
	sheetData := make([][]string, 0)
	for _, data := range result {
		if data.SystemPartition != "" {
			partition := []response.Partition{}
			if err := json.Unmarshal([]byte(data.SystemPartition), &partition); err != nil {
				panic(constant.BuildInvalidArgumentWithArgs("系统分区解析json错误", err.Error()))
			}
			systemPartition := ""
			for _, v := range partition {
				systemPartition = systemPartition + v.MountPoint + "：" + strconv.Itoa(v.Size/1024) + "GiB," + v.FsType + "；"
			}
			data.SystemPartition = systemPartition
		}

		sheetData = append(sheetData, []string{
			data.ImageName,
			data.ImageID,
			data.SourceName,
			data.Architecture,
			data.OsType,
			data.OsVersion,
			data.Format,
			data.BootMode,
			data.SystemPartition,
			data.Description,
			strconv.Itoa(data.DeviceTypeNum),
			data.CreatedTime,
		})
	}
	exportImageHeader := BaseLogic.ExportImageHeader
	if language == BaseLogic.EN_US {
		exportImageHeader = BaseLogic.ExportImageHeaderEn
	}
	if err := e.SetSheetData("Sheet1", exportImageHeader, sheetData); err != nil {
		logger.Warn("create excel error:", err.Error())
		return "", "", err
	}
	e.File.SetActiveSheet(index)
	rand.Seed(time.Now().UnixNano())
	downloadFileName := "image_list_" + time.Now().Format(BaseLogic.FormatDate) + "_" + strconv.Itoa(rand.Intn(900000)+100000) + ".xlsx" //生成100000到100000+900000之间的随机数，左闭右开
	fileName, err := e.SaveFile(downloadFileName)
	if err != nil {
		logger.Warn("保存excel到本地失败" + err.Error())
		return "", "", err
	}
	return fileName, downloadFileName, nil
}

//【镜像管理】【修改镜像描述】
func ModifyImage(logger *log.Logger, param *requestTypes.ModifyImageRequest, imageId string) (*response.CommonResponse, error) {

	logid := logger.GetPoint("logid").(string)
	req := sdkParameters.NewModifyImageParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	req.WithTraceID(logid)
	req.WithImageID(imageId)

	body := sdkModels.ModifyImageRequest{
		Description: param.Description,
	}
	req.WithBody(&body)

	c, err := sdk.SdkClient.Image.ModifyImage(req, nil)
	if err != nil {
		if err != nil {
			respErr, b := util.GetOpenapiError(err)
			if !b {
				logger.Warnf("unknown error in ModifyImage:%s", err.Error())
				return nil, errors.New("unknown error") //未知错误不暴露给用户
			}
			// respErr.Code
			// respErr.Status
			// respErr.Message
			logger.Warnf("ModifyImage error:%s", respErr.Message)
			return nil, errors.New(respErr.Message)
		}
	}

	return &response.CommonResponse{
		Success: c.Payload.Result.Success,
	}, nil
}

// 【镜像管理】【删除镜像】
func DeleteImage(logger *log.Logger, imageId string) (*response.CommonResponse, error) {
	logid := logger.GetPoint("logid").(string)
	userId := logger.GetPoint("userId").(string)
	reqInfo := sdkParameters.NewDescribeImagesParams()
	reqInfo.WithImageID(&imageId)
	reqInfo.WithBmpUserID(&userId)
	logger.Info("DescribeImages openapi sdk req:", util.InterfaceToJson(reqInfo))
	clientInfo, errInfo := sdk.SdkClient.Image.DescribeImages(reqInfo, nil)
	logger.Info("DescribeImages openapi sdk resp:", util.InterfaceToJson(clientInfo))
	if errInfo != nil {
		if errInfo != nil {
			respErr, b := util.GetOpenapiError(errInfo)
			if !b {
				logger.Warnf("unknown error in CreateImage:", errInfo.Error())
				return nil, errors.New("unknown error") //未知错误不暴露给用户
			}
			// respErr.Code
			// respErr.Status
			// respErr.Message
			logger.Warnf("DeleteImage error:", respErr.Message)
			return nil, errors.New(respErr.Message)
		}
	}

	req := sdkParameters.NewDeleteImageParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	req.WithBmpUserID(&userId)
	req.WithImageID(imageId)
	logger.Info("DeleteImage openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.Image.DeleteImage(req, nil)
	logger.Info("DeleteImage openapi sdk resp:", util.InterfaceToJson(client))

	if err != nil {
		if err != nil {
			respErr, b := util.GetOpenapiError(err)
			if !b {
				logger.Warnf("unknown error in CreateImage:", err.Error())
				return nil, errors.New("unknown error") //未知错误不暴露给用户
			}
			// respErr.Code
			// respErr.Status
			// respErr.Message
			logger.Warnf("DeleteImage error:", respErr.Message)
			return nil, errors.New(respErr.Message)
		}
	}
	res := &response.CommonResponse{}
	if client.Payload.Result.Success {
		arr := strings.Split(clientInfo.Payload.Result.Images[0].URL, "/")
		if len(arr) >= 2 {
			fileName := arr[len(arr)-1]
			uploadPath, _ := beego.AppConfig.String("upload.path")
			if err := os.Remove(uploadPath + "/" + fileName); err != nil {
				return nil, nil //即便删除文件失败，也不报错 errors.New("删除上传的镜像文件失败")
			}
		}
		res.Success = true
	} else {
		res.Success = false
	}
	return res, nil
}
func UploadImage(logger *log.Logger, param *requestTypes.UploadImageRequest, header *multipart.FileHeader, requestId string) (*response.ImageUpload, error) {
	//fmt.Println("文件大小", header.Size, util.GetMd5("qq123"), util.GetMd5("./log/test.txt"), util.Md5sum("./log/test.txt"))
	//os.Exit(1)
	//if header.Size > 16*1024 { //16K
	//	logger.Warn("file large exceed, size:", header.Size)
	//	return "", errors.New("文件大小超出范围")
	//}
	arr := strings.Split(header.Filename, ".")
	//suffix := strings.ToLower(arr[len(arr)-1])
	suffix := strings.ToLower(arr[len(arr)-2]) + "." + strings.ToLower(arr[len(arr)-1])
	fmt.Println(suffix)
	//AvailableSuffix := []string{"qcow2", "xz"}
	if suffix != "tar.xz" && strings.ToLower(arr[len(arr)-1]) != "qcow2" { //util.InArray(suffix, AvailableSuffix) {
		logger.Warn("file format not support:", suffix)
		return nil, errors.New("文件类型不符合,必须是tar.xz或者qcow2类型")
	} else {
		if strings.ToLower(arr[len(arr)-1]) == "qcow2" {
			suffix = "qcow2"
		} else {
			suffix = "tar"
		}
	}
	dir, err := util.GetBaseDir()
	if err != nil {
		logger.Warn("获取当前根目录失败" + err.Error())
		return nil, err
	}
	uploadPath, _ := beego.AppConfig.String("upload.path")
	util.CreateDirIfNotExist(uploadPath)                                       //如果目录不存在，创建一个目录
	localFileName := path.Join(dir, uploadPath, requestId+"_"+header.Filename) //prefix+"-"+header.Filename
	if err = util.SaveToLocal(localFileName, param.ImageFile); err != nil {
		logger.Warn("保存excel到本地失败" + err.Error())
		return nil, err
	}
	imageUrl, _ := beego.AppConfig.String("image.url")
	res := &response.ImageUpload{
		Format:   suffix,
		FileName: header.Filename,
		Url:      imageUrl + requestId + "_" + header.Filename,
		Hash:     util.Md5sum(localFileName),
	}
	return res, nil
}

// 【镜像管理】【上传镜像】
func CreateImage(logger *log.Logger, param *requestTypes.CreateImageRequest) (response.ImageId, error) {
	logid := logger.GetPoint("logid").(string)
	req := sdkParameters.NewCreateImageParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)

	body := sdkModels.CreateImageRequest{
		ImageName:       &param.ImageName,
		Architecture:    &param.Architecture,
		OsType:          &param.OsType,
		Version:         &param.Version,
		Format:          &param.Format,
		BootMode:        &param.BootMode,
		Filename:        &param.FileName,
		URL:             &param.Url,
		Hash:            &param.Hash,
		Source:          BaseLogic.USER_DEFINED, //用户上传的，默认是自定义镜像
		Description:     param.Description,
		SystemPartition: param.SystemPartition,
		DataPartition:   param.DataPartition,
	}
	req.WithBody(&body)
	logger.Info("CreateImage openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.Image.CreateImage(req, nil)
	logger.Info("CreateImage openapi sdk resp:", util.InterfaceToJson(client))

	if err != nil {
		respErr, b := util.GetOpenapiError(err)
		if !b {
			logger.Warnf("unknown error in CreateImage:", err.Error())
			return response.ImageId{}, errors.New("unknown error") //未知错误不暴露给用户
		}
		// respErr.Code
		// respErr.Status
		// respErr.Message
		logger.Warnf("CreateImage error:", respErr.Message)
		return response.ImageId{}, errors.New(respErr.Message)
	}
	if client == nil || client.Payload == nil || client.Payload.Result == nil {
		return response.ImageId{}, errors.New("create image failed")
	}
	return response.ImageId{
		ImageId: client.Payload.Result.ImageID,
	}, nil
}

func DescribeImageDeviceTypes(logger *log.Logger, param *requestTypes.QueryImageDeviceTypesRequest) (*response.DeviceTypePage, error) {

	logid := logger.GetPoint("logid").(string)
	req := sdkParameters.NewDescribeImageDeviceTypesParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	req.WithIsAll(&param.IsAll)
	req.WithPageSize(&param.PageSize)
	req.WithPageNumber(&param.PageNumber)
	req.WithArchitecture(&param.Architecture)
	req.WithImageID(param.ImageID)
	req.WithIsBind(&param.IsBind)
	logger.Info("DescribeImageDeviceTypes openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.Image.DescribeImageDeviceTypes(req, nil)
	logger.Info("DescribeImageDeviceTypes openapi sdk resp:", util.InterfaceToJson(client))

	if err != nil || client.Payload == nil {
		logger.Info("openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}
	resp := &response.DeviceTypePage{}
	if len(client.Payload.Result.DeviceTypes) == 0 {
		resp.DeviceTypes = make([]response.DeviceType, 0)
		return resp, nil
	}
	for _, v := range client.Payload.Result.DeviceTypes {
		obj := response.DeviceType{} //不能放在外层！！！
		obj.DeviceTypeID = v.DeviceTypeID
		obj.Name = v.Name
		obj.IDcID = v.IDcID
		obj.IdcName = v.IDcName
		obj.IdcNameEn = v.IDcNameEn
		obj.DeviceType = v.DeviceType
		obj.Architecture = v.Architecture
		obj.BootMode = v.BootMode
		obj.DeviceSeries = v.DeviceSeries
		obj.DeviceSeriesName = v.DeviceSeriesName
		obj.Description = v.Description

		obj.CreatedBy = v.CreatedBy
		obj.CreatedTime = util.UtcToTime(v.CreatedTime)
		obj.UpdatedBy = v.UpdatedBy
		obj.UpdatedTime = util.UtcToTime(v.UpdatedTime)
		obj.Architecture = v.Architecture

		obj.InstanceStatus = v.InstanceStatus

		resp.DeviceTypes = append(resp.DeviceTypes, obj)
	}
	resp.PageNumber = client.Payload.Result.PageNumber
	resp.PageSize = client.Payload.Result.PageSize
	resp.TotalCount = client.Payload.Result.TotalCount
	return resp, nil
}
func AssociatedDeviceType(logger *log.Logger, param *requestTypes.AssociateImageDeviceTypeRequest) (bool, error) {

	logid := logger.GetPoint("logid").(string)
	req := sdkParameters.NewAssociatedDeviceTypeParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	array := []string{}
	if len(param.DeviceTypeIDs) != 0 {
		array = strings.Split(param.DeviceTypeIDs, ",")
	}
	body := sdkModels.AssociateDeviceTypeRequest{
		ImageID:       &param.ImageId,
		DeviceTypeIDs: array,
	}
	req.WithBody(&body)
	logger.Info("AssociatedDeviceType openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.Image.AssociatedDeviceType(req, nil)
	logger.Info("AssociatedDeviceType openapi sdk resp:", util.InterfaceToJson(client))

	if err != nil || client.Payload == nil {
		logger.Info("openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return false, err
		}
		return false, errors.New(respErr.Message)
	}
	return client.Payload.Result.Success, err
}

func DissociatedDeviceType(logger *log.Logger, param *requestTypes.DisassociateDeviceTypeImageRequest) (bool, error) {

	logid := logger.GetPoint("logid").(string)
	req := sdkParameters.NewDissociatedDeviceTypeParams()
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	req.WithTraceID(logid)
	userId := logger.GetPoint("userId").(string)
	req.WithBmpUserID(&userId)
	body := sdkModels.DissociatedImageRequest{
		DeviceTypeID: &param.DeviceTypeID,
		ImageID:      &param.ImageId,
	}
	req.WithBody(&body)
	logger.Info("DissociatedDeviceType openapi sdk req:", util.InterfaceToJson(req))
	client, err := sdk.SdkClient.Image.DissociatedDeviceType(req, nil)
	logger.Info("DissociatedDeviceType openapi sdk resp:", util.InterfaceToJson(client))

	if err != nil || client.Payload == nil {
		logger.Info("openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return false, err
		}
		return false, errors.New(respErr.Message)
	}
	return client.Payload.Result.Success, err
}

func QueryImagesByDeviceType(logger *log.Logger, param requestTypes.QueryImagesByDeviceTypeRequest) (map[string][]*sdkModels.Image, error) {

	logid := logger.GetPoint("logid").(string)
	userId := logger.GetPoint("userId").(string)
	req := sdkDeviceTypeParameters.NewDescribeDeviceTypeImagesParams()
	req.WithTraceID(logid)
	req.WithBmpUserID(&userId)
	req.WithDeviceTypeID(&param.DeviceTypeID)
	language := logger.GetPoint("language").(string)
	req.WithBmpLanguage(&language)
	// req.WithIdcID(&param.IdcID)
	logger.Info("DescribeImages openapi request:", util.InterfaceToJson(req))
	responseApi, err := sdk.SdkClient.DeviceType.DescribeDeviceTypeImages(req, nil)
	logger.Info("DescribeImages openapi response:", util.InterfaceToJson(responseApi))
	if err != nil {
		logger.Info("DescribeImages openapi error:", err.Error())
		respErr, b := util.GetOpenapiError(err)
		if !b {
			return nil, err
		}
		return nil, errors.New(respErr.Message)
	}
	res := map[string][]*sdkModels.Image{}
	images := responseApi.Payload.Result.Images
	for _, v := range images {
		if _, ok := res[v.OsType]; !ok {
			res[v.OsType] = []*sdkModels.Image{v}
		} else {
			res[v.OsType] = append(res[v.OsType], v)
		}
	}
	return res, nil
}
