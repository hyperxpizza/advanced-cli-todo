package db

import (
	"database/sql"
	"time"

	"github.com/hyperxpizza/advanced-cli-todo/internal/models"
)

// Inserting a new task into the sqlite databas
func (db *Database) InsertTask(t models.NewTaskRequest) (int64, error) {
	db.logger.Debugf("Inserting a new task with title: %s", t.Title)
	stmt, err := db.db.Prepare(`insert into tasks(title, description, done, priority, dueDate, created, updated) values ($1, $2, $3, $4, $5, $6, $7)`)
	if err != nil {
		return 0, err
	}

	db.mutex.Lock()
	res, err := stmt.Exec(t.Title, t.Description, false, t.Priority, t.DueDate, time.Now(), time.Now())
	db.mutex.Unlock()
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
	db.mutex.Unlock()
	if err != nil {
		return nil, err
	}

	if due.Valid {
		task.DueDate = due.Time
	}

	return &task, nil
}

//Selects all the tasks from the database
func (db *Database) GetAllTasks(order string) ([]*models.Task, error) {
	var tasks []*models.Task

	db.mutex.Lock()
	rows, err := db.db.Query(`select * from tasks`)
	db.mutex.Unlock()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		due := sql.NullTime{}
		var task models.Task
		err := rows.Scan(
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
			task.DueDate = due.Time
		}

		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func (db *Database) UpdateDone(id int, done bool) error {

	stmt, err := db.db.Prepare(`update tasks set done = $1 where id = $2`)
	if err != nil {
		return err
	}

	db.mutex.Lock()
	_, err = stmt.Exec(done, id)
	db.mutex.Unlock()
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) DeleteTask() error {
	return nil
}
