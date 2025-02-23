package monitorRuleLogic

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"coding.jd.com/aidc-bmp/bmp-openapi/constant"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/instanceDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/monitorRuleDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/projectDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/rMonitorRulesInstanceDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/dao/userDao"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/baseLogic"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/instanceLogic"
	"coding.jd.com/aidc-bmp/bmp-openapi/logic/messageLogic"
	"coding.jd.com/aidc-bmp/bmp-openapi/types/request"
	"coding.jd.com/aidc-bmp/bmp-openapi/types/response"
	"coding.jd.com/aidc-bmp/bmp-openapi/util"
	log "coding.jd.com/aidc-bmp/bmp_log"
	beego "github.com/beego/beego/v2/server/web"
)

var monitorRuleSendUrl string
var monitorRuleDelUrl string
var monitorRuleListUrl string

func Init() error {
	bmp_pronoea_host, err := beego.AppConfig.String("bmp_pronoea_host")
	if err != nil {
		return err
	}
	bmp_pronoea_port, _ := beego.AppConfig.String("bmp_pronoea_port")
	if err != nil {
		return err
	}
	monitorRuleSendUrl = fmt.Sprintf("http://%s:%s/api/rules/write", bmp_pronoea_host, bmp_pronoea_port)
	monitorRuleDelUrl = fmt.Sprintf("http://%s:%s/api/rules/delete", bmp_pronoea_host, bmp_pronoea_port)
	monitorRuleListUrl = fmt.Sprintf("http://%s:%s/api/rules/list", bmp_pronoea_host, bmp_pronoea_port)
	return nil
}

var MapMetric2Table map[string]string = map[string]string{
	"bmp.disk.bytes.read":         "bmp_monitor_counter",
	"bmp.disk.bytes.write":        "bmp_monitor_counter",
	"bmp.disk.counts.read":        "bmp_monitor_counter",
	"bmp.disk.counts.write":       "bmp_monitor_counter",
	"bmp.network.bytes.ingress":   "bmp_monitor_counter",
	"bmp.network.bytes.egress":    "bmp_monitor_counter",
	"bmp.network.packets.ingress": "bmp_monitor_counter",
	"bmp.network.packets.egress":  "bmp_monitor_counter",
}

var DimensionMapZh map[string]string = map[string]string{
	"disk":       "盘符",
	"mountpoint": "挂载点",
	"nic":        "网卡",
}

var RuleStatusMapZh map[int64]string = map[int64]string{
	1: "正常",
	2: "已禁用",
	3: "报警",
}

var RuleStatusMapEn map[int64]string = map[int64]string{
	1: "Normal",
	2: "Disabled",
	3: "Alert",
}

// 1表示一般，2表示严重，3表示紧急
var RuleAlertLevelMapZh map[int64]string = map[int64]string{
	1: "一般",
	2: "严重",
	3: "紧急",
}

var RuleAlertLevelMapEn map[int64]string = map[int64]string{
	1: "General",
	2: "Serious",
	3: "Emergency",
}

var RuleTriggerCalculationEn map[string]string = map[string]string{
	"min": "Min",
	"max": "Max",
	"avg": "Avg",
	"sum": "Sum",
}

var RuleTriggerCalculationZh map[string]string = map[string]string{
	"min": "最小值",
	"max": "最大值",
	"avg": "平均值",
	"sum": "总和",
}

var RuleTriggerOperation map[string]string = map[string]string{
	"gt":  ">",
	"gte": ">=",
	"lt":  "<",
	"lte": "<=",
	"eq":  "==",
	"neq": "!=",
}

var ResourceMapZh map[string]string = map[string]string{
	"instance": "实例",
}

type DeleteRuleToPronoeaRequest struct {
	RuleIDs string `json:"ruleIds"`
	Source  string `json:"source"`
}

type SendRuleToPronoeaRequest struct {
	Source string           `json:"source"`
	Rules  []*response.Rule `json:"rules"`
}

func Trigger2Description(logger *log.Logger, trigger *response.RuleTrigger) string {

	var tpl string //统计周期5分钟 CPU使用率 平均值 >= 1% 持续1个周期
	var metricName string
	var staticName string
	var operationName string
	if logger.GetPoint("language").(string) != baseLogic.EN_US {
		tpl = "统计周期 %d分钟 %s %s %s %s%s 持续 %d个周期"
		metricName = messageLogic.Types2ShowZh[trigger.Metric]
		staticName = RuleTriggerCalculationZh[trigger.Calculation]
	} else {
		tpl = "static period %dminutes %s %s %s %s%s continue %dperiods"
		metricName = messageLogic.Types2ShowEn[trigger.Metric]
		staticName = RuleTriggerCalculationEn[trigger.Calculation]
	}
	operationName = RuleTriggerOperation[trigger.Operation]
	if operationName == "" {
		operationName = trigger.Operation
	}
	var threshold string
	if float64(int(trigger.Threshold)) == trigger.Threshold { //直接取整
		threshold = strconv.Itoa(int(trigger.Threshold))
	} else {
		threshold = strconv.FormatFloat(trigger.Threshold, 'f', 2, 64)
	}

	return fmt.Sprintf(tpl, trigger.Period, metricName, staticName, operationName, threshold, trigger.CalculationUnit, trigger.Times)
}

func RuleEntity2Rule(logger *log.Logger, entity *monitorRuleDao.MonitorRules, idDetail bool) *response.Rule {

	res := &response.Rule{
		RuleID:        entity.RuleID,
		RuleName:      entity.RuleName,
		Dimension:     entity.Dimension,
		DimensionName: entity.Dimension,
		Resource:      entity.Resource,
		ResourceName:  entity.Resource,
		Status:        int64(entity.Status),
	}
	if logger.GetPoint("language").(string) != baseLogic.EN_US {
		res.DimensionName = DimensionMapZh[entity.Dimension]
		res.ResourceName = ResourceMapZh[entity.Resource]
	}

	trOption := []*response.RuleTrigger{}
	if err := json.Unmarshal([]byte(entity.TriggerOption), &trOption); err != nil {
		logger.Warnf("RuleEntity2Rule.Unmarshal error, rule_id:%s, error:%s", entity.RuleID, err.Error())
	} else {
		res.TriggerOption = trOption
	}
	triggerDescriptions := []string{}
	for idx, v := range trOption {
		item := Trigger2Description(logger, v)
		triggerDescriptions = append(triggerDescriptions, item)
		trOption[idx].Description = item
		trOption[idx].TableName = MapMetric2Table[v.Metric]
		if trOption[idx].TableName == "" {
			trOption[idx].TableName = "bmp_monitor_gauge"
		}
	}
	res.TriggerDescription = triggerDescriptions

	notifyOption := &response.RuleNotice{}
	if err := json.Unmarshal([]byte(entity.NoticeOption), notifyOption); err != nil {
		logger.Warnf("RuleEntity2Rule.Unmarshal error, rule_id:%s, error:%s", entity.RuleID, err.Error())
	} else {
		res.NoticeOption = notifyOption
	}

	q := map[string]interface{}{
		"is_del":  0,
		"rule_id": entity.RuleID,
	}
	rRuleInstanceEntities, err := rMonitorRulesInstanceDao.GetAllRMonitorRulesInstance(logger, q)
	if err != nil {
		logger.Warnf("DescribeRule.GetAllRMonitorRulesInstance error, ruleId:%s, error:%s", entity.RuleID, err.Error())
	}
	instanceIds := []string{}
	for _, v := range rRuleInstanceEntities {
		instanceIds = append(instanceIds, v.InstanceID)
	}
	res.InstanceIds = instanceIds
	res.InstanceCount = int64(len(instanceIds))

	if idDetail {
		iq := map[string]interface{}{
			"is_del":         0,
			"instance_id.in": instanceIds,
		}
		instanceEntities, err := instanceDao.GetAllInstance(logger, iq)
		if err != nil {
			logger.Warnf("RuleEntity2Rule.GetAllInstance error:%s", err.Error())
		} else {
			instances := []*response.Instance{}
			for _, v := range instanceEntities {
				instances = append(instances, instanceLogic.InstanceEntity2Instance(logger, v, nil))
			}
			res.Instances = instances
		}

		if res.NoticeOption != nil {
			noticeUserId := res.NoticeOption.UserID
			userEntity, err := userDao.GetUserById(logger, noticeUserId)
			if err != nil {
				logger.Warnf("RuleEntity2Rule.GetUserById error, userid:%s, error:%s", noticeUserId, err.Error())
			}
			res.NoticeOption.UserName = userEntity.UserName
		}

	}

	relatedResourceCount := int64(len(rRuleInstanceEntities))
	res.RelatedResourceCount = relatedResourceCount
	if relatedResourceCount > 0 {
		res.DeviceTag = rRuleInstanceEntities[0].Tags
	}

	if logger.GetPoint("language").(string) == baseLogic.EN_US {
		res.StatusName = RuleStatusMapEn[res.Status]

	} else {
		res.StatusName = RuleStatusMapZh[res.Status]
	}

	return res

}

func AddRule(logger *log.Logger, param *request.AddRuleRequest) (bool, error) {
	instanceIds := param.InstanceIds

	q := map[string]interface{}{
		"is_del":         0,
		"instance_id.in": instanceIds,
	}
	instanceEntities, err := instanceDao.GetAllInstance(logger, q)
	if err != nil {
		logger.Warnf("AddRule.GetAllInstance error, instance_ids:%v, error:%s", instanceIds, err.Error())
		return false, err
	}
	if len(instanceEntities) < len(instanceIds) {
		logger.Warnf("AddRule.GetAllInstance number lt, instance_ids:%v", instanceIds)
		return false, errors.New("instanceIds param error")
	}
	entity := &monitorRuleDao.MonitorRules{
		RuleID:    util.GetUuid("mr_"),
		Status:    1,
		RuleName:  param.RuleName,
		Dimension: param.Dimension,
		Resource:  param.Resource,
		ProjectID: param.ProjectID,
		// TriggerOption: ,
		// NoticeOption: ,
		CreatedTime: int(time.Now().Unix()),
		UpdatedTime: int(time.Now().Unix()),
		UserID:      logger.GetPoint("userId").(string),
		UserName:    logger.GetPoint("username").(string),
	}
	v, _ := json.Marshal(param.TriggerOption)
	entity.TriggerOption = string(v)
	v1, _ := json.Marshal(param.NoticeOption)
	entity.NoticeOption = string(v1)
	if _, err := monitorRuleDao.AddMonitorRules(logger, entity); err != nil {
		logger.Warnf("AddRule.dao.AddMonitorRules error:%s", err.Error())
		return false, err
	} else {
		rEntities := []*rMonitorRulesInstanceDao.RMonitorRulesInstance{}
		for _, instanceEntity := range instanceEntities {
			rEntity := &rMonitorRulesInstanceDao.RMonitorRulesInstance{
				RuleID:       entity.RuleID,
				RuleName:     entity.RuleName,
				InstanceID:   instanceEntity.InstanceID,
				InstanceName: instanceEntity.InstanceName,
				Tags:         param.DeviceTag,
				CreatedTime:  int(time.Now().Unix()),
				UpdatedTime:  int(time.Now().Unix()),
				ProjectID:    param.ProjectID,
			}
			rEntities = append(rEntities, rEntity)
		}
		if len(rEntities) > 0 {
			if err := rMonitorRulesInstanceDao.AddMultiRMonitorRulesInstance(logger, rEntities); err != nil {
				logger.Warnf("AddRule.AddMultiRMonitorRulesInstance error:%s", err.Error())
				return false, err
			}
		}
	}

	rule := RuleEntity2Rule(logger, entity, false)
	if err := sendRuleToPronoea(logger, rule); err != nil {
		logger.Warnf("AddRule.sendRuleToPronoea error, rule_id:%s, error:%s", entity.RuleID, err.Error())
		return false, err
	}
	return true, nil

}

func DescribeRule(logger *log.Logger, ruleId string) (*response.Rule, error) {
	ruleEntity, err := monitorRuleDao.GetMonitorRuleById(logger, ruleId)
	if err != nil {
		logger.Warnf("DescribeRule.GetMonitorRuleById error, ruleId:%s, error:%s", ruleId, err.Error())
		return nil, err
	}
	res := RuleEntity2Rule(logger, ruleEntity, true)
	return res, nil

}

func DescribeRules(logger *log.Logger, param *request.DescribeRulesRequest, p util.Pageable) (*response.RuleList, error) {

	if param.ProjectID != "" {
		projectEntity, err := projectDao.GetProjectById(logger, param.ProjectID)
		if err != nil {
			logger.Warnf("DescribeRules.GetProjectById error, projectId:%s, error:%s", param.ProjectID, err.Error())
			return nil, err
		}
		if projectEntity.UserID != logger.GetPoint("userId").(string) {
			logger.Warnf("DescribeRules.project.userid not equal traffic.userId, pass")
			return nil, nil
		}
	}

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
	if param.RuleName != "" {
		query["rule_name.like"] = "%" + param.RuleName + "%"
	}
	if param.Status != 0 {
		query["status"] = param.Status
	}
	if param.ProjectID != "" {
		query["project_id"] = param.ProjectID
	}
	count, err := monitorRuleDao.GetMonitorRulesCount(logger, query)
	if err != nil {
		logger.Warnf("GetSshkeyCount dao error, query:%s, error:%s", util.ObjToJson(query), err.Error())
		return nil, err
	}
	entityList := []*monitorRuleDao.MonitorRules{}
	if param.IsAll == baseLogic.IS_ALL {
		entityList, err = monitorRuleDao.GetMultiMonitorRules(logger, query, []string{}, []string{"id"}, []string{"desc"}, 0, 0)
	} else {
		entityList, err = monitorRuleDao.GetMultiMonitorRules(logger, query, []string{}, []string{"id"}, []string{"desc"}, offset, limit)
	}
	if err != nil {
		logger.Warn("DescribeRules.GetMultiMonitorRules sql error:", err.Error())
		return nil, err
	}
	items := []*response.Rule{}
	for _, ruleEntity := range entityList {
		item := RuleEntity2Rule(logger, ruleEntity, false)
		items = append(items, item)
	}
	resp := &response.RuleList{
		Rules:      items,
		PageNumber: p.PageNumber,
		PageSize:   p.PageSize,
		TotalCount: count,
	}

	return resp, nil
}

func EditRule(logger *log.Logger, param *request.EditRuleRequest) (bool, error) {

	ruleId := param.RuleId
	entity, err := monitorRuleDao.GetMonitorRuleById(logger, ruleId)
	if err != nil {
		logger.Warnf("EditRule.GetMonitorRuleById error, ruleId:%s, error:%s", ruleId, err.Error())
		return false, err
	}

	instanceIds := param.InstanceIds

	q := map[string]interface{}{
		"is_del":         0,
		"instance_id.in": instanceIds,
	}
	instanceEntities, err := instanceDao.GetAllInstance(logger, q)
	if err != nil {
		logger.Warnf("EditRule.GetAllInstance error, instance_ids:%v, error:%s", instanceIds, err.Error())
		return false, err
	}
	if len(instanceEntities) < len(instanceIds) {
		logger.Warnf("EditRule.GetAllInstance number lt, instance_ids:%v", instanceIds)
		return false, errors.New("instanceIds param error")
	}

	// entity.Status = 1

	// alertEntities, _ := monitorAlertDao.GetAllMonitorAlertsByRuleId(logger, param.RuleId)
	// if len(alertEntities) > 0 {
	// 	entity.Status = 3 //已告警
	// }

	entity.RuleName = param.RuleName
	entity.Dimension = param.Dimension
	entity.Resource = param.Resource
	entity.UpdatedTime = int(time.Now().Unix())
	entity.UserID = logger.GetPoint("userId").(string)
	v, _ := json.Marshal(param.TriggerOption)
	entity.TriggerOption = string(v)
	v1, _ := json.Marshal(param.NoticeOption)
	entity.NoticeOption = string(v1)
	if err := monitorRuleDao.UpdateMonitorRules(logger, entity); err != nil {
		logger.Warnf("EditRule.dao.UpdateMonitorRules error:%s", err.Error())
		return false, err
	} else {
		q := map[string]interface{}{
			"is_del":  0,
			"rule_id": param.RuleId,
		}
		u := map[string]interface{}{
			"is_del": 1,
		}
		if err := rMonitorRulesInstanceDao.UpdateMultiRMonitorRulesInstance(logger, q, u); err != nil {
			logger.Warnf("EditRule.UpdateMultiRMonitorRulesInstance error, rule_id:%s, error:%s", param.RuleId, err.Error())
		}

		rEntities := []*rMonitorRulesInstanceDao.RMonitorRulesInstance{}
		for _, instanceEntity := range instanceEntities {
			rEntity := &rMonitorRulesInstanceDao.RMonitorRulesInstance{
				RuleID:       entity.RuleID,
				RuleName:     entity.RuleName,
				InstanceID:   instanceEntity.InstanceID,
				InstanceName: instanceEntity.InstanceName,
				Tags:         param.DeviceTag,
				CreatedTime:  int(time.Now().Unix()),
				UpdatedTime:  int(time.Now().Unix()),
			}
			rEntities = append(rEntities, rEntity)
		}
		if len(rEntities) > 0 {
			if err := rMonitorRulesInstanceDao.AddMultiRMonitorRulesInstance(logger, rEntities); err != nil {
				logger.Warnf("EditRule.AddMultiRMonitorRulesInstance error:%s", err.Error())
				return false, err
			}
		}
	}
	rule := RuleEntity2Rule(logger, entity, true)
	if err := sendRuleToPronoea(logger, rule); err != nil {
		logger.Warnf("EditRule.sendRuleToPronoea error, rule_id:%s, error:%s", param.RuleId, err.Error())
		return false, err
	}
	return true, nil

}

func EnableRule(logger *log.Logger, param *request.EnableRuleRequest) (bool, error) {

	ruleId := param.RuleId
	entity, err := monitorRuleDao.GetMonitorRuleById(logger, ruleId)
	if err != nil {
		logger.Warnf("EnableRule.GetMonitorRuleById error, ruleId:%s, error:%s", ruleId, err.Error())
		return false, err
	}
	if entity.Status != 2 {
		logger.Warnf("EnableRule.rule enablestatus now %d not allowed enable", entity.Status)
		panic(constant.NOT_SUPPORTED)
	}
	entity.Status = 1 //如果之前已经告警了呢？
	entity.UpdatedTime = int(time.Now().Unix())
	if err := monitorRuleDao.UpdateMonitorRules(logger, entity); err != nil {
		logger.Warnf("EnableRule.UpdateMonitorRules error, ruleId:%s, error:%s", ruleId, err.Error())
		return false, err
	}
	rule := RuleEntity2Rule(logger, entity, true)
	if err := sendRuleToPronoea(logger, rule); err != nil {
		logger.Warnf("EnableRule.sendRuleToPronoea error, rule_id:%s, error:%s", param.RuleId, err.Error())
		return false, err
	}
	return true, nil

}

func DisableRule(logger *log.Logger, param *request.DisableRuleRequest) (bool, error) {

	ruleId := param.RuleId
	entity, err := monitorRuleDao.GetMonitorRuleById(logger, ruleId)
	if err != nil {
		logger.Warnf("DisableRule.GetMonitorRuleById error, ruleId:%s, error:%s", ruleId, err.Error())
		return false, err
	}
	if entity.Status == 2 { //已经是禁用状态
		logger.Warnf("DisableRule.rule enablestatus now %d not allowed enable", entity.Status)
		panic(constant.NOT_SUPPORTED)
	}
	entity.Status = 2
	entity.UpdatedTime = int(time.Now().Unix())
	if err := monitorRuleDao.UpdateMonitorRules(logger, entity); err != nil {
		logger.Warnf("DisableRule.UpdateMonitorRules error, ruleId:%s, error:%s", ruleId, err.Error())
		return false, err
	}
	if err := DelRuleToPronoea(logger, param.RuleId); err != nil {
		logger.Warnf("DisableRule.DelRuleToPronoea error, rule_id:%s, error:%s", param.RuleId, err.Error())
		return false, err
	}
	return true, nil

}

func DeleteRule(logger *log.Logger, param *request.DeleteRuleRequest) (bool, error) {

	ruleId := param.RuleId
	entity, err := monitorRuleDao.GetMonitorRuleById(logger, ruleId)
	if err != nil {
		logger.Warnf("DeleteRule.GetMonitorRuleById error, ruleId:%s, error:%s", ruleId, err.Error())
		return false, err
	}
	if entity.Status != 2 {
		panic(constant.BuildInvalidArgumentWithArgs("禁用状态的规则才允许删除", "enabled rule not allow delete"))
	}

	entity.IsDel = 1
	if err := monitorRuleDao.UpdateMonitorRules(logger, entity); err != nil {
		logger.Warnf("DeleteRule.UpdateMonitorRules error, ruleId:%s, error:%s", ruleId, err.Error())
		return false, err
	}
	// if err := DelRuleToPronoea(logger, param.RuleId); err != nil {
	// 	logger.Warnf("DeleteRule.DelRuleToPronoea error, rule_id:%s, error:%s", param.RuleId, err.Error())
	// 	return false, err
	// }
	return true, nil

}

func sendRuleToPronoea(logger *log.Logger, rule *response.Rule) error {
	rules := []*response.Rule{rule}
	return sendRulesToPronoea(logger, rules)

}

func sendRulesToPronoea(logger *log.Logger, rules []*response.Rule) error {
	url := monitorRuleSendUrl
	req := SendRuleToPronoeaRequest{
		Source: "bmpInBand",
		Rules:  rules,
	}
	logger.Infof("sendRuleToPronoea.request:%s", util.InterfaceToJson(req))
	header := map[string]string{
		"Traceid": logger.GetPoint("logid").(string),
	}
	resp, err := util.DoHttpPost(logger, url, header, req)
	if err != nil {
		logger.Warnf("sendRuleToPronoea post error, query:%s, error:%s", util.InterfaceToJson(rules), err.Error())
		return err
	}
	logger.Infof("sendRuleToPronoea resp:%s", string(resp))
	return nil

}

func DelRuleToPronoea(logger *log.Logger, ruleId string) error {
	url := monitorRuleDelUrl
	req := DeleteRuleToPronoeaRequest{
		RuleIDs: ruleId,
		Source:  "bmpInBand",
	}
	logger.Infof("DelRuleToPronoea.request:%s", util.InterfaceToJson(req))
	header := map[string]string{
		"Traceid": logger.GetPoint("logid").(string),
	}
	resp, err := util.DoHttpPost(logger, url, header, req)
	if err != nil {
		logger.Warnf("DelRuleToPronoea post error, rule_id:%s, error:%s", ruleId, err.Error())
		return err
	}
	logger.Infof("DelRuleToPronoea resp:%s", string(resp))
	return nil

}

func DelMonitorRulesByProjectID(logger *log.Logger, projectId string) error {
	q := map[string]interface{}{
		"is_del":     0,
		"project_id": projectId,
	}
	ruleEntities, err := monitorRuleDao.GetAllMonitorRules(logger, q)
	if err != nil {
		logger.Warnf("DelMonitorRulesByProjectID.GetAllMonitorRules error, query:%s, error:%s", util.InterfaceToJson(q), err.Error())
		return err
	}
	if len(ruleEntities) == 0 {
		logger.Infof("DelMonitorRulesByProjectID rule empty, projectId:%s", projectId)
		return nil
	}
	ruleIds := []string{}
	for _, ruleEntity := range ruleEntities {
		ruleIds = append(ruleIds, ruleEntity.RuleID)
	}
	q1 := map[string]interface{}{
		"is_del":     0,
		"rule_id.in": ruleIds,
	}
	u := map[string]interface{}{
		"is_del": 1,
	}
	if err := monitorRuleDao.UpdateMultiMonitorRules(logger, q1, u); err != nil {
		logger.Warnf("DelMonitorRulesByProjectID.UpdateMultiMonitorRules error, query:%s, error:%s", util.InterfaceToJson(q1), err.Error())
		return err
	}

	//TODO 下发配置到proea
	ruleidStr := strings.TrimRight(strings.Join(ruleIds, ","), ",")
	if err := DelRuleToPronoea(logger, ruleidStr); err != nil {
		logger.Warnf("DelMonitorRulesByProjectID.DelRuleToPronoea error, ruleIds:%s, err:%s", ruleidStr, err.Error())
	}
	return nil
}

//先删除规则，再重新下发规则
func DelRMonitorRuleInstances(logger *log.Logger, projectId string, instanceIds string) error {

	q1 := map[string]interface{}{
		"is_del":         0,
		"project_id":     projectId,
		"instance_id.in": strings.Split(instanceIds, ","),
	}
	rRuleInstanceEntities, err := rMonitorRulesInstanceDao.GetAllRMonitorRulesInstance(logger, q1)
	if err != nil {
		logger.Warnf("DelRMonitorRuleInstances.GetAllRMonitorRulesInstance error, query:%s, error:%s", util.InterfaceToJson(q1), err.Error())
	}

	ri := map[string][]string{}
	for _, rRuleInstanceEntity := range rRuleInstanceEntities {
		_, ok := ri[rRuleInstanceEntity.RuleID]
		if ok {
			ri[rRuleInstanceEntity.RuleID] = append(ri[rRuleInstanceEntity.RuleID], rRuleInstanceEntity.InstanceID)
		} else {
			ri[rRuleInstanceEntity.RuleID] = []string{rRuleInstanceEntity.InstanceID}
		}
	}

	ruleIds := []string{}
	for k, _ := range ri {
		ruleIds = append(ruleIds, k)
	}
	qAll := map[string]interface{}{
		"is_del":     0,
		"project_id": projectId,
		"rule_id.in": ruleIds,
	}
	rRuleInstanceEntitiesAll, err := rMonitorRulesInstanceDao.GetAllRMonitorRulesInstance(logger, qAll)
	if err != nil {
		logger.Warnf("DelRMonitorRuleInstances.GetAllRMonitorRulesInstance2 error, query:%s, error:%s", util.InterfaceToJson(q1), err.Error())
	}

	//数据库动作
	u := map[string]interface{}{
		"is_del": 1,
	}
	if err := rMonitorRulesInstanceDao.UpdateMultiRMonitorRulesInstance(logger, q1, u); err != nil {
		logger.Warnf("DelRMonitorRuleInstances.UpdateMultiRMonitorRulesInstance error, query:%s, error:%s", util.InterfaceToJson(q1), err.Error())
		return err
	}

	riAll := map[string][]string{}
	for _, rRuleInstanceEntity := range rRuleInstanceEntitiesAll {
		_, ok := riAll[rRuleInstanceEntity.RuleID]
		if ok {
			riAll[rRuleInstanceEntity.RuleID] = append(riAll[rRuleInstanceEntity.RuleID], rRuleInstanceEntity.InstanceID)
		} else {
			riAll[rRuleInstanceEntity.RuleID] = []string{rRuleInstanceEntity.InstanceID}
		}
	}

	remainRuleInstances := map[string]struct{}{}
	for k, v := range riAll {
		v1 := ri[k]
		if len(v1) == len(v) { //删除规则
			continue
		} else { //规则里面还有其他实例，需要删除后重新下发
			remainRuleInstances[k] = struct{}{}
		}
	}

	//下发配置到proea
	ruleidsStr := strings.Join(ruleIds, ",")
	if err := DelRuleToPronoea(logger, ruleidsStr); err != nil {
		logger.Warnf("DelRMonitorRuleInstances.DelRuleToPronoea error, ruleIds:%s, err:%s", ruleidsStr, err.Error())
	}

	toSendRules := []*response.Rule{}
	for k, _ := range remainRuleInstances {
		ruleEntity, err := monitorRuleDao.GetMonitorRuleById(logger, k)

		if err != nil {
			logger.Warnf("DelRMonitorRuleInstances.GetMonitorRuleById error, ruleId:%s, error:%s", k, err.Error())
			continue
		}
		rule := RuleEntity2Rule(logger, ruleEntity, false) //这里会查ruleInstance表，所以上面的删除ruleInstance动作要放上面
		toSendRules = append(toSendRules, rule)
	}
	if len(toSendRules) > 0 {
		if err := sendRulesToPronoea(logger, toSendRules); err != nil {
			logger.Warnf("DelRMonitorRuleInstances.sendRulesToPronoea error, rules:%s, error:%s", util.InterfaceToJson(toSendRules), err.Error())
		}
	}

	return nil
}

type PronoeaRuleListRequest struct {
	Source string `json:"source"`
}

type PronoeaRuleListResponse struct {
	Code      int              `json:"Code"`
	RequestID string           `json:"RequestId"`
	Message   string           `json:"Message"`
	Result    PronoeaRulesList `json:"Result"`
}

type PronoeaRulesList struct {
	Rules []PronoeaRuleListItem `json:"rules"`
}

type PronoeaRuleListItem struct {
	RuleID      string   `json:"ruleId"`
	InstanceIDs []string `json:"instanceIds"`
}

type RuleDiffResponseItem struct {
	PronoeaExist bool `json:"pronoeaExist"`
	DaoExist     bool `json:"daoExist"`
}

func CheckDiffFromPronoea(logger *log.Logger) (map[string]map[string]RuleDiffResponseItem, error) {
	url := monitorRuleListUrl
	req := PronoeaRuleListRequest{
		Source: "bmpInBand",
	}
	logger.Infof("CheckDiffFromPronoea.PronoeaRuleListRequest:%s", util.InterfaceToJson(req))
	header := map[string]string{
		"Traceid": logger.GetPoint("logid").(string),
	}
	resp, err := util.DoHttpGet(logger, url, header, req)
	if err != nil {
		logger.Warnf("sendRuleToPronoea post error:%s", err.Error())
		return nil, err
	}
	logger.Infof("sendRuleToPronoea resp:%s", string(resp))
	pronoeaRes := &PronoeaRuleListResponse{}
	if err := json.Unmarshal(resp, pronoeaRes); err != nil {
		logger.Warnf("CheckDiffFromPronoea.PronoeaRuleListResponse unmarshal error:%s", err.Error())
		return nil, err
	}

	finalRes := make(map[string]map[string]RuleDiffResponseItem)
	for _, rule := range pronoeaRes.Result.Rules {
		_, ok := finalRes[rule.RuleID]
		if !ok {
			finalRes[rule.RuleID] = make(map[string]RuleDiffResponseItem)
		}
		for _, instanceId := range rule.InstanceIDs {
			v, ok := finalRes[rule.RuleID][instanceId]
			if !ok {
				finalRes[rule.RuleID][instanceId] = RuleDiffResponseItem{
					PronoeaExist: true,
				}
			} else {
				v.PronoeaExist = true
				finalRes[rule.RuleID][instanceId] = v
			}
		}
	}

	rMonitorRulesInstanceEntities, err := rMonitorRulesInstanceDao.GetAllRMonitorRulesInstanceEffective(logger)
	if err != nil {
		logger.Warnf("CheckDiffFromPronoea.GetAllRMonitorRulesInstance error:%s", err.Error())
		return nil, err
	}
	for _, rMonitorRulesInstanceEntity := range rMonitorRulesInstanceEntities {
		_, ok := finalRes[rMonitorRulesInstanceEntity.RuleID]
		if !ok {
			finalRes[rMonitorRulesInstanceEntity.RuleID] = make(map[string]RuleDiffResponseItem)
		}
		instanceId := rMonitorRulesInstanceEntity.InstanceId
		v, ok := finalRes[rMonitorRulesInstanceEntity.RuleID][instanceId]
		if !ok {
			finalRes[rMonitorRulesInstanceEntity.RuleID][instanceId] = RuleDiffResponseItem{
				DaoExist: true,
			}
		} else {
			v.DaoExist = true
			finalRes[rMonitorRulesInstanceEntity.RuleID][instanceId] = v
		}
	}
	return finalRes, nil
}
