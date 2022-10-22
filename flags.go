package main

import (
	"flag"
	"fmt"
	"os"
)

type config struct {
	abstract string
	help     bool
	version  bool
}

func parseFlags() (config, error) {
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
		fmt.Println("paf 0.2")
		os.Exit(0)
	}

	if len(args) == 0 {
		return config{}, fmt.Errorf("no argument passed")
	} else if len(args) > 1 {
		return config{}, fmt.Errorf("too many arguments passed")
	}

	conf.abstract = args[0]

	return conf, nil
}

func showHelp() {
	fmt.Print("Usage: paf {abstract} [-h] [-v]\n\n")
	fmt.Print("Formats the passed abstract to be a valid BibTeX value.\n\n")
	fmt.Println("Optional parameter:")
	fmt.Println("  -h\tDisplays this help message and exits")
	fmt.Println("  -v\tDisplays program's version number and exists")
}
