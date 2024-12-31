package request

type DesrcibeAgentStatusRequest struct {
	// InstanceId,多个用逗号分隔
	// required: true
	InstanceId string `json:"instanceId" validate:"required"`
}

type DesrcibeTagsRequest struct {
	// instanceId
	// required: true
	InstanceID string `json:"instanceId" validate:"required"`
	// tagName [disk mountpoint nic]
	// required: true
	TagName string `json:"tagName" validate:"required,oneof=disk mountpoint nic"`
}
