package utils

import (
	"castle/constants"
	"castle/parse"
	"fmt"
	"os"
	"strings"
)

func MagicCmds(cmd string, dir string, config parse.T) (string, bool) {
	if strings.HasPrefix(cmd, "SETDIR! ") {
		dir = cmd[8:]
	} else if strings.HasPrefix(cmd, "GETDIR!") {
		dir = "."
		cwd, _ := os.Getwd()
		fmt.Println(constants.LIGHTGREY, "\bCurrent directory:", cwd, constants.RESET)
	} else if strings.HasPrefix(cmd, "ECHO! ") {
		fmt.Println(constants.LIGHTGREY, "\b"+cmd[6:], constants.RESET)
	} else if strings.HasPrefix(cmd, "SET! ") {
		split := strings.Split(cmd[5:], " ")
		if len(split) != 2 {
			fmt.Println(constants.LIGHTRED, "\bInvalid SET! command. Expected SET! <key> <value>", constants.RESET)
			os.Exit(1)
		}

		os.Setenv(split[0], split[1])
	} else if strings.HasPrefix(cmd, "GET! ") {
		val := cmd[5:]
		if len(val) < 1 {
			fmt.Println(constants.LIGHTRED, "\bInvalid GET! command. Expected GET! <key>", constants.RESET)
			os.Exit(1)
		}

		fmt.Println(constants.LIGHTGREY, "\b"+val+": ", os.Getenv(val), constants.RESET)
	} else if strings.HasPrefix(cmd, "TASK! ") {
		v := cmd[6:]
		if len(v) < 1 {
			fmt.Println(constants.LIGHTRED, "\bInvalid TASK! command. Expected TASK! <task>", constants.RESET)
			os.Exit(1)
		}

		RunTask(config, v)
	} else if strings.HasPrefix(cmd, "EXIT!") {
		os.Exit(0)
	} else {
		return dir, false
	}
	return dir, true
}
