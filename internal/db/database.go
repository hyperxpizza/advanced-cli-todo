package db

import (
	"database/sql"
	"errors"
	"os"

	"github.com/hyperxpizza/advanced-cli-todo/internal/common"
	"github.com/hyperxpizza/advanced-cli-todo/internal/customErrors"
)

type Database struct {
	db *sql.DB
}

func NewDatabase(path string) (*Database, error) {

	err := common.CheckIfFileExists(path)
	if err != nil {
		if !errors.Is(err, customErrors.Wrap(customErrors.ErrFileNotFound)) {
			return nil, err
		}
		//if database file does not exist, create one
		return createNewDB(path)
	}

	database, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	db := Database{database}

	return &db, nil
}

func createNewDB(path string) (*Database, error) {
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	database, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	db := Database{database}
	//load schema from file
	if err = db.loadSchema(); err != nil {
		return nil, err
	}

	return &db, nil
}

func (db *Database) loadSchema() error {
	return nil
}

func (db *Database) Close() {
	db.db.Close()
}
