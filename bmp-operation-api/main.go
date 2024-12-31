package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"coding.jd.com/aidc-bmp/bmp-operation-api/controllers"
	redisUtil "coding.jd.com/aidc-bmp/bmp-operation-api/services/redis"
	sessionUtil "coding.jd.com/aidc-bmp/bmp-operation-api/services/session"
	"github.com/beego/beego/v2/core/logs"

	_ "coding.jd.com/aidc-bmp/bmp-operation-api/routers"

	"coding.jd.com/aidc-bmp/bmp-operation-api/services/openApi"
	_ "coding.jd.com/aidc-bmp/bmp-operation-api/swaggermodels"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/beego/beego/v2/server/web/session/redis"
	_ "github.com/go-sql-driver/mysql"
)


const (
	defaultConfigFile = "conf/bmp-operation.ini"
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

	log.Println("[bmp-operation] start...")
	// logs.SetLogger(logs.Adapteroperation)
	logs.SetLevel(logs.LevelTrace)
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)
	logs.Info("============== Start========")

	// initial Config From file
	if err := beego.LoadAppConfig("ini", *configFile); err != nil {
		fmt.Println("LoadAppConfig Error:", err.Error())
		return
	}
	//if err := initGormDB(); err != nil {
	//	fmt.Println("initGormDB() Error:", err.Error())
	//	return
	//}

	//kafka, 创建实例付费成功接收msg
	// kafkaOrderHost, _ := beego.AppConfig.String("kafka.order.bootstrap-server")
	// kafkaOrderTopic, _ := beego.AppConfig.String("kafka.order.topic.order_produce")
	// go kafka.Consume(kafkaOrderHost, kafkaOrderTopic)

	// TODO 监听付款成功的kafka消息，然后请求innerCreateInstance
	controllers.Init()
	openApi.Init()
	redisUtil.Init()
	sessionUtil.Init()

	beego.BConfig.MaxUploadSize = 1 << 35 //32G
	beego.Run()
}

// func initGormDB() error {
// 	// initial mysql
// 	sqlconn, _ := beego.AppConfig.String("jdbc.url") //mysql
// 	return dao.InitGormDb(sqlconn)
// }
