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

func NewCLI(c *config.Config, logger logrus.FieldLogger, database *db.Database) *CLI {
	return &CLI{db: database, logger: logger}
}

func (c *CLI) Run() {}
