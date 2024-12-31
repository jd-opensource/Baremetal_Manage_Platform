package resourceLogic

import (
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/deviceTypeDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/imageDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/userDao"
	requestTypes "coding.jd.com/aidc-bmp/bmp-openapi/types/request"
	"coding.jd.com/aidc-bmp/bmp-openapi/util"
	log "coding.jd.com/aidc-bmp/bmp_log"
	"fmt"
)

func DescribeResources(logger *log.Logger, param *requestTypes.QueryResourcesRequest) (int, error) {
	query := map[string]interface{}{}
	if param.Name != "" {
		query["name"] = param.Name
		list, err := deviceTypeDao.GetAllDeviceType(logger, query)
		if err != nil {
			logger.Warnf("GetAllDeviceType error, query:%s, error:%s", util.ObjToJson(query), err.Error())
			return 0, err
		}
		fmt.Println("查询到的数量", len(list))
		return len(list), err
	}
	if param.DeviceType != "" {
		query["device_type"] = param.DeviceType
		list, err := deviceTypeDao.GetAllDeviceType(logger, query)
		if err != nil {
			logger.Warnf("GetAllDeviceType error, query:%s, error:%s", util.ObjToJson(query), err.Error())
			return 0, err
		}
		fmt.Println("查询到的数量", len(list))
		return len(list), err
	}
	if param.ImageName != "" {
		query["image_name"] = param.ImageName
		list, err := imageDao.GetAllImage(logger, query)
		if err != nil {
			logger.Warnf("GetAllImage error, query:%s, error:%s", util.ObjToJson(query), err.Error())
			return 0, err
		}
		fmt.Println("查询到的数量", len(list))
		return len(list), err
	}
	if param.UserName != "" {
		query["user_name"] = param.UserName
		list, err := userDao.GetAllUser(logger, query)
		if err != nil {
			logger.Warnf("GetAllUser error, query:%s, error:%s", util.ObjToJson(query), err.Error())
			return 0, err
		}
		fmt.Println("查询到的数量", len(list))
		return len(list), err
	}
	return 0, nil
}
