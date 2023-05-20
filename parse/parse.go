package parse

import (
	"castle/base"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func Parse(filename string) Schema {
	file := open(filename)

	var schema Schema
	decoder := yaml.NewDecoder(file)
	base.LogInfo("Parsing config file.")
	err := decoder.Decode(&schema)
	if err != nil {
		base.LogError("Failed to parse config file.")
		log.Fatal(err)
	}

	schema.SetUnsetDefaults()
	schema.SchemaSetTaskArgs()
	return schema
}

func open(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		base.LogError("Failed to open config file.")
		log.Fatal(err)
	}
	return file
}
