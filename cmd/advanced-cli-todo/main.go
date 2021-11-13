package main

import (
	"flag"
	"log"
	"os"
	"runtime"

	"github.com/hyperxpizza/advanced-cli-todo/internal/common"
	"github.com/hyperxpizza/advanced-cli-todo/internal/config"
	"github.com/hyperxpizza/advanced-cli-todo/internal/runner"
)

var mode = flag.String("mode", "default", "mode of running the program. availiable options: api|cli|default")
var configPathPtr = flag.String("config", "", "Path to file containing config.yml")
var loglevel = flag.String("loglevel", "info", "level of logger|default:info")

func main() {
	flag.Parse()

	//init a new logger instance
	logger := common.NewLogger(*loglevel)
	logger.Infof("Running the app in mode: %s", *mode)

	//load config from file
	c, err := config.NewConfig(*configPathPtr)
	if err != nil {
		log.Fatalf("Could not load config from file: %s error: %s", *configPathPtr, err.Error())
		os.Exit(1)
	}

	switch *mode {
	case "default":
		runner.RunInDefaultMode()
	case "web":
		runner.RunAPI()
	case "cli":
		runner.RunCli()
	default:
		logger.Errorf("Mode: %s unknown! Aborting...", *mode)
		runtime.Goexit()
	}

}
