package monitorAlertLogic

import (
	"encoding/json"

	"coding.jd.com/aidc-bmp/bmp-openapi/constant"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/instanceDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/monitorAlertDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/monitorRuleDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/baseLogic"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/instanceLogic"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/monitorRuleLogic"
	"coding.jd.com/aidc-bmp/bmp-openapi/types/request"
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
	"coding.jd.com/aidc-bmp/bmp-openapi/util"
	log "coding.jd.com/aidc-bmp/bmp_log"
)

func AlertEntity2Alert(logger *log.Logger, entity *monitorAlertDao.MonitorAlerts, isDetail bool) *response.Alert {

	rt := &response.RuleTrigger{}
	if err := json.Unmarshal([]byte(entity.Trigger), rt); err != nil {
		logger.Warnf("AlertEntity2Alert unmarshal Trigger error, value:%s, error:%s", entity.Trigger, err.Error())
		return nil
	}

	res := &response.Alert{
		AlertID:   entity.AlertID,
		AlertTime: int64(entity.AlertTime),
		Resource:  entity.Resource,
		// 资源id,目前就是实例id
		ResourceID: entity.ResourceID,
		// 资源名称,目前就是实例名称
		ResourceName: entity.ResourceName,
		// 触发条件,接口需要翻译
		Trigger:            entity.Trigger,
		TriggerDescription: monitorRuleLogic.Trigger2Description(logger, rt),
		// 报警值
		AlertValue: entity.AlertValue + entity.CalculationUnit,
		// 1表示一般，2表示严重，3表示紧急
		AlertLevel: int64(entity.AlertLevel),
		// 告警持续时间(分钟为单位)
		AlertPeriod: int64(entity.AlertPeriod),
		// 通知对象 userid
		UserID: entity.UserID,
		// 通知对象 用户名
		UserName: entity.UserName,
		// 创建时间戳
		CreatedTime: int64(entity.CreatedTime),
	}
	if logger.GetPoint("language").(string) != baseLogic.EN_US {
		//TODO trans trigger\level ...
		res.Resource = monitorRuleLogic.ResourceMapZh[entity.Resource]
		res.AlertLevelDescription = monitorRuleLogic.RuleAlertLevelMapZh[res.AlertLevel]
	} else {

		res.AlertLevelDescription = monitorRuleLogic.RuleAlertLevelMapEn[res.AlertLevel]
	}
	if isDetail {
		ruleEntity, err := monitorRuleDao.GetMonitorRuleById(logger, entity.RuleID)
		if err != nil {
			logger.Warnf("AlertEntity2Alert.GetMonitorRuleById error, ruleId:%s, error:%s", entity.RuleID, err.Error())
		} else {
			res.Rule = monitorRuleLogic.RuleEntity2Rule(logger, ruleEntity, true)
		}

		instanceEntity, err := instanceDao.GetInstanceById(logger, entity.ResourceID)
		if err != nil {
			logger.Warnf("AlertEntity2Alert.GetInstanceById error, instanceId:%s, error:%s", entity.ResourceID, err.Error())
		} else {
			res.Instance = instanceLogic.InstanceEntity2Instance(logger, instanceEntity, nil)
		}

	} else {
		res.Rule = &response.Rule{
			RuleID:   entity.RuleID,
			RuleName: entity.RuleName,
		}
	}

	return res
}

func DescribeAlert(logger *log.Logger, AlertId string) (*response.Alert, error) {
	alertEntity, err := monitorAlertDao.GetMonitorAlertsById(logger, AlertId)
	if err != nil {
		logger.Warnf("DescribeAlert.GetMonitorAlertsById error, AlertId:%s, error:%s", AlertId, err.Error())
		panic(constant.BuildInvalidArgumentWithArgs(constant.NOT_FOUND.Messagech, constant.NOT_FOUND.MessageEn))
	}
	res := AlertEntity2Alert(logger, alertEntity, true)
	return res, nil

}

func DescribeAlerts(logger *log.Logger, param *request.DescribeAlertsRequest, p util.Pageable) (*response.AlertList, error) {
	offset, limit := p.Pageable2offset()
	query := map[string]interface{}{
		"is_del": 0,
	}
	if param.UserID != "" {
		query["user_id"] = param.UserID
	}
	if param.UserName != "" {
		query["user_name"] = param.UserName
	}
	if param.RuleID != "" {
		query["rule_id"] = param.RuleID
	}
	if param.ResourceID != "" {
		query["resource_id"] = param.ResourceID
	}
	if param.StartTime != 0 {
		query["alert_time.>="] = param.StartTime
	}
	if param.EndTime != 0 {
		query["alert_time.<="] = param.EndTime
	}
	count, err := monitorAlertDao.GetMonitorAlertsCount(logger, query)
	if err != nil {
		logger.Warnf("DescribeAlerts.GetMonitorAlertsCount error, query:%s, error:%s", util.ObjToJson(query), err.Error())
		return nil, err
	}
	entityList := []*monitorAlertDao.MonitorAlerts{}
	if param.IsAll == baseLogic.IS_ALL {
		entityList, err = monitorAlertDao.GetMultiMonitorAlerts(logger, query, []string{}, []string{"id"}, []string{"desc"}, 0, 0)
	} else {
		entityList, err = monitorAlertDao.GetMultiMonitorAlerts(logger, query, []string{}, []string{"id"}, []string{"desc"}, offset, limit)
	}
	if err != nil {
		logger.Warn("DescribeAlerts.GetMultiMonitorAlerts sql error:", err.Error())
		return nil, err
	}
	items := []*response.Alert{}
	for _, alertEntity := range entityList {
		item := AlertEntity2Alert(logger, alertEntity, false)
		items = append(items, item)
	}
	resp := &response.AlertList{
		Alerts:     items,
		PageNumber: p.PageNumber,
		PageSize:   p.PageSize,
		TotalCount: count,
	}

	return resp, nil
}

func DeleteAlert(logger *log.Logger, param *request.DeleteAlertRequest) (bool, error) {

	alertId := param.AlertId
	entity, err := monitorAlertDao.GetMonitorAlertsById(logger, alertId)
	if err != nil {
		logger.Warnf("DeleteAlert.GetMonitorAlertsById error, alertId:%s, error:%s", alertId, err.Error())
		return false, err
	}

	entity.IsDel = 1
	if err := monitorAlertDao.UpdateMonitorAlerts(logger, entity); err != nil {
		logger.Warnf("DeleteAlert.UpdateMonitorAlerts error, ruleId:%s, error:%s", alertId, err.Error())
		return false, err
	}
	return true, nil

}
