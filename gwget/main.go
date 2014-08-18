// main.go
// gwget main package
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
	"gosix/gwget/fetch"
)

func Init() {
	usage := flag.Usage
	flag.Usage = func() {
		fmt.Println("gwget - automated network fetcher (lite wget clone)")
		fmt.Println("Synopsis: gwget [OPTION]... [URL]...")
		usage()
	}
}

func main() {
	var help bool
	flag.BoolVar(&help, "help", false, "Print a help message, which might be useful")
	flag.Parse()
	urls := flag.Args()

	if help {
		flag.Usage()
		return
	}

	// TODO: handle other options

	sync_chan := make(chan bool, len(urls))
	for _, url := range urls {
		// regular raw download
		go fetch.Fetch(sync_chan, url)
	}

	// sync with completed goroutines
	for i := 0; i < len(urls); i++ {
		<-sync_chan // ignore result
	}
}
