package db

import (
	"database/sql"
	"os"
)

type Database struct {
	file *os.File
	db   *sql.DB
}
