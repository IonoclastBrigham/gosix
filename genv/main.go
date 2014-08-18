// main.go
// genv main package
//
// Copyright © 2014 Brigham Toskin
// This software is part of the GOSIX source distribution. It is distributable
// under the terms of a modified MIT License. You should have received a copy of
// the license in the file LICENSE. If not, see:
// <https://github.com/IonoclastBrigham/gosix/blob/master/LICENSE>
//
// Formatting:
//	utf-8 ; unix ; 80 cols ; tabwidth 4
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sort"
	"strings"
)

func cli_err(msg string) {
	fmt.Fprintln(os.Stderr, "Error: ", msg)
	flag.Usage()
	os.Exit(1)
}

type SymTab map[string]string

func enmap(kv_strings []string, kv_map SymTab) {
	for _, def := range kv_strings {
		kv := strings.Split(def, "=")
		if len(kv) != 2 {
			msg := fmt.Sprintf("Invalid variable definition (%s)", def)
			cli_err(msg)
		}
		kv_map[kv[0]] = kv[1]
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

	// environment vars
	var env []string
	kv_map := make(SymTab)
	if !*ignore {
		env = os.Environ()
	}
	for len(args) > 0 {
		arg := args[0]
		if strings.Index(arg, "=") < 1 {
			break // not of form "name=val", preserve tail and bail
		}
		args = args[1:] // pop front
		env = append(env, arg)
	}
	enmap(env, kv_map)
	env = []string{}
	for k, v := range kv_map {
		env = append(env, fmt.Sprintf("%s=%s", k, v))
	}
	sort.Strings(env)

	// list or exec
	if len(args) == 0 { // null case: print environment
		for _, v := range env {
			fmt.Println(v)
		}
	} else { // remaining args: exec util with any args
		util := exec.Command(args[0])
		util.Args = args
		util.Stdin = os.Stdin
		util.Stdout = os.Stdout
		util.Stderr = os.Stderr
		util.Env = env
		err := util.Run()
		if err != nil {
			log.Fatal(err)
		}
	}
}
