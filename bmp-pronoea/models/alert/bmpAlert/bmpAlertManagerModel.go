package bmpAlert

import (
	"encoding/json"
	"fmt"
	"git.jd.com/bmp-pronoea/common/util"
	"git.jd.com/bmp-pronoea/dao/api/bmpApi"
	"git.jd.com/bmp-pronoea/dao/api/bmpApi/exec"
	"git.jd.com/bmp-pronoea/models"
	"git.jd.com/bmp-pronoea/models/alert"
	"git.jd.com/bmp-pronoea/models/rules/bmpRules"
	"git.jd.com/bmp-pronoea/types/request"
	"github.com/astaxie/beego/logs"
	"net/url"
	"strconv"
)

type BmpAlertManagerModel struct {
	alert.BaseAlert
}

func (b *BmpAlertManagerModel) ReportAlerts(requestId string, alertMsg *request.AlertReceiverRequest) error {
	err := b.BaseAlert.ReportAlerts(requestId, alertMsg)
	if err != nil {
		return err
	}

	bmpAlertMsgDetailList, err := b.turnReportAlertInfoList(requestId, alertMsg)
	if err != nil {
		logs.Info(requestId, fmt.Sprintf("turn bmp report alert info error %v", err))
		return fmt.Errorf("turn bmp report alert info error %v", err)
	}
	if bmpAlertMsgDetailList == nil || len(bmpAlertMsgDetailList) <= 0 {
		logs.Info(requestId, fmt.Sprintf("bmp report alert info do not have firing %v", err))
		return nil
	}
	bmpAlertMsg := &bmpApi.BmpAlertRequest{
		Alerts : bmpAlertMsgDetailList,
	}
	return  exec.ReportAlertInfoToBmp(requestId, bmpAlertMsg)
}

func (b *BmpAlertManagerModel) turnReportAlertInfoList(requestId string, alertMsg *request.AlertReceiverRequest) ([]*bmpApi.BmpAlertDetail, error) {
	if alertMsg == nil || alertMsg.Alerts == nil || len(alertMsg.Alerts) <= 0 {
		return nil, fmt.Errorf("alert receive info empty ！")
	}

	currentTime := util.GetCurrentTimestamp()
	bmpAlertMsg := make([]*bmpApi.BmpAlertDetail, 0)
	for _, alertInfo := range alertMsg.Alerts {
		if alertInfo.Status != request.ALERT_RECEIVE_STATUS_FIRING {
			continue
		}
		if alertInfo.Labels == nil || len(alertInfo.Labels) < 0 {
			continue
		}

		tiggerTmp, err := b.turnReportAlertInfoTrigger(requestId,alertInfo.Labels)
		if err != nil {
			logs.Info(requestId, fmt.Sprintf("turnReportAlertInfoList fail, turn tigger error %v", err))
			continue
		}
		triggerJsonByte, err := json.Marshal(tiggerTmp)
		if err != nil {
			logs.Info(requestId, fmt.Sprintf("alert info labels json.Marshal error %v %v", err, alertInfo.Labels))
			continue
			//return nil, fmt.Errorf("alert info labels json.Marshal error %v", err)
		}

		tmpMsg := &bmpApi.BmpAlertDetail{}
		tmpMsg.RuleName = alertInfo.Labels[bmpRules.SYSTEM_BMPLABELS_ALERTNAME]
		tmpMsg.RuleID   = alertInfo.Labels[bmpRules.CUSTOM_BMPLABELS_RULEID]
		tmpMsg.InstanceID = alertInfo.Labels[bmpRules.CUSTOM_BMPLABELS_INSTANCEID]
		tmpMsg.CalculationUnit = alertInfo.Labels[bmpRules.CUSTOM_BMPLABELS_CALCULATION_UNIT]
		tmpMsg.Trigger = url.QueryEscape(string(triggerJsonByte))

		alertValue := alertInfo.Annotations[bmpRules.CUSTOM_BMPLABELS_CURRENTVALUE]
		alertValueFloat64, err := util.TurnStringToFloat64(alertValue, 2)
		if err != nil {
			logs.Info(requestId, fmt.Sprintf("turn alert value error %v %v", err, alertValue))
			continue
		}

		if metricTmp, exist := alertInfo.Labels[bmpRules.CUSTOM_BMPLABELS_METRIC]; exist {
			if unitConversion, exist := models.MetricNameUnitConversionMap[metricTmp]; exist {
				alertValueFloat64 = alertValueFloat64 / unitConversion
			}
		}
		alertValue = strconv.FormatFloat(alertValueFloat64, 'f', 2, 64)

		tmpMsg.AlertValue = alertValue
		tmpMsg.AlertTimestamp = currentTime

		tmpMsg.AlertPeriod = tiggerTmp.Period * tiggerTmp.Times

		bmpAlertMsg = append(bmpAlertMsg, tmpMsg)
	}
	return bmpAlertMsg, nil
}

func (b *BmpAlertManagerModel) turnReportAlertInfoTrigger(requestId string, labelsMap map[string]string) (*request.RuleTrigger, error) {
	if labelsMap == nil || len(labelsMap) <= 0 {
		return nil, fmt.Errorf("labels info empty! ")
	}

	trigger := &request.RuleTrigger{}
	trigger.Metric = labelsMap[bmpRules.CUSTOM_BMPLABELS_METRIC]
	trigger.Period = 0
	if _, exist := labelsMap[bmpRules.CUSTOM_BMPLABELS_PERIOD]; exist {
		tmpInt, err := strconv.Atoi(labelsMap[bmpRules.CUSTOM_BMPLABELS_PERIOD])
		if err == nil {
			trigger.Period = tmpInt
		} else {
			logs.Info(requestId, fmt.Sprintf("turn labelsMap[period] to int error %v", err))
		}
	}

	trigger.Times = 0
	if _, exist := labelsMap[bmpRules.CUSTOM_BMPLABELS_TIMES]; exist {
		tmpInt, err := strconv.Atoi(labelsMap[bmpRules.CUSTOM_BMPLABELS_TIMES])
		if err == nil {
			trigger.Times = tmpInt
		} else {
			logs.Info(requestId, fmt.Sprintf("turn labelsMap[times] to int error %v", err))
		}
	}

	trigger.NoticeLevel = 0
	if _, exist := labelsMap[bmpRules.CUSTOM_BMPLABELS_NOTICELEVEL]; exist {
		tmpInt, err := strconv.Atoi(labelsMap[bmpRules.CUSTOM_BMPLABELS_NOTICELEVEL])
		if err == nil {
			trigger.NoticeLevel = tmpInt
		} else {
			logs.Info(requestId, fmt.Sprintf("turn labelsMap[noticeLevel] to int error %v", err))
		}
	}

	trigger.Threshold = 0
	if _, exist := labelsMap[bmpRules.CUSTOM_BMPLABELS_THRESHOID]; exist {
		tmpFloat64, err := strconv.ParseFloat(labelsMap[bmpRules.CUSTOM_BMPLABELS_THRESHOID], 64)
		if err == nil {
			trigger.Threshold = tmpFloat64
		} else {
			logs.Info(requestId, fmt.Sprintf("turn labelsMap[threshold] to float error %v", err))
		}
	}

	trigger.Calculation     = labelsMap[bmpRules.CUSTOM_BMPLABELS_CALCULATION]
	trigger.CalculationUnit = labelsMap[bmpRules.CUSTOM_BMPLABELS_CALCULATION_UNIT]
	trigger.Operation       = labelsMap[bmpRules.CUSTOM_BMPLABELS_OPERATION]

	return trigger, nil
}