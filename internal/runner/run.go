package runner

import (
	"sync"

	"github.com/hyperxpizza/advanced-cli-todo/internal/api"
	"github.com/hyperxpizza/advanced-cli-todo/internal/cli"
	"github.com/hyperxpizza/advanced-cli-todo/internal/config"
	"github.com/hyperxpizza/advanced-cli-todo/internal/db"
	"github.com/sirupsen/logrus"
)

type Runner struct {
	wg              sync.WaitGroup
	c               *config.Config
	logger          logrus.FieldLogger
	db              *db.Database
	closeApiChannel chan bool
	closeCLiChannel chan bool
}

func NewRunner(c *config.Config, logger logrus.FieldLogger) (*Runner, error) {

	database, err := db.NewDatabase(c, logger)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return &Runner{
		wg:              sync.WaitGroup{},
		c:               c,
		logger:          logger,
		db:              database,
		closeApiChannel: make(chan bool),
		closeCLiChannel: make(chan bool),
	}, nil
}

//Running both api and cli
func (r *Runner) RunInDefaultMode() {
	r.wg = sync.WaitGroup{}

	r.wg.Add(1)
	go func() {
		r.RunAPI()
	}()

	go func() {
		if err := r.RunCli(); err != nil {
			r.wg.Done()
		}

	}()

	r.wg.Wait()
}

//Running only cli
func (r *Runner) RunCli() error {
	r.logger.Info("Starting CLI mode...")
	c := cli.NewCLI(r.c, r.logger, r.db)
	err := c.Run()
	if err != nil {
		return err
	}

	return nil
}

//Running only api
func (r *Runner) RunAPI() {
	r.logger.Info("Starting API mode...")
	a := api.NewAPI(r.c, r.logger, r.db)
	//start the api server
	a.Run()
}

func (r *Runner) Close() {
	r.logger.Println("Exiting runner")
	r.wg.Done()
	r.wg.Done()
}
