package response

type Apikey struct {
	ID          int    `json:"id"`
	ApikeyID    string `json:"apiKeyId"`    // 秘钥uuid
	Name        string `json:"name"`        // 秘钥名称
	ReadOnly    int8   `json:"readOnly"`    //是否支持只读，read_only =1 的时候说明这个key是只读key，不能访问写方法。
	Token       string `json:"token"`       // 32位字符令牌
	Type        string `json:"type"`        // Token类型system/user
	UserID      string `json:"userId"`      // 所属用户
	CreatedBy   string `json:"createdBy"`   // 创建者
	UpdatedBy   string `json:"updatedBy"`   // 更新者
	CreatedTime string `json:"createdTime"` // 创建时间戳
	UpdatedTime string `json:"updatedTime"` // 更新时间戳
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
	ApikeyId string `json:"apikeyId"`
}
