package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	conf, err := parseFlags()
	if err != nil {
		fmt.Printf("paf: error: %s\n", err)
		os.Exit(1)
	}

	out := format(conf.abstract)
	// Append a newline if abstract does not end with one
	if out[len(out)-1] != '\n' {
		out += "\n"
	}

	fmt.Print(out)
}

func format(abstract string) string {
	out := ""
	previous := ""

	// Replace tabs since they would otherwise break the program
	abstract = strings.ReplaceAll(abstract, "\t", "  ")

	for i := 0; i < len(abstract); i++ {
		letter := string(abstract[i])

		// Quote percentage signs
		if letter == "%" && previous != "\\" {
			out += "\\"
		}

		// Remove single newlines
		if letter == "\n" && previous != "\n" {
			if i+1 < len(abstract) {
				// We are not at the end of the string => Check if there is a newline following
				if abstract[i+1] != '\n' {
					continue
				}
			} else {
				// We are at the end of the string => No newline can follow => Current newline is single
				break
			}
		}

		// Remove repeated spaces
		if letter == " " {
			// Skip all additional spaces
			for i < len(abstract) && abstract[i] == ' ' {
				i++
			}

			// We have to move i back since it is pointing to the first non-space character
			// If we don't move it back, the outer for loop would skip the first non-space character
			i--
			letter = string(abstract[i])
		}

		previous = letter
		out += letter
	}

	return out
}
