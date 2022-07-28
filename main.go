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
	LIGHTGREY   = "\033[1;37m"
	LIGHTPURPLE = "\033[1;35m"
	LIGHTRED    = "\033[1;31m"
	LIGHTYELLOW = "\033[1;33m"
	RESET       = "\033[0m"

	VERSION = "v0.1.0"
)

var (
	filename    = flag.String("config", "castle.yaml", "Config YAML file to parse.")
	showVersion = flag.Bool("version", false, "Show version and exit.")
	shouldBuild = flag.Bool("build", false, "Build the project.")
	shouldRun   = flag.Bool("run", false, "Run the project.")
	shouldTest  = flag.Bool("test", false, "test the project.")
)

func init() {
	flag.StringVar(filename, "c", "castle.yml", "Config YAML file to parse.")
	flag.BoolVar(showVersion, "v", false, "Show version.")
	flag.BoolVar(shouldBuild, "b", false, "Build the project.")
	flag.BoolVar(shouldRun, "r", false, "Run the project.")
	flag.BoolVar(shouldTest, "t", false, "Test the project.")
}

func building() { fmt.Println(LIGHTBLUE, "\bBuilding... 🔨", RESET) }
func running()  { fmt.Println(LIGHTCYAN, "\bRunning... 🚀", RESET) }
func testing()  { fmt.Println(LIGHTYELLOW, "\bTesting... 🧪", RESET) }
func none() bool {
	return !*shouldBuild && !*shouldRun && !*shouldTest && !*showVersion
}

func main() {
	flag.Parse()

	fmt.Println(LIGHTPURPLE, "\bCastle", VERSION, RESET)

	if *showVersion {
		os.Exit(0)
	}

	config := parse.Parse(*filename)
	if none() {
		// Remove the -c and -config and following argument from os.Args
		for i, v := range os.Args {
			if v == "-c" || v == "-config" {
				os.Args = append(os.Args[:i], os.Args[i+2:]...)
			}
		}

		if os.Args[1] == "task" {
			fmt.Println(LIGHTRED, "\bTask:", os.Args[2], "... 📝", RESET)
			RunSection(config.Tasks[os.Args[2]])
		}
	}

	if *shouldBuild {
		building()
		RunSection(config.Build)
	}

	if *shouldRun {
		running()
		RunSection(config.Run)
	}

	if *shouldTest {
		testing()
		RunSection(config.Test)
	}

	if !*shouldBuild && !*shouldRun && !*shouldTest {
		building()
		RunSection(config.Build)
		running()
		RunSection(config.Run)
	}

	os.Exit(0)
}

func RunSection(iter []string) {
	for _, cmd := range iter {
		fmt.Println(LIGHTGREY, "\b→", RESET, cmd)

		c := strings.Split(cmd, " ")
		cmd := exec.Command(c[0], c[1:]...)

		out, err := cmd.CombinedOutput()
		fmt.Println(string(out))

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	}
}
