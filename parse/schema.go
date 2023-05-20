package parse

import (
	"castle/base"
	"strings"
)

type Schema struct {
	Config struct {
		Workers      int    `yaml:"workers"`
		LogLevel     string `yaml:"log_level"`
		LogShellCmds bool   `yaml:"log_shell_cmds"`
	} `yaml:"config"`
	Tasks    map[string][]string `yaml:"tasks"`
	TaskArgs map[string][][]string
}

func (s *Schema) SetUnsetDefaults() {
	if s.Config.Workers == 0 {
		s.Config.Workers = 1
	}

	if s.Config.LogLevel == "" {
		s.Config.LogLevel = "hide"
		base.Level = base.Levels[s.Config.LogLevel]
	}

	if s.Config.LogShellCmds == false {
		s.Config.LogShellCmds = true
	}
}

// The following functions process and parse the loaded YAML (that aligns with the schema).

// Convert the string of a command into a slice of strings, for every task and substep.
func (s *Schema) SchemaSetTaskArgs() {
	s.TaskArgs = make(map[string][][]string)
	for task, steps := range s.Tasks {
		s.TaskArgs[task] = make([][]string, len(steps))
		for i, subtask := range steps {
			split := strings.Split(subtask, " ")
			split = RejoinStringArgs(split)
			s.TaskArgs[task][i] = split
		}
	}
}

var (
	STRING_STARTERS = [3]string{"\"", "'", "`"} // Recognised string parenthesis (single, double, backtick)
)

// RejoinStringArgs takes a slice of strings and rejoins any strings that were split by spaces.
func RejoinStringArgs(args []string) []string {
	var output []string
	var current string
	var stringChar string
	var inString bool

	for _, arg := range args {
		if inString {
			// Exit string
			if strings.HasSuffix(arg, stringChar) {
				output = append(output, current+arg) // add string as we removed them with split
				current = ""
				inString = false
			} else { // Continue string
				current += arg + " "
			}
		} else {
			// Enter string
			if HasPrefixSlice(arg, STRING_STARTERS) {
				stringChar = string(arg[0])
				current += arg + " "
				inString = true
			} else {
				output = append(output, arg)
			}
		}
	}
	return output
}

// HasPrefixSlice checks if a string has any of the prefixes in a slice.
func HasPrefixSlice(s string, prefixes [3]string) bool {
	for _, p := range prefixes {
		if strings.HasPrefix(s, p) {
			return true
		}
	}
	return false
}
