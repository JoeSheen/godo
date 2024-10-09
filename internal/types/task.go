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
	Urgent
)

type CategoryType string

type Task struct {
	ID                 int64
	Title              string
	Priority           Priority
	Completed          bool
	Category           CategoryType
	CreatedTimestamp   time.Time
	CompletedTimestamp *time.Time
	Deadline           *time.Time
}
