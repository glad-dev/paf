package main

import (
	"strings"
	"testing"
)

func TestRemoveRepeatedSpaces(t *testing.T) {
	valueMap := map[string]string{
		// Single space
		" ": " ",
		// Only spaces
		"     ": " ",
		// Tab
		"\t": " ",
		// Tab and space at the end
		"\t ": " ",
		// Tab and space at the beginning
		" \t": " ",
		// Tab and spaces at the end and beginning
		" \t ": " ",
		// Tab and chars
		"A\tB":     "A B",
		"\tA B":    " A B",
		"\tA\tB":   " A B",
		"A B\t":    "A B ",
		"A\tB\t":   "A B ",
		"\tA\tB\t": " A B ",
		// Spaces in the middle
		"A      B                 C D        E": "A B C D E",
		// Spaces at the beginning
		"       A B": " A B",
		// Spaces at the end
		"A B              ": "A B ",
		// Spaces at the beginning and end
		"              A B              ": " A B ",
		// Single space at the end
		"A B ": "A B ",
		// Single space at the beginning
		" A B": " A B",
		// Single space at the beginning and end
		" A B ": " A B ",
		// Single newline (which should be replaced with a space) and a single space
		" \nA":  " A",
		"A\n ":  "A ",
		"A\n B": "A B",
		"A \nB": "A B",
	}

	for got, want := range valueMap {
		got = format(got)

		if got != want {
			t.Errorf("got: '%s', wanted '%s'", got, want)
		}
	}
}

func TestRemoveNewlineAndQuotePercentage(t *testing.T) {
	valueMap := map[string]string{
		"%\n":   "\\%",
		"\\\n%": "\\ \\%",
		"\n\\%": " \\%",
		"":      "",
	}

	for got, want := range valueMap {
		got = format(got)

		if got != want {
			t.Errorf("got: '%s', wanted '%s'", got, want)
		}
	}
}

func TestQuotePercentage(t *testing.T) {
	// Map with structure: Input => Expected output
	valueMap := map[string]string{
		"%":                           "\\%",
		"%%%%%":                       "\\%\\%\\%\\%\\%",
		"\\%":                         "\\%",
		"%\\%":                        "\\%\\%",
		"\\%%":                        "\\%\\%",
		"%\\%%":                       "\\%\\%\\%",
		"This contains an unquoted %": "This contains an unquoted \\%",
		"This contains an unquoted % and quoted \\%": "This contains an unquoted \\% and quoted \\%",
	}

	for got, want := range valueMap {
		got = format(got)

		if got != want {
			t.Errorf("got: '%s', wanted '%s'", got, want)
		}
	}
}

func TestRemoveSingleNewlineButKeepDouble(t *testing.T) {
	valueMap := map[string]string{
		"A\n\nB":    "A\n\nB",
		"\nA\n\nB":  " A\n\nB",
		"A\n\nB\n":  "A\n\nB",
		"A\n\nB\nC": "A\n\nB C",
	}

	for got, want := range valueMap {
		got = format(got)

		if got != want {
			t.Errorf("got: '%s', wanted '%s'", got, want)
		}
	}
}

func TestRemoveSingleNewline(t *testing.T) {
	inputs := []string{
		"This\ncontains\nmany\nnewlines",
		"\nThis started with a newline",
		"This ends with a newline\n",
		"\nThis starts and ends with a newline\n",
		"\n",
	}

	for _, val := range inputs {
		// Replace newlines with spaces
		wanted := strings.ReplaceAll(val, "\n", " ")
		// Remove trailing space
		wanted = strings.TrimSuffix(wanted, " ")

		got := format(val)

		if wanted != got {
			t.Errorf("got: '%s', wanted '%s'", got, wanted)
		}
	}
}

func TestNoChange(t *testing.T) {
	inputs := []string{
		"This line requires no modifications",
		"This line contains an escaped percentage sign: \\%",
		"This line contains two newlines (\n\n)",
		"\n\nThis\n\nline\n\ncontains\n\nnewlines\n\ninstead\n\nof\n\nspaces\n\n",
		"\n\n",
		" ",
		"",
	}

	for _, got := range inputs {
		formatted := format(got)
		if formatted != got {
			t.Errorf("got '%s', wanted '%s'", formatted, got)
		}
	}
}
