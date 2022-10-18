package main

import (
	"flag"
	"fmt"
)

func parseFlags() (string, error) {
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		return "", fmt.Errorf("no argument passed")
	} else if len(args) > 1 {
		return "", fmt.Errorf("too many arguments passed")
	}

	return args[0], nil
}
