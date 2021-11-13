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
func NewAPI(c *config.Config, logger logrus.FieldLogger, db *db.Database) *API {
	var api API
	api.db = db
	api.c = c
	api.logger = logger
	api.CloseChannel = make(chan bool)

	return &api
}

//Runs the API router
func (a *API) Run() {
	a.setupAndRunRouter(a.c)
}

func (a *API) ForceClose() {}
