package utils

import (
	"castle/constants"
	"fmt"
	"os"
	"strings"
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
func Running()    { fmt.Println(constants.LIGHTCYAN, "\bRunning... 🏎️", constants.RESET) }
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
