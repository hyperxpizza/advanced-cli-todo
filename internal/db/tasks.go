package db

import (
	"time"

	"github.com/hyperxpizza/advanced-cli-todo/internal/models"
)

// Inserting a new task into the sqlite database
func (db *Database) InsertTask(title, description string, priority int, due *time.Time) (int64, error) {
	db.logger.Debugf("Inserting a new task with title: %s", title)
	stmt, err := db.db.Prepare(`insert into tasks(id, title, description, priority, dueDate, created, updated) values(DEFAULT, $1, $2, $3, $4, $5, $6, $7)`)
	if err != nil {
		return 0, err
	}

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

func (db *Database) GetTaskByID() (*models.Task, error) {
	var task models.Task

	return &task, nil
}

func (db *Database) GetAllTasks(order string) ([]*models.Task, error) {
	var tasks []*models.Task

	return tasks, nil
}
