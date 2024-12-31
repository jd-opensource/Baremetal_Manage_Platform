package parser

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strings"
	"time"

	"coding.jd.com/aidc-bmp/oob-log-agent/object"
	"coding.jd.com/aidc-bmp/oob-log-agent/util"

	log "coding.jd.com/aidc-bmp/bmp_log"
	"github.com/axgle/mahonia"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

const LastCollectTimeKey = "lastCollecttTime"
const timeFormat = "01/02/2006 15:04:05"

// Parser interface to collect and parse text
type Parser interface {
	ParseFRU(*log.Logger) error

	ParseLog(*log.Logger) error
}

func fruCmd() string {
	ipmi, _ := beego.AppConfig.String("oob.ipmi-path")
	return ipmi + " -I lanplus -H %s -U %s -P %s fru print 0"
}

func selListCmd() string {
	ipmi, _ := beego.AppConfig.String("oob.ipmi-path")
	return ipmi + " -I lanplus -H %s -U %s -P %s sel list "
}

// BaseParser Abstract Parser
type BaseParser struct {
	SN      string
	Brand   string
	IloIP   string
	IloUser string
	IloPass string
}

func (b BaseParser) ParseFRU(logger *log.Logger) error {
	logger.Infof(">>>>>>>>BaseParser.parseFRU(), sn:%s, ilo_ip:%s", b.SN, b.IloIP)
	defer logger.Infof("<<<<<<<<BaseParser.parseFRU(), sn:%s, ilo_ip:%s", b.SN, b.IloIP)

	var (
		res string
		err error
	)

	if runtime.GOOS == "windows" {
		res, _, err = util.ExecCommand(logger, "CMD", "/C", fmt.Sprintf(fruCmd(), b.IloIP, b.IloUser, b.IloPass))
	} else {
		res, _, err = util.ExecCommand(logger, "sh", "-c", fmt.Sprintf(fruCmd(), b.IloIP, b.IloUser, b.IloPass))
	}

	if err != nil { //带外不通

		code := object.MonitorOOBError
		message := "ParseFRU Error: " + err.Error()
		saveResult(b.SN, code, message)

		logger.Warnf("Run Command %s error:%s", fmt.Sprintf(fruCmd(), b.IloIP, b.IloUser, b.IloPass), err.Error())
		// 带外错误，增加计数，如果达到配置阈值，告警
		count, err1 := util.HIncrObjectToRedis(object.ErrorCountField, b.SN, 1)
		if err1 != nil {
			logger.Warnf("%+v. HIncrObjectToRedis Error : %s", b, err.Error())
		}

		// 大于配置的告警次数阈值，发送告警
		if count > beego.AppConfig.DefaultInt("oob.error-alert-count", 3) && count <= (beego.AppConfig.DefaultInt("oob.error-alert-count", 3)+beego.AppConfig.DefaultInt("oob.error-alert-times", 3)) {
			err := util.PubEventToRedis(object.OOBErrorChannel+":"+b.SN, err.Error())
			if err != nil {
				logger.Warnf("%+v. PubEventToRedis Error : %s", b, err.Error())
			}
		}
		return err
	} else {
		// 带外连接成功，重置错误计数器
		if err := util.HDelObjectFromRedis(object.ErrorCountField, []string{b.SN}); err != nil {
			logger.Warnf("%+v. HDelObjectFromRedis ErrorCountField Error : %s", b, err.Error())
		}
	}
	// logs.Debug("Collect():" + res)

	var retStr string
	if runtime.GOOS == "windows" {
		enc := mahonia.NewDecoder("GBK")
		retStr = enc.ConvertString(res)
	} else {
		retStr = res
	}

	var obj object.FRUObject

	lines := strings.Split(retStr, "\n")

	for _, line := range lines {

		tmp := strings.Split(line, ":")
		if strings.Contains(tmp[0], "Product Manufacturer") {
			obj.Brand = strings.TrimSpace(tmp[1])
		}
		if strings.Contains(tmp[0], "Product Serial") {
			obj.SN = strings.TrimSpace(tmp[1])
		}
		if strings.Contains(tmp[0], "Product Name") {
			obj.Model = strings.TrimSpace(tmp[1])
		}
	}

	// 比较Manufacturer、SN，如果不同，发告警
	isDiff := false
	var diffMsg string
	// if !strings.EqualFold(obj.Manufacturer, b.Manufacturer) { //Manufacturer不一致
	// 	isDiff = true
	// 	diffMsg = diffMsg + fmt.Sprintf("oob Manufacturer:%s", b.Manufacturer)
	// }
	if !strings.EqualFold(obj.SN, b.SN) { //sn不一致
		isDiff = true
		diffMsg = diffMsg + fmt.Sprintf("oob SN:%s", b.SN)
	}

	if isDiff {
		message := fmt.Sprintf("Read FRU Different: SN = %s, Brand = %s", obj.SN, obj.Brand) + ", " + diffMsg
		err := util.PubEventToRedis(object.OOBErrorChannel+":"+b.SN, message)
		if err != nil {
			logger.Warnf("%+v. PubEventToRedis Error : %s", b, err.Error())
		}
	}

	logger.Infof("BaseParser.parseFRU():obj=%+v", obj)

	return nil

}

func (p BaseParser) ParseLog(logger *log.Logger) error {

	logger.Infof(">>>>>>>>BaseParser.ParseLog(), sn:%s, ilo_ip:%s", p.SN, p.IloIP)
	defer logger.Infof("<<<<<<<<BaseParser.ParseLog(), sn:%s, ilo_ip:%s", p.SN, p.IloIP)

	code := object.MonitorSuccess
	message := "Success"

	var (
		res string
		err error
	)
	if runtime.GOOS == "windows" {
		res, _, err = util.ExecCommand(logger, "CMD", "/C", fmt.Sprintf(selListCmd(), p.IloIP, p.IloUser, p.IloPass))
	} else {
		res, _, err = util.ExecCommand(logger, "sh", "-c", fmt.Sprintf(selListCmd(), p.IloIP, p.IloUser, p.IloPass))
	}

	if err != nil {
		logger.Warnf("parseLog() oob connection Error, sn:%s, err:%s ", p.SN, err.Error())
		code = object.MonitorOOBError
		message = fmt.Sprintf("parseLog() Error=%s", err.Error())
		saveResult(p.SN, code, message)
		return err
	}

	// logs.Debug(p.SN, " parseSelInfo Collect():", []byte(res))

	var retStr string
	if runtime.GOOS == "windows" {
		enc := mahonia.NewDecoder("GBK")
		retStr = enc.ConvertString(res)
	} else {
		retStr = res
	}

	logger.Info("sel list result,", p.SN, retStr)

	lines := strings.Split(retStr, "\n")

	var updatedLastCollectTime string
	var lastCollectTime time.Time = getLastCollectTime(p.SN)

	toPushMessages := []string{}
	for i := 0; i < len(lines); i++ {
		// 1fb | 07/05/2021 | 15:22:14 | System Firmwares #0x1d | System boot initiated | Asserted
		line := strings.TrimSpace(lines[i])
		items := strings.Split(line, "|")
		if len(items) != 6 {
			logger.Info(p.SN, " Sel list item invalid:", line)
			continue
		}
		if strings.TrimSpace(items[1]) == "Pre-Init" {
			//d7c |  Pre-Init  |0000028859| Drive Slot / Bay #0xbc | Drive Present | Asserted
			// 这种带外日志无意义
			continue
		}
		eventDate := strings.TrimSpace(items[1])
		eventTime := strings.TrimSpace(items[2])

		eventT, err := time.Parse(timeFormat, fmt.Sprintf("%s %s", eventDate, eventTime))
		if err != nil {
			logger.Warnf("eventT time parse error,sn:%s, line:%s, eventDate:%s, eventTime:%s ", p.SN, line, eventDate, eventTime)
		}
		if !eventT.After(lastCollectTime) {
			continue
		}

		updatedLastCollectTime = fmt.Sprintf("%s %s", eventDate, eventTime)
		// eventSensor := strings.Split(strings.TrimSpace(items[3]), "#")[0]
		// eventDesc := strings.TrimSpace(items[4])
		// eventDetail := strings.TrimSpace(items[5])
		toPushMessages = append(toPushMessages, line)
	}

	if updatedLastCollectTime != "" { //根据最后一条日志的时间更新redis里面此sn的采集截止时间
		if err := util.HSetObjectToRedis(LastCollectTimeKey, p.SN, updatedLastCollectTime); err != nil {
			logger.Warn("set getLastCollectTime err,", p.SN, updatedLastCollectTime, err.Error())
		} else {
			logger.Info("set getLastCollectTime success,", p.SN, updatedLastCollectTime)
		}
	}

	if len(toPushMessages) > 0 {

		msg := object.LogObject{
			SeqNumber: p.SN,
			Messages:  toPushMessages,
		}
		b, err := json.Marshal(msg)
		if err != nil {
			logger.Warnf("toPushMessages marshal error: %s", err.Error())
			return err
		}
		err = util.PubEventToRedis(object.LogChannel+":"+p.SN, string(b))
		if err != nil {
			info := fmt.Sprintf("%s LogAlert PubEventToRedis ERROR, MSG: %s, ERROR:%s", p.SN, string(b), err.Error())
			logger.Warn(info)
			return err
		} else {
			info := fmt.Sprintf("%s LogAlert PubEventToRedis Succeed, MSG: %s", p.SN, string(b))
			logger.Info(info)
		}
	}
	saveResult(p.SN, code, message)

	return nil

}

func getLastCollectTime(sn string) time.Time {
	res, err := util.HGetObjectFromRedis(LastCollectTimeKey, sn)
	if err != nil {
		logs.Warn("getLastCollectTime error, sn :", sn, err.Error())
		return time.Now().AddDate(-10, 0, 0)
	}
	t, err := time.Parse(timeFormat, res)
	if err != nil {
		logs.Warn("Parse LastCollectTime error, sn :%s, time: %s", sn, res)
		return time.Now().AddDate(-10, 0, 0)
	}
	return t

}

func saveResult(key string, code int, msg string) error {
	fields := map[string]interface{}{
		object.MessageField: msg,
		object.CodeField:    code,
	}

	// 成功，重置错误计数
	if code == object.MonitorSuccess {
		fields[object.ErrorCountField] = 0
	}

	err := util.HMSetObjectToRedis(key, fields)
	if err != nil {
		logs.Warn("Write %+v. Message error:%s", key, err.Error())
		return err
	}
	return nil
}

func GetParser(logger *log.Logger, m object.CPSRecord) Parser {
	logger.Infof("GetOobUserPasswd CPSRecord:%s", util.Obj2String(m))
	var oobUser, oobPasswd = util.GetOobUserPasswd(m)

	return BaseParser{SN: m.SN, IloIP: m.IloIP, IloUser: oobUser, IloPass: oobPasswd, Brand: m.Brand}

}
