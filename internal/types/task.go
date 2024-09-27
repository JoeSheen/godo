package types

import (
	"time"
)

type Priority int

const (
	Lowest Priority = iota
	Low
	Medium
	High
	Highest
	Urgent // come up with a better name for this value
)

type CategoryType string

type Task struct {
	ID                 string
	Title              string
	Priority           Priority
	Completed          bool
	Category           CategoryType
	CreatedTimestamp   time.Time
	CompletedTimestamp *time.Time
	Deadline           *time.Timer
}
