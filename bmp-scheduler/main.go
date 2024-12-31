package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"coding.jd.com/aidc-bmp/bmp-scheduler/dao"
	processor "coding.jd.com/aidc-bmp/bmp-scheduler/processor"
	scheduler "coding.jd.com/aidc-bmp/bmp-scheduler/scheduler"
	email "coding.jd.com/aidc-bmp/bmp-scheduler/service/email"
	rabbitMq "coding.jd.com/aidc-bmp/bmp-scheduler/service/rabbit_mq"
	redis "coding.jd.com/aidc-bmp/bmp-scheduler/service/redis"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

var (
	VERSION    = "should be compiled with make."
	BUILD_TIME = "should be compiled with make."
)

const (
	defaultConfigFile = "conf/bmp-scheduler.ini"
)

var configFile = flag.String("c", defaultConfigFile, "specify config file")
var versionFlag = flag.Bool("v", false, "show version")

func main() {

	flag.Parse()

	if *versionFlag {
		fmt.Println("commit_id:", VERSION, "build_time:", BUILD_TIME)
		os.Exit(0)
	}

	fmt.Println(time.Now(), "[bmp-scheduler] start...")
	defer fmt.Println(time.Now(), "[bmp-scheduler] end...")

	initConfig()
	scheduler.Run()
}

func initConfig() {
	// initial Config From file
	if err := beego.LoadAppConfig("ini", *configFile); err != nil {
		panic(fmt.Sprintf("LoadAppConfig Error:%s", err.Error()))
	}

	if err := initRedis(); err != nil {
		fmt.Println(time.Now(), "initRedis Error:", err.Error())
	}

	initGormDB()

	if err := initRabbitMqTemplate(); err != nil {
		panic(fmt.Sprintf("initRabbitMqTemplate Error:%s", err.Error()))
	}

	// if err := initSdnApiConfig(); err != nil {
	// 	panic(fmt.Sprintf("initSdnApiConfig Error:%s", err.Error()))
	// }

	if err := initIloConfig(); err != nil {
		panic(fmt.Sprintf("initIloConfig Error:%s", err.Error()))
	}

	if err := initEmailConfig(); err != nil {
		panic(fmt.Sprintf("initEmailConfig Error:%s", err.Error()))
	}

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

	host1, _ := beego.AppConfig.String("bmp_mq_host")
	port1, _ := beego.AppConfig.String("bmp_mq_port")
	user1, _ := beego.AppConfig.String("bmp_mq_user")
	password1, _ := beego.AppConfig.String("bmp_mq_password")
	vhost1, _ := beego.AppConfig.String("bmp_mq_vhost")
	mqUrl1 := "amqp://" + user1 + ":" + password1 + "@" + host1 + ":" + port1 + "/" + vhost1
	return rabbitMq.InitMqTemplate(mqUrl, mqUrl1)
}

// func initSdnApiConfig() error {
// 	cn_east_1_ip, _ := beego.AppConfig.String("sdn.configs.cn-east-1.ip")
// 	cn_east_jdr1_ip, _ := beego.AppConfig.String("sdn.configs.cn-east-jdr1.ip")
// 	cn_north_1_ip, _ := beego.AppConfig.String("sdn.configs.cn-north-1.ip")
// 	cn_east_tz1_ip, _ := beego.AppConfig.String("sdn.configs.cn-east-tz1.ip")
// 	cn_east_jn1_ip, _ := beego.AppConfig.String("sdn.configs.cn-east-jn1.ip")
// 	cn_east_10_ip, _ := beego.AppConfig.String("sdn.configs.cn-east-10.ip")
// 	cn_east_11_ip, _ := beego.AppConfig.String("sdn.configs.cn-east-11.ip")

// 	cn_east_1_nameservers, _ := beego.AppConfig.Strings("sdn.configs.cn-east-1.nameservers")
// 	cn_east_jdr1_nameservers, _ := beego.AppConfig.Strings("sdn.configs.cn-east-jdr1.nameservers")
// 	cn_north_1_nameservers, _ := beego.AppConfig.Strings("sdn.configs.cn-north-1.nameservers")
// 	cn_east_tz1_nameservers, _ := beego.AppConfig.Strings("sdn.configs.cn-east-tz1.nameservers")
// 	cn_east_jn1_nameservers, _ := beego.AppConfig.Strings("sdn.configs.cn-east-jn1.nameservers")
// 	cn_east_10_nameservers, _ := beego.AppConfig.Strings("sdn.configs.cn-east-10.nameservers")
// 	cn_east_11_nameservers, _ := beego.AppConfig.Strings("sdn.configs.cn-east-11.nameservers")

// 	c := map[string]sdnApi.Cfg{
// 		"cn-east-1":    sdnApi.Cfg{cn_east_1_ip, cn_east_1_nameservers},
// 		"cn-east-jdr1": sdnApi.Cfg{cn_east_jdr1_ip, cn_east_jdr1_nameservers},
// 		"cn-north-1":   sdnApi.Cfg{cn_north_1_ip, cn_north_1_nameservers},
// 		"cn-east-tz1":  sdnApi.Cfg{cn_east_tz1_ip, cn_east_tz1_nameservers},
// 		"cn-east-jn1":  sdnApi.Cfg{cn_east_jn1_ip, cn_east_jn1_nameservers},
// 		"cn-east-10":   sdnApi.Cfg{cn_east_10_ip, cn_east_10_nameservers},
// 		"cn-east-11":   sdnApi.Cfg{cn_east_11_ip, cn_east_11_nameservers},
// 	}
// 	return sdnApi.InitSdnApiConfig(c)

// }

func initIloConfig() error {
	processor.IloUsername, _ = beego.AppConfig.String("ilo.username")
	processor.IloPassword, _ = beego.AppConfig.String("ilo.password")
	return nil
}

func initEmailConfig() error {
	host, _ := beego.AppConfig.String("mail.server.host")
	port, _ := beego.AppConfig.String("mail.server.port")
	port_n, _ := strconv.Atoi(port)
	username, _ := beego.AppConfig.String("mail.server.username")
	password, _ := beego.AppConfig.String("mail.server.password")
	from, _ := beego.AppConfig.String("mail.server.email")
	c := email.EmailConfig{
		Host:     host,
		Port:     port_n,
		Username: username,
		Password: password,
		From:     from,
	}
	return email.InitEmailConfig(c)
}
