package main

import "flag"

//flags required for running tests
var configPathPtr = flag.String("config", "", "Path to file containing config.yml")
var filePtr = flag.String("file", "", "Path to file")
var loglevel = flag.String("loglevel", "debug", "level of logger|default:info")
