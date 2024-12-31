package api

type ApiMessage struct {
	RequestId string `json:"request_id"`
}

func (c ApiMessage) ClazzName() string {
	return "com.jcloud.cps.ironic.event.api.ApiMessage"
}
