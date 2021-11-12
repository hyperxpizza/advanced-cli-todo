package models

import "time"

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	Priority    int       `json:"priority"`
	DueDate     time.Time `json:"dueDate"`
	Created     time.Time `json:"created"`
	Updated     time.Time `json:"updated"`
}
