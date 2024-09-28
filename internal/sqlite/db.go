package sqlite

import (
	"context"
	"database/sql"
	"time"

	"github.com/JoeSheen/godo/internal/types"
	_ "github.com/mattn/go-sqlite3"
)

const (
	driverName = "sqlite3"
)

type DatabaseContext struct {
	db *sql.DB
}

func OpenDBConnection(fileName string) (*DatabaseContext, error) {
	db, err := sql.Open(driverName, fileName)
	if err != nil {
		return nil, err
	}
	return &DatabaseContext{db: db}, nil
}

func (d *DatabaseContext) CreateTask(priority int, deadline *time.Time, title, category string) (int64, error) {
	stmt, err := d.db.Prepare("INSERT INTO tasks(title, priority, completed, category, created_timestamp, deadline) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		return -1, err
	}
	result, err := stmt.ExecContext(context.Background(), title, priority, false, category, time.Now(), deadline)
	if err != nil {
		return -1, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (d *DatabaseContext) GetAllTasks() ([]types.Task, error) {
	rows, err := d.db.QueryContext(context.Background(), "SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}

	var tasks []types.Task
	for rows.Next() {
		var t types.Task
		err = rows.Scan(
			&t.ID,
			&t.Title,
			&t.Priority,
			&t.Completed,
			&t.Category,
			&t.CreatedTimestamp,
			&t.CompletedTimestamp,
			&t.Deadline,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	return tasks, nil
}

func (d *DatabaseContext) GetTaskByID(id int) (types.Task, error) {
	var t types.Task
	row := d.db.QueryRowContext(context.Background(), "SELECT * FROM tasks WHERE id = ?", id)
	err := row.Scan(
		&t.ID,
		&t.Title,
		&t.Priority,
		&t.Completed,
		&t.Category,
		&t.CreatedTimestamp,
		&t.CompletedTimestamp,
		&t.Deadline,
	)
	return t, err
}

func (d *DatabaseContext) GetAllUncompletedTasks(completed bool) ([]types.Task, error) {
	rows, err := d.db.QueryContext(context.Background(), "SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}

	var tasks []types.Task
	for rows.Next() {
		var t types.Task
		err = rows.Scan(
			&t.ID,
			&t.Title,
			&t.Priority,
			&t.Completed,
			&t.Category,
			&t.CreatedTimestamp,
			&t.CompletedTimestamp,
			&t.Deadline,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	return tasks, nil
}
