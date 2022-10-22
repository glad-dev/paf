package main

import (
	"fmt"
	"strings"
)

func main() {
	conf := parseFlags()
	out := format(conf.abstract)
	// Append a newline if abstract does not end with one
	if out[len(out)-1] != '\n' {
		out += "\n"
	}

	fmt.Print(out)
}

func format(abstract string) string { //nolint:gocognit
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
					// Replace single newline with a space except when previous char is a space or next char is "-"
					if previous == " " || previous == "-" {
						continue
					}

					previous = " "
					out += " "

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

			// If the previous letter and the current letter is a space, we do not want to add another space.
			// This case should only happen when a newline and a single space are combined.
			if previous == " " && letter == " " {
				continue
			}
		}

		previous = letter
		out += letter
	}

	return out
}
