package api

import (
	"github.com/hyperxpizza/advanced-cli-todo/internal/config"
	"github.com/hyperxpizza/advanced-cli-todo/internal/db"
	"github.com/sirupsen/logrus"
)

type API struct {
	db           *db.Database
	c            *config.Config
	logger       logrus.FieldLogger
	CloseChannel chan (bool)
}

//Initializes and returns a new API structure
func NewAPI(c *config.Config, logger logrus.FieldLogger) (*API, error) {
	var api API

	db, err := db.NewDatabase(c, logger)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	api.db = db
	api.c = c
	api.logger = logger
	api.CloseChannel = make(chan bool)

	return &api, nil
}

//Runs the API router
func (a *API) Run() {
	setupAndRunRouter(a.c)
}

func (a *API) ForceClose() {}
