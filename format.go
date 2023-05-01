package main

import (
	"strings"
)

func format(abstract string) string {
	abstract = strings.TrimSpace(abstract)
	abstract = strings.ReplaceAll(abstract, "\t", " ")

	abstract = quotePercentageSigns(abstract)
	abstract = removeSingleNewline(abstract)
	abstract = condenseSpaces(abstract)

	return abstract
}

func quotePercentageSigns(abstract string) string {
	s := strings.Builder{}

	for i, letter := range abstract {
		if letter == '%' {
			if i == 0 || abstract[i-1] != '\\' {
				s.WriteRune('\\')

			}
		}

		s.WriteRune(letter)
	}

	return s.String()
}

func removeSingleNewline(abstract string) string {
	s := strings.Builder{}

	for i, letter := range abstract {
		// Remove single newlines
		if letter == '\n' && abstract[i-1] != '\n' {
			if i == len(abstract)-1 {
				// We are at the end of the string => No newline can follow => Current newline is single
				break
			}

			// Check if there is a newline following
			if abstract[i+1] != '\n' {
				// Replace single newline with a space except when previous char is a space or next char is "-"
				if abstract[i-1] == ' ' || abstract[i-1] == '-' {
					continue
				}

				s.WriteRune(' ')
				continue
			}
		}

		s.WriteRune(letter)
	}

	return s.String()
}

func condenseSpaces(abstract string) string {
	s := strings.Builder{}

	split := strings.Split(abstract, "") // Splits strings into UTF-8 chars
	for i := 0; i < len(split); i++ {
		if split[i] == " " {
			// Skip all following spaces
			for i < len(split) && split[i] == " " {
				i++
			}

			// Write the condensed space
			s.WriteString(" ")
		}

		s.WriteString(split[i])
	}

	return s.String()
}
