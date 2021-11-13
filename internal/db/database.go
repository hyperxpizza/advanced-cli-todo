package db

import (
	"database/sql"
	"errors"
	"os"
	"sync"

	"github.com/hyperxpizza/advanced-cli-todo/internal/common"
	"github.com/hyperxpizza/advanced-cli-todo/internal/config"
	"github.com/hyperxpizza/advanced-cli-todo/internal/customErrors"
	"github.com/sirupsen/logrus"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	db     *sql.DB
	logger logrus.FieldLogger
	mutex  sync.Mutex
}

//Initializes a new Database object, creates the file is does not exis and fills the schema
func NewDatabase(c *config.Config, logger logrus.FieldLogger) (*Database, error) {
	logger.Debug("Initializing a new database connection")

	err := common.CheckIfFileExists(c.Database.Path)
	if err != nil {
		logger.Debug("Database file not found")
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

	db := Database{db: database, logger: logger, mutex: sync.Mutex{}}

	//check if tables exist
	_, err = database.Query(`select * from tasks`)
	if err != nil {
		if err.Error() == "no such table: tasks" {
			db.logger.Debug("Required tables do not exist, creating")
			if err = db.loadSchema(c.Database.Schema); err != nil {
				return nil, err
			}
			logger.Debug("A new database object has been created")
		} else {
			return nil, err
		}
	}

	return &db, nil
}

//Creates a new Database file
func createNewDB(path, schemaPath string, logger logrus.FieldLogger) (*Database, error) {
	logger.Debug("Creating a new database file")
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	database, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	db := Database{db: database, logger: logger, mutex: sync.Mutex{}}
	//load schema from file
	if err = db.loadSchema(schemaPath); err != nil {
		return nil, err
	}

	return &db, nil
}

//loads schema from sql file and executes it creating new tables
func (db *Database) loadSchema(schemaPath string) error {
	data, err := common.ReadFile(schemaPath)
	if err != nil {
		return err
	}
	sqlData := string(data)
	db.logger.Debugf("SQL to execute = %s", sqlData)
	stmt, err := db.db.Prepare(sqlData)
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
