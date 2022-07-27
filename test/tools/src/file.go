package main

import (
	"fmt"
	"os"
)

func main() {
	if os.Args[1] == "build" {
		fmt.Println("Built 🧱")
	} else if os.Args[1] == "run" {
		fmt.Println("Running, running, running, running... 🏃‍♀️💨")
	} else {
		fmt.Println(os.Args[1])
	}
}
