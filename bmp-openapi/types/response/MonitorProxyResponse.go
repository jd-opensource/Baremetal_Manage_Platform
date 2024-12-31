package response

type AgentStatusResponse struct {
	AgentStatus []AgentStatusItem `json:"agentStatus"`
}

type AgentStatusItem struct {
	InstanceID   string `json:"instance_id"`
	SN           string `json:"sn"`
	AgentVersion string `json:"agent_version"`
	Timestamp    int64  `json:"timestamp"`
	// 1->running 2->stopped 0->unknown
	Status int `json:"status"`
	//status 运行 停止 未知
	StatusName string `json:"statusName"`
}

type TagsResponse struct {
	Tags []string `json:"tags"`
}
