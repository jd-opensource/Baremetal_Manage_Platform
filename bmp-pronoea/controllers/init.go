package controllers

import (
	"git.jd.com/bmp-pronoea/models/alert/bmpAlert"
	"git.jd.com/bmp-pronoea/models/rules/bmpRules"
)

// BmpInBandRuleModeles 获取规则
var BmpInBandRuleModeles *bmpRules.BmpInBandMonitorRules

// BmpAlertManagerModels 告警规则处理处理
var BmpAlertManagerModels *bmpAlert.BmpAlertManagerModel


func init() {
	BmpInBandRuleModeles = &bmpRules.BmpInBandMonitorRules{}
	BmpAlertManagerModels = &bmpAlert.BmpAlertManagerModel{}
}

