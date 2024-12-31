package imageLogic

import (
	"fmt"
	"time"

	"coding.jd.com/aidc-bmp/bmp-openapi/dao/deviceTypeDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/idcDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/instanceDao"

	"coding.jd.com/aidc-bmp/bmp-openapi/constant"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/rDeviceTypeImageDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/baseLogic"

	imageDao "coding.jd.com/aidc-bmp/bmp-openapi/dao/imageDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/osLogic"
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
	responseTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/response"
	"coding.jd.com/aidc-bmp/bmp-openapi/util"
	log "coding.jd.com/aidc-bmp/bmp_log"
	namespacePrefix "git.jd.com/cps-golang/ironic-common/ironic/common/NamespacePrefixs"
	commonUtil "git.jd.com/cps-golang/ironic-common/ironic/util"
)

func DImageEntity2Image(d *imageDao.DImage, tz string) *responseTypes.Image {
	return &responseTypes.Image{
		ID:              d.ID,
		ImageID:         d.ImageID,
		ImageName:       d.ImageName,
		Format:          d.Format,
		Filename:        d.Filename,
		URL:             d.URL,
		Hash:            d.Hash,
		Source:          d.Source,
		Description:     d.Description,
		SystemPartition: d.SystemPartition,
		DataPartition:   d.DataPartition,
		//Platform:        d.Platform,
		OsVersion:   d.OsVersion,
		CreatedTime: util.TimestampToString(int64(d.CreatedTime), tz),

		//Os: responseTypes.OS{
		//	Id:           d.OsId,
		//	NameEn:       d.OsNameEn,
		//	NameZh:       d.OsNameZh,
		//	Version:      d.OsVersion,
		//	Type:         d.OsType,
		//	Bits:         int(d.OsBits),
		//	Architecture: d.OsArchitecture,
		//	Platform:     d.OsPlatform,
		//},
	}
}

func ImageEntityWithOs2Image(logger *log.Logger, d *imageDao.DImage, deviceTypeNum int, isBind bool, tz string) *responseTypes.Image {
	image := &responseTypes.Image{
		ID:              d.ID,
		ImageID:         d.ImageID,
		ImageName:       d.ImageName,
		Format:          d.Format,
		BootMode:        d.BootMode,
		Filename:        d.Filename,
		URL:             d.URL,
		Hash:            d.Hash,
		Source:          d.Source,
		SourceName:      baseLogic.Source[d.Source],
		Architecture:    d.Architecture,
		Description:     d.Description,
		SystemPartition: d.SystemPartition,
		DataPartition:   d.DataPartition,
		OsID:            d.OsID,
		OsType:          d.OsType,
		OsName:          d.OsName,
		OsVersion:       d.OsVersion,
		CreatedBy:       d.CreatedBy,
		UpdatedBy:       d.UpdatedBy,
		CreatedTime:     util.TimestampToString(int64(d.CreatedTime), tz),
		UpdatedTime:     util.TimestampToString(int64(d.UpdatedTime), tz),
		DeviceTypeNum:   deviceTypeNum,
		IsBind:          isBind,
	}
	if logger.GetPoint("language").(string) == baseLogic.EN_US {
		image.SourceName = baseLogic.SourceEn[d.Source]
	}
	return image
}

func CreateImage(logger *log.Logger, param *requestTypes.CreateImageRequest) (string, error) {
	//_, err := osLogic.GetAndCheckById(logger, param.OsID)
	//if err != nil {
	//	logger.Warnf("Os not exist:osId:%s, error:%s", param.OsID, err.Error())
	//	return "", err
	//}
	list, _ := imageDao.GetAllImage(logger, map[string]interface{}{
		"image_name": param.ImageName,
		"is_del":     baseLogic.IS_NOT_DEL,
	})
	if len(list) != 0 {
		logger.Warn("imageName exist:", param.ImageName)
		panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("镜像名称 %s 已存在", param.ImageName), fmt.Sprintf("imageName %s exist", param.ImageName)))
	}
	bits := 64
	if param.Architecture == "i386" { //特殊处理
		bits = 32
	}
	osId, err := osLogic.CreateOS(logger, &requestTypes.CreateOSRequest{
		OsName:       param.ImageName,
		OsType:       param.OsType,
		Architecture: param.Architecture,
		Bits:         bits, //默认64位,i38是32位的！
		OsVersion:    param.Version,
		SysUser:      "root", //默认root用户
	})
	if err != nil {
		panic(constant.BuildInvalidArgumentWithArgs("os创建失败:"+err.Error(), "os created failed:"+err.Error()))
	}
	entity := &imageDao.Image{
		ImageName: param.ImageName,
		//ImageVersion:    param.ImageVersion,
		OsID:            osId,
		Format:          param.Format,
		BootMode:        param.BootMode,
		Filename:        param.Filename,
		URL:             param.Url,
		Hash:            param.Hash,
		Source:          param.Source,
		Description:     param.Description,
		SystemPartition: param.SystemPartition,
		DataPartition:   param.DataPartition,
		CreatedBy:       logger.GetPoint("username").(string),
		CreatedTime:     int(time.Now().Unix()),
		UpdatedBy:       logger.GetPoint("username").(string),
		UpdatedTime:     int(time.Now().Unix()),
	}
	entity.ImageID = commonUtil.GetRandString("i-", namespacePrefix.IMAGE_ID_LENGTH, false, true, true)
	if _, err := imageDao.AddImage(logger, entity); err != nil {
		logger.Warnf("CreateImage AddImage sql error, entity:%s, error:%s", util.ObjToJson(entity), err.Error())
		return "", err
	}
	return entity.ImageID, nil
}

func DescribeImages(logger *log.Logger, param *requestTypes.QueryImagesRequest, p util.Pageable) ([]*responseTypes.Image, int64, error) {
	tz := logger.GetPoint("timezone").(string)
	offset, limit := p.Pageable2offset()
	query := map[string]interface{}{
		"image_id":   param.ImageID,
		"image_name": param.ImageName,
		"version":    param.Version,
		"os_id":      param.OsID,
		//"image_ids":   strings.Join(param.ImageIds, ","),
		"source":         param.Source,
		"architecture":   param.Architecture,
		"os_type":        param.OsType,
		"is_del":         baseLogic.IS_NOT_DEL,
		"device_type_id": param.DeviceTypeID,
	}
	//queryList := query
	//query["count"] = true
	_, count, err := imageDao.GetImageAndOsList(logger, query, true, 0, 0)
	if err != nil {
		logger.Warnf("DescribeDeviceTypes_GetDeviceTypeCount dao error, query:%s, error:%s", util.ObjToJson(query), err.Error())
		return nil, 0, err
	}
	entityList := []*imageDao.DImage{}
	if param.IsAll == baseLogic.IS_ALL {
		//entityList, err = imageDao.GetMultiImage(logger, query, []string{}, []string{"id"}, []string{"desc"}, 0, 0)
		fmt.Println("获取全部", query)
		entityList, _, err = imageDao.GetImageAndOsList(logger, query, false, 0, 0)
	} else {
		//entityList, err = imageDao.GetMultiImage(logger, query, []string{}, []string{"id"}, []string{"desc"}, offset, limit)
		entityList, _, err = imageDao.GetImageAndOsList(logger, query, false, offset, limit)
	}
	if err != nil {
		logger.Warn("imageDao.QueryByDeviceOsType sql error, query:%s, error:%s", util.ObjToJson(query), err.Error())
		return nil, 0, err
	}
	images := []*responseTypes.Image{}
	for _, image := range entityList {
		isBind := false
		if param.DeviceTypeID != "" {
			deviceTypeImageNum, _ := rDeviceTypeImageDao.QueryDeviceTypeImagesCount(logger, map[string]interface{}{
				"image_id":       image.ImageID,
				"device_type_id": param.DeviceTypeID,
				"is_del":         baseLogic.IS_NOT_DEL,
			})
			if deviceTypeImageNum != 0 {
				isBind = true
			}
		}

		//os, err := osLogic.GetByOsId(logger, image.OsID)
		//if err != nil && os == nil {
		//	os = &responseTypes.Os{} //如果没有找到os，默认为空
		//}
		deviceTypeNum, err := rDeviceTypeImageDao.QueryDeviceTypeImagesCount(logger, map[string]interface{}{
			"image_id": image.ImageID,
			"is_del":   baseLogic.IS_NOT_DEL,
		})
		if err != nil {
			logger.Warn("获取镜像对应的机型个数错误", image.OsID, err.Error())
			return nil, 0, err
		}
		images = append(images, ImageEntityWithOs2Image(logger, image, int(deviceTypeNum), isBind, tz))
	}
	return images, count, nil
}

func GetByImageId(logger *log.Logger, imageId string) (*responseTypes.Image, error) {
	tz := logger.GetPoint("timezone").(string)
	entity, _, err := imageDao.GetImageAndOsList(logger, map[string]interface{}{"image_id": imageId}, false, 0, 1)
	if err != nil {
		logger.Warn("GetImageByUuid sql error:", imageId, err.Error())
		return nil, err
	}
	if len(entity) == 0 {
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}
	//os, err := osLogic.GetByOsId(logger, entity.OsID)
	//if err != nil {
	//	logger.Warn("GetByOsId sql error:", entity.OsID, err.Error())
	//	return nil, err
	//}
	deviceTypeNum, err := rDeviceTypeImageDao.QueryDeviceTypeImagesCount(logger, map[string]interface{}{
		"image_id": imageId,
		"is_del":   baseLogic.IS_NOT_DEL,
	})
	if err != nil {
		logger.Warn("获取镜像对应的机型个数错误", imageId, err.Error())
		return nil, err
	}
	return ImageEntityWithOs2Image(logger, entity[0], int(deviceTypeNum), false, tz), nil
}

func ModifyImage(logger *log.Logger, param *requestTypes.ModifyImageRequest, imageId string) error {
	//先废弃，暂时用不到这个接口
	//if param.ImageName != "" {
	//	image, _ := imageDao.GetAllImage(logger, map[string]interface{}{
	//		"image_name": param.ImageName,
	//		"is_del":     baseLogic.IS_NOT_DEL,
	//	})
	//	if len(image) != 0 && image[0].ImageName != param.ImageName {
	//		logger.Warn("imageName exist:", param.ImageName)
	//		panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("镜像名称 %s 已存在", param.ImageName), fmt.Sprintf("imageName %s exist", param.ImageName)))
	//	}
	//}
	entity, err := imageDao.GetImageByUuid(logger, imageId)
	if err != nil {
		logger.Warnf("GetImageByUuid sql error, image_id:%s, error:%s", imageId, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}
	//if param.ImageName != "" {
	//	if err := osLogic.ModifyOS(logger, &requestTypes.ModifyOSRequest{
	//		OsName: param.ImageName,
	//	}, entity.OsID); err != nil {
	//		logger.Warnf("ModifyOS UpdateOsById sql error, entity:%s, error:%s", util.ObjToJson(entity), err.Error())
	//		return err
	//	}
	//}
	//entity.ImageName = param.ImageName
	////entity.ImageVersion = param.ImageVersion
	//entity.OsID = param.OsID
	//entity.Format = param.Format
	//entity.Filename = param.Filename
	//entity.URL = param.Url
	//entity.Hash = param.Hash
	//entity.Source = param.Source
	//entity.SystemPartition = param.SystemPartition
	//	entity.DataPartition = param.DataPartition
	if param.Description != nil {
		entity.Description = *param.Description
	}

	entity.UpdatedBy = logger.GetPoint("username").(string)
	entity.UpdatedTime = int(time.Now().Unix())
	if err := imageDao.UpdateImageById(logger, entity); err != nil {
		logger.Warnf("ModifyImage UpdateImageById sql error, entity:%s, error:%s", util.ObjToJson(entity), err.Error())
		return err
	}
	return nil
}

func DeleteImage(logger *log.Logger, imageId string) error {
	entity, err := imageDao.GetImageByUuid(logger, imageId)
	if err != nil {
		logger.Warnf("DeleteImage %s GetImageByUuid sql error:", imageId, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}
	if entity.Source == baseLogic.COMMON {
		panic(constant.BuildInvalidArgumentWithArgs("预置镜像不能删除", "preImage not allowed delete"))
	}
	entity.IsDel = 1
	entity.UpdatedBy = logger.GetPoint("username").(string)
	entity.UpdatedTime = int(time.Now().Unix())
	entity.DeletedTime = int(time.Now().Unix())
	if err := imageDao.UpdateImageById(logger, entity); err != nil {
		logger.Warnf("DeleteImage %s UpdateImageById sql error:", imageId, err.Error())
		return err
	}
	//同时删除os表
	if err := osLogic.DeleteOS(logger, entity.OsID); err != nil {
		logger.Warnf("DeleteOs %s UpdateOsById sql error:", entity.OsID, err.Error())
		return err
	}
	return nil
}
func QueryImageDeviceTypes(logger *log.Logger, param *requestTypes.QueryImageDeviceTypesRequest, p util.Pageable) ([]*responseTypes.DeviceType, int64, error) {
	//tz := logger.GetPoint("timezone").(string)
	language := logger.GetPoint("language").(string)
	offset, limit := p.Pageable2offset()
	query := map[string]interface{}{
		"image_id":     param.ImageID,
		"architecture": param.Architecture,
		"is_del":       baseLogic.IS_NOT_DEL,
		"is_bind":      param.IsBind,
	}
	//queryList := query
	//query["count"] = true
	_, count, err := rDeviceTypeImageDao.GetImageDeviceTypeList(logger, query, true, 0, 0)
	//os.Exit(1)
	if err != nil {
		logger.Warnf("DescribeDeviceTypes_GetDeviceTypeCount dao error, query:%s, error:%s", util.ObjToJson(query), err.Error())
		return nil, 0, err
	}
	entityList := []*deviceTypeDao.DeviceType{}
	if param.IsAll == baseLogic.IS_ALL {
		//entityList, err = imageDao.GetMultiImage(logger, query, []string{}, []string{"id"}, []string{"desc"}, 0, 0)
		fmt.Println("获取全部", query)
		entityList, _, err = rDeviceTypeImageDao.GetImageDeviceTypeList(logger, query, false, 0, 0)
	} else {
		//entityList, err = imageDao.GetMultiImage(logger, query, []string{}, []string{"id"}, []string{"desc"}, offset, limit)
		entityList, _, err = rDeviceTypeImageDao.GetImageDeviceTypeList(logger, query, false, offset, limit)
	}
	if err != nil {
		logger.Warn("imageDao.QueryByDeviceOsType sql error, query:%s, error:%s", util.ObjToJson(query), err.Error())
		return nil, 0, err
	}
	fmt.Println(entityList)
	deviceTypes := []*responseTypes.DeviceType{}
	for _, entity := range entityList {
		//isBind := false
		//if param.DeviceTypeID != "" {
		//	deviceTypeImageNum, _ := rDeviceTypeImageDao.QueryDeviceTypeImagesCount(logger, map[string]interface{}{
		//		"image_id":       image.ImageID,
		//		"device_type_id": param.DeviceTypeID,
		//		"is_del":         baseLogic.IS_NOT_DEL,
		//	})
		//	if deviceTypeImageNum != 0 {
		//		isBind = true
		//	}
		//}
		//deviceTypeNum, err := rDeviceTypeImageDao.QueryDeviceTypeImagesCount(logger, map[string]interface{}{
		//	"image_id": image.ImageID,
		//	"is_del":   baseLogic.IS_NOT_DEL,
		//})
		//if err != nil {
		//	logger.Warn("获取镜像对应的机型个数错误", image.OsID, err.Error())
		//	return nil, 0, err
		//}
		//deviceType.DeviceTypeID
		idc, err := idcDao.GetIdcById(logger, entity.IDcID) //获取机房信息
		if err != nil {
			idc = &idcDao.Idc{}
		}
		idcName := idc.Name
		if language == baseLogic.EN_US {
			idcName = idc.NameEn
		}
		instanceList, err := instanceDao.GetAllInstance(logger, map[string]interface{}{
			"device_type_id": entity.DeviceTypeID,
			"image_id":       param.ImageID,
		})
		instanceStatus := []string{}
		for _, v := range instanceList {
			instanceStatus = append(instanceStatus, v.Status)
		}
		deviceSeries := baseLogic.DeviceTypeSeries[entity.DeviceSeries]
		if language == baseLogic.EN_US {
			deviceSeries = baseLogic.DeviceTypeSeriesEn[entity.DeviceSeries]
		}
		deviceType := &responseTypes.DeviceType{
			ID:               entity.ID,
			IDcID:            entity.IDcID,
			IDcName:          idcName,
			IDcNameEn:        idc.NameEn,
			DeviceType:       entity.DeviceType,
			DeviceTypeID:     entity.DeviceTypeID,
			Name:             entity.Name,
			Description:      entity.Description,
			DeviceSeries:     entity.DeviceSeries,
			DeviceSeriesName: deviceSeries,
			Architecture:     entity.Architecture,
			InstanceStatus:   instanceStatus,
		}
		fmt.Println(deviceType)
		deviceTypes = append(deviceTypes, deviceType) //deviceTypeLogic.DeviceTypeEntity2DeviceType(logger, deviceType)
	}
	return deviceTypes, count, nil
}
