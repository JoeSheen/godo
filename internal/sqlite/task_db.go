package sqlite

import (
	"time"

	"github.com/JoeSheen/godo/internal/types"
	_ "github.com/mattn/go-sqlite3"
)

func (dc *DatabaseContext) CreateTask(title string, priority int, category string, deadline *time.Time) (int64, error) {
	result, err := dc.db.Exec(
		"INSERT INTO tasks(title, priority, completed, category, created_timestamp, deadline) VALUES(?, ?, ?, ?, ?, ?)",
		title,
		priority,
		false,
		category,
		time.Now(),
		deadline,
	)
	if err != nil {
		return -1, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (dc *DatabaseContext) GetAllTasks() ([]types.Task, error) {
	rows, err := dc.db.Query("SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}

	tasks := []types.Task{}
	for rows.Next() {
		t := types.Task{}
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

func (dc *DatabaseContext) GetTaskById(id int) (types.Task, error) {
	t := types.Task{}
	row := dc.db.QueryRow("SELECT * FROM tasks WHERE id = ?", id)
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

func (dc *DatabaseContext) GetAllTasksByCompletedStatus(completed bool) ([]types.Task, error) {
	rows, err := dc.db.Query("SELECT * FROM tasks WHERE completed = ?", completed)
	if err != nil {
		return nil, err
	}

	tasks := []types.Task{}
	for rows.Next() {
		t := types.Task{}
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

func (dc *DatabaseContext) ToggleTaskCompleted(id int) error {
	t, err := dc.updateCompletedValue(id)
	if err != nil {
		return err
	}
	_, err = dc.db.Exec(
		"UPDATE tasks SET completed = ?, completed_timestamp = ? WHERE id = ?",
		t.Completed,
		t.CompletedTimestamp,
		t.ID,
	)
	return err
}

func (dc *DatabaseContext) updateCompletedValue(id int) (*types.Task, error) {
	t, err := dc.GetTaskById(id)
	if err != nil {
		return nil, err
	}
	t.Completed = !t.Completed
	if !t.Completed {
		t.CompletedTimestamp = nil
	} else {
		var timestamp = time.Now()
		t.CompletedTimestamp = &timestamp
	}
	return &t, nil
}

func (dc *DatabaseContext) DeleteTaskById(id int) error {
	if _, err := dc.db.Exec("DELETE FROM tasks WHERE id = ?", id); err != nil {
		return err
	}
	return nil
}

func (dc *DatabaseContext) DeleteAllTasks() error {
	if _, err := dc.db.Exec("DELETE FROM tasks"); err != nil {
		return err
	}
	return nil
}
