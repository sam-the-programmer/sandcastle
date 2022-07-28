package main

import (
	"castle/parse"
	"flag"
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

	VERSION = "v0.0.1"
)

var (
	filename    = flag.String("config", "castle.yaml", "Config YAML file to parse.")
	showVersion = flag.Bool("version", false, "Show version and exit.")
	shouldBuild = flag.Bool("build", false, "Build the project.")
	shouldRun   = flag.Bool("run", false, "Run the project.")
)

func init() {
	flag.StringVar(filename, "c", "castle.yml", "Config YAML file to parse.")
	flag.BoolVar(showVersion, "v", false, "Show version.")
	flag.BoolVar(shouldBuild, "b", false, "Build the project.")
	flag.BoolVar(shouldRun, "r", false, "Run the projecte.")
}

func main() {
	flag.Parse()

	if *showVersion {
		fmt.Println(LIGHTPURPLE, "\bCastle", VERSION, RESET)
		os.Exit(0)
	}

	fmt.Println(LIGHTPURPLE, "\bCastle", VERSION, RESET)
	config := parse.Parse(*filename)

	if *shouldBuild {
		fmt.Println("Building...")
		RunSection(config.Build)
		os.Exit(0)
	}

	if *shouldRun {
		fmt.Println("Running...")
		RunSection(config.Run)
		os.Exit(0)
	}

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
