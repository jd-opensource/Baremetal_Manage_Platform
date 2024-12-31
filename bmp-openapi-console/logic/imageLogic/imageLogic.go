package imageLogic

import (
	"fmt"

	"coding.jd.com/aidc-bmp/bmp-openapi/dao/deviceTypeDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/idcDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/instanceDao"

	"coding.jd.com/aidc-bmp/bmp-openapi/constant"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/rDeviceTypeImageDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/baseLogic"

	imageDao "coding.jd.com/aidc-bmp/bmp-openapi/dao/imageDao"
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
	responseTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/response"
	"coding.jd.com/aidc-bmp/bmp-openapi/util"
	log "coding.jd.com/aidc-bmp/bmp_log"
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
