package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func cut(linechan chan *string, re *regexp.Regexp, delim string, fieldindices []int) {
	for {
		line := <-linechan
		if line == nil {
			break
		}
		fields := re.Split(*line, -1)
		printed := 0
		for _, f := range fieldindices {
			if f > len(fields) {
				continue
			}
			if printed > 0 {
				fmt.Print(delim)
			}
			fmt.Print(fields[f-1])
			printed++
		}
		fmt.Print("\n")
	}
}

func main() {
	delim := flag.String("d", "\t", "Character to use as the field delimiter instead of tab")
	fieldnum := flag.String("f", "", "Required: field(s) to echo, separated by commas\n\texamples:\n\t\tgcut -f 2\n\t\tgcut -f 2,3,7")
	showhelp := flag.Bool("h", false, "Show commandline flags help.")
	flag.Parse()
	if *showhelp {
		fmt.Println("gcut v1\nExtract fields from delimited lines.\nUsage:")
		flag.PrintDefaults()
		os.Exit(0)
	}
	if len(*delim) != 1 || len(*fieldnum) < 1 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// extract the field indices the user wants
	fieldstrlist := regexp.MustCompile(",").Split(*fieldnum, -1)
	fieldindices := make([]int, len(fieldstrlist))
	for i := 0; i < len(fieldstrlist); i++ {
		fnum, err := strconv.ParseInt(fieldstrlist[i], 0, 0)	
		fieldindices[i] = int(fnum)
		if fieldindices[i] < 1 {
			fmt.Printf("Invalid field number specified: %d\n", fieldindices[i])
			os.Exit(1)
		}
		if err != nil {
			fmt.Printf("Error reading field (-f) value(s): %v\n", err)
			flag.PrintDefaults()
			os.Exit(1)
		}
	}

	re := regexp.MustCompile(*delim)
	scanner := bufio.NewScanner(os.Stdin)
	linechan := make(chan *string)
	go cut(linechan, re, *delim, fieldindices)
	for scanner.Scan() {
		line := scanner.Text()
		linechan <- &line
	}
	linechan <- nil // sync with goroutine
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error: "+ err.Error())
		os.Exit(1)
	}
}
