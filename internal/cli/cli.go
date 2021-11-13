package cli

import (
	"fmt"

	"github.com/hyperxpizza/advanced-cli-todo/internal/config"
	"github.com/hyperxpizza/advanced-cli-todo/internal/db"
	"github.com/jroimartin/gocui"
	"github.com/sirupsen/logrus"
)

type CLI struct {
	db     *db.Database
	logger logrus.FieldLogger
}

func NewCLI(c *config.Config, logger logrus.FieldLogger, database *db.Database) *CLI {
	return &CLI{db: database, logger: logger}
}

func (c *CLI) Run() error {
	gui, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		c.logger.Error(err)
		return err
	}

	defer gui.Close()

	gui.SetManagerFunc(layout)

	if err := gui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}

	if err := gui.MainLoop(); err != nil && err != gocui.ErrQuit {
		return err
	}

	return nil
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("hello", maxX/2-7, maxY/2, maxX/2+7, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "Hello world!")
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
