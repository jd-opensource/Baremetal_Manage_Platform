package event

import (
	"time"

	"git.jd.com/cps-golang/ironic-common/ironic/util"
)

type Event struct {
	EventUuid string `json:"event_uuid"`
	ClazzName string `json:"clazz_name"`
	Body      string `json:"body"`
	Time      int64  `json:"time"`
	UserID    string `json:"user_id"`
}

type CpsRabbitMsg interface {
	ClazzName() string
}

func NewEvent(msg CpsRabbitMsg, logid string, userId string) (*Event, error) {
	b, err := util.Convert2String(msg)
	if err != nil {
		return nil, err
	}
	return &Event{
		EventUuid: logid,
		ClazzName: msg.ClazzName(),
		Body:      string(b),
		Time:      time.Now().UnixNano() / 1e6,
		UserID:    userId,
	}, nil
}
