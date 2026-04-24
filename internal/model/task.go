package model

import "time"

type Task struct {
	ID          int64
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
