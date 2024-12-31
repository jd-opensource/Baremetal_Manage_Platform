package scheduler

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	beego "github.com/beego/beego/v2/server/web"

	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/commandLogic"
	"coding.jd.com/aidc-bmp/bmp-scheduler/logic/deviceLogic"
	"coding.jd.com/aidc-bmp/bmp-scheduler/processor"
	rabbitMq "coding.jd.com/aidc-bmp/bmp-scheduler/service/rabbit_mq"
	"git.jd.com/cps-golang/ironic-common/ironic/enums/CommandStatus"
	ironicAgentEvent "git.jd.com/cps-golang/ironic-common/ironic/event/command"
	commonUtil "git.jd.com/cps-golang/ironic-common/ironic/util"
	log "git.jd.com/cps-golang/log"
)

func doReceiveFromIronicAgent(errChan chan int) {

	fmt.Println(time.Now().String() + "start doReceiveFromIronicAgent......")
	ch, err := rabbitMq.ReceiveFromIronicAgentDriver()
	if err != nil {
		fmt.Println(time.Now().String()+"rabbitMq.ReceiveFromIronicAgentDriver error:", err.Error())
		errChan <- 1
		return
	}
	for v := range ch {
		go dispatchIronicAgentMqMessage(v.Body)
	}
	fmt.Println(time.Now().String() + "doReceiveFromIronicAgent.chan closed")
	//ch关闭就表示异常
	errChan <- 1
}

// ironic-agent/driver 消息派发
func dispatchIronicAgentMqMessage(msg []byte) {

	logPath, _ := beego.AppConfig.String("log.path")
	logger := log.New(logPath + "/bmp-scheduler-from-driver_or_agent.log")
	logger.SetStableFields([]string{"method", "uri", "header", "request", "response"})
	//agent过来的消息只有sn，没有requestid，临时生成
	logger.Point("logid", commonUtil.GenerateRandUuid())
	logger.TimeStart("all_t")
	logger.TimeStart("self_t")
	defer logger.Flush()
	defer logger.TimeEnd("self_t")
	defer logger.TimeEnd("all_t")
	logger.Infof("receive msg from driver/agent: %s", string(msg))
	fmt.Printf(time.Now().String()+"receive msg from driver/agent: %s \n", string(msg))
	event := ironicAgentEvent.CallbackCommandMessage{}
	if err := json.Unmarshal(msg, &event); err == nil {

		sn := event.CommandMessage.Sn
		if event.Action != "Heart" {
			_, err := deviceLogic.GetBySn(logger, sn)
			if err != nil {
				logger.Warnf("dispatchIronicAgentMqMessage.GetBySn error, sn:%s, error:%s", sn, err.Error())
				return
			}

			//logger.Point("region", entity.Region)
			//logger.Point("az", entity.Az)
		}

		if event.Action == "Ping" {
			//尤其是ping指令，很可能因为重试等原因，同时接收到前后两次的响应，导致后面的指令执行两次，所以这里随机sleep一段时间，或者加锁
			num := rand.Int31n(60)
			time.Sleep(time.Duration(num) * time.Second)
		}

		//接受来自driver的带外检测等信息，如果正确，将设备置为已上架??暂时去掉这个逻辑
		if event.Action == "CheckInitConfig" {
			if err := deviceLogic.CheckDevice(logger, event); err != nil {
				logger.Warnf("update device manageStatus,reason, sn:%s, error:%s", event.Sn, err.Error())
				return
			}
			return
		}
		logger.Point("action", event.Action)
		command, err := commandLogic.GetBySnAndActionAndStatus(logger, event.Sn, event.Action, CommandStatus.RUNNING)
		if err != nil {
			// logger.Warn("dispatchIronicAgentMqMessage GetBySnAndActionAndStatus sql error:", event.Sn, event.Action, err.Error())
			return
		}
		if command == nil {
			return
		}
		v, _ := json.Marshal(command)
		logger.Info("receive driver/agent for command:", string(v))
		processor.ProcessorMap[command.Action].Callback(logger, command, event)
	} else {
		logger.Warn("unknown format message from agent: ", string(msg))
	}
}
