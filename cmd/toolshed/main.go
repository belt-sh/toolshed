package main

import (
	"flag"
	"fmt"
	"os"
)

const helpText = `usage: toolshed [options]

`

var (
	version = flag.Bool("version", false, "")
)

func usage() {
	fmt.Fprintf(os.Stderr, helpText)
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if *version {
		fmt.Fprintf(os.Stdout, "toolshed %s (%s)\n", Version, Commit)
		os.Exit(0)
	}
}