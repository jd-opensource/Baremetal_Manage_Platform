package service

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strconv"

	"coding.jd.com/aidc-bmp/oob-log-alert/object"

	"coding.jd.com/aidc-bmp/oob-log-alert/util"

	"github.com/beego/beego/v2/core/logs"
)

// DailReportEvent process daily report, get cps monitor status from redis and count errors. send mail
func DailReportEvent() error {

	defer func() {
		if r := recover(); r != nil {
			t := make([]byte, 1<<16)
			runtime.Stack(t, true)
			// t, _ := json.Marshal(r)
			logs.Warn("DailReportEvent.panic msg,", string(t))
		}
	}()

	logs.Trace(">>>>>>>>DailReportEvent()")
	defer logs.Trace("<<<<<<<<DailReportEvent()")
	// get cps
	cpsList, err := GetCPSFromRedis()
	if err != nil {
		logs.Warn("DailReportEvent() GetCPSFromRedis Error:%s", err.Error())
		return err
	}

	logs.Debug("DailReportEvent.GetCPSFromRedis res, ", util.ToString(cpsList))
	// end get cps

	idcId := ""
	// totalCount := len(cpsList.CPSs)
	successCount := 0
	oobErrorCount := 0
	otherErrorCount := 0
	var cpsRes []cpsResult
	for _, v := range cpsList.CPSs {
		idcId = v.IdcId

		var res cpsResult
		res.CPS = v
		// fmt.Printf("DailReportEvent sn: %s\n", v.SN)
		logs.Debug("HMGetObjectFromRedis redis:", v.SN, object.CodeField, object.MessageField)
		result, err := util.HMGetObjectFromRedis(v.SN, object.CodeField, object.MessageField)
		logs.Debug("HMGetObjectFromRedis redis:", v.SN, util.ToString(result))
		if err != nil {
			res.Code = -1
			res.Msg = "Get Result From Redis return nil"
			otherErrorCount++
		} else {
			if len(result) < 2 {
				logs.Warn("DailReportEvent.data invalid, passed,", v.SN, result)
				continue
			}

			if result[0] == nil || result[1] == nil { //有些新录入进来的设备，还没有做过带外监控，在redis里面没有数据，所以要跳过
				logs.Warn("DailReportEvent.data invalid, passed,", v.SN, result)
				continue
			}
			var err error
			c, ok := result[0].(string)
			if !ok {
				logs.Warn("DailReportEvent.data code invalid, passed,", v.SN, result)
				continue
			}
			res.Code, err = strconv.Atoi(c)
			if err != nil {
				logs.Warn("DailReportEvent.data code invalid, passed,", v.SN, result)
				continue
			}
			res.Msg, ok = result[1].(string)
			if !ok {
				logs.Warn("DailReportEvent.data msg invalid, passed,", v.SN, result)
				continue
			}

			// res.LastMonitorTimeField = result[2].(string)
			if res.Code == object.MonitorSuccess {
				successCount++
			} else if res.Code == object.MonitorOOBError {
				oobErrorCount++
			} else {
				otherErrorCount++
			}
		}

		cpsRes = append(cpsRes, res)

	}

	// send mail
	mailObject := dailyReportMailTpl{
		IdcId:           idcId,
		TotalCount:      successCount + oobErrorCount + otherErrorCount,
		SuccessCount:    successCount,
		OOBErrorCount:   oobErrorCount,
		OtherErrorCount: otherErrorCount,
		CPSResult:       cpsRes,
	}

	v, _ := json.Marshal(mailObject)
	fmt.Println("dailyReport origin content, ", string(v))
	logs.Debug("dailyReport origin content, ", string(v))

	// logs.Debug("mailObject:%+v", mailObject)

	// begin send mail
	// errMail := util.SendMailByTpl(dailyReportMailTemplate, dailyReportMailSubject, mailObject)
	// if errMail != nil {
	// 	fmt.Println("DailReportEvent() Send Mail Error:%s", errMail.Error())
	// 	logs.Warn("DailReportEvent() Send Mail Error:%s", errMail.Error())
	// 	return errMail
	// }
	// end send mail

	return nil
}
