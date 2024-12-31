package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"coding.jd.com/aidc-bmp/bmp-openapi/dao"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/monitorDataLogic"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/monitorProxyLogic"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/monitorRuleLogic"
	_ "coding.jd.com/aidc-bmp/bmp-openapi/routers"
	rabbitMq "coding.jd.com/aidc-bmp/bmp-openapi/service/rabbit_mq"
	redis "coding.jd.com/aidc-bmp/bmp-openapi/service/redis"
	_ "coding.jd.com/aidc-bmp/bmp-openapi/swaggermodels"
	"coding.jd.com/aidc-bmp/bmp-openapi/types/request"

	"log"

	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"

	schedulerCron "coding.jd.com/aidc-bmp/bmp-openapi/scheduler/cron"

	_ "coding.jd.com/aidc-bmp/bmp-openapi/constant"
)

//////

const (
	defaultConfigFile = "conf/bmp-openapi.ini"
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

	hour := time.Now().Local().Hour()
	minute := time.Now().Local().Minute()
	second := time.Now().Local().Second()
	fmt.Println("=============time now is:", hour, minute, second)

	log.Println("[ironic-api] start...")

	// initial Config From file
	if err := beego.LoadAppConfig("ini", *configFile); err != nil {
		fmt.Println(time.Now(), "LoadAppConfig Error:", err.Error())
		return
	}
	if err := initRedis(); err != nil {
		fmt.Println(time.Now(), "initRedis Error: ", err.Error())
		// return
	}

	initGormDB()

	if err := initRabbitMqTemplate(); err != nil {
		fmt.Println(time.Now(), "initRabbitMqTemplate Error:", err.Error())
		// return
	}
	defer rabbitMq.DestroyTemplate()

	request.InitValidator()

	// if err := initIdcApiConfig(); err != nil {
	// 	fmt.Println("initIdcApiConfig Error:", err.Error())
	// 	return
	// }

	// if err := initSdnApiConfig(); err != nil {
	// 	fmt.Println("initSdnApiConfig Error:", err.Error())
	// 	return
	// }

	if err := monitorDataLogic.Init(); err != nil {
		fmt.Println(time.Now(), "monitorDataLogic Init Error:", err.Error())

	}
	if err := monitorRuleLogic.Init(); err != nil {
		fmt.Println(time.Now(), "monitorRuleLogic Init Error:", err.Error())

	}
	if err := monitorProxyLogic.Init(); err != nil {
		fmt.Println(time.Now(), "monitorProxyLogic Init Error:", err.Error())

	}
	go schedulerCron.Run()

	beego.Run()
}

func initRedis() error {
	// initial redis
	var namespace, host, port, password string
	var dbNum int
	namespace, _ = beego.AppConfig.String("bmp_redis_namespace")
	host, _ = beego.AppConfig.String("bmp_redis_host")
	port, _ = beego.AppConfig.String("bmp_redis_port")
	password, _ = beego.AppConfig.String("bmp_redis_password")
	dbNum = beego.AppConfig.DefaultInt("bmp_redis_db", 0)
	addr := fmt.Sprintf("%s:%s", host, port)
	return redis.InitRedis(namespace, addr, password, dbNum)
}

func initGormDB() {
	// initial mysql
	host, _ := beego.AppConfig.String("bmp_db_host")
	port, _ := beego.AppConfig.String("bmp_db_port")
	user, _ := beego.AppConfig.String("bmp_db_user")
	password, _ := beego.AppConfig.String("bmp_db_password")
	name, _ := beego.AppConfig.String("bmp_db_name")
	//root:admin@tcp(10.226.192.113:3306)/aidc_bmp?charset=utf8&parseTime=true&charset=utf8&loc=Asia%2FShanghai
	sqlconn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + name + "?charset=utf8&parseTime=true&loc=Local"
	dao.InitGormDb(sqlconn)
}

func initRabbitMqTemplate() error {
	host, _ := beego.AppConfig.String("bmp_mq_host")
	port, _ := beego.AppConfig.String("bmp_mq_port")
	user, _ := beego.AppConfig.String("bmp_mq_user")
	password, _ := beego.AppConfig.String("bmp_mq_password")
	vhost, _ := beego.AppConfig.String("bmp_mq_vhost")
	mqUrl := "amqp://" + user + ":" + password + "@" + host + ":" + port + "/" + vhost
	// exchange, _ := beego.AppConfig.String("rabbit.mq.ironic.exchange")
	// receive_routing_key, _ := beego.AppConfig.String("rabbit.mq.ironic.receive_routing_key")
	return rabbitMq.InitMqTemplate(mqUrl)
}

// func initIdcApiConfig() error {
// 	host, _ := beego.AppConfig.String("idc.host")
// 	user, _ := beego.AppConfig.String("idc.user")
// 	token, _ := beego.AppConfig.String("idc.token")
// 	return idcApi.InitIdcApiConfig(host, user, token)
// }

// func initSdnApiConfig() error {
// 	cn_east_1, _ := beego.AppConfig.String("sdn.configs.cn-east-1.ip")
// 	cn_east_jdr1, _ := beego.AppConfig.String("sdn.configs.cn-east-jdr1.ip")
// 	cn_north_1, _ := beego.AppConfig.String("sdn.configs.cn-north-1.ip")
// 	cn_east_tz1, _ := beego.AppConfig.String("sdn.configs.cn-east-tz1.ip")
// 	cn_east_jn1, _ := beego.AppConfig.String("sdn.configs.cn-east-jn1.ip")
// 	cn_east_11, _ := beego.AppConfig.String("sdn.configs.cn-east-11.ip")

// 	c := map[string]string{
// 		"cn-east-1":    cn_east_1,
// 		"cn-east-jdr1": cn_east_jdr1,
// 		"cn-north-1":   cn_north_1,
// 		"cn-east-tz1":  cn_east_tz1,
// 		"cn-east-jn1":  cn_east_jn1,
// 		"cn-east-11":   cn_east_11,
// 	}
// 	return sdnApi.InitSdnApiConfig(c)

// }
