package runner

import (
	"sync"

	"github.com/hyperxpizza/advanced-cli-todo/internal/api"
	"github.com/hyperxpizza/advanced-cli-todo/internal/config"
	"github.com/sirupsen/logrus"
)

type Runner struct {
	wg     sync.WaitGroup
	c      *config.Config
	logger logrus.FieldLogger
}

func NewRunner(c *config.Config, logger logrus.FieldLogger) *Runner {
	return &Runner{
		wg:     sync.WaitGroup{},
		c:      c,
		logger: logger,
	}
}

//Running both api and cli
func (r *Runner) RunInDefaultMode() {
	r.wg = sync.WaitGroup{}
	r.wg.Add(2)

	go func() {
		err := r.RunCli()
		if err != nil {
			r.wg.Done()
			return
		}
	}()

	go func() {
		err := r.RunAPI()
		if err != nil {
			r.wg.Done()
			return
		}
	}()

	r.wg.Wait()
}

//Running only cli
func (r *Runner) RunCli() error {
	r.logger.Info("Starting CLI mode...")
	return nil
}

//Running only api
func (r *Runner) RunAPI() error {
	r.logger.Info("Starting API mode...")
	a, err := api.NewAPI(r.c, r.logger)
	if err != nil {
		r.logger.Error(err)
		return err
	}

	//start the api server
	a.Run()

	return nil
}

func (r *Runner) Close() {

}
