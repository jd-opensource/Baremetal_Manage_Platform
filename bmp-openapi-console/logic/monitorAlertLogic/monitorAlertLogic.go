package monitorAlertLogic

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"coding.jd.com/aidc-bmp/bmp-openapi/constant"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/deviceDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/instanceDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/messageDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/monitorAlertDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/monitorRuleDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/userDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/baseLogic"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/instanceLogic"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/mailLogic"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/messageLogic"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/monitorRuleLogic"
	"coding.jd.com/aidc-bmp/bmp-openapi/types/request"
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
	"coding.jd.com/aidc-bmp/bmp-openapi/util"
	log "coding.jd.com/aidc-bmp/bmp_log"
	messageType "git.jd.com/cps-golang/ironic-common/ironic/enums/MessageType"
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

func AddAlert(logger *log.Logger, paramOrigin *request.AddAlertRequest) (bool, error) {

	for _, param := range paramOrigin.Alerts {

		ruleId := param.RuleID
		ruleEntity, err := monitorRuleDao.GetMonitorRuleById(logger, ruleId)
		if err != nil {
			logger.Warnf("AddAlert.GetMonitorRuleById error, ruleId:%s, error:%s", ruleId, err.Error())
			panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("ruleId %s 不存在", ruleId), fmt.Sprintf("ruleId  %s notFound", ruleId)))
		}
		//对接的第三方服务，需要在这里补齐一些信息
		logger.Point("userId", ruleEntity.UserID)
		logger.Point("username", ruleEntity.UserName)

		rt := request.RuleTrigger{}
		ro := request.RuleNotice{}
		//trigger里面有%>=!=等各种特殊字符，约定赛迪对trigger做urlencode，所以这里需要做一次urldecode操作
		param.Trigger, err = url.QueryUnescape(param.Trigger)
		if err != nil {
			logger.Warnf("AddAlert.QueryUnescape error:%s", err.Error())
			return false, err
		}
		if err := json.Unmarshal([]byte(param.Trigger), &rt); err != nil {
			logger.Warnf("AddAlert param.Trigger unmarshal error, value:%s, error:%s", param.Trigger, err.Error())
			panic(constant.BuildInvalidArgumentWithArgs("param.Trigger 格式错误", "param.Trigger invalid"))
		}
		if err := json.Unmarshal([]byte(ruleEntity.NoticeOption), &ro); err != nil {
			logger.Warnf("AddAlert NoticeOption unmarshal error, value:%s, error:%s", ruleEntity.NoticeOption, err.Error())
			panic(constant.BuildInvalidArgumentWithArgs("ruleEntity.NoticeOption 格式错误", "ruleEntity.NoticeOption invalid"))
		}

		instanceId := param.InstanceID
		instanceEntity, err := instanceDao.GetInstanceById(logger, instanceId)
		if err != nil {
			logger.Warnf("AddAlert.GetInstanceById error, instanceId:%s, error:%s", instanceId, err.Error())
			panic(constant.BuildInvalidArgumentWithArgs(fmt.Sprintf("instanceId %s 不存在", instanceId), fmt.Sprintf("instanceId  %s notFound", instanceId)))
		}
		alertEntity := &monitorAlertDao.MonitorAlerts{
			AlertID:         util.GetUuid("ma_"),
			AlertTime:       int(param.AlertTimestamp),
			RuleID:          param.RuleID,
			RuleName:        ruleEntity.RuleName,
			Resource:        "instance",
			ResourceID:      param.InstanceID,
			ResourceName:    instanceEntity.InstanceName,
			Trigger:         param.Trigger,
			AlertValue:      param.AlertValue,
			CalculationUnit: param.CalculationUnit,
			AlertLevel:      int8(rt.NoticeLevel),
			AlertPeriod:     param.AlertPeriod,
			UserID:          ruleEntity.UserID,
			UserName:        ruleEntity.UserName,
			CreatedTime:     int(time.Now().Unix()),
			IsRecover:       int8(param.IsRecover),
			ProjectID:       ruleEntity.ProjectID,
		}
		if param.IsRecover == 0 { //0表示告警，1表示恢复
			if _, err := monitorAlertDao.AddMonitorAlerts(logger, alertEntity); err != nil {
				logger.Warnf("AddAlert.AddMonitorAlerts error:%s", err.Error())
				return false, err
			}

			ruleEntity.Status = 3
			if err := monitorRuleDao.UpdateMonitorRules(logger, ruleEntity); err != nil {
				logger.Warnf("AddAlert.UpdateMonitorRules error:%s", err.Error())
			}
		}

		if !judgeIsPublish(logger, ro) {
			logger.Warn("judgeIsPublish passed!")
			continue
		}

		if err := NoticeAlert(logger, alertEntity, ruleEntity, instanceEntity); err != nil {
			logger.Warnf("AddAlert.NoticeAlert error, alertId:%s, error:%s", alertEntity.AlertID, err.Error())
			return false, err
		}
	}

	return true, nil
}

//规则制定了哪个时间段告警需要通知，不在此时间段的告警不做通知。true表示需要通知，false不需要
func judgeIsPublish(logger *log.Logger, ro request.RuleNotice) bool {
	tz := logger.GetPoint("timezone").(string)
	if tz == "" {
		tz = "Asia/Shanghai"
	}
	//2024-07-15 09:27:23只取09:27:23
	timeStr := strings.Split(util.TimestampToString(time.Now().Unix(), tz), " ")[1]
	start := ro.EffectiveIntervalStart
	end := ro.EffectiveIntervalEnd

	//
	timeStrItems := strings.Split(timeStr, ":")
	startItems := strings.Split(start, ":")
	endItems := strings.Split(end, ":")

	logger.Infof("judgeIsPublish info, nowTime:%s, noticeStartTime:%s, noticeEndTime:%s", timeStrItems, startItems, endItems)

	t1, _ := strconv.Atoi(timeStrItems[0])
	t2, _ := strconv.Atoi(timeStrItems[1])
	t3, _ := strconv.Atoi(timeStrItems[2])
	tReal := t1*3600 + t2*60 + t3

	s1, _ := strconv.Atoi(startItems[0])
	s2, _ := strconv.Atoi(startItems[1])
	s3, _ := strconv.Atoi(startItems[2])
	sReal := s1*3600 + s2*60 + s3

	e1, _ := strconv.Atoi(endItems[0])
	e2, _ := strconv.Atoi(endItems[1])
	e3, _ := strconv.Atoi(endItems[2])
	eReal := e1*3600 + e2*60 + e3

	if tReal >= sReal && tReal <= eReal {
		return true
	}
	return false
}

func NoticeAlert(logger *log.Logger, alertEntity *monitorAlertDao.MonitorAlerts, ruleEntity *monitorRuleDao.MonitorRules, instanceEntity *instanceDao.Instance) error {
	logger.Infof("addAlert.DEBUG, ruleEntity:%s", util.InterfaceToJson(ruleEntity))

	rt := request.RuleTrigger{}
	ro := request.RuleNotice{}
	if err := json.Unmarshal([]byte(alertEntity.Trigger), &rt); err != nil {
		logger.Warnf("NoticeAlert param.Trigger unmarshal error, value:%s, error:%s", alertEntity.Trigger, err.Error())
		return nil
	}
	if err := json.Unmarshal([]byte(ruleEntity.NoticeOption), &ro); err != nil {
		logger.Warnf("NoticeAlert NoticeOption unmarshal error, value:%s, error:%s", ruleEntity.NoticeOption, err.Error())
		return nil
	}

	deviceEntity, err := deviceDao.GetDeviceById(logger, instanceEntity.DeviceID)
	if err != nil {
		logger.Warnf("NoticeAlert.GetDeviceById error, deviceId:%s, error:%s", instanceEntity.DeviceID, err.Error())
		return nil
	}

	noticeWays := ro.NoticeWay
	alertR := AlertEntity2Alert(logger, alertEntity, true)
	userEntity, err := userDao.GetUserById(logger, ruleEntity.UserID)
	logger.Infof("addAlert.DEBUG, user:%s", util.InterfaceToJson(userEntity))
	if err != nil {
		logger.Warnf("NoticeAlert.GetUserById error, userId:%s, error:%s", ruleEntity.UserID, err.Error())
	}

	if util.InArray(int64(1), noticeWays) { //1表示站内信
		//专门为带内监控定制的消息detail格式
		type OT struct {
			Trigger     string `json:"trigger"`
			AlertTime   int    `json:"alertTime"`
			AlertPeriod int    `json:"alertPeriod"`
			AlertValue  string `json:"alertValue"`
		}
		d := OT{
			Trigger:     alertEntity.Trigger,
			AlertTime:   alertEntity.AlertTime,
			AlertPeriod: alertEntity.AlertPeriod,
			AlertValue:  alertEntity.AlertValue + alertEntity.CalculationUnit,
		}
		detail := fmt.Sprintf("触发规则:%s, 告警值:%s, 告警持续时间:%d分钟", alertR.TriggerDescription, d.AlertValue, d.AlertPeriod)

		if userEntity != nil && userEntity.Language == baseLogic.EN_US {
			detail = fmt.Sprintf("trigger condition:%s, alert value:%s, alert continue time:%dminutes", alertR.TriggerDescription, d.AlertValue, d.AlertPeriod)
		}

		if alertEntity.IsRecover == 1 { //0表示告警，1表示恢复
			detail = "【恢复】" + detail
		}

		webMsg := &messageDao.WebMessage{
			MessageID:      util.GetUuid("m-"),
			SN:             deviceEntity.Sn,
			MessageType:    messageType.MessageTypeInBondAlert,
			MessageSubType: rt.Metric,
			ResourceType:   "inBondAlert",
			FinishTime:     d.AlertTime,
			Detail:         detail,
			Content:        detail,
			UserID:         ruleEntity.UserID,
			UserName:       ruleEntity.UserName,
			InstanceID:     instanceEntity.InstanceID,
			InstanceName:   instanceEntity.InstanceName,
			Result:         "success",
			RuleID:         ruleEntity.RuleID,
			RuleName:       ruleEntity.RuleName,
			IsRecover:      alertEntity.IsRecover,
		}
		if webMsg.IsRecover == 0 { //0表示告警，1表示恢复
			if err := messageLogic.AddWebMessage(logger, webMsg); err != nil {
				logger.Warnf("NoticeAlert.AddWebMessage error:%s", err.Error())
			}
		}

	}
	if util.InArray(int64(2), noticeWays) { //2表示邮件

		//发送邮件通知
		obj := &mailLogic.InbondAlertMailMessage{
			MessageType:    messageLogic.Types2ShowZh[messageType.MessageTypeInBondAlert],
			MessageSubType: messageLogic.Types2ShowZh[rt.Metric],
			ResourceType:   "instance",
			AlertTime:      time.Unix(int64(alertEntity.AlertTime), 0).Format("2006-01-02 15:04:05"),
			AlertPeriod:    alertEntity.AlertPeriod,
			Trigger:        alertR.TriggerDescription,
			AlertValue:     alertEntity.AlertValue + alertEntity.CalculationUnit,
			SN:             deviceEntity.Sn,
			UserID:         ruleEntity.UserID,
			UserName:       ruleEntity.UserName,
			InstanceID:     instanceEntity.InstanceID,
			InstanceName:   instanceEntity.InstanceName,
			IsRecover:      int(alertEntity.IsRecover),
			RuleName:       ruleEntity.RuleName,
			RuleID:         ruleEntity.RuleID,
		}
		subject := fmt.Sprintf("%s-%s DEBUG", messageLogic.Types2ShowZh[messageType.MessageTypeInBondAlert], messageLogic.Types2ShowZh[rt.Metric])
		if obj.IsRecover == 1 { //0表示告警，1表示恢复
			subject = subject + "-恢复"
		}
		logger.Infof("addAlert.DEBUG, subject:%s, email:%s", util.InterfaceToJson(userEntity), userEntity.Email)
		err = mailLogic.SendMailByTpl(logger, mailLogic.MessageInBondAlertTemplate, subject, obj, userEntity.Email)
		if err != nil {
			logger.Warnf("addAlert.SendMail error:%s", err.Error())
		} else {
			logger.Infof("addAlert.SendMail success, receiver:%s", userEntity.Email)
		}
	}

	return nil
}

func DescribeAlert(logger *log.Logger, AlertId string) (*response.Alert, error) {
	alertEntity, err := monitorAlertDao.GetMonitorAlertsById(logger, AlertId)
	if err != nil {
		logger.Warnf("DescribeAlert.GetMonitorAlertsById error, AlertId:%s, error:%s", AlertId, err.Error())
		return nil, err
	}
	res := AlertEntity2Alert(logger, alertEntity, true)
	return res, nil

}

func DescribeAlerts(logger *log.Logger, param *request.DescribeAlertsRequest, p util.Pageable) (*response.AlertList, error) {
	userId := logger.GetPoint("userId").(string)
	offset, limit := p.Pageable2offset()
	query := map[string]interface{}{
		"is_del": 0,
	}
	if param.UserID != "" {
		query["user_id"] = param.UserID
	} else {
		query["user_id"] = userId
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
	if param.ProjectID != "" {
		query["project_id"] = param.ProjectID
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
