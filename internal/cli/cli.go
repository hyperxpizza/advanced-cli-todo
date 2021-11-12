package cli

import (
	"flag"

	"github.com/hyperxpizza/advanced-cli-todo/internal/common"
	"github.com/hyperxpizza/advanced-cli-todo/internal/config"
	"github.com/hyperxpizza/advanced-cli-todo/internal/db"
)

type CLI struct {
	db *db.Database
}

func NewCLI(c *config.Config) (*CLI, error) {
	logger := common.NewLogger(*loglevel)
	database, err := db.NewDatabase(c, logger)
	if err != nil {
		return nil, err
	}

	defer database.Close()

	return &CLI{db: database}, nil
}

func Run() {
	flag.Parse()
	printFlags()
}
