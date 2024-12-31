package scheduler

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"coding.jd.com/aidc-bmp/bmp-scheduler/apiActor"
	rabbitMq "coding.jd.com/aidc-bmp/bmp-scheduler/service/rabbit_mq"
	ironicApiEvent "git.jd.com/cps-golang/ironic-common/ironic/event"
	ironicAgentEvent "git.jd.com/cps-golang/ironic-common/ironic/event/command"
	log "git.jd.com/cps-golang/log"
	beego "github.com/beego/beego/v2/server/web"
)

func doReceiveFromIronicApi(errChan chan int) {

	fmt.Println(time.Now().String() + "start doReceiveFromIronicApi......")
	// defer func() {
	// 	fmt.Println("doReceiveFromIronicApi ch error1 ......")
	// 	errChan <- 1
	// 	fmt.Println("doReceiveFromIronicApi ch error2 ......")
	// }()

	ch, err := rabbitMq.ReceiveFromIronicApi2Scheduler()
	if err != nil {
		fmt.Println(time.Now().String()+"rabbitMq.ReceiveFromIronicApi2Scheduler error:", err.Error())
		errChan <- 1
		return
	}
	for v := range ch {
		go dispatchIronicApiMqMessage(v.Body)
	}
	fmt.Println(time.Now().String() + "doReceiveFromIronicApi.chan closed")
	//ch关闭就表示异常
	errChan <- 1
}

// ironic-api 消息派发
func dispatchIronicApiMqMessage(msg []byte) {
	logPath, _ := beego.AppConfig.String("log.path")
	logger := log.New(logPath + "/bmp-scheduler-from-api.log")
	logger.SetStableFields([]string{"method", "uri", "header", "request", "response"})
	logger.TimeStart("all_t")
	logger.TimeStart("self_t")
	defer logger.Flush()
	defer logger.TimeEnd("self_t")
	defer logger.TimeEnd("all_t")

	event := ironicApiEvent.Event{}
	fmt.Println(time.Now(), "receive msg from bmp-openapi: ", string(msg))
	if err := json.Unmarshal(msg, &event); err == nil && event.ClazzName != "" {
		//EventUuid就是ironicApide的requestId
		logger.Point("logid", event.EventUuid)
		logger.Info("receive msg from ironic-api: ", string(msg))
		logger.Point("user_id", event.UserID)
		if strings.HasPrefix(event.ClazzName, "com.jcloud.cps.ironic.event.api") {

			items := strings.Split(event.ClazzName, ".")
			msgType := strings.TrimSuffix(items[len(items)-1], "Message")
			logger.Point("action", msgType)
			logger.Point("event_uuid", event.EventUuid)
			if actor, ok := apiActor.ApiActorMap[msgType]; ok {
				go actor.Do(logger, event.Body, msgType)
			} else {
				logger.Warn("unsupported msgtype:", msgType, "msg body:", event.Body)
			}

		} else if strings.EqualFold(event.ClazzName, "com.jcloud.cps.ironic.event.command.CallbackCommandMessage") {
			//不知道为啥从ironic-api发过来的重试command指令被归类为类似于agent发过来的callback指令，在这里特殊处理一下
			callbackMsg := ironicAgentEvent.CallbackCommandMessage{}
			if err := json.Unmarshal([]byte(event.Body), &callbackMsg); err != nil {
				logger.Warn("dispatchIronicApiMqMessage.CallbackCommandMessage unmarshal error:", err.Error())
				return
			}
			if strings.EqualFold(callbackMsg.Action, "start") {
				apiActor.GetAndStartFirstCommand(logger, callbackMsg.Sn)
			}

		} else {
			logger.Warn("unsupported msg from ironic-api, msg:", string(msg))
		}
	} else {
		fmt.Println(time.Now(), "receive unknown msg from ironic-api: ", string(msg))
	}
}
