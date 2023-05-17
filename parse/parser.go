package parse

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type ConfigObj struct {
	BatchSize     int      `yaml:"batch-size"`
	ParallelCmds  []string `yaml:"parallel"`
	ParallelTasks []string `yaml:"parallel-tasks"`
	LogLevel      string   `yaml:"log-level"`
}

func (c *ConfigObj) IsParallelTask(name string) bool {
	o := false
	for _, v := range c.ParallelTasks {
		if v == name {
			o = true
			break
		}
	}
	return o
}

func (c *ConfigObj) IsParallelCmd(name string) bool {
	o := false
	for _, v := range c.ParallelCmds {
		if v == name {
			o = true
			break
		}
	}
	return o
}

type T struct {
	Config ConfigObj           `yaml:"config"`
	Tasks  map[string][]string `yaml:"tasks"`
	Build  []string            `yaml:"build"`
	Deploy []string            `yaml:"deploy"`
	Format []string            `yaml:"format"`
	Run    []string            `yaml:"run"`
	Test   []string            `yaml:"test"`
}

func Parse(filename string) T {
	fmt.Println("Config File → ", filename)
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
