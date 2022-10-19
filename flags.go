package main

import (
	"flag"
	"fmt"
	"os"
)

type config struct {
	abstract string
	help     bool
}

func parseFlags() (config, error) {
	conf := config{
		abstract: "",
		help:     false,
	}

	flag.BoolVar(&conf.help, "h", false, "Show help")

	flag.Parse()
	args := flag.Args()

	if conf.help {
		fmt.Print("Usage: paf {abstract}\n\n")
		fmt.Println("Formats the passed abstract to be a valid BibTeX value.")
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
