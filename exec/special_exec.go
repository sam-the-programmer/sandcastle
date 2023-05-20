package exec

import (
	"castle/base"
	"os"
)

const (
	DefaultCastelFileName     = "castle.yaml"
	DefaultCastleFileContents = `config:
  workers: 1
  log_level: debug

tasks:
  build:
    - echo "Building... 🧱"
  run:
	- echo "Running... ⏲️"`
)

var (
	SpecialCommands = map[string]func([]string){
		"init": Init,
	}
)

func Init([]string) {
	if _, err := os.Stat(DefaultCastelFileName); !os.IsNotExist(err) {
		base.LogError("Castle file already exists.")
		os.Exit(1)
	}

	file, err := os.Create(DefaultCastelFileName)
	if err != nil {
		base.LogError("Failed to create castle file.")
		os.Exit(1)
	}
	defer file.Close()

	file.WriteString(DefaultCastleFileContents)
}

func IsSpecialCommand(command string) bool {
	_, ok := SpecialCommands[command]
	return ok
}
