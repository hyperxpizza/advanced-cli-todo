package db

import "github.com/hyperxpizza/advanced-cli-todo/internal/models"

func (db *Database) InsertTask() (int, error) {
	var id int

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
