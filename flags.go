package main

import (
	"flag"
	"fmt"
	"os"
)

const version = "0.3"

func parseFlags() string {
	help := flag.Bool("h", false, "Show help")
	showVersion := flag.Bool("v", false, "Show the version of paf")

	flag.Parse()
	args := flag.Args()

	if *help { //nolint: gocritic
		showHelp()
		os.Exit(0)
	} else if *showVersion {
		fmt.Printf("paf %s\n", version)
		os.Exit(0)
	}

	if len(args) == 0 {
		fmt.Print("No argument passed\n\n")
		showHelp()
		os.Exit(1)
	} else if len(args) > 1 {
		fmt.Print("Too many argument passed\n\n")
		showHelp()
		os.Exit(1)
	}

	return args[0]
}

func showHelp() {
	fmt.Print("Usage: paf {abstract} [-h] [-v]\n\n")
	fmt.Print("Formats the passed abstract to be a valid BibTeX value.\n\n")
	fmt.Println("Optional parameter:")
	fmt.Println("  -h\tDisplays this help message and exits")
	fmt.Println("  -v\tDisplays program's version number and exists")
}
