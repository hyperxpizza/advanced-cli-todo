package cli

import (
	"flag"
	"fmt"
)

var loglevel = flag.String("loglevel", "debug", "level of logger|default:info")
var interactive = flag.Bool("interactive", false, "run in interactive mode")

func printFlags() {
	fmt.Println("Flags:")
	fmt.Printf("	-loglevel=%s\n", *loglevel)
	fmt.Printf("	-interactive=%v\n", *interactive)
}

func printUsage() {}
