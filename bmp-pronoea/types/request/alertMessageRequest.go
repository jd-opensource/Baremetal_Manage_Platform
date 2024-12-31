package request
const (
	ALERT_RECEIVE_STATUS_FIRING = "firing"
	ALERT_RECEIVE_STATUS_RESOLVED = "resolved"
)

type AlertReceiverRequest struct {
	Receiver      string `json:"receiver"`
	Status        string `json:"status"`
	Alerts        []*AlertInfo `json:"alerts"`
}

type AlertInfo struct {
	Status        string `json:"status"`
	Labels        map[string]string `json:"labels"`
	Annotations   map[string]string `json:"annotations"`
	StartsAt	  string `json:"startsAt"`
	EndsAt	      string `json:"endsAt"`
}
