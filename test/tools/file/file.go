package main

import (
	"fmt"
	"os"
)

func main() {
	switch os.Args[1] {
	case "build":
		fmt.Println("Built 🧱")
	case "run":
		fmt.Println("Running, running, running, running... 🏃💨")
	case "err":
		os.Stderr.Write([]byte("Oh no!"))
	default:
		fmt.Println(os.Args[1])
	}
}
