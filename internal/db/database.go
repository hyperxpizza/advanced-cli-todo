package db

import (
	"database/sql"
	"errors"
	"os"

	"github.com/hyperxpizza/advanced-cli-todo/internal/common"
	"github.com/hyperxpizza/advanced-cli-todo/internal/config"
	"github.com/hyperxpizza/advanced-cli-todo/internal/customErrors"
	"github.com/sirupsen/logrus"
)

type Database struct {
	db     *sql.DB
	logger logrus.FieldLogger
}

func NewDatabase(c *config.Config, logger logrus.FieldLogger) (*Database, error) {

	err := common.CheckIfFileExists(c.Database.Path)
	if err != nil {
		if errors.Is(err, customErrors.Wrap(customErrors.ErrFileNotFound)) {
			return nil, err
		}
		//if database file does not exist, create one
		return createNewDB(c.Database.Path, c.Database.Schema, logger)
	}

	database, err := sql.Open("sqlite3", c.Database.Path)
	if err != nil {
		return nil, err
	}

	db := Database{db: database, logger: logger}

	return &db, nil
}

func createNewDB(path, schemaPath string, logger logrus.FieldLogger) (*Database, error) {
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	database, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	db := Database{db: database, logger: logger}
	//load schema from file
	if err = db.loadSchema(schemaPath); err != nil {
		return nil, err
	}

	return &db, nil
}

func (db *Database) loadSchema(schemaPath string) error {
	data, err := common.ReadFile(schemaPath)
	if err != nil {
		return err
	}
	sql := string(data)
	stmt, err := db.db.Prepare(sql)
	if err != nil {
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) Close() {
	db.db.Close()
}
