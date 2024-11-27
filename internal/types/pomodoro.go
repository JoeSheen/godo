package types

import "time"

type Pomodoro struct {
	ID        int64
	StartTime time.Time
}
