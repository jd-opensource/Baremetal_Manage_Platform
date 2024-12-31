package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"coding.jd.com/aidc-bmp/oob-log-alert/crontab"
	"coding.jd.com/aidc-bmp/oob-log-alert/dao"
	"coding.jd.com/aidc-bmp/oob-log-alert/util"

	_ "coding.jd.com/aidc-bmp/oob-log-alert/router"
	"coding.jd.com/aidc-bmp/oob-log-alert/service"

	"github.com/beego/beego/v2/adapter/toolbox"
	"github.com/beego/beego/v2/core/logs"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const (
	defaultConfigFile = "conf/alert-config.ini"
)

var configFile = flag.String("c", defaultConfigFile, "specify config file")
var helpFlag = flag.Bool("h", false, "print help")

func init() {
	flag.Usage = usage
}

func main() {

	// parse flag
	flag.Parse()

	if *helpFlag {
		flag.Usage()
		exit(0)
	}

	// initial Config From file
	if err := web.LoadAppConfig("ini", *configFile); err != nil {
		fmt.Println("LoadAppConfig() Error:" + err.Error())
		exit(1)
	}

	// init log
	logConfig := "{\"filename\":\"" + web.AppConfig.DefaultString("app.alert-log-file", "/var/log/bmp/bmp-oob-alert/bmp-log-alert.log") + "\",\"level\":" + web.AppConfig.DefaultString("app.log-level", "7") + "}"
	logs.SetLogger(logs.AdapterFile, logConfig)
	logs.EnableFuncCallDepth(true)

	logs.Info("========CPS-OOB-ALERT Start========")

	// initial redis
	// initial redis
	rNamespace, _ := web.AppConfig.String("bmp_redis_namespace")
	rIdc, _ := web.AppConfig.String("app.idc")
	rhost, _ := web.AppConfig.String("bmp_redis_host")
	rport, _ := web.AppConfig.String("bmp_redis_port")
	rAddr := fmt.Sprintf("%s:%s", rhost, rport)
	rPasswd, _ := web.AppConfig.String("bmp_redis_password")
	rDb, _ := web.AppConfig.Int("bmp_redis_db")

	if err := util.InitRedis(rNamespace, rIdc, rAddr, rPasswd, rDb); err != nil {
		logs.Warn("Redis Connect Error, Ping result : %s", err.Error())
		os.Exit(0)
	}
	defer util.RedisUtil.Close()

	// init Database
	host, _ := web.AppConfig.String("bmp_db_host")
	port, _ := web.AppConfig.String("bmp_db_port")
	user, _ := web.AppConfig.String("bmp_db_user")
	password, _ := web.AppConfig.String("bmp_db_password")
	name, _ := web.AppConfig.String("bmp_db_name")
	//root:admin@tcp(10.226.192.113:3306)/aidc_bmp?charset=utf8&parseTime=true&charset=utf8&loc=Asia%2FShanghai
	mysqlUrn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + name + "?charset=utf8&parseTime=true&loc=Local"
	if err := orm.RegisterDataBase("default", "mysql", mysqlUrn); err != nil {
		logs.Warn("Mysql Connect Error : %s", err.Error())
		exit(0)
	}
	// create table
	// orm.RunSyncdb("default", false, true)
	orm.Debug = true
	dao.InitGormDb(mysqlUrn)

	// register exit func
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	go func() {
		<-signalChan
		exit(0)
	}()

	// fmt.Println("[DEBUG]", web.AppConfig.String("redis.addr"))

	// init HealthCheck
	// toolbox.AddHealthCheck("Redis", &service.RedisCheck{})
	// toolbox.AddHealthCheck("Mysql", &service.MySQLCheck{})

	// init powerstatus check task
	// dailyInstanceStatusCron, _ := web.AppConfig.String("alert.daily-instance-status-cron")
	// powerStatusTask := toolbox.NewTask("PowerStatus", dailyInstanceStatusCron, service.DoStaticCpsPowerState)
	// err := powerStatusTask.Run()
	// if err != nil {
	// 	logs.Error("powerStatusTask() Error : " + err.Error())
	// }
	// toolbox.AddTask("PowerStatus", powerStatusTask)

	// init daily report task
	// dailyReportCron, _ := web.AppConfig.String("alert.daily-report-cron")
	// dailReportTask := toolbox.NewTask("DailyReport", dailyReportCron, service.DailReportEvent)
	// err = dailReportTask.Run()
	// if err != nil {
	// 	logs.Error("dailReportTask() Error : " + err.Error())
	// }
	// toolbox.AddTask("DailyReport", dailReportTask)

	//定时任务同步device数据到redis，供agent使用
	deviceSyncCron, _ := web.AppConfig.String("app.sync-device-cron")
	syncDeviceTask := toolbox.NewTask("SyncDeviceTask", deviceSyncCron, crontab.SyncDeviceToRedis)
	err := syncDeviceTask.Run()
	if err != nil {
		logs.Error("syncDeviceTask() Error : " + err.Error())
	}
	toolbox.AddTask("SyncDeviceTask", syncDeviceTask)

	toolbox.StartTask()
	defer toolbox.StopTask()

	// init event subscribe
	ipmiLogpath, _ := web.AppConfig.String("alert.save-ipmi-logpath")
	if err := service.InitAlert(
		web.AppConfig.DefaultBool("alert.save-ipmi-log", false),
		ipmiLogpath); err != nil {

		logs.Critical("Alert Init Error")
		exit(3)
	}

	util.InitOpenapi()

	// start web
	web.Run()

}

func usage() {

	fmt.Printf(`Usage: oob-log-alert [-c configfile]
Options:

`)
	flag.PrintDefaults()
}

func exit(result int) {

	util.RedisUtil.Close()

	logs.Info("========CPS-OOB-ALERT Exit========")

	os.Exit(result)
}
