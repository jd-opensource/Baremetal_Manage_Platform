package global

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"path"
	"runtime"
	"strings"
	"syscall"

	"coding.jd.com/bmp/cps-agent/util"

	"github.com/astaxie/beego/logs"
	"github.com/gofrs/flock"
)

type Config struct {
	Monitor struct {
		TimestampAlign int `json:"timestamp.align"`
	} `json:"monitor"`
	Agent struct {
		Collect struct {
			Scheduler string `json:"scheduler"`
			Timeout   int    `json:"timeout"`
		} `json:"collect"`
		Heartbeat struct {
			Scheduler string `json:"scheduler"`
		} `json:"heartbeat"`
		Pid string `json:"pid"`
		Log string `json:"log"`
	} `json:"agent"`
	Proxy struct {
		URL string `json:"url"`
	} `json:"proxy"`
}

var (
	SerialNumber string
	AgentConfig  Config
	fileLock     *flock.Flock
	Debug        bool
	SignalChan   chan os.Signal
)

func init() {
	AgentConfig = Config{}
	SignalChan = make(chan os.Signal, 1)
}

func DefaultString(v, d string) string {
	if strings.TrimSpace(v) != "" {
		return v
	}
	return d
}

func DefaultInt(v, d int) int {
	if v != 0 {
		return v
	}
	return d
}

func DefaultInt64(v int, d int64) int64 {
	if v != 0 {
		return int64(v)
	}
	return d
}

func Init() {

	if strings.TrimSpace(SerialNumber) == "" {
		if hostInfo, err := util.SN(); err != nil || hostInfo.SerialNumber == "" {
			fmt.Fprintln(os.Stderr, "Read Host SerialNumber Error, Exit...")
			os.Exit(1)
		} else {
			SerialNumber = hostInfo.SerialNumber
		}
	}

}

func InitApp(conf string) {
	runtime.GOMAXPROCS(1)

	initConfig(conf)
	if runtime.GOOS == "linux" {
		initLog()
		initPid()
	}

	signal.Ignore(
		syscall.SIGHUP, syscall.SIGQUIT,
		syscall.SIGILL, syscall.SIGTRAP, syscall.SIGABRT,
		syscall.SIGBUS, syscall.SIGFPE, syscall.SIGSEGV,
		syscall.SIGPIPE, syscall.SIGALRM, syscall.SIGTERM)
	signal.Notify(SignalChan, syscall.Signal(0xa), syscall.SIGINT)
}

func UninitApp() {
	if fileLock != nil && fileLock.Locked() {
		fileLock.Unlock()
		os.Remove(fileLock.Path())
	}
	util.RemovePID()
}

func initPid() {

	if !Debug {
		//util.Daemon()
	}

	if err := os.MkdirAll(path.Dir(AgentConfig.Agent.Pid), os.ModePerm); err != nil {
		fmt.Fprintln(os.Stderr, "Create Directory Failed, Exit...", err.Error())
		os.Exit(1)
	}

	fileLock = flock.New(AgentConfig.Agent.Pid + ".lock")
	locked, err := fileLock.TryLock()

	if err != nil || !locked || fileLock == nil {
		fmt.Fprintln(os.Stderr, "Init Failed, already running, Exit...")
		os.Exit(0)
	}
	util.NewPID(AgentConfig.Agent.Pid)
}

func initConfig(conf string) {

	c := `{
		"monitor": {
			"timestamp.align": 60
		},
		"agent": {
			"collect": {
				"scheduler": "0/60 * * * * *",
				"timeout": 3
			},
			"heartbeat": {
				"scheduler": "0/60 * * * * *"
			},
			"pid": "/var/run/bmp-agent.pid",
			"log": "/var/log/bmp-agent.log"
		},
		"proxy": {
			"url": "http://bmp-agent-proxy:8805/api/v1/agent/"
		}
	}`

	if err := json.Unmarshal([]byte(c), &AgentConfig); err != nil {
		fmt.Fprintln(os.Stderr, "Config Format Error, Exit...")
		os.Exit(1)
	}
	// // TODO remove
	// logs.Debug("Config %+v", AgentConfig)

}

func initLog() {
	if Debug {
		logs.SetLogger(logs.AdapterConsole, `{"level":7,"color":true}`)
		logs.EnableFuncCallDepth(true)
	} else {
		logConfig := `{"filename":"` + DefaultString(AgentConfig.Agent.Log, "./jdcps-agent.log") + `","level":5}`
		logs.SetLogger(logs.AdapterFile, logConfig)
		logs.EnableFuncCallDepth(false)
	}
}
