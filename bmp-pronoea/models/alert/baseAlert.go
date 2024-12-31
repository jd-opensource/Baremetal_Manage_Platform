package alert

import (
	"fmt"
	"git.jd.com/bmp-pronoea/types/request"
)

const (
	RECEIVER_MODELS_BMP = "bmpAlertReceiver"
)

type IAlert interface {
	ReportAlerts(requestId string, alertInfo *request.AlertReceiverRequest) error
}

type BaseAlert struct {
	IAlert
}

func (b *BaseAlert) ReportAlerts (requestId string, alertInfo *request.AlertReceiverRequest) error {
	if alertInfo == nil || alertInfo.Alerts == nil || len(alertInfo.Alerts) <= 0 {
		return fmt.Errorf("alert info empty ! ")
	}
	return nil
}
