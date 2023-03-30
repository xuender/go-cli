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
	fmt.Fprintf(os.Stderr, "{{ .Name }}\n\n")
	fmt.Fprintf(os.Stderr, "TODO.\n\n")
	fmt.Fprintf(os.Stderr, "Usage: %s [flags]\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}
