package CustomInfoLogic

import (
	"coding.jd.com/aidc-bmp/bmp-console-api/logic/BaseLogic"
	rdsUtil "coding.jd.com/aidc-bmp/bmp-console-api/thirdpart/redis"
	"coding.jd.com/aidc-bmp/bmp-console-api/types/response"
	"fmt"

	//"coding.jd.com/bmp/ironic-console-jdstack/thirdpart/redis"
	"encoding/json"

	requestTypes "coding.jd.com/aidc-bmp/bmp-console-api/types/request"
	"coding.jd.com/aidc-bmp/bmp-console-api/util"
	log "coding.jd.com/aidc-bmp/bmp_log"
	"git.jd.com/cps-golang/ironic-common/exception"
	"github.com/go-redis/redis"
)

func SetCustomInfo(logger *log.Logger, param requestTypes.SetCustomInfoRequest) *response.CommonResponse {
	userId := logger.GetPoint("userId").(string)
	redisValue := make(map[string]map[string]map[string]bool, 0)
	redisValue["customInfo"] = param.PageValue
	err := rdsUtil.RedisCli.Set("custom_info_console_"+userId+":"+param.PageKey, util.InterfaceToJson(redisValue), 0).Err()
	if err != nil {
		panic(exception.Exception{Msg: "redis set " + param.PageKey + " error", Status: BaseLogic.ERROR_CODE})
	}
	res := response.CommonResponse{
		Success: true,
	}
	return &res
}
func GetCustomInfo(logger *log.Logger, param requestTypes.QueryCustomInfoRequest) *map[string]map[string]map[string]bool {
	userId := logger.GetPoint("userId").(string)
	if param.Reload == "true" {
		rdsUtil.RedisCli.Del("custom_info_console_" + userId + ":" + param.PageKey)
	}
	result, err := rdsUtil.RedisCli.Get("custom_info_console_" + userId + ":" + param.PageKey).Result()
	if err == redis.Nil { //如果key不存在，取默认值
		fmt.Println("meiqudao")
		result = util.InterfaceToJson(BaseLogic.HeaderList[param.PageKey])
		//panic(exception.Exception{Msg: "redis name " + param.PageKey + " does not exist", Status: BaseLogic.ERROR_CODE})
	} else if err != nil {
		panic(exception.Exception{Msg: "get redis name " + param.PageKey + " failed", Status: BaseLogic.ERROR_CODE})
	}
	obj := make(map[string]map[string]map[string]bool, 0) //responseWeb.GetCustomInfoResponse{}
	err = json.Unmarshal([]byte(result), &obj)
	return &obj
}
