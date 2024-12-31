package scheduler

import (
	"encoding/json"
	"runtime"
	"time"

	log "coding.jd.com/aidc-bmp/bmp_log"
	"coding.jd.com/aidc-bmp/oob-log-agent/object"
	"coding.jd.com/aidc-bmp/oob-log-agent/util"
	beego "github.com/beego/beego/v2/server/web"
)

// 定时任务
func CheckOoblog() error {
	logPath, _ := beego.AppConfig.String("log.path")
	logger := log.New(logPath + "/CheckOoblogCron.log")
	logger.SetStableFields([]string{"method", "uri", "header", "request", "response"})
	logger.TimeStart("all_t")
	logger.TimeStart("self_t")
	logger.Point("task_type", "CheckOoblogCron")
	//没有requestid，临时生成
	logger.Point("logid", util.GenerateRandUuid())
	defer logger.Flush()
	defer logger.TimeEnd("self_t")
	defer logger.TimeEnd("all_t")

	defer func() {
		if r := recover(); r != nil {
			logger.Warnf("CheckOoblog.Cron catch exception, error:%v, stack:%s", r, util.GetCurrentGoroutineStack())
		}
	}()

	routineCount := beego.AppConfig.DefaultInt("app.goroutine-num", runtime.NumCPU()*10)

	// get CPS list from redis
	cpsList, err := GetCPSFromRedis(logger)
	if err != nil {
		logger.Warnf("GetCPSListFromRedis Error : %s", err.Error())
		return err
	}

	// init gopool to avoid to many goroutines
	// set runtime option
	runtime.GOMAXPROCS(runtime.NumCPU())
	pool := util.NewPool(routineCount)
	for _, cps := range cpsList.CPSs {

		pool.Add()

		go func(cps object.CPSRecord) {
			defer pool.Done()

			c := make(chan int)
			go func(cps object.CPSRecord, c chan int) {
				job := Job{cps}
				job.Run(logger)
				c <- 1
			}(cps, c)
			select {
			case <-c:
				logger.Infof("sn:%s, iloip:%s oob job done", cps.SN, cps.IloIP)
				return
			case <-time.After(3 * time.Minute): //线上统计，onecli命令挂住最长达10min左右，但是onecli不好控制，在入口这里控制3min超时时间
				logger.Warnf("sn:%s, iloip:%s oob job timeout, maybe onecli,please check it", cps.SN, cps.IloIP)
				return
			}

		}(cps)

	}

	// wait
	pool.Wait()

	// write to
	return nil
}

// GetCPSFromRedis read CPS list from redis, according region and az in config file
func GetCPSFromRedis(logger *log.Logger) (object.CPSObject, error) {

	var cpsObject object.CPSObject

	cpsMap, err := util.HGetAllFromRedis("")
	if err != nil {
		logger.Warnf("get cps list error. %s", err.Error())
	}

	logger.Infof("GetCPSFromRedis origin cpss:%s", util.Obj2String(cpsMap))

	for _, v := range cpsMap {
		var cpsInfo object.CPSRecord
		err := json.Unmarshal([]byte(v), &cpsInfo)
		if err != nil {
			logger.Warnf("convert json to CPSRecord error, error:%s", err.Error())
			continue
		}

		cpsObject.CPSs = append(cpsObject.CPSs, cpsInfo)
		cpsObject.Region = cpsInfo.Region
		cpsObject.AZ = cpsInfo.AZ
	}

	//logs.Debug(cpsObject)

	return cpsObject, nil
}
