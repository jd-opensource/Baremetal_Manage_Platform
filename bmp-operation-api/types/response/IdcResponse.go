package response

import "encoding/json"

type Idc struct {
	// address
	Address string `json:"address"`

	// create time
	CreateTime string `json:"createTime"`

	// created by
	CreatedBy string `json:"createdBy"`

	// ID
	ID int64 `json:"id"`

	// i dc ID
	IDcID string `json:"idcId"`

	// ilo password
	IloPassword string `json:"iloPassword"`

	// ilo user
	IloUser string `json:"iloUser"`

	// level
	Level string `json:"level"`

	// name
	Name string `json:"name"`

	// name en
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

	// update time
	UpdateTime string `json:"updateTime"`

	// updated by
	UpdatedBy string `json:"updatedBy"`
}

type IdcPage struct {
	Idcs []Idc `json:"idcs"`
	PagingResponse
}

type IdcInfo struct {
	Idc *Idc `json:"idc"`
}

type CreateIdcResult struct {
	IdcId string `json:"idcId"`
}

func (d Idc) MarshalJSON() ([]byte, error) {
	type Alias Idc
	return json.Marshal(struct {
		Alias
		//CreateTime string `json:"create_time"`
		//UpdateTime string `json:"update_time"`
	}{
		Alias: Alias(d),
		//CreateTime: d.CreateTime.Format("2006-01-02 15:04:05"),
		//UpdateTime: d.UpdateTime.Format("2006-01-02 15:04:05"),
	})
}
