package response

import (
	"encoding/json"
)

type AuditLog struct {
	ID int `json:"id"`
	// 日志uuid
	LogID string `json:"logid"`
	// 操作
	Operation string `json:"operation"`
	// 操作名称
	OperationName string `json:"operationName"`
	// 操作人
	UserName string `json:"userName"`
	// 操作人
	UserID string `json:"userID"`
	//sn
	Sn string `json:"sn"`
	//device_id
	DeviceID string `json:"deviceId"`
	//instance_id
	InstanceID string `json:"instanceId"`
	//结果 success/fail
	Result string `json:"result"`
	//失败原因
	FailReason string `json:"failReason"` // reason for fail
	// 操作时间
	OperateTime int `json:"operateTime"`
	// 完成时间
	FinishTime int `json:"finishTime"`
}

type AuditLogList struct {
	// 机型列表
	AuditLogs  []*AuditLog `json:"auditLogs"`
	PageNumber int64       `json:"pageNumber"`
	PageSize   int64       `json:"pageSize"`
	TotalCount int64       `json:"totalCount"`
}

func (d AuditLog) MarshalJSON() ([]byte, error) {
	type Alias AuditLog
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

type AuditLogsType struct {
	// 数据库中存储的operation
	Operation string `json:"operation"`
	// name，自动适配中英文
	Name string `json:"name"`
}
