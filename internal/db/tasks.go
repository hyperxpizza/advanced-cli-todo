package db

import (
	"database/sql"
	"time"

	"github.com/hyperxpizza/advanced-cli-todo/internal/models"
)

// Inserting a new task into the sqlite databas
func (db *Database) InsertTask(title, description string, priority int, due *time.Time) (int64, error) {
	db.logger.Debugf("Inserting a new task with title: %s", title)
	stmt, err := db.db.Prepare(`insert into tasks(id, title, description, priority, dueDate, created, updated) values(DEFAULT, $1, $2, $3, $4, $5, $6, $7)`)
	if err != nil {
		return 0, err
	}

	db.mutex.Lock()
	defer db.mutex.Unlock()
	res, err := stmt.Exec(title, description, false, priority, due, time.Now(), time.Now())
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (db *Database) GetTaskByID(id int) (*models.Task, error) {
	var task models.Task

	db.mutex.Lock()
	defer db.mutex.Unlock()
	due := sql.NullTime{}
	err := db.db.QueryRow(`select * from tasks where id = $1`).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Done,
		&task.Priority,
		&due,
		&task.Created,
		&task.Updated,
	)
	if err != nil {
		return nil, err
	}

	if due.Valid {
		task.DueDate = &due.Time
	} else {
		task.DueDate = nil
	}

	return &task, nil
}

func (db *Database) GetAllTasks(order string) ([]*models.Task, error) {
	var tasks []*models.Task

	return tasks, nil
}

func (db *Database) MarkTaskAsDone(id int) error {
	return nil
}

func (db *Database) DeleteTask() error {
	return nil
}
