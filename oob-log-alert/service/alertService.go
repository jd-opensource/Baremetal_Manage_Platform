package service

import (
	"fmt"
	"strings"
	"time"

	log "coding.jd.com/aidc-bmp/bmp_log"

	"coding.jd.com/aidc-bmp/oob-log-alert/object"

	"coding.jd.com/aidc-bmp/oob-log-alert/util"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/go-redis/redis"
)

// Alert for send Alert main
type Alert struct {
	pubsub    *redis.PubSub
	ch        <-chan *redis.Message
	IsSaveLog bool
	LogPath   string
}

// InitAlert initial Alert ,subscribe to Redis,
func InitAlert(isSaveLog bool, logPath string) error {

	a := &Alert{IsSaveLog: isSaveLog, LogPath: logPath}
	// init subscribe
	a.pubsub = util.PSubEventToRedis(object.LogChannel, object.OOBErrorChannel)

	_, err := a.pubsub.Receive()
	if err != nil {
		logs.Critical("psubscribe error:%s", err.Error())
		return err
	}

	logs.Debug("InitAlert config", a)
	a.Run()
	return nil

}

// Run Alert Module run in a goroutine
func (a Alert) Run() {

	// Consume messages.
	go func(a *Alert) {

		a.ch = a.pubsub.Channel()
		logs.Debug("psubcribe success:", a.pubsub)
		logs.Debug("waiting message...")

		for msg := range a.ch {

			tmp := strings.Split(msg.Channel, ":")
			channel := tmp[len(tmp)-2]
			deviceSN := tmp[len(tmp)-1]

			// check
			if msg.Payload == "" {
				continue
			}

			if dailyMoreThanOnce(deviceSN, msg.Payload) { //同一设备同类型的报警，每天只报一次
				continue
			}

			logPath, _ := beego.AppConfig.String("log.path")
			logger := log.New(logPath + "/msg-from-oob_agent.log")
			logger.SetStableFields([]string{"method", "uri", "header", "request", "response"})
			//agent过来的消息只有sn，没有requestid，临时生成
			logger.Point("logid", util.GenerateRandUuid())
			logger.TimeStart("all_t")
			logger.TimeStart("self_t")
			defer logger.Flush()
			defer logger.TimeEnd("self_t")
			defer logger.TimeEnd("all_t")

			// switch which type of message
			var err error
			switch channel {
			case object.LogChannel:
				logger.Point("task", "LogChannel")
				err = LogEvent(logger, deviceSN, msg.Payload, a.IsSaveLog, a.LogPath)
			case object.OOBErrorChannel:
				logger.Point("task", "OOBErrorChannel")
				err = CPSEvent(deviceSN, msg.Payload)
			default:
				continue
			}

			if err != nil {
				logs.Error("Alert Receive ERROR MSG:", msg, err.Error())
			}
		}
	}(&a)

}

func dailyMoreThanOnce(sn, payload string) bool {
	if len(payload) > 1024 {
		payload = payload[:1024]
	}
	k := fmt.Sprintf("%s_day_only_once_%s", sn, payload)
	if v, err := util.GetObjectFromRedis(k); err == nil && v == "true" { //已经报过
		logs.Warn("Alert more than once,sn:%s, msg:%s", sn, payload)
		return true
	}

	current := time.Now()
	t, _ := time.ParseInLocation("2006-01-02", current.Format("2006-01-02"), time.Local)
	e := t.Unix() + 86400
	c := current.Unix()
	ttl := int(e - c)
	if err := util.SetObjectToRedisWithExpire(k, "true", ttl); err != nil {
		logs.Warn("Alert set once daily info error, sn:%s, msg:%s, err:%+v", sn, payload, err)
	}
	return false

}
