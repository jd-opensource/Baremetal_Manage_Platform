package response

type Raid struct {
	// raid uuid
	RaidID string `json:"raidId"`
	// raid 名称
	Name string `json:"name"`
	//中文描述
	DescriptionEn string `json:"descriptionEn"`
	// description
	DescriptionZh string `json:"descriptionZh"`
}
type RaidList struct {
	//raid实体列表
	Raids []*Raid `json:"raids"`
	//PageNumber int64   `json:"pageNumber"`
	//PageSize   int64   `json:"pageSize"`
	//TotalCount int64   `json:"totalCount"`
}
type RaidInfo struct {
	//raid实体
	Raid *Raid `json:"raid"`
}
type RaidId struct {
	//raid uuid
	RaidId string `json:"raidId"`
}
