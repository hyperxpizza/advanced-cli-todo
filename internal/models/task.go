package models

import "time"

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	Priority    uint      `json:"priority"`
	DueDate     time.Time `json:"dueDate"`
	Created     time.Time `json:"created"`
	Updated     time.Time `json:"updated"`
}

type NewTaskRequest struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Priority    int       `json:"priority"`
	DueDate     time.Time `json:"dueDate"`
}

//Creates a new task structure
func NewTask(id int, title, description string, done bool, priority uint, due, created, updated time.Time) Task {
	return Task{ID: id, Title: title, Description: description, Done: done, Priority: priority, DueDate: due, Created: created, Updated: updated}
}
