package util

import (
	"github.com/astaxie/beego/logs"
	"time"
)

func GetCurrentTimeStrWithlayout(layout string) string {
	return time.Now().Format(layout)
}

func GetCurrentTimeStr() string {
	return GetCurrentTimeStrWithlayout("2006-01-02 15:04:05")
}

func GetTimestampByStr(str string) int64 {
	layout := "2006-01-02 15:04:05"
	return GetTimestampByStrWithLayout(str, layout)
}

func GetTimestampByStrWithLayout(str, layout string) int64 {
	loc, err := time.LoadLocation("Local")
	if err != nil {
		logs.Warn("load local location error:%v", err)
		return 0
	}

	//layout := "2006-01-02 15:04:05"
	tm, err := time.ParseInLocation(layout, str, loc)
	if err != nil {
		logs.Debug("get timestamp by str:%v error:%v", str, err)
		return 0
	}

	t := tm.Unix()

	return t
}

func GetCurrentTimestamp() int64 {
	tNowStr := GetCurrentTimeStr()
	return GetTimestampByStr(tNowStr)
}

func GetTimeStrByTimestamp(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 15:04:05")
}

func GetUtcTimeByTimeStr(timeStr string) (string, error) {
	t, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)
	if err != nil {
		return "", err
	}
	utcTime := t.UTC().Format("2006-01-02T15:04:05Z")
	return utcTime, nil
}

func GetUtcTimeByTimeStamp(timestamp int64) (string, error) {
	timeStr := GetTimeStrByTimestamp(timestamp)
	return GetUtcTimeByTimeStr(timeStr)
}