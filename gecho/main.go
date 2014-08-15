// main.cpp
// gecho main package
//
// Copyright © 2014 Brigham Toskin
// This software is part of the GOSIX source distribution. It is distributable
// under the terms of a modified MIT License. You should have received a copy of
// the license in the file LICENSE. If not, see:
// <http://code.google.com/p/rogue-op/wiki/LICENSE>
//
// Formatting:
//	utf-8 ; unix ; 80 cols ; tabwidth 4
////////////////////////////////////////////////////////////////////////////////

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
