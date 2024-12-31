package service

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	log "coding.jd.com/aidc-bmp/bmp_log"
	"coding.jd.com/aidc-bmp/oob-log-alert/object"
	"coding.jd.com/aidc-bmp/oob-log-alert/service/mailLogic"
	"coding.jd.com/aidc-bmp/oob-log-alert/util"

	"coding.jd.com/aidc-bmp/oob-log-alert/dao"
	messageType "git.jd.com/cps-golang/ironic-common/ironic/enums/MessageType"

	"github.com/beego/beego/v2/core/logs"
)

const timeFormat = "01/02/2006 15:04:05"

// LogEvent process Logs event, filter and send mail
func LogEvent(logger *log.Logger, sn, message string, isSaveLog bool, logPath string) error {

	log := object.LogObject{}

	errJSON := json.Unmarshal([]byte(message), &log)
	if errJSON != nil {
		logs.Warn("%s LogEvent %s JSONUmarshal Error %s", sn, errJSON.Error())
		return errJSON
	}
	// end check

	// get cps
	cps, err := GetCPS(sn)
	if err != nil {
		return err
	}
	// end get cps
	mailAlertItems := []MailAlertLog{}
	mailMessageItems := []*dao.WebMessage{}
	logItems := []*dao.CpsLogItems{}

	// GetFilter
	filters, err := dao.GetFaultUsedRules(logger)
	if err != nil {
		logs.Warn("%s GetFaultUsedRules err: %s", sn, err.Error())
		return err
	} else {
		// filter
		for i := len(log.Messages) - 1; i >= 0; i-- {
			line := log.Messages[i]

			items := strings.Split(line, "|")
			if len(items) != 6 {
				logs.Debug(sn, " Sel list item invalid:", line)
				continue
			}
			logid := strings.TrimSpace(items[0])
			eventTime, err := time.Parse(timeFormat, fmt.Sprintf("%s %s", strings.TrimSpace(items[1]), strings.TrimSpace(items[2])))
			if err != nil {
				logs.Warn("%s Sel list item time invalid:%s", sn, line)
			}
			logcondition := strings.TrimSpace(items[3])
			logdesc := strings.TrimSpace(strings.Join(items[4:], "|"))
			eventDetail := strings.TrimSpace(strings.Join(items[3:], "|"))

			logger.Infof("logcondition is %s, logdesc is :%s\n", logcondition, logdesc)

			isMatch := false

			var selectedf *dao.CpsFaultRules

			for _, f := range filters {
				if strings.HasPrefix(logcondition, strings.TrimSpace(f.Condition)) && strings.HasPrefix(logdesc, strings.TrimSpace(f.Threshold)) {
					isMatch = true
					selectedf = f
					break
				}

			}

			collectTime := time.Now()

			if isMatch {

				cid, err := dao.InsertAndGetCollectionID(logger, sn, eventDetail, selectedf.ID, collectTime, eventTime, selectedf.Level)
				if err != nil {
					logger.Warnf("InsertAndGetCollectionID error: %s", err.Error())
					return nil
				} else {
					logger.Infof("InsertAndGetCollectionID success, sn: %s, line:%s, cid:%d", sn, line, cid)
				}

				item := dao.CpsLogItems{
					Sn:           sn,
					CollectionID: int(cid),
					EventTime:    eventTime,
					CollectTime:  collectTime,
				}
				logItems = append(logItems, &item)
				mailAlertItems = append(mailAlertItems, MailAlertLog{
					EventTime:   eventTime,
					FaultTypeZh: selectedf.FaultType,
					FaultNameZh: selectedf.Desc,
				})

				//插入消息到message表
				c := fmt.Sprintf("%s/%s", cps.InstanceName, sn)
				if cps.InstanceName == "" {
					c = sn
				}
				message := &dao.WebMessage{
					MessageID:      util.GetUuid("m-"),
					ResourceType:   "oob-alert-log",
					FinishTime:     int(time.Now().Unix()),
					MessageType:    messageType.MessageTypeOobMonitor,
					MessageSubType: selectedf.FaultType,
					SN:             item.Sn,
					FaultType:      selectedf.FaultType,
					Content:        eventDetail,
					Detail:         c, //都赋值，
					LogID:          logid,
					AlertTime:      int(item.EventTime.Unix()),
					AlertCount:     1,
				}
				instanceEntity, err := dao.GetInstanceBySn(logger, item.Sn)
				if err != nil {
					logger.Warnf("InsertMessage.GetInstanceBySn error, sn:%s, error:%s", item.Sn, err.Error())
				}
				if instanceEntity != nil {
					message.InstanceID = instanceEntity.InstanceID
					message.InstanceName = instanceEntity.InstanceName
					message.UserID = instanceEntity.UserID

					user, err := dao.GetUserById(logger, instanceEntity.UserID)
					if err != nil {
						logger.Warnf("InsertMessage.GetUserById error, user_id:%s, error:%s", instanceEntity.UserID, err.Error())
					}
					if user != nil {
						message.UserName = user.UserName
					}
				}
				mailMessageItems = append(mailMessageItems, message)
				err = dao.InsertMessage(logger, message)
				if err != nil {
					logger.Warnf("sn %s insert message err, error:%s, content:%s ", cps.SN, err.Error(), util.ToString(item))
				} else {
					logger.Infof("sn %s insert message succeed, content:%s", cps.SN, util.ToString(item))
				}

			}
			if !isMatch {
				//告警日志被过滤掉时，记录日志方便排查
				info := fmt.Sprintf("%s Log alert filtered, line: %s", sn, line)
				logger.Warnf(info)
			} else {
				info := fmt.Sprintf("%s Log alert selected, line: %s, rule:%s", sn, line, selectedf.Desc)
				logger.Info(info)
			}
		}
		// end filter

	}

	if len(logItems) == 0 {
		return nil
	}

	// insert to Database
	err = dao.BatchInsertLogItems(logger, logItems)
	if err != nil {
		logger.Warnf("sn %s insert oob-log-items err, error:%s, content:%s ", cps.SN, err.Error(), util.ToString(logItems))
	} else {
		logger.Infof("sn %s insert oob-log-items succeed, content:%s", cps.SN, util.ToString(logItems))
	}

	// if dailyMoreThanOnce(sn, message) { //同一设备同类型的报警，每天只邮件报一次
	// 	return nil
	// }

	// send mail
	// mailObject := logMailTpl{CPS: cps, Logs: mailAlertItems}
	// errMail := util.SendMailByTpl(logMailTemplate, logMailSubject, mailObject)
	// if errMail != nil {
	// 	logs.Warn(cps.SN + " Send LogEvent Mail failed: " + errMail.Error())
	// 	return errMail
	// } else {
	// 	logs.Debug(cps.SN + " Send LogEvent Mail succeed")
	// }
	// end send mail
	if len(mailMessageItems) > 0 {
		items := []mailLogic.MailMessage{}
		for _, v := range mailMessageItems {
			items = append(items, mailLogic.MailMessage{
				MessageID:      v.MessageID,
				ResourceType:   v.ResourceType,
				FinishTime:     time.Unix(int64(v.FinishTime), 0).Format("2006-01-02 15:04:05"),
				MessageType:    v.MessageType,
				MessageSubType: v.MessageSubType,
				SN:             v.SN,
				FaultType:      v.FaultType,
				Content:        v.Content,
				Detail:         v.Detail,
				LogID:          v.LogID,
				AlertTime:      time.Unix(int64(v.AlertTime), 0).Format("2006-01-02 15:04:05"),
				AlertCount:     fmt.Sprintf("%d", v.AlertCount),

				InstanceID:   v.InstanceID,
				InstanceName: v.InstanceName,
				UserID:       v.UserID,
				UserName:     v.UserName,
			})
		}
		subject := fmt.Sprintf("故障消息-【%s】提醒", mailMessageItems[0].MessageSubType)
		mailTpl := mailLogic.MessageOobTemplate
		err = mailLogic.SendMailByTpl(logger, mailTpl, subject, items)
		if err != nil {
			logger.Warnf("oob message SendMailByTpl error:%s", err.Error())
		} else {
			logger.Info("oob message SendMailByTpl success")
		}
	}

	return nil
}

// CPSEvent process CPS event for oob access error, send mail
func CPSEvent(sn, message string) error {

	if message == "" {
		return nil
	}

	// get cps
	// cps, err := GetCPS(sn)
	// if err != nil {
	// 	return err
	// }
	// end get cps

	// send mail
	// mailObject := alertOOBMailTpl{CPS: cps, Err: message}
	// logs.Debug(sn, "mailObject:", mailObject)

	// errMail := util.SendMailByTpl(alertOOBMailTemplate, alertOOBMailSubject, mailObject)
	// if errMail != nil {
	// 	logs.Warn(sn, "Send CPSEvent Mail Error:%s", errMail.Error())
	// 	return errMail
	// } else {
	// 	logs.Warn(sn, "Send CPSEvent Mail Succeed")
	// }
	// end send mail

	return nil

}
