package crontab

import (
	"fmt"
	"git.jd.com/bmp-pronoea/config"
	"github.com/astaxie/beego/logs"
	"github.com/beego/beego/v2/task"
)

func InitCrontab() error {

	if config.SysConfig.Crontabs == nil || len(config.SysConfig.Crontabs) <= 0 {
		//logs.Info("InitCrontab", " not exist crontab tasks ！")
		return fmt.Errorf("not exist crontab tasks ！")
	}

	if _, exist := config.SysConfig.Crontabs[config.CRONTAB_WIPE_PUSHGATEWAY]; exist {
		wipePushgatewayTask := task.NewTask(config.CRONTAB_WIPE_PUSHGATEWAY, config.SysConfig.Crontabs[config.CRONTAB_WIPE_PUSHGATEWAY], WipeAllPushgatewayData)
		task.AddTask("wipePushgatewayTask", wipePushgatewayTask)
	}

	task.StartTask()
	logs.Info("InitCrontab", " crontab started ！")
	return nil
}
