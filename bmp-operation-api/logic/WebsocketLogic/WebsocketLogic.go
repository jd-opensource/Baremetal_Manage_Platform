package WebsocketLogic

import (
	"fmt"
	"sync"
	"time"

	"coding.jd.com/aidc-bmp/bmp-operation-api/util"
	log "git.jd.com/cps-golang/log"
	"github.com/gorilla/websocket"
)

const (
	MessageTypeHeartbeatFromClient    = "heartbeat"    // 心跳
	MessageTypeHeartbeatFromServer    = "ok"           // 心跳回复成功
	MessageTypeRegisterFromClient     = "register"     // 注册
	MessageTypeRegisterSuccFromServer = "register ok"  // 注册成功
	MessageTypeRegisterFailFromServer = "register err" // 注册失败
	MessageTypeNewNofifyFromServer    = "nofifyfromserver"

	HeartbeatCheckTime = 5   // 心跳检测几秒检测一次
	HeartbeatTime      = 300 // 心跳距离上一次的最大时间

	ChanBufferRegister   = 100 // 注册chan缓冲
	ChanBufferUnregister = 100 // 注销chan大小
)

// ReceiveMessage 客户端请求的报文格式(心跳or业务消息)
type ReceiveMessage struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

// NofifyMessage 回复给客户端的报文格式(心跳回复or业务消息)
type NofifyMessage struct {
	Type     string `json:"type"`
	SubType  string `json:"subType"`
	Message  string `json:"message"`
	MetaData string `json:"metaData"` // 元数据
}

// var Nofify = make(chan NofifyMessage)

// 客户端连接信息
type Client struct {
	ID            string          // 连接ID
	UserId        string          // 账号id, 一个账号可能有多个连接
	Socket        *websocket.Conn // socket连接
	HeartbeatTime int64           // 前一次心跳时间
}

// 客户端管理
type ClientManager struct {
	Clients  map[string]*Client  // 保存连接
	Accounts map[string][]string // 账号和连接关系,map的key是账号id即：AccountId，这里主要考虑到一个账号多个连接
	mu       *sync.Mutex
}

// 定义一个管理Manager
var Manager = ClientManager{
	Clients:  make(map[string]*Client),  // 参与连接的用户，出于性能的考虑，需要设置最大连接数
	Accounts: make(map[string][]string), // 账号和连接关系
	mu:       new(sync.Mutex),
}

var (
	RegisterChan   = make(chan *Client, ChanBufferRegister)   // 注册
	UnregisterChan = make(chan *Client, ChanBufferUnregister) // 注销
)

//有新消息，在这里通知到客户端
func init() {

	/*
		go dealRegister()
		go dealHeartbeat()

		go MSendInterval()
	*/

	// for {
	// 	msg := <-Nofify

	// 	logPath, _ := beego.AppConfig.String("log.path")
	// 	logger := log.New(logPath + "/bmp-msg-notify.log")
	// 	logger.SetStableFields([]string{"time", "type", "subtype", "message"})
	// 	//没有requestid，临时生成
	// 	logger.Point("logid", util.GenerateRandUuid())
	// 	logger.Point("time", time.Now().String())
	// 	logger.Point("type", msg.Type)
	// 	logger.Point("subtype", msg.SubType)
	// 	logger.Point("message", msg.Message)

	// 	// for _, client := range Manager.Clients {
	// 	// 	err := client.Socket.WriteJSON(msg)
	// 	// 	if err != nil {
	// 	// 		logger.Warnf("client.WriteJSON error: %v", err)
	// 	// 		client.Socket.Close()
	// 	// 		delete(Manager.Clients, client)
	// 	// 	}
	// 	// }
	// 	logger.Flush()
	// }
}

// 注册注销
func dealRegister() {
	for {
		select {
		case conn := <-RegisterChan: // 新注册，新连接
			// 加入连接,进行管理
			connBind(conn)

			// 回复消息
			content := NofifyMessage{
				Type:    "register",
				SubType: "register",
				Message: "register success",
			}
			_ = conn.Socket.WriteJSON(content)

		case conn := <-UnregisterChan: // 注销，或者没有心跳
			fmt.Println("UnregisterChan come...")
			// 删除Client
			connUnBind(conn)
		}
	}
}

// 新加连接
func connBind(c *Client) {
	Manager.mu.Lock()
	defer Manager.mu.Unlock()

	// 加入到连接
	Manager.Clients[c.ID] = c

	// 加入到绑定
	if _, ok := Manager.Accounts[c.UserId]; ok { // 该账号已经有绑定，就追加一个绑定
		Manager.Accounts[c.UserId] = append(Manager.Accounts[c.UserId], c.ID)
	} else { // 没有就新增一个账号的绑定切片
		Manager.Accounts[c.UserId] = []string{c.ID}
	}
}

// 删除连接并关闭
func connUnBind(c *Client) {
	Manager.mu.Lock()
	defer Manager.mu.Unlock()

	defer c.Socket.Close()

	// 取消连接
	delete(Manager.Clients, c.ID)

	// 取消绑定
	if len(Manager.Accounts[c.UserId]) > 0 {
		for k, clientId := range Manager.Accounts[c.UserId] {
			if clientId == c.ID { // 找到绑定客户端Id
				Manager.Accounts[c.UserId] = append(Manager.Accounts[c.UserId][:k], Manager.Accounts[c.UserId][k+1:]...)
			}
			if len(Manager.Accounts[c.UserId]) == 0 {
				delete(Manager.Accounts, c.UserId)
			}
		}
	}
}

// 维持心跳
func dealHeartbeat() {
	for {
		// 获取所有的Clients
		Manager.mu.Lock()
		// clients := make([]*Client, len(Manager.Clients))
		clients := []*Client{}
		for _, c := range Manager.Clients {
			clients = append(clients, c)
		}
		Manager.mu.Unlock()
		fmt.Println(time.Now().String(), "debug... clients is:", clients)
		fmt.Println(time.Now().String(), "debug... Manager.Clients is:", Manager.Clients)
		fmt.Println(time.Now().String(), "debug... Manager.Accounts is:", Manager.Accounts)
		for _, c := range clients {
			// fmt.Println("debug... c is:", c)
			if time.Now().Unix()-c.HeartbeatTime > HeartbeatTime {
				connUnBind(c)
			}
		}

		time.Sleep(time.Second * HeartbeatCheckTime)
	}
}

// 有新消息时，业务侧通过这个func发送消息到客户端
func SendMsg2Client(logger *log.Logger, userIds []string, message NofifyMessage) error {

	for _, userId := range userIds {
		// 获取连接id
		clients := GetClient(userId)
		// 发送消息
		for _, c := range clients {
			err := c.Socket.WriteJSON(message)
			if err != nil {
				logger.Infof("SendMsg2Client failed, userId:%s, message:%s, error:%s", userId, util.ObjToJsonString(message), err.Error())
				connUnBind(c)
			} else {
				logger.Infof("SendMsg2Client success, userId:%s, message:%s", userId, util.ObjToJsonString(message))
			}
		}
	}

	return nil
}

// 根据账号获取连接
func GetClient(userId string) []*Client {
	clients := make([]*Client, 0)

	Manager.mu.Lock()
	defer Manager.mu.Unlock()

	if len(Manager.Accounts[userId]) > 0 {
		for _, clientId := range Manager.Accounts[userId] {
			if c, ok := Manager.Clients[clientId]; ok {
				clients = append(clients, c)
			}
		}
	}

	return clients
}

func MSendInterval() {

	message := NofifyMessage{
		Type:    "test",
		SubType: "send_interval",
		Message: "hello socket",
	}
	for {
		for _, c := range Manager.Clients {
			if c == nil {
				fmt.Println("socket client is nil ...., check it")
				continue
			}
			err := c.Socket.WriteJSON(message)
			if err != nil {
				fmt.Printf(time.Now().String(), "SendMsg2Client failed, message:%s, error:%s\n", util.ObjToJsonString(message), err.Error())
				connUnBind(c)
			} else {
				fmt.Printf(time.Now().String(), "SendMsg2Client success, userId:%s, message:%s\n", util.ObjToJsonString(message))
			}

		}
		time.Sleep(20 * time.Second)
	}

}
