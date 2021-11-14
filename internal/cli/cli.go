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
	gui    *gocui.Gui
}

func NewCLI(c *config.Config, logger logrus.FieldLogger, database *db.Database) (*CLI, error) {
	gui, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return &CLI{db: database, logger: logger, gui: gui}, nil
}

func (c *CLI) Run() error {

	defer c.gui.Close()

	c.gui.SetManagerFunc(layout)
	if err := c.setKeyBindings(); err != nil {
		return err
	}

	if err := c.gui.MainLoop(); err != nil && err != gocui.ErrQuit {
		return err
	}

	return nil
}

func (c *CLI) setKeyBindings() error {
	if err := c.gui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
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
