package types

type CreateCommandRequest struct {
	Action          string
	CommandType     string
	RequestId       string
	Sn              string
	InstanceId      string
	ParentCommandId int64
}
