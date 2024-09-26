package types

import "github.com/google/uuid"

type Priority int

const (
	Low Priority = iota
	Medium
	High
)

type Task struct {
	Id       uuid.UUID
	Name     string
	Priority Priority
}
