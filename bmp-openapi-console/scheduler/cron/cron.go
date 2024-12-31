package cron

import (
	"github.com/beego/beego/v2/adapter/toolbox"
)

func Run() {

	c1 := toolbox.NewTask("CleanLogCron", "0 5 */1 * * *", doCleanLogCron)
	toolbox.AddTask("CleanLogCron", c1)

	toolbox.StartTask()
	defer toolbox.StopTask()
}
