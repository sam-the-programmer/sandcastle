package utils

import (
	"castle/constants"
	"castle/parse"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
)

func RemoveFlags() {
	for i, v := range os.Args {
		if strings.HasPrefix(v, "-") {
			os.Args = append(os.Args[:i], os.Args[i+2:]...)
		}
	}
}

func Version() {
	fmt.Println(constants.LIGHTPURPLE, "\bSandCastle", constants.VERSION, constants.RESET)
}

func Building()   { fmt.Println(constants.LIGHTBLUE, "\bBuilding... 🔨", constants.RESET) }
func Running()    { fmt.Println(constants.LIGHTCYAN, "\bRunning... ⏲️", constants.RESET) }
func Formatting() { fmt.Println(constants.LIGHTGREEN, "\bFormatting... 🎨", constants.RESET) }
func Testing()    { fmt.Println(constants.LIGHTYELLOW, "\bTesting... 🧪", constants.RESET) }
func Deploying()  { fmt.Println(constants.LIGHTRED, "\bDeploying... 🚀", constants.RESET) }

func Task(name string) {
	fmt.Println(constants.LIGHTRED, "\bTask:", name, "... 📝", constants.RESET)
}

func RunningCmd(cmd any) {
	fmt.Println(constants.LIGHTGREY, "\b→", constants.RESET, cmd)
}

func UnknownCmd(v any) {
	fmt.Println(constants.LIGHTRED, "Unknown command:", v, constants.RESET)
}

func RunTask(config parse.T, taskName string) {
	if config.Config.LogLevel != "none" {
		Task(taskName)
	}

	if config.Config.IsParallelTask(taskName) {
		var batchSize int
		if config.Config.BatchSize == 0 {
			batchSize = -1
		} else {
			batchSize = config.Config.BatchSize
		}
		RunSectionParallel(config.Tasks[taskName], batchSize, config)
	} else {
		RunSection(config.Tasks[taskName], config)
	}
}

func RunSection(iter []string, config parse.T) {
	dir := "."
	shouldContinue := false
	for _, cmd := range iter {
		// Magic commands.
		shouldContinue = MagicCmds(cmd, &dir, config)
		if shouldContinue {
			continue
		}

		RunCmdInner(cmd, dir, config)
	}
}

func RunSectionParallel(iter []string, batchSize int, config parse.T) {
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
				defer wg.Done()

				dir := "."
				if !MagicCmds(cmd, &dir, config) {
					RunCmdInner(cmd, ".", config)
				}

			}(cmd)
		}

		wg.Wait()
	}
}

// Runs an argument command, e.g. "castle build"
func RunArgCmd(names []string, cmds [][]string, logFns []func(), config parse.T) {
	for i, v := range names {
		switch config.Config.LogLevel {
		default:
			logFns[i]()
		}
		if config.Config.IsParallelCmd(v) {
			var batchSize int
			if config.Config.BatchSize == 0 {
				batchSize = -1
			} else {
				batchSize = config.Config.BatchSize
			}
			RunSectionParallel(cmds[i], batchSize, config)
		} else {
			RunSection(cmds[i], config)
		}
	}
}

func RunCmdInner(cmd string, directory string, config parse.T) {
	c := strings.Split(cmd, " ")
	shCmd := exec.Command(c[0], c[1:]...)

	switch config.Config.LogLevel {
	case "none":
	case "cmd":
		RunningCmd(cmd)
	default:
		RunningCmd(cmd)
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
