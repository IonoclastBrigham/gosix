// main.cpp
// genv main package
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
	"flag"
	"fmt"
	"os"
	//	"os/exec"
	"sort"
	"strings"
)

func cli_err(msg string) {
	fmt.Fprintln(os.Stderr, "Error: ", msg)
	flag.Usage()
	os.Exit(1)
}

type SymTab map[string]string

// this has a potential data race, accessing the env map from a goroutine.
// right now, only one goroutine accesses it, so it's not an issue.
func define(defsIn <-chan string, env SymTab) {
	for def := range defsIn {
		kv := strings.Split(def, "=")
		if len(kv) != 2 {
			msg := fmt.Sprintf("Invalid variable definition (%s)", def)
			cli_err(msg)
		}
		env[kv[0]] = kv[1]
	}
}

func init() {
	usage := flag.Usage
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "genv - print, modify, and export environment variables.")
		usage()
		fmt.Fprintln(os.Stderr, `Examples:
  genv -i CC=/usr/bin/clang GOPATH=/home/userx/go go build
  genv EDITOR=emacs git commit -a
  #!/<GOPATH>/bin/genv lua -l strict # at the top of a script
  genv # prints environment variables`)
	}
}

func main() {
	ignore := flag.Bool("i", false, "Causes genv to completely ignore the environment it inherits.")
	help_short := flag.Bool("h", false, "Display usage message.")
	help_long := flag.Bool("help", false, "Display usage message.")
	flag.Parse()
	args := flag.Args()

	if *help_short || *help_long {
		flag.Usage()
		return
	}

	defsChan := make(chan string)
	env := make(SymTab)
	go define(defsChan, env)

	// environment vars
	if !*ignore {
		for _, def := range os.Environ() {
			defsChan <- def
		}
		//		fmt.Printf("%d vars in environment:\n", i)
	}

	// vars defined on command line
	for i, arg := range args {
		if !strings.Contains(arg, "=") {
			// slice off tail after var=val pairs
			args = args[i:]
			break
		}
		defsChan <- arg
	}
	close(defsChan)

	// list or exec
	if len(args) == 0 { // null case: print environment
		// grab a slice of sorted keys so we can print sorted
		names := make([]string, len(env))
		i := 0
		for n := range env {
			names[i] = n
			i++
		}
		sort.Strings(names)
		for _, n := range names {
			fmt.Printf("%s=%s\n", n, env[n])
		}
	} else { // remaining args: exec util with any args

		// TODO
	}
}
