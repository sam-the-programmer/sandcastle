package parse

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type T struct {
	Build []string `yaml:"build"`
	Run   []string `yaml:"run"`
}

func Parse(filename string) T {
	fmt.Println("File → ", filename)
	fmt.Println()

	t := T{}

	// Read the file
	f, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := yaml.Unmarshal(f, &t); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return t
}
