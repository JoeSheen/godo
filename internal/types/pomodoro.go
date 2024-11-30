package types

import "time"

type Pomodoro struct {
	ID        int64
	Name      string
	Active    bool
	StartTime time.Time
	EndTime   time.Time
}
