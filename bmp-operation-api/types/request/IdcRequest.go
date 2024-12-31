package request

// 机房管理 模块请求字段
type CreateIdcRequest struct {
	Address string `json:"address"`
	// ilo password
	IloPassword string `json:"iloPassword"`
	// ilo user
	IloUser string `json:"iloUser"`
	// level
	Level string `json:"level"`
	// name
	Name string `json:"name"`
	// nameEn
	NameEn string `json:"nameEn"`
	// shortname
	Shortname string `json:"shortname"`
	// switch password1
	SwitchPassword1 string `json:"switchPassword1"`
	// switch password2
	SwitchPassword2 string `json:"switchPassword2"`
	// switch user1
	SwitchUser1 string `json:"switchUser1"`
	// switch user2
	SwitchUser2 string `json:"switchUser2"`
}

// 【机房管理】【编辑机房】
type ModifyIdcRequest struct {
	//ID                        int    `gorm:"primaryKey;column:id" json:"-"`                                        // 主键
	IdcID string `json:"idcId"`
	// address
	Address *string `json:"address"`
	// ilo password
	IloPassword *string `json:"iloPassword"`
	// ilo user
	IloUser *string `json:"iloUser"`
	// level
	Level *string `json:"level"`
	// name
	Name *string `json:"name"`
	// name en
	NameEn *string `json:"nameEn"`
	// shortname
	Shortname *string `json:"shortname"`
	// switch password1
	SwitchPassword1 *string `json:"switchPassword1"`
	// switch password2
	SwitchPassword2 *string `json:"switchPassword2"`
	// switch user1
	SwitchUser1 *string `json:"switchUser1"`
	// switch user2
	SwitchUser2 *string `json:"switchUser2"`
}

// 【机房管理】【机房列表】
type QueryIdcListRequest struct {
	BaseRequest
	PagingRequest
	Name  string `json:"name"`
	Level string `json:"level"`
	IsAll string `json:"isAll"`
	Show  string `json:"show"`
	// 非空表示导出
	ExportType string `json:"exportType"`
}

// 【机房管理】【机房详情信息】
type QueryIdcRequest struct {
	Show  string `json:"show"`
	IdcID string `json:"idcId"`
}

// 【机房管理】【删除】
type DeleteIdcRequest struct {
	IdcID string `json:"idcId"`
}
