package db

import (
	"database/sql"
	"os"
)

type Database struct {
	file *os.File
	db   *sql.DB
}

func NewDatabase(path string) (*Database, error) {
	var db Database

	return &db, nil
}

func (db *Database) loadSchema() (string, error) {

}
