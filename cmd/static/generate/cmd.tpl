package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Usage = usage
	flag.Parse()

	// TODO
}

func usage() {
	fmt.Fprintf(flag.CommandLine.Output(), "{{ .Name }}\n\n")
	fmt.Fprintf(flag.CommandLine.Output(), "TODO.\n\n")
	fmt.Fprintf(flag.CommandLine.Output(), "usage: %s [path ...]\n", os.Args[0])
	flag.PrintDefaults()
}
