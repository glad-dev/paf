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

	continueAt := 0 // We can't use a for-i loop since that would break unicode chars
	for i, letter := range abstract {
		if i < continueAt {
			continue
		}

		if abstract[i] == ' ' {
			// Skip all following spaces
			for i < len(abstract) && abstract[i] == ' ' {
				i++
			}

			// Write the condensed space
			s.WriteRune(' ')

			continueAt = i
			continue
		}

		s.WriteRune(letter)
	}

	return s.String()
}
