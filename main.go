package main

import (
	"castle/base"
	"castle/exec"
	"castle/parse"
	"flag"
	"fmt"
	"os"
)

var (
	configFile  string
	flagVersion bool
)

func init() {
	base.Level = 10 // Hide all base logs

	flag.StringVar(&configFile, "f", "castle.yaml", "Config file to use.")
	flag.BoolVar(&flagVersion, "v", false, "Print version and exit.")
	flag.Parse()
}

func main() {
	// Version and exit
	if flagVersion {
		fmt.Println(base.VERSION)
		os.Exit(0)
	}

	config := parse.Parse(configFile)
	command := flag.Arg(0)

	// No command
	if command == "" {
		fmt.Println("Please specify a command.")
		os.Exit(1)
	}

	// e.g. castle init
	if exec.IsSpecialCommand(command) {
		exec.SpecialCommands[command](flag.Args()[1:])
		os.Exit(0)
	}

	// Normal command
	exec.ExecuteTask(config.TaskArgs[command], command, config.Config.LogShellCmds)
}
