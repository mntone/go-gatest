package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/mntone/go-gatest/internal/model"
)

const (
	taskAddStmt    = "INSERT INTO tasks(description,created_at,updated_at)VALUES(?1,?2,?2)"
	taskAllStmt    = "SELECT * FROM tasks ORDER BY updated_at"
	taskFindStmt   = "SELECT * FROM tasks WHERE description LIKE '%'||?||'%' ORDER BY updated_at"
	taskRemoveStmt = "DELETE FROM tasks WHERE id=?"
)

type TaskRepositoryImpl struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) TaskRepositoryImpl {
	return TaskRepositoryImpl{
		db: db,
	}
}

func (repo TaskRepositoryImpl) Add(
	ctx context.Context,
	description string,
	createdAt time.Time,
) (created bool, err error) {
	result, err := repo.db.ExecContext(ctx, taskAddStmt, description, createdAt)
	if err != nil {
		return false, err
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowAffected == 1, nil
}

func (repo TaskRepositoryImpl) All(
	ctx context.Context,
) (tasks []model.Task, err error) {
	rows, err := repo.db.QueryContext(ctx, taskAllStmt)
	if err != nil {
		return
	}
	defer func() {
		closeError := rows.Close()
		if closeError != nil {
			err = errors.Join(err, closeError)
		}
	}()

	for rows.Next() {
		var task model.Task
		var createdAtText, updatedAtText string
		if err := rows.Scan(
			&task.ID,
			&task.Description,
			&createdAtText,
			&updatedAtText,
		); err != nil {
			return nil, err
		}

		task.CreatedAt, err = time.Parse(time.RFC3339Nano, createdAtText)
		if err != nil {
			return nil, err
		}

		task.UpdatedAt, err = time.Parse(time.RFC3339Nano, updatedAtText)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return
}

func (repo TaskRepositoryImpl) Find(
	ctx context.Context,
	keyword string,
) (tasks []model.Task, err error) {
	rows, err := repo.db.QueryContext(ctx, taskFindStmt, keyword)
	if err != nil {
		return
	}
	defer func() {
		closeError := rows.Close()
		if closeError != nil {
			err = errors.Join(err, closeError)
		}
	}()

	for rows.Next() {
		var task model.Task
		var createdAtText, updatedAtText string
		if err := rows.Scan(
			&task.ID,
			&task.Description,
			&createdAtText,
			&updatedAtText,
		); err != nil {
			return nil, err
		}

		task.CreatedAt, err = time.Parse(time.RFC3339Nano, createdAtText)
		if err != nil {
			return nil, err
		}

		task.UpdatedAt, err = time.Parse(time.RFC3339Nano, updatedAtText)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return
}

func (repo TaskRepositoryImpl) RemoveByID(
	ctx context.Context,
	taskID int64,
) (deleted bool, err error) {
	result, err := repo.db.ExecContext(ctx, taskRemoveStmt, taskID)
	if err != nil {
		return false, err
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowAffected == 1, nil
}
