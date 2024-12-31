package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"coding.jd.com/aidc-bmp/bmp-console-api/logic/WebsocketLogic"
	response "coding.jd.com/aidc-bmp/bmp-console-api/types/response"
	"github.com/gorilla/websocket"
)

// WebSocketController operations for Msg
type WebSocketController struct {
	BaseController
}

func (wsc *WebSocketController) Send() {
	userId := wsc.logPoints.GetPoint("userId").(string)
	err := WebsocketLogic.SendMsg2Client(wsc.logPoints, []string{userId}, WebsocketLogic.NofifyMessage{
		Type:    "test",
		SubType: "sendTestSocketmsg",
		Message: "test text",
	})
	success := true
	if err != nil {
		success = false
	}

	wsc.Res.Result = response.CommonResponse{
		Success: success,
	}

}

func (wsc *WebSocketController) Get() {

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		// 解决跨域问题
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	ws, err := upgrader.Upgrade(wsc.Ctx.ResponseWriter, wsc.Ctx.Request, nil)
	if err != nil {
		log.Fatal(err)
	}

	ConnId := wsc.logPoints.GetPoint("logid").(string)
	c := &WebsocketLogic.Client{
		ID:            ConnId, // 连接id
		UserId:        wsc.logPoints.GetPoint("userId").(string),
		HeartbeatTime: time.Now().Unix(),
		Socket:        ws,
	}

	//  defer ws.Close()
	WebsocketLogic.RegisterChan <- c

	defer func() {
		WebsocketLogic.UnregisterChan <- c
	}()
	for {
		// 读取消息
		_, body, err := c.Socket.ReadMessage()
		if err != nil {
			wsc.logPoints.Warnf("c.Socket.ReadMessage error:%s", err.Error())
			break
		}

		fmt.Println("message from client:", string(body))

		msg := WebsocketLogic.ReceiveMessage{}
		err = json.Unmarshal(body, &msg)
		if err != nil {
			wsc.logPoints.Warnf("unmarshal error:%s", err.Error())
			continue
		}

		if msg.Type == WebsocketLogic.MessageTypeHeartbeatFromClient { // 维持心跳消息
			// 刷新连接时间
			c.HeartbeatTime = time.Now().Unix()

			// 回复心跳
			replyMsg := WebsocketLogic.NofifyMessage{
				Type: WebsocketLogic.MessageTypeHeartbeatFromServer,
			}
			err = c.Socket.WriteJSON(replyMsg)
			if err != nil {
				fmt.Println("c.Socket.WriteJSON error:", err.Error())
				wsc.logPoints.Warnf("c.Socket.WriteJSON error:%s", err.Error())
			}
			//
		} else {
			//目前nothing
		}
	}

}
