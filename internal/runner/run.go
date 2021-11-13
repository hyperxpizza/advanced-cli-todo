package runner

import (
	"sync"

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
	return nil
}

//Running only api
func (r *Runner) RunAPI() error {
	return nil
}

func (r *Runner) Close() {
}
