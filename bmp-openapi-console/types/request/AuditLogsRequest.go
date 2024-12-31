package request

import log "coding.jd.com/aidc-bmp/bmp_log"

type DescribeAuditLogsRequest struct {
	// instance_id
	InstanceID string `json:"instance_id" validate:"required"`
	// 操作名称
	Operation string `json:"operation"`
	// 操作人
	UserName string `json:"username"`
	// result
	Result string `json:"result"`
	// 操作时间下限
	StartTime int `json:"startTime"`
	// 操作时间上限
	EndTime int `json:"endTime"`
	// 是否显示所有
	IsAll string `json:"isAll"`
	Pageable
}

func (req *DescribeAuditLogsRequest) Validate(logger *log.Logger) {
	validate(req, logger)
}
