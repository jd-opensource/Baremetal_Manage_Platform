package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"coding.jd.com/aidc-bmp/bmp-driver/command"
	"coding.jd.com/aidc-bmp/bmp-driver/cron"
	mq "coding.jd.com/aidc-bmp/bmp-driver/service/rabbit_mq"
	beego "github.com/beego/beego/v2/server/web"
)

var run_mode string

type Config struct {
	Heart string
}

var (
	VERSION    = "should be compiled with make."
	BUILD_TIME = "should be compiled with make."
	PERIOD     = 300
)

const (
	defaultConfigFile = "conf/bmp-driver.ini"
)

var configFile = flag.String("c", defaultConfigFile, "specify config file")
var versionFlag = flag.Bool("v", false, "show version")

func main() {

	flag.Parse()

	// initial Config From file
	if err := beego.LoadAppConfig("ini", *configFile); err != nil {
		panic(fmt.Sprintf("LoadAppConfig Error:%s", err.Error()))
	}

	if *versionFlag {
		fmt.Println("commit_id:", VERSION, "build_time:", BUILD_TIME)
		os.Exit(0)
	}

	fmt.Println("[bmp-driver] start...")
	defer fmt.Println("[bmp-driver] endding...")

	if err := mq.InitMqTemplate(); err != nil {
		fmt.Println("initMqTemplate error:", err.Error())
		panic(err)
	}
	defer mq.DestroyTemplate()

	//go func() {
	//	ticker := time.NewTicker(time.Duration(300) * time.Second)
	//	defer ticker.Stop()
	//	for range ticker.C {
	//		command.Heartbeat()
	//	}
	//}()

	go cron.Run()

	//从ironic-api获取任务
	var errChan chan int = make(chan int)
	go doReceiveFromIronicScheduler(errChan)
	for range errChan {
		go doReceiveFromIronicScheduler(errChan)
		time.Sleep(10 * time.Second)
	}
}

func doReceiveFromIronicScheduler(errChan chan int) {

	// defer func() {
	// 	fmt.Println("doReceiveFromIronicScheduler ch error ......")
	// 	errChan <- 1
	// }()

	ch, err := mq.ReceiveFromIronicScheduler()
	if err != nil {
		fmt.Println(time.Now().String()+"ReceiveFromIronicScheduler.ch error:", err.Error())
		errChan <- 1
	}
	for v := range ch {
		fmt.Println(time.Now().String()+"receive message from ironic-scheduler:", string(v.Body))
		go command.Run(v.Body)
	}
	//这里chan已经关闭，表示异常
	errChan <- 1
}

//func initMqConf() error {
//	//run_mode = os.Getenv("IRONIC_DRIVER_MODE")
//	//
//	//if run_mode == "" {
//	//	run_mode = "dev"
//	//}
//	//fmt.Println("IRONIC_DRIVER_MODE is:", run_mode)
//	//b, err := ioutil.ReadFile(fmt.Sprintf("./conf/config-%s.yml", run_mode))
//	//if err != nil {
//	//	return err
//	//}
//	//
//	//if err := yaml.Unmarshal(b, &mq.C); err != nil {
//	//	return err
//	//}
//	return mq.InitMqTemplate()
//}
