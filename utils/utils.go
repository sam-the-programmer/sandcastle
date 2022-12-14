package utils

import (
	"castle/constants"
	"castle/parse"
	"fmt"
	"log"
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

func Building()   { fmt.Println(constants.LIGHTBLUE, "\bBuilding... ๐จ", constants.RESET) }
func Running()    { fmt.Println(constants.LIGHTCYAN, "\bRunning... ๐๏ธ", constants.RESET) }
func Formatting() { fmt.Println(constants.LIGHTGREEN, "\bFormatting... ๐จ", constants.RESET) }
func Testing()    { fmt.Println(constants.LIGHTYELLOW, "\bTesting... ๐งช", constants.RESET) }
func Deploying()  { fmt.Println(constants.LIGHTRED, "\bDeploying... ๐", constants.RESET) }

func Task(name string) {
	fmt.Println(constants.LIGHTRED, "\bTask:", name, "... ๐", constants.RESET)
}

func RunningCmd(cmd any) {
	fmt.Println(constants.LIGHTGREY, "\bโ", constants.RESET, cmd)
}

func UnknownCmd(v any) {
	fmt.Println(constants.LIGHTRED, "Unknown command:", v, constants.RESET)
}

func RunTask(config parse.T, taskName string) {
	Task(taskName)

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

		RunCmdInner(cmd, dir)
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
					RunCmdInner(cmd, ".")
				}

			}(cmd)
		}

		wg.Wait()
	}
}

// Runs an argument command, e.g. "castle build"
func RunArgCmd(names []string, cmds [][]string, logFns []func(), config parse.T) {
	for i, v := range names {
		logFns[i]()
		if config.Config.IsParallelCmd(v) {
			log.Println("Running in parallel")
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

func RunCmdInner(cmd string, directory string) {
	RunningCmd(cmd)

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
