package response

import (
	"encoding/json"
	"time"
)

type Command struct {
	Id              int64     `json:"id"`
	RequestId       string    `json:"request_id"`
	Sn              string    `json:"sn"`
	InstanceId      string    `json:"instance_id"`
	Action          string    `json:"action"`
	Type            string    `json:"type"`
	Status          string    `json:"status"`
	ParentCommandId int64     `json:"parent_command_id"`
	ExecuteCount    int       `json:"execute_count"`
	Result          string    `json:"result,omitempty"`
	CreateTime      time.Time `json:"create_time"`
	UpdateTime      time.Time `json:"update_time"`
}

type QueryCommandsResult struct {
	Commands   []*Command `json:"commands"`
	PageNumber int64      `json:"pageNumber"`
	PageSize   int64      `json:"pageSize"`
	TotalCount int64      `json:"totalCount"`
}

//解决time.Time类型序列化时默认格式为2018-12-25T19:43:35+08:00而不是2006-01-02 15:04:05的问题
func (d Command) MarshalJSON() ([]byte, error) {
	type Alias Command
	return json.Marshal(struct {
		Alias
		CreateTime string `json:"create_time"`
		UpdateTime string `json:"update_time"`
	}{
		Alias:      Alias(d),
		CreateTime: d.CreateTime.Format("2006-01-02 15:04:05"),
		UpdateTime: d.UpdateTime.Format("2006-01-02 15:04:05"),
	})
}
