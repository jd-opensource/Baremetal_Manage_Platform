package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"coding.jd.com/bmp/agent-proxy-jdstack/logic"
	_ "coding.jd.com/bmp/agent-proxy-jdstack/routers"
	"coding.jd.com/bmp/agent-proxy-jdstack/util"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

const (
	defaultConfigFile = "conf/proxy.ini"
)

type LoggerConfig struct {
	FileName string `json:"filename"` //将日志保存到的文件名及路径
	Level    int    `json:"level"`    // 日志保存的时候的级别，默认是 Trace 级别
	Maxlines int    `json:"maxlines"` // 每个文件保存的最大行数，若文件超过maxlines，则将日志保存到下个文件中，为0表示不设置。默认值 1000000
	Maxsize  int    `json:"maxsize"`  // 每个文件保存的最大尺寸，若文件超过maxsize，则将日志保存到下个文件中，为0表示不设置。默认值是 1 << 28, //256 MB
	// Daily               bool   `json:"daily"`    // 设置日志是否每天分割一次，默认是 true
	Hourly bool `json:"hourly"` // 设置日志是否每小时分割一次，默认是 true
	// Maxdays             int    `json:"maxdays"`  // 设置保存最近几天的日志文件，超过天数的日志文件被删除，为0表示不设置，默认保存 7 天
	MaxHours            int    `json:"maxhours"`
	Rotate              bool   `json:"rotate"` // 是否开启 logrotate，默认是 true
	Perm                string `json:"perm"`   // 日志文件权限
	RotatePerm          string `json:"rotateperm"`
	EnableFuncCallDepth bool   `json:"-"` // 输出文件名和行号
	LogFuncCallDepth    int    `json:"-"` // 函数调用层级
}

func main() {

	var configFile = flag.String("c", defaultConfigFile, "specify config file")
	var helpFlag = flag.Bool("h", false, "print help")

	flag.Parse()

	if *helpFlag {
		flag.PrintDefaults()
		os.Exit(0)
	}

	web.LoadAppConfig("ini", *configFile)

	// logConfig := "{\"filename\":\"" + beego.AppConfig.DefaultString("app.log-file", "log/cps-ops-web.log") + "\",\"level\":" + beego.AppConfig.DefaultString("app.log-level", "7") + "}"
	// logs.SetLogger(logs.AdapterFile, logConfig)
	// logs.SetLogger(logs.AdapterConsole, `{"level":8,"color":true}`)
	// logs.EnableFuncCallDepth(true)
	LogsInit()

	logs.Info("========BMP-MONITOR-PROXY Start========")
	initApp()

	// initial redis
	rNamespace, _ := web.AppConfig.String("bmp_redis_namespace")
	rIdc, _ := web.AppConfig.String("app.idc")
	rhost, _ := web.AppConfig.String("bmp_redis_host")
	rport, _ := web.AppConfig.String("bmp_redis_port")
	rAddr := fmt.Sprintf("%s:%s", rhost, rport)
	rPasswd, _ := web.AppConfig.String("bmp_redis_password")
	rDb, _ := web.AppConfig.Int("bmp_redis_db")

	if err := util.InitRedis(rNamespace, rIdc, rAddr, rPasswd, rDb); err != nil {
		logs.Warn("Redis Connect Error, Ping result : %s", err.Error())
		os.Exit(0)
	}
	defer util.RedisUtil.Close()

	if err := logic.Init(); err != nil {
		logs.Warn("prometheus.gateway init Error%s", err.Error())
	}

	web.Run()

}

// LogsInit 日志init
func LogsInit() {
	logCfg := LoggerConfig{
		FileName:            "log/proxy.log",
		Level:               logs.LevelDebug,
		Hourly:              true,
		MaxHours:            7 * 24,
		EnableFuncCallDepth: true,
		LogFuncCallDepth:    3,
		RotatePerm:          "777",
		Perm:                "777",
	}

	// 设置beego log库的配置
	b, _ := json.Marshal(logCfg)
	_ = logs.SetLogger(logs.AdapterFile, string(b))
	// logs.Async(1e3) //异步输出允许设置缓冲 chan 的大小
}

func initApp() {

	// Signal
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	go func() {
		<-signalChan
		exit(0)
	}()

}

func exit(result int) {

	logs.Info("========BMP-MONITOR-PROXY Exit========")

	os.Exit(result)
}
