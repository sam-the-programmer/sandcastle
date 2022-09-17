package main

import (
	"castle/parse"
	"castle/utils"
	"flag"
	"os"
)

var (
	filename    = flag.String("config", "castle.yaml", "Config YAML file to parse.")
	showVersion = flag.Bool("version", false, "Show version and exit.")
	showLog     = flag.Bool("log", true, "Show console output output.")
)

func init() {
	flag.StringVar(filename, "c", "castle.yaml", "Config YAML file to parse.")
	flag.BoolVar(showVersion, "v", false, "Show version.")
	flag.BoolVar(showLog, "l", true, "Show console output output.")
}

func main() {
	flag.Parse()

	utils.Version()

	if *showVersion {
		os.Exit(0)
	}

	config := parse.Parse(*filename)
	// Remove the flags
	utils.RemoveFlags()

	if len(os.Args) > 1 {
		// If using a keyword like "build", "run", "test", etc.
		switch os.Args[1] {
		case "task":
			utils.RunTask(config, os.Args[2])
		default:
			for _, v := range os.Args[1:] {
				switch v {
				case "build":
					utils.RunArgCmd([]string{"build"}, [][]string{config.Build}, []func(){utils.Building}, config)
				case "run":
					utils.RunArgCmd([]string{"run"}, [][]string{config.Run}, []func(){utils.Running}, config)
				case "format":
					utils.RunArgCmd([]string{"format"}, [][]string{config.Format}, []func(){utils.Formatting}, config)
				case "test":
					utils.RunArgCmd([]string{"test"}, [][]string{config.Test}, []func(){utils.Testing}, config)
				case "deploy":
					utils.RunArgCmd([]string{"deploy"}, [][]string{config.Deploy}, []func(){utils.Deploying}, config)
				case "all":
					utils.RunArgCmd([]string{"build", "run", "format", "test", "deploy"},
						[][]string{config.Build, config.Run, config.Format, config.Test, config.Deploy},
						[]func(){utils.Building, utils.Running, utils.Formatting, utils.Testing, utils.Deploying}, config)
				default:
					os.Exit(1)
				}
			}
		}
	} else {
		// Just "castle"
		utils.Building()
		utils.RunSection(config.Build, config)
		utils.Running()
		utils.RunSection(config.Run, config)
	}

	os.Exit(0)
}
