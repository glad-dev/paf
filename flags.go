package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const version = "0.2.4"

type config struct {
	abstract       string
	help           bool
	version        bool
	checkForUpdate bool
}

func parseFlags() config {
	conf := config{
		abstract:       "",
		help:           false,
		version:        false,
		checkForUpdate: false,
	}

	flag.BoolVar(&conf.help, "h", false, "Show help")
	flag.BoolVar(&conf.version, "v", false, "Show the version of paf")
	flag.BoolVar(&conf.checkForUpdate, "u", false, "Check if there is a newer version available")

	flag.Parse()
	args := flag.Args()

	if conf.help { //nolint: gocritic
		showHelp()
		os.Exit(0)
	} else if conf.version {
		fmt.Printf("paf %s\n", version)
		os.Exit(0)
	} else if conf.checkForUpdate {
		checkForUpdate()
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

	conf.abstract = args[0]

	return conf
}

func showHelp() {
	fmt.Print("Usage: paf {abstract} [-h] [-v] [-u]\n\n")
	fmt.Print("Formats the passed abstract to be a valid BibTeX value.\n\n")
	fmt.Println("Optional parameter:")
	fmt.Println("  -h\tDisplays this help message and exits")
	fmt.Println("  -v\tDisplays program's version number and exists")
	fmt.Println("  -u\tChecks if there is a newer version available")
}

func checkForUpdate() {
	type responseStruct struct {
		Name string `json:"name"`
		// The other fields are of no interest
	}

	fmt.Println("Checking for updates")

	// Since we do not need any custom headers, we can use http.Get instead of http.NewRequest
	res, err := http.Get("https://api.github.com/repos/glad-dev/paf/tags") //nolint: noctx
	if err != nil {
		fmt.Printf("Could not connect to the server: %s\n", err)
		os.Exit(1)
	} else if res.StatusCode > 299 { //nolint: gomnd
		fmt.Printf("Response failed with status code %d\n", res.StatusCode)
		os.Exit(1)
	}

	body, err := io.ReadAll(res.Body)
	_ = res.Body.Close()
	if err != nil { //nolint: wsl
		fmt.Printf("Could not read the response: %s\n", err)
		os.Exit(1)
	}

	var responseStructArray []responseStruct

	err = json.Unmarshal(body, &responseStructArray)
	if err != nil {
		fmt.Printf("Could not unmarshal the response: %s\n", err)
		os.Exit(1)
	}

	if len(responseStructArray) == 0 {
		fmt.Println("No tags were found")
		os.Exit(1)
	}

	// The loop implicitly relies on the latest tag being listed first.
	for _, elem := range responseStructArray {
		if strings.HasPrefix(elem.Name, "v") { //nolint: wsl
			// This will match any tags that start with "v", not just version tags.
			// However, using a RegEx seems like overkill.

			if fmt.Sprintf("v%s", version) == elem.Name {
				fmt.Println("The program is up to date")
			} else {
				fmt.Println("The program is out of date")
				fmt.Printf("You are using version %s, while the lastes version is %s\n", version, elem.Name[1:])
				fmt.Println("Run 'go install github.com/glad-dev/paf@latest'")
			}

			os.Exit(0)
		}
	}

	// No tag starts with "v"
	fmt.Println("No version tag was found.")
	fmt.Println("Please file an issue at https://github.com/glad-dev/paf")
	os.Exit(1)
}
