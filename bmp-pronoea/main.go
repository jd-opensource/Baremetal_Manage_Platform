package main

import (
	"fmt"
	"git.jd.com/bmp-pronoea/config"
	_ "git.jd.com/bmp-pronoea/routers"
	"git.jd.com/bmp-pronoea/service/crontab"
	"github.com/astaxie/beego/logs"
	"github.com/beego/beego/v2/task"
	"os"

	"github.com/astaxie/beego"
)

func main() {

	if err := initLog() ; err != nil {
		fmt.Printf("init log error: %s", err.Error())
		os.Exit(1)
	}

	err := config.ParseConfigAppConf()
	if err != nil {
		logs.Info("system init error %v ", err)
		fmt.Printf("system init error %v ", err)
		os.Exit(1)
	}

	logs.Info("bmp-pronoea runing... ... ")

	crontab.InitCrontab()
	defer task.StopTask()

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

func initLog() error {
	level, err := beego.AppConfig.Int("log.level")
	if err != nil {
		return fmt.Errorf("get log.level error : %v", err)
	}

	filePath := beego.AppConfig.String("log.file")

	maxDays, err := beego.AppConfig.Int("log.days")
	if err != nil {
		return fmt.Errorf("get log.days error : %v", err)
	}

	logPerm := beego.AppConfig.String("log.perm")
	logRotatePerm := beego.AppConfig.String("log.rotateperm")

	logConfigStr := fmt.Sprintf("{\"filename\": \"%s\", \"level\":%d, \"maxdays\":%d, \"perm\":\"%s\", \"rotateperm\":\"%s\" }",
		filePath, level, maxDays, logPerm, logRotatePerm)
	logs.SetLogger("file", logConfigStr)

	return nil
}
