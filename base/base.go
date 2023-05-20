package base

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	BLUE        = "\033[34m"
	CYAN        = "\033[36m"
	GREEN       = "\033[32m"
	GREY        = "\033[37m"
	LIGHTBLUE   = "\033[1;34m"
	LIGHTCYAN   = "\033[1;36m"
	LIGHTGREEN  = "\033[1;32m"
	LIGHTGREY   = "\033[1;37m"
	LIGHTPURPLE = "\033[1;35m"
	LIGHTRED    = "\033[1;31m"
	LIGHTYELLOW = "\033[1;33m"
	PURPLE      = "\033[35m"
	RED         = "\033[31m"
	RESET       = "\033[0m"
	YELLOW      = "\033[33m"

	VERSION = "v0.2.0"
)

var (
	Levels = map[string]int8{
		"debug":   0,
		"info":    1,
		"warning": 2,
		"error":   3,
		"none":    4,
	}

	Level = Levels["info"]
)

func Colourize(color string, msg string) string {
	return color + msg + RESET
}

func Timestamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func LogError(msg string) {
	if Level <= Levels["error"] {
		fmt.Println(Timestamp(), Colourize(LIGHTRED, "[ Error ]"), Colourize(LIGHTGREY, msg))
	}
}

func LogInfo(msg string) {
	if Level <= Levels["info"] {
		fmt.Println(Timestamp(), Colourize(LIGHTCYAN, "[ Info ]"), Colourize(LIGHTGREY, msg))
	}
}

func LogSuccess(msg string) {
	if Level <= Levels["info"] {
		fmt.Println(Timestamp(), Colourize(LIGHTGREEN, "[ Success ]"), Colourize(LIGHTGREY, msg))
	}
}

func LogWarning(msg string) {
	if Level <= Levels["warning"] {
		fmt.Println(Timestamp(), Colourize(LIGHTYELLOW, "[ Warning ]"), Colourize(LIGHTGREY, msg))
	}
}

func LogRun(command []string) {
	fmt.Println(
		Colourize(LIGHTGREY, "→ "),
		strings.Join(command, " "),
	)
}

func LogTaskStart(task string) {
	fmt.Println(
		Colourize(LIGHTBLUE, "Task: "+strings.ToUpper(string(task[0]))+task[1:]),
	)
}

// Log an error from an error and exit
func LogPanic(err error) {
	// if Level <= Levels["error"] {
	fmt.Println(Colourize(LIGHTRED, "Error: "), Colourize(LIGHTGREY, err.Error()))
	// }
	os.Exit(1)
}

// Log an error from a string and exit
func LogExit(msg string) {
	fmt.Println(Colourize(LIGHTRED, "Error: "), Colourize(LIGHTGREY, msg))
	os.Exit(1)
}
