package scheduler

import (
	"fmt"
	"strings"
	"time"

	beego "github.com/beego/beego/v2/server/web"

	"coding.jd.com/aidc-bmp/bmp-scheduler/dao/instanceDao"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/instanceLogic"
	"coding.jd.com/aidc-bmp/bmp-scheduler/service/redis"
	"git.jd.com/cps-golang/ironic-common/ironic/common/Constants"
	commonUtil "git.jd.com/cps-golang/ironic-common/ironic/util"
	log "git.jd.com/cps-golang/log"
)

func doUpdateInstanceStatusCron() error {
	logPath, _ := beego.AppConfig.String("log.path")
	logger := log.New(logPath + "/bmp-scheduler-updateInstanceCron.log")
	logger.SetStableFields([]string{"method", "uri", "header", "request", "response"})
	logger.TimeStart("all_t")
	logger.TimeStart("self_t")
	logger.Point("task_type", "UpdateInstanceStatusCron")
	//没有requestid，临时生成
	logger.Point("logid", commonUtil.GenerateRandUuid())
	defer logger.Flush()
	defer logger.TimeEnd("self_t")
	defer logger.TimeEnd("all_t")

	key := "SchedulerLock:cps-ironic:update_instance_status_cron"
	is_set, err := redis.SetNxObjectToRedis(key, "1", 60)
	if err != nil {
		logger.Warn("CheckTimeoutCommand SetNxObjectToRedis error:", key, err.Error())
	}
	if is_set {
		check_instances, err := redis.Keys(fmt.Sprintf("%s*", Constants.REDIS_KEY_CHECK))
		if err != nil {
			logger.Warn("redis.Keys command error:", err.Error())
		}
		logger.Info("check_instances is:", check_instances)
		for _, key_check_instance := range check_instances {
			check_parts := strings.Split(key_check_instance, ":")
			if len(check_parts) < 2 {
				logger.Warn("invalid check key:", key_check_instance)
				continue
			}
			instance_id := check_parts[1]
			instance_key := fmt.Sprintf("%s:%s", Constants.REDIS_KEY_INSTANCE, instance_id)
			r, err := redis.Keys(instance_key)

			if err != nil {
				logger.Warn("redis.Keys command error:", instance_id, err.Error())
			}
			if len(r) > 0 { //Constants.REDIS_KEY_INSTANCE_instance_id 存在时，说明命令还没执行完，本次不操作，等待下一次cron
				continue
			}
			instance_entity, err := instanceLogic.GetByInstanceId(logger, instance_id)
			if err != nil {
				logger.Warn("GetByInstanceId error:", instance_id, err.Error())
				continue
			}

			status, err := redis.GetObjectFromRedis(key_check_instance)
			if err != nil {
				logger.Warn("GetObjectFromRedis error:", key_check_instance, err.Error())
				continue
			}
			instance_entity.Status = strings.ToLower(status)
			instance_entity.UpdatedTime = int(time.Now().Unix())
			if err := instanceDao.UpdateInstanceById(logger, instance_entity); err != nil {
				logger.Warn("UpdateInstanceById sql error:", instance_id, err.Error())
				continue
			}

			if err := redis.DelObjectFromRedis([]string{key_check_instance}); err != nil {
				logger.Warn("DelObjectFromRedis error:", key_check_instance, err.Error())
				continue
			}
			logger.Infof("changed instanceId:%s to status: %s success", instance_id, status)
		}
	}
	return nil
}
