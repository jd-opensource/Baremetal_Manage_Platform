package response


type PrometheusRuleListResponse struct {
	Rules []*PrometheusRuleListItem `json:"rules"`
}


type PrometheusRuleListItem struct {
	RuleId      string   `json:"ruleId"`
	InstanceIds []string `json:"instanceIds"`
}
