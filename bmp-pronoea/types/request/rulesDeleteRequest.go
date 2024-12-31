package request

type RulesDeleteRequest struct {
	BaseRequest
	RuleIds string   `json:"ruleIds" valid:"Required"`     //多个id，用英文逗号分隔
	Source  string   `json:"source" valid:"Required"`
}
