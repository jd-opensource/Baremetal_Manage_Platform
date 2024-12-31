package service

import (
	"coding.jd.com/aidc-bmp/oob-log-alert/dao"
	"coding.jd.com/aidc-bmp/oob-log-alert/util"

	"encoding/json"

	log "coding.jd.com/aidc-bmp/bmp_log"
)

func SetCustomInfo(logger *log.Logger, param dao.SetCustomInfoRequest) dao.CommonResponse {
	userId := logger.GetPoint("userId").(string)
	key := "custom_info_operation_" + userId + ":" + param.PageKey
	redisValue := make(map[string]map[string]map[string]bool, 0)
	redisValue["custom_info"] = param.PageValue
	err := util.SetObjectToRedis(key, util.ToString(redisValue))
	if err != nil {
		logger.Warnf("SetCustomInfo to redis error, key:%s", key)
	}
	res := dao.CommonResponse{
		Success: true,
	}
	return res
}
func GetCustomInfo(logger *log.Logger, param dao.QueryCustomInfoRequest) map[string]map[string]map[string]bool {
	userId := logger.GetPoint("userId").(string)
	key := "custom_info_operation_" + userId + ":" + param.PageKey

	result, err := util.GetObjectFromRedis(key)
	if err != nil {
		logger.Warnf("GetCustomInfo error, err:%s", err.Error())
	}

	if param.Reload == "true" || result == "" {
		result := util.ToString(util.HeaderList[param.PageKey])
		go func() {
			err := util.SetObjectToRedis(key, result)
			if err != nil {
				logger.Warnf("GetCustomInfo.reload.reset to redis error, key:%s", key)
			}
		}()
		return util.HeaderList[param.PageKey]
	}

	obj := make(map[string]map[string]map[string]bool, 0) //responseWeb.GetCustomInfoResponse{}
	if err = json.Unmarshal([]byte(result), &obj); err != nil {
		logger.Warnf("GetCustomInfo unmarshal error, key:%s, content:%s, error:%s", key, result, err.Error())
		return nil
	}
	return obj
}
