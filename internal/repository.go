package internal

import (
	"context"
	"time"

	"github.com/mntone/go-gatest/internal/model"
)

type TaskRepository interface {
	Add(
		ctx context.Context,
		description string,
		createdAt time.Time,
	) (created bool, err error)

	All(
		ctx context.Context,
	) (tasks []model.Task, err error)

	Find(
		ctx context.Context,
		keyword string,
	) (tasks []model.Task, err error)

	RemoveByID(
		ctx context.Context,
		taskID int64,
	) (deleted bool, err error)
}
