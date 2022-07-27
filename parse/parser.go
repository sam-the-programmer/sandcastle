package parse

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type T struct {
	Build []string `yaml:"build"`
	Run   []string `yaml:"run"`
}

func Parse(filename string) T {
	t := T{}

	// Read the file
	f, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	if err := yaml.Unmarshal(f, &t); err != nil {
		log.Fatal(err)
	}

	return t
}
