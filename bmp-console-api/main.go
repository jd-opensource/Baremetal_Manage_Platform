package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"coding.jd.com/aidc-bmp/bmp-console-api/controllers"
	sessionUtil "coding.jd.com/aidc-bmp/bmp-console-api/thirdpart/session"
	"github.com/beego/beego/v2/core/logs"

	_ "coding.jd.com/aidc-bmp/bmp-console-api/routers"
	rds "coding.jd.com/aidc-bmp/bmp-console-api/thirdpart/redis"
	_ "github.com/beego/beego/v2/server/web/session/redis"
	_ "github.com/go-sql-driver/mysql"

	beego "github.com/beego/beego/v2/server/web"

	"coding.jd.com/aidc-bmp/bmp-console-api/thirdpart/openApi"
)

const (
	defaultConfigFile = "conf/bmp-console.ini"
)

var (
	VERSION    = "should be compiled with make."
	BUILD_TIME = "should be compiled with make."
)

var configFile = flag.String("c", defaultConfigFile, "specify config file")
var versionFlag = flag.Bool("v", false, "show version")

func main() {

	time.LoadLocation("Asia/Shanghai")
	flag.Parse()

	if *versionFlag {
		fmt.Println("commit_id:", VERSION, "build_time:", BUILD_TIME)
		os.Exit(0)
	}

	log.Println("[bmp-console] start...")
	// logs.SetLogger(logs.AdapterConsole)
	logs.SetLevel(logs.LevelTrace)
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)
	logs.Info("============== Start========")

	// initial Config From file
	if err := beego.LoadAppConfig("ini", *configFile); err != nil {
		fmt.Println("LoadAppConfig Error:", err.Error())
		return
	}
	rds.Init()
	//cpsApiHost, err := beego.AppConfig.String("cps-api.host")
	//if err != nil {
	//	fmt.Println("parse cps-api.host from confile error:", err.Error())
	//	return
	//}
	//if err := cpsApi.InitCpsApiConfig(cpsApiHost); err != nil {
	//	fmt.Println("init cpsApi conf error:", err.Error())
	//	// return
	//}

	//if err := initGormDB(); err != nil {
	//	fmt.Println("initGormDB() Error:", err.Error())
	//	return
	//}

	//kafka, 创建实例付费成功接收msg
	// kafkaOrderHost, _ := beego.AppConfig.String("kafka.order.bootstrap-server")
	// kafkaOrderTopic, _ := beego.AppConfig.String("kafka.order.topic.order_produce")
	// go kafka.Consume(kafkaOrderHost, kafkaOrderTopic)

	// TODO 监听付款成功的kafka消息，然后请求innerCreateInstance
	openApi.Init()
	sessionUtil.Init()

	controllers.Init() //手动调用初始化

	beego.Run()
}

// func initGormDB() error {
// 	// initial mysql
// 	sqlconn, _ := beego.AppConfig.String("jdbc.url") //mysql
// 	return dao.InitGormDb(sqlconn)
// }
