package main

import (
	"fmt"
	"os"
)

func main() {
	newline := true
	args := os.Args[1:]
	if len(args) < 1 {
		return
	}
	if args[0] == "-n" {
		newline = false
		args = args[1:]
	}
	for _, arg := range args {
		fmt.Printf("%s ", arg)
	}
	if newline {
		fmt.Printf("%s", "\n")
	}
}
