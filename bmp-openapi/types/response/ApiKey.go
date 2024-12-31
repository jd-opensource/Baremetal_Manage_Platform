package response

type Apikey struct {
	ID int `json:"id"`
	// apikey uuid
	// required: true
	ApikeyID string `json:"apiKeyId"`
	// 名称
	// required: true
	Name string `json:"name"`
	// 是否支持只读, [0/1], 1表示只读
	// required: true
	ReadOnly int8 `json:"readOnly"`
	// 32位字符令牌，使用token来独立访问openapi
	// required: true
	Token string `json:"token"`
	// apikey的类型，[system/user]
	// required: true
	Type string `json:"type"`
	// 所属用户uuid
	// required: true
	UserID string `json:"userId"`
	// 创建者
	CreatedBy string `json:"createdBy"`
	// 更新者
	UpdatedBy string `json:"updatedBy"`
	// 创建时间
	CreatedTime string `json:"createdTime"`
	// 更新时间
	UpdatedTime string `json:"updatedTime"`
}
type ApikeyList struct {
	Apikeys    []*Apikey `json:"apikeys"`
	PageNumber int64     `json:"pageNumber"`
	PageSize   int64     `json:"pageSize"`
	TotalCount int64     `json:"totalCount"`
}
type ApikeyInfo struct {
	Apikey *Apikey `json:"apikey"`
}
type ApikeyId struct {
	// apikey uuid
	// required: true
	ApikeyId string `json:"apikeyId"`
}
