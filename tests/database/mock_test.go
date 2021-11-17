package main

import (
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

// Creates a new sqlmock database object for testing
func createNewMockDB() (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	return db, mock, nil
}

// Get tasks table columns for mock
func getColumns() []string {
	return []string{"id", "title", "description", "done", "priority", "dueDate", "created", "updated"}
}

func TestGetAllTasks(t *testing.T) {
	db, mock, err := createNewMockDB()
	assert.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows(getColumns()).AddRow(1, "task 1", "description 1", false, 5, time.Now().Add(time.Hour*1), time.Now(), time.Now()).AddRow(2, "task 2", "description 2", false, 5, time.Now().Add(time.Hour*1), time.Now(), time.Now()).AddRow(3, "task 3", "description 3", false, 5, time.Now().Add(time.Hour*1), time.Now(), time.Now())

	mock.ExpectQuery("SELECT * FROM tasks").WillReturnRows(rows)
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
