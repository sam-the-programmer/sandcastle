package main

import (
	"castle/parse"
	"castle/utils"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"sync"

	"strings"
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
			runTask(config)
		default:
			for _, v := range os.Args[1:] {
				switch v {
				case "build":
					runArgCmd([]string{"build"}, [][]string{config.Build}, []func(){utils.Building}, config)
				case "run":
					runArgCmd([]string{"run"}, [][]string{config.Run}, []func(){utils.Running}, config)
				case "format":
					runArgCmd([]string{"format"}, [][]string{config.Format}, []func(){utils.Formatting}, config)
				case "test":
					runArgCmd([]string{"test"}, [][]string{config.Test}, []func(){utils.Testing}, config)
				case "deploy":
					runArgCmd([]string{"deploy"}, [][]string{config.Deploy}, []func(){utils.Deploying}, config)
				case "all":
					runArgCmd([]string{"build", "run", "format", "test", "deploy"},
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
		RunSection(config.Build)
		utils.Running()
		RunSection(config.Run)
	}

	os.Exit(0)
}

// Running commands and tasks
func runTask(config parse.T) {
	utils.Task(os.Args[2])

	if config.Config.IsParallelTask(os.Args[2]) {
		var batchSize int
		if config.Config.BatchSize == 0 {
			batchSize = -1
		} else {
			batchSize = config.Config.BatchSize
		}
		RunSectionParallel(config.Tasks[os.Args[2]], batchSize)
	} else {
		RunSection(config.Tasks[os.Args[2]])
	}
	os.Exit(0)
}

func RunSection(iter []string) {
	dir := "."
	shouldContinue := false
	for _, cmd := range iter {
		// Magic commands.
		dir, shouldContinue = utils.MagicCmds(cmd, dir)
		if shouldContinue {
			continue
		}

		runCmdInner(cmd, dir)
	}
}

func RunSectionParallel(iter []string, batchSize int) {
	if batchSize == -1 {
		batchSize = len(iter)
	}

	var batches [][]string
	var current []string
	for _, cmd := range iter {
		current = append(current, cmd)
		if len(current) == batchSize {
			batches = append(batches, current)
			current = []string{}
		}
	}

	if len(current) > 0 {
		batches = append(batches, current)
	}

	for _, batch := range batches {
		var wg sync.WaitGroup
		for _, cmd := range batch {
			wg.Add(1)
			go func(cmd string) {
				runCmdInner(cmd, ".")
				wg.Done()
			}(cmd)
		}

		wg.Wait()
	}
}

// Runs an argument command, e.g. "castle build"
func runArgCmd(names []string, cmds [][]string, logFns []func(), config parse.T) {
	for i, v := range names {
		logFns[i]()
		if config.Config.IsParallelCmd(v) {
			var batchSize int
			if config.Config.BatchSize == 0 {
				batchSize = -1
			} else {
				batchSize = config.Config.BatchSize
			}
			RunSectionParallel(cmds[i], batchSize)
		} else {
			RunSection(cmds[i])
		}
	}
}

func runCmdInner(cmd string, directory string) {
	utils.RunningCmd(cmd)

	c := strings.Split(cmd, " ")
	shCmd := exec.Command(c[0], c[1:]...)
	if *showLog {
		shCmd.Stdout = os.Stdout
		shCmd.Stderr = os.Stdout
	}
	shCmd.Dir = directory

	err := shCmd.Run()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
