package cli

import (
	"github.com/hyperxpizza/advanced-cli-todo/internal/config"
	"github.com/hyperxpizza/advanced-cli-todo/internal/db"
	"github.com/sirupsen/logrus"
)

type CLI struct {
	db     *db.Database
	logger logrus.FieldLogger
}

func NewCLI(c *config.Config, logger logrus.FieldLogger) (*CLI, error) {
	database, err := db.NewDatabase(c, logger)
	if err != nil {
		return nil, err
	}

	defer database.Close()

	return &CLI{db: database, logger: logger}, nil
}

func (c *CLI) Run() {}
