package main

import (
	"castle/parse"
	"fmt"
	"os"
	"os/exec"

	"strings"
)

const (
	LIGHTBLUE   = "\033[1;34m"
	LIGHTCYAN   = "\033[1;36m"
	LIGHTPURPLE = "\033[1;35m"
	RESET       = "\033[0m"
)

func main() {
	fmt.Println(LIGHTPURPLE, "\bCastle v0.0.1\n", RESET)

	config := parse.Parse(os.Args[1])

	fmt.Println(LIGHTBLUE, "\bBuilding... 🔨", RESET)
	RunSection(config.Build)
	fmt.Println(LIGHTCYAN, "\bRunning... 🚀", RESET)
	RunSection(config.Run)
}

func RunSection(iter []string) {
	for _, cmd := range iter {
		fmt.Println("→ ", cmd)

		c := strings.Split(cmd, " ")
		d := strings.Join(c[1:], " ")
		cmd := exec.Command(c[0], d)

		out, err := cmd.CombinedOutput()
		fmt.Println(string(out))

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	}
}
