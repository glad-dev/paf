package main

import (
	"flag"
	"fmt"
	"os"
)

const version = "0.2.1"

type config struct {
	abstract string
	help     bool
	version  bool
}

func parseFlags() config {
	conf := config{
		abstract: "",
		help:     false,
		version:  false,
	}

	flag.BoolVar(&conf.help, "h", false, "Show help")
	flag.BoolVar(&conf.version, "v", false, "Show the version of paf")

	flag.Parse()
	args := flag.Args()

	if conf.help {
		showHelp()
		os.Exit(0)
	} else if conf.version {
		fmt.Printf("paf %s\n", version)
		os.Exit(0)
	}

	if len(args) == 0 {
		fmt.Println("paf: error: no argument passed")
		showHelp()
		os.Exit(1)
	} else if len(args) > 1 {
		fmt.Println("paf: error: too many argument passed")
		showHelp()
		os.Exit(1)
	}

	conf.abstract = args[0]

	return conf
}

func showHelp() {
	fmt.Print("Usage: paf {abstract} [-h] [-v]\n\n")
	fmt.Print("Formats the passed abstract to be a valid BibTeX value.\n\n")
	fmt.Println("Optional parameter:")
	fmt.Println("  -h\tDisplays this help message and exits")
	fmt.Println("  -v\tDisplays program's version number and exists")
}
