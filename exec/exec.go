package exec

import (
	"castle/base"
	"os"
	"os/exec"
	"strings"
)

var (
	HasShell    bool
	ChosenShell string
	ValidShells = []string{"bash", "zsh", "sh", "csh", "ksh", "tcsh", "dash", "fish"}
)

func init() {
	for _, v := range ValidShells {
		if _, err := exec.LookPath(v); err == nil {
			HasShell = true
			ChosenShell = v
			break
		}
	}
}

func Execute(command string, args []string, fullCmd []string, logShellCmds bool) {
	if IsShellCommand(command) {
		if logShellCmds {
			base.LogRun(fullCmd)
		}
		ExecuteShellCommand(command, args)
	} else {
		base.LogRun(fullCmd)
		ExecuteCommand(command, args)
	}
}

func ExecuteCommand(command string, args []string) {
	cmd := exec.Command(command, args...)
	constructAndRunCmd(cmd)
}

func ExecuteShellCommand(command string, args []string) {
	if !HasShell {
		base.LogExit("No shell found. Valid shells are " + strings.Join(ValidShells, ", "))
	}

	cmd := exec.Command(ChosenShell, "-c", command+" "+strings.Join(args, " "))
	constructAndRunCmd(cmd)
}

func constructAndRunCmd(cmd *exec.Cmd) {
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		base.LogPanic(err)
	}
}

func IsShellCommand(command string) bool {
	_, err := exec.LookPath(command)
	if err != nil {
		return true
	}
	return false
}

func ExecuteTask(task [][]string, taskName string, logShellCmds bool) {
	base.LogTaskStart(taskName)
	for _, command := range task {
		Execute(command[0], command[1:], command, logShellCmds)
	}
}
