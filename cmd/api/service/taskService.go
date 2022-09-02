package service

import (
	"database/sql"
	"todolist/cmd/api/model"
)

type Task struct {
	db *sql.DB
}

func CreateTaskService(db *sql.DB) *Task {
	return &Task{db}
}

func (service Task) CreateTask(task model.Task) (uint64, error) {
	stmt, err := service.db.Prepare(
		"INSERT INTO tasks (title, description, start_date, end_date, priority, done, assignee) VALUES (?, ?, ?, ?, ?, ?, ?)",
	)

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(task.Title, task.Description, task.StartDate, task.EndDate, task.Priority, task.Done, task.Assignee)

	if err != nil {
		return 0, err
	}

	lastInsertId, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return uint64(lastInsertId), nil
}

func (service Task) GetTasks() ([]model.Task, error) {
	rows, err := service.db.Query("SELECT * FROM tasks")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []model.Task

	for rows.Next() {
		var task model.Task

		if err := rows.Scan(
			&task.Title,
			&task.Description,
			&task.StartDate,
			&task.EndDate,
			&task.Priority,
			&task.Done,
			&task.Assignee,
		); err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (service Task) GetTask(id uint64) (model.Task, error) {
	row, err := service.db.Query("SELECT * FROM tasks WHERE id = ?", id)

	if err != nil {
		return model.Task{}, err
	}

	defer row.Close()

	var task model.Task

	if row.Next() {
		if err := row.Scan(
			&task.Title,
			&task.Description,
			&task.StartDate,
			&task.EndDate,
			&task.Priority,
			&task.Done,
			&task.Assignee,
		); err != nil {
			return model.Task{}, err
		}
	}

	return task, nil
}

func (service Task) UpdateTask(id uint64, task model.Task) error {
	stmt, err := service.db.Prepare(
		"UPDATE tasks SET title = ?, description = ?, start_date = ?, end_date = ?, priority = ?, done = ?, assignee = ? WHERE id = ?",
	)

	if err != nil {
		return err
	}

	defer stmt.Close()

	if _, err := stmt.Exec(
		&task.Title,
		&task.Description,
		&task.StartDate,
		&task.EndDate,
		&task.Priority,
		&task.Done,
		&task.Assignee,
		id,
	); err != nil {
		return err
	}

	return nil
}

func (service Task) DeleteTask(id uint64) error {
	stmt, err := service.db.Prepare("DELETE from tasks WHERE id = ?")

	if err != nil {
		return err
	}

	defer stmt.Close()

	if _, err := stmt.Exec(id); err != nil {
		return err
	}

	return nil
}
