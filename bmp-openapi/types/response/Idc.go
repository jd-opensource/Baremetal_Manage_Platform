package response

type Idc struct {
	ID int `json:"id"`
	// 机房uuid
	IDcID string `json:"idcId"`
	// 机房名称
	Name string `json:"name"`
	// 机房名称en
	NameEn string `json:"nameEn"`
	// shortname
	Shortname string `json:"shortname"`
	// 机房等级
	Level string `json:"level"`
	// 机房地址
	Address string `json:"address"`
	// 机房公用带外管理user
	IloUser string `json:"iloUser"`
	// 机房公用带外管理password
	IloPassword string `json:"iloPassword"`
	// 交换机用户名1
	SwitchUser1 string `json:"switchUser1"`
	// 交换机密码1
	SwitchPassword1 string `json:"switchPassword1"`
	// 交换机用户名2
	SwitchUser2 string `json:"switchUser2"`
	// 交换机密码2
	SwitchPassword2 string `json:"switchPassword2"`
	// 创建时间
	CreateTime string `json:"createTime"`
	// 修改时间
	UpdateTime string `json:"updateTime"`
	// 创建者
	CreatedBy string `json:"createdBy"`
	// 修改用户
	UpdatedBy string `json:"updatedBy"`
}
type IdcList struct {
	// 机房列表
	Idcs []*Idc `json:"idcs"`
	// 页数
	PageNumber int64 `json:"pageNumber"`
	// 页大小
	PageSize int64 `json:"pageSize"`
	// 总条数
	TotalCount int64 `json:"totalCount"`
}
type IdcInfo struct {
	// 机房信息
	Idc *Idc `json:"idc"`
}
type IdcID struct {
	// 机房id
	IdcID string `json:"idcId"`
}

//func (d Idc) MarshalJSON() ([]byte, error) {
//	type Alias Idc
//	return json.Marshal(struct {
//		Alias
//		CreateTime string `json:"createTime"`
//		UpdateTime string `json:"updateTime"`
//	}{
//		Alias: Alias(d),
//		//CreateTime: time.Unix(d.CreateTime,0).Format("2006-01-02 15:04:05"),
//		//UpdateTime: time.Unix(d.UpdateTime,0).Format("2006-01-02 15:04:05"),
//	})
//}
