package bmpRules

import (
	"encoding/json"
	"fmt"
	"git.jd.com/bmp-pronoea/common/util"
	"git.jd.com/bmp-pronoea/config"
	"git.jd.com/bmp-pronoea/models"
	"git.jd.com/bmp-pronoea/models/rules"
	"git.jd.com/bmp-pronoea/types/request"
	"git.jd.com/bmp-pronoea/types/response"
	"github.com/astaxie/beego/logs"
	"gopkg.in/yaml.v2"
	"reflect"
	"strings"
)

type BmpInBandMonitorRules struct {
	rules.BaseRule
}

// DeleteRule 删除规则
func (b *BmpInBandMonitorRules) DeleteRules (requestId, ruleIds, source string) error {
	return b.BaseRule.DeleteRules(requestId, ruleIds, source)
}

func (b *BmpInBandMonitorRules) WriteRule(requestId, ruleContentJson string) error {
	err := b.BaseRule.WriteRule(requestId, ruleContentJson)
	if err != nil {
		return err
	}

	bmpRuleInfo := &request.RulesRequest{}
	err = json.Unmarshal([]byte(ruleContentJson), &bmpRuleInfo)
	if err != nil {
		return err
	}

	for _, ruleTmp := range bmpRuleInfo.Rules {
		prometheusRule, err := b.createRuleContent(bmpRuleInfo.Source, ruleTmp)
		if err != nil {
			return err
		}
		filePath := fmt.Sprintf("%v/%v_%v.yml",
			                   config.SysConfig.PrometheusRulePath,
			                   bmpRuleInfo.Source,
			                   ruleTmp.RuleId)

		ruleByte, err := yaml.Marshal(prometheusRule)
		if err != nil {
			return fmt.Errorf("yaml marshal error %v", err)
		}

		// 同时写多个规则，写入时不重启，最后统一重新启动
		err = b.BaseRule.WriteRuleAndReload(requestId, filePath, ruleByte, false)
		if err != nil {
			logs.Info(requestId, fmt.Sprintf("write bmp rule error %v, %v", err, filePath))
			return err
		}
	}

	err = b.ReloadPromethus(requestId)
	if err != nil {
		logs.Info(requestId, fmt.Sprintf("write bmp rule success, but reload prometheus error %v, %v", err, ruleContentJson))
		return err
	}

	return nil
}

func (b *BmpInBandMonitorRules) createRuleContent(source string, req *request.RulesDetail) (*PrometheusRule, error) {
	if err := req.Validate(req); err != nil {
		return nil, err
	}

	prometheusRuleGroupDetail := new(PrometheusRuleGroupDetail)
	ruleName := fmt.Sprintf("%v_%v", source, req.RuleId)
	prometheusRuleGroupDetail.Name = ruleName

	for _, instanceId := range req.InstanceIds {
		if util.IsEmpty(instanceId) {
			continue
		}
		for _, tigger := range req.TriggerOption {
			//source_alarmId_instanceId_metricname_device
			alertName := fmt.Sprintf(ALERT_NAME_TEMPLATE,
				source, req.RuleId, instanceId, tigger.Metric, req.DeviceTag)

			if _, exist := AlertExprMethodMap[tigger.Calculation]; !exist {
				continue
			}
			if _, exist := AlertExprOperatorMap[tigger.Operation]; !exist {
				continue
			}

			var expr string
			if util.IsEmpty(req.DeviceTag) {
				expr = fmt.Sprintf(ALERT_EXPR_DEVICE_WITHOUT, tigger.TableName, instanceId, tigger.Metric)
			} else {
				expr = fmt.Sprintf(ALERT_EXPR_DEVICE, tigger.TableName, instanceId, tigger.Metric, req.DeviceTag)
			}

			//基础计算法
			if tmp, exist := SampleDefaultMethodMap[tigger.Metric]; exist {
				expr = fmt.Sprintf(tmp, expr, fmt.Sprintf("%vm", tigger.Period + 1))
			}

			expr = fmt.Sprintf(AlertExprMethodMap[tigger.Calculation], expr)

			threshold := tigger.Threshold
			if unitConversion, exist := models.MetricNameUnitConversionMap[tigger.Metric]; exist {
				threshold = threshold * unitConversion
			}

			expr = fmt.Sprintf("%v%v%v", expr, AlertExprOperatorMap[tigger.Operation], threshold)

			//持续时长
			forTimeLong := fmt.Sprintf("%vm", tigger.Times * tigger.Period)

			labels := make(map[string]interface{}, 0)
			labels[CUSTOM_LABELS_NOTICE_PERIOD] = NOTICE_PERIOD_60_LABEL    // 默认通知周期
			if tmpNoticePeriodLabel, exist := NoticePeriodMap[req.NoticeOption.NoticePeriod]; exist {
				labels[CUSTOM_LABELS_NOTICE_PERIOD] = tmpNoticePeriodLabel
			}
			labels[CUSTOM_BMPLABELS_INSTANCEID] = instanceId
			labels[CUSTOM_BMPLABELS_RULEID]     = req.RuleId

			if tigger != nil {
				tTigger := reflect.TypeOf(*tigger)
				vTigger := reflect.ValueOf(*tigger)
				for i := 1; i < tTigger.NumField(); i++ {
					field := tTigger.Field(i)
					value := vTigger.Field(i)
					labels[util.HeadToLower(field.Name)] = value.Interface()
					/*
					switch value.Kind().String() {
					case "int":
						labels[field.Name] = fmt.Sprintf("%v", value.Int())
						break
					case "float64":
						labels[field.Name] = fmt.Sprintf("%v", value.Float())
						break
					default:
						labels[field.Name] = value.String()
						break
					}
					 */
				}
			}

			tmpRule := &PrometheusRuleDetail{
				Alert  : alertName,
				Expr   : expr,
				For    : forTimeLong,
				Labels : labels,
			}
			tmpRule.Annotations.CurrentValue = "{{ $value }}"
			prometheusRuleGroupDetail.Rules = append(prometheusRuleGroupDetail.Rules, tmpRule)
		}
	}

	prometheusRule := &PrometheusRule{
		Groups: make([]*PrometheusRuleGroupDetail, 0),
	}
	prometheusRule.Groups = append(prometheusRule.Groups, prometheusRuleGroupDetail)

	return prometheusRule, nil
}

func (b *BmpInBandMonitorRules) RulesList(requestId, source string) (interface{}, error){
	rulesFileList, err := b.BaseRule.GetRulesFileList(requestId, source)
	if err != nil {
		logs.Info(requestId, fmt.Sprintf("get rule file list error %v %v", source, err))
		return nil, err
	}

	ruleList := &response.PrometheusRuleListResponse{}
	ruleList.Rules = make([]*response.PrometheusRuleListItem, 0)

	if rulesFileList == nil || len(rulesFileList) <= 0 {
		return ruleList, nil
	}

	for ruleFilePath, ruleByte := range rulesFileList {
		ruleYmlObj := &PrometheusRule{}
		err = yaml.Unmarshal(ruleByte, ruleYmlObj)
		if err != nil {
			logs.Info(requestId,fmt.Sprintf("Invalid rule %v, %v", ruleFilePath, err))
		}

		for _, detail := range ruleYmlObj.Groups {
			if util.IsEmpty(detail.Name) || detail.Rules == nil || len(detail.Rules) <= 0 {
				continue
			}
			ruleId := strings.Replace(detail.Name, fmt.Sprintf("%v_", source), "", 1)
			instanceIds := make([]string, 0)
			for _, tmpRule := range detail.Rules {
				instanceId := util.TurnInterfaceToString(tmpRule.Labels[CUSTOM_BMPLABELS_INSTANCEID])
				instanceIds = append(instanceIds, instanceId)
			}

			if instanceIds == nil || len(instanceIds) <= 0 {
				continue
			}

			ruleListDetail := &response.PrometheusRuleListItem{
				RuleId: ruleId,
				InstanceIds : instanceIds,
			}

			ruleList.Rules = append(ruleList.Rules, ruleListDetail)
		}
	}

	return ruleList, nil
}