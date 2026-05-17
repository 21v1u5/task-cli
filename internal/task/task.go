package task

import "time"

type Status string

const (
	StatusPending Status = "pending"
	StatusDone    Status = "done"
)

type Task struct{
	ID          int64
	Title       string
	Description string
	Status      Status
	CreatedAt   time.Time
	UpdatedAt   time.Time
}