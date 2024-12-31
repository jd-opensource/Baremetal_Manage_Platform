package main

import (
	"flag"
	"fmt"
	"os"

	"coding.jd.com/aidc-bmp/oob-log-agent/util"

	"coding.jd.com/aidc-bmp/oob-log-agent/scheduler"

	"github.com/beego/beego/v2/adapter/toolbox"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

const (
	defaultConfigFile = "conf/agent-config.ini"
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
		os.Exit(0)
	}

	// initial Config From file
	beego.LoadAppConfig("ini", *configFile)

	// init log
	logConfig := "{\"filename\":\"" + beego.AppConfig.DefaultString("app.agent-log-file", "/var/log/bmp/bmp-oob-agent/bmp-log-agent.log") + "\",\"level\":" + beego.AppConfig.DefaultString("app.log-level", "7") + "}"
	logs.SetLogger(logs.AdapterConsole, logConfig)
	logs.EnableFuncCallDepth(true)

	logs.Info("========CPS-OOB-AGENT Start========")
	defer logs.Info("========CPS-OOB-AGENT Finish========")

	// initial redis
	rNamespace, _ := beego.AppConfig.String("bmp_redis_namespace")
	rIdc, _ := beego.AppConfig.String("app.idc")
	rhost, _ := beego.AppConfig.String("bmp_redis_host")
	rport, _ := beego.AppConfig.String("bmp_redis_port")
	rAddr := fmt.Sprintf("%s:%s", rhost, rport)
	rPasswd, _ := beego.AppConfig.String("bmp_redis_password")
	rDb, _ := beego.AppConfig.Int("bmp_redis_db")

	if err := util.InitRedis(rNamespace, rIdc, rAddr, rPasswd, rDb); err != nil {
		logs.Warn("Redis Connect Error, Ping result : %s", err.Error())
		os.Exit(0)
	}
	defer util.RedisUtil.Close()

	// 每10分钟执行一次带外日志检测任务
	oobCheckTral, _ := beego.AppConfig.String("job.check-ooblog-cron")
	c1 := toolbox.NewTask("oobCheck", oobCheckTral, scheduler.CheckOoblog)
	err := c1.Run()
	if err != nil {
		logs.Warn("CheckOoblog() Error : " + err.Error())
		// logs.Error("CheckOoblog() Error : " + err.Error())
	}
	toolbox.AddTask("oobCheck", c1)

	toolbox.StartTask()
	defer toolbox.StopTask()

	//是否需要select??
	select {}

}

func usage() {

	fmt.Printf(`Usage: oob-log-agent [-c configfile]
Options:

`)
	flag.PrintDefaults()

}
