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

const (
	LIGHTBLUE   = "\033[1;34m"
	LIGHTCYAN   = "\033[1;36m"
	LIGHTGREEN  = "\033[1;32m"
	LIGHTGREY   = "\033[1;37m"
	LIGHTPURPLE = "\033[1;35m"
	LIGHTRED    = "\033[1;31m"
	LIGHTYELLOW = "\033[1;33m"
	RED         = "\033[31m"
	RESET       = "\033[0m"

	VERSION = "v0.1.3"
)

var (
	filename    = flag.String("config", "castle.yaml", "Config YAML file to parse.")
	showVersion = flag.Bool("version", false, "Show version and exit.")
)

func init() {
	flag.StringVar(filename, "c", "castle.yaml", "Config YAML file to parse.")
	flag.BoolVar(showVersion, "v", false, "Show version.")
}

func building()   { fmt.Println(LIGHTBLUE, "\bBuilding... 🔨", RESET) }
func running()    { fmt.Println(LIGHTCYAN, "\bRunning... 🏎️", RESET) }
func formatting() { fmt.Println(LIGHTGREEN, "\bFormatting... 🎨", RESET) }
func testing()    { fmt.Println(LIGHTYELLOW, "\bTesting... 🧪", RESET) }
func deploying()  { fmt.Println(LIGHTRED, "\bDeploying... 🚀", RESET) }

func runKeyCmd(names []string, cmds [][]string, logFns []func(), config parse.T) {
	for i, v := range names {
		logFns[i]()
		if config.Config.IsParallelCmd(v) {
			fmt.Println("RUN P")
			var batchSize int
			if config.Config.BatchSize == 0 {
				batchSize = -1
			} else {
				batchSize = config.Config.BatchSize
			}
			RunSectionParallel(cmds[i], batchSize)
		} else {
			fmt.Println("RUN NORMAL")
			RunSection(cmds[i])
		}
	}
}

func main() {
	flag.Parse()

	fmt.Println(LIGHTPURPLE, "\bSandCastle", VERSION, RESET)

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
					runKeyCmd([]string{"build"}, [][]string{config.Build}, []func(){building}, config)
				case "run":
					runKeyCmd([]string{"run"}, [][]string{config.Run}, []func(){running}, config)
				case "format":
					runKeyCmd([]string{"format"}, [][]string{config.Format}, []func(){formatting}, config)
				case "test":
					runKeyCmd([]string{"test"}, [][]string{config.Test}, []func(){testing}, config)
				case "deploy":
					runKeyCmd([]string{"deploy"}, [][]string{config.Deploy}, []func(){deploying}, config)
				case "all":
					runKeyCmd([]string{"build", "run", "format", "test", "deploy"},
						[][]string{config.Build, config.Run, config.Format, config.Test, config.Deploy},
						[]func(){building, running, formatting, testing, deploying}, config)
				default:
					fmt.Println(LIGHTRED, "Unknown command:", v, RESET)
					os.Exit(1)
				}
			}
		}
	} else {
		// Just "castle"
		building()
		RunSection(config.Build)
		running()
		RunSection(config.Run)
	}

	os.Exit(0)
}

// Running commands and tasks
func runTask(config parse.T) {
	fmt.Println(LIGHTRED, "\bTask:", os.Args[2], "... 📝", RESET)

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
		dir, shouldContinue = magicCmds(cmd, dir)
		if shouldContinue {
			continue
		}

		runCmdInner(cmd, dir)
	}
}

func magicCmds(cmd string, dir string) (string, bool) {
	if strings.HasPrefix(cmd, "SETDIR! ") {
		dir = cmd[8:]
	} else if strings.HasPrefix(cmd, "GETDIR!") {
		dir = "."
		cwd, _ := os.Getwd()
		fmt.Println(LIGHTGREY, "\bCurrent directory:", cwd, RESET)
	} else if strings.HasPrefix(cmd, "ECHO! ") {
		fmt.Println(LIGHTGREY, "\b"+cmd[6:], RESET)
	} else {
		return dir, false
	}
	return dir, true
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

func runCmdInner(cmd string, directory string) {
	fmt.Println(LIGHTGREY, "\b→", RESET, cmd)

	c := strings.Split(cmd, " ")
	shCmd := exec.Command(c[0], c[1:]...)
	shCmd.Stdout = os.Stdout
	shCmd.Stderr = os.Stdout
	shCmd.Dir = directory

	err := shCmd.Run()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
