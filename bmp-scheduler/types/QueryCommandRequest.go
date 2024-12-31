package types

import "time"

type QueryCommandRequest struct {
	StartTime time.Time
	EndTime   time.Time
}
