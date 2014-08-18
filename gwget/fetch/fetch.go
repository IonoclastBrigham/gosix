// fetch.go
// gwget url fetching routines
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

package fetch

import (
	"log"
	//	"net"
	"os"
	"strings"
)

const CONN_PROTO = "tcp"

func print_err(msg string) {
	log.Fprintln(os.Stderr, msg)
}

func parse_url(url string) (proto, host, port, path string) {
	var addr string
	split := strings.Split(url, "://")
	if len(split) == 2 {
		proto = split[0]
		addr = split[1]
	} else if len(split) == 1 {
		proto = "http"
		addr = split[0]
	} else {
		print_err("Invalid url: " + url)
	}

	idx := strings.Index(addr, "/")
	if idx > 0 {
		host = addr[:idx]
		path = addr[idx:]
	} else if idx == -1 {
		host = addr
		path = "/"
	} else {
		print_err("Invalid host spec: " + addr)
	}

	idx = strings.Index(host, ":")
	if idx > 0 {
		host = host[:idx]
		port = host[idx:]
	} else if idx == -1 {
		port = proto
	} else {
		print_err("Invalid host specification: " + host)
	}

	return
}

func Fetch(sync_chan chan<- bool, url string) {
	proto, host, port, path := parse_url(url)
	log.Println(proto, host, port, path)
	sync_chan <- true
}
