package scheduler

import (
	"time"

	rabbitMq "coding.jd.com/aidc-bmp/bmp-scheduler/service/rabbit_mq"
	"github.com/beego/beego/v2/adapter/toolbox"
)

//ironic-scheduler模块四种任务请求都在这里做派发

// dispatchIronicApi 从ironic-api投递过来的任务请求
// dispatchIronicAgent 从ironic-agent投递过来的任务请求
// CheckTimeoutCommandCron 定时任务1,对执行超时的任务做重试等
// UpdateInstanceStatusCron 定时任务2,执行任务完成/失败后，更新实例状态

func Run() {
	defer rabbitMq.DestroyTemplate()
	// cron1 命令执行超时定时检测任务
	c1 := toolbox.NewTask("checkTimeoutCommand", "0 */1 * * * *", doCheckTimeoutCommandCron)
	toolbox.AddTask("checkTimeoutCommand", c1)

	// cron2 命令执行完后实例状态更新任务
	c2 := toolbox.NewTask("updateInstanceStatus", "*/30 * * * * *", doUpdateInstanceStatusCron)
	toolbox.AddTask("updateInstanceStatus", c2)

	c3 := toolbox.NewTask("doCleanAuditLogCron", "0 0 */1 * * *", doCleanAuditLogCron)
	toolbox.AddTask("doCleanAuditLogCron", c3)

	c4 := toolbox.NewTask("CleanLogCron", "0 5 */1 * * *", doCleanLogCron)
	toolbox.AddTask("CleanLogCron", c4)

	toolbox.StartTask()
	defer toolbox.StopTask()

	//从ironic-api获取任务
	var ironicApiErrChan chan int = make(chan int, 1)
	//从ironic-agent获取任务
	var ironicAgentErrChan chan int = make(chan int, 1)
	go doReceiveFromIronicApi(ironicApiErrChan)
	go doReceiveFromIronicAgent(ironicAgentErrChan)
	for {
		select {
		case <-ironicApiErrChan:
			go doReceiveFromIronicApi(ironicApiErrChan)
		case <-ironicAgentErrChan:
			go doReceiveFromIronicAgent(ironicAgentErrChan)
		}
		time.Sleep(10 * time.Second)
	}
}
