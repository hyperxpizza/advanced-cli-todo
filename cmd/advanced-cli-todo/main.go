package main

import (
	"flag"

	"github.com/hyperxpizza/advanced-cli-todo/internal/common"
)

var mode = flag.String("mode", "default", "mode of running the program. availiable options: web|cli|default")
var configPathPtr = flag.String("config", "", "Path to file containing config.yml")
var loglevel = flag.String("loglevel", "info", "level of logger|default:info")

func main() {
	flag.Parse()

	logger := common.NewLogger(*loglevel)
	logger.Infof("Running the app in mode: %s", *mode)

	switch *mode

}
