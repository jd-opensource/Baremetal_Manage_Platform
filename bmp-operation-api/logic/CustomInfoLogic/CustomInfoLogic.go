package CustomInfoLogic

import (
	"strings"

	"coding.jd.com/aidc-bmp/bmp-operation-api/logic/BaseLogic"
	"coding.jd.com/aidc-bmp/bmp-operation-api/logic/RaidLogic"
	rdsUtil "coding.jd.com/aidc-bmp/bmp-operation-api/services/redis"
	"coding.jd.com/aidc-bmp/bmp-operation-api/types/response"

	//"coding.jd.com/bmp/ironic-console-jdstack/thirdpart/redis"
	"encoding/json"

	requestTypes "coding.jd.com/aidc-bmp/bmp-operation-api/types/request"
	"coding.jd.com/aidc-bmp/bmp-operation-api/util"
	"git.jd.com/cps-golang/ironic-common/exception"
	log "git.jd.com/cps-golang/log"
	"github.com/go-redis/redis"
)

func SetCustomInfo(logger *log.Logger, param requestTypes.SetCustomInfoRequest) *response.CommonResponse {
	userId := logger.GetPoint("userId").(string)
	redisValue := make(map[string]map[string]map[string]bool, 0)
	redisValue["customInfo"] = param.PageValue
	err := rdsUtil.RedisCli.Set("custom_info_operation_"+userId+":"+param.PageKey, util.InterfaceToJson(redisValue), 0).Err()
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
		rdsUtil.RedisCli.Del("custom_info_operation_" + userId + ":" + param.PageKey)
	}
	result, err := rdsUtil.RedisCli.Get("custom_info_operation_" + userId + ":" + param.PageKey).Result()
	if err == redis.Nil { //如果key不存在，取默认值
		result = util.InterfaceToJson(BaseLogic.HeaderList[param.PageKey])
		//panic(exception.Exception{Msg: "redis name " + param.PageKey + " does not exist", Status: BaseLogic.ERROR_CODE})
	} else if err != nil {
		panic(exception.Exception{Msg: "get redis name " + param.PageKey + " failed", Status: BaseLogic.ERROR_CODE})
	}
	obj := make(map[string]map[string]map[string]bool, 0) //responseWeb.GetCustomInfoResponse{}
	err = json.Unmarshal([]byte(result), &obj)
	return &obj
}

func GetRaidconf(logger *log.Logger) map[int]map[string]map[string]string {
	raids, _ := RaidLogic.QueryRaidsAll(logger)
	m := map[string]string{}
	for _, raid := range raids.Raids {
		m[strings.ToUpper(raid.Name)] = raid.RaidID
	}
	r := map[int]map[string]map[string]string{}
	if logger.GetPoint("language").(string) == BaseLogic.EN_US {
		//r[2]["RAID"] = map[string]string{"RAID0": m["RAID0"], "RAID1": m["RAID1"], "RAID0-stripping": m["NO RAID"]}
		r = map[int]map[string]map[string]string{
			1: map[string]map[string]string{
				"NO RAID": map[string]string{"NO RAID": m["NO RAID"]},
			},
			2: map[string]map[string]string{
				"RAID": map[string]string{"RAID0": m["RAID0"], "RAID1": m["RAID1"], "RAID5": m["RAID5"], "RAID10": m["RAID10"]},
			},
			3: map[string]map[string]string{
				"RAID0-stripping": map[string]string{"RAID0": m["RAID0"]},
			},
		}
	} else {
		r = map[int]map[string]map[string]string{
			1: map[string]map[string]string{
				"NO RAID": map[string]string{"NO RAID": m["NO RAID"]},
			},
			2: map[string]map[string]string{
				//"NO RAID": map[string]string{"NO RAID": m["NO RAID"]},
				"RAID": map[string]string{"RAID0": m["RAID0"], "RAID1": m["RAID1"], "RAID5": m["RAID5"], "RAID10": m["RAID10"]},
			},
			3: map[string]map[string]string{
				"单盘RAID0": map[string]string{"RAID0": m["RAID0"]},
			},
		}
	}
	return r
}
