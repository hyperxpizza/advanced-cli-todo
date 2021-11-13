package api

import (
	"github.com/hyperxpizza/advanced-cli-todo/internal/config"
	"github.com/hyperxpizza/advanced-cli-todo/internal/db"
	"github.com/sirupsen/logrus"
)

type API struct {
	db     *db.Database
	logger logrus.FieldLogger
}

func NewAPI(c *config.Config, logger logrus.FieldLogger) (*API, error) {
	var api API

	db, err := db.NewDatabase(c, logger)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	api.db = db
	api.logger = logger

	return &api, nil
}
