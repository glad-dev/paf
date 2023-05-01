package main

import (
	"testing"
)

func TestRemoveRepeatedSpaces(t *testing.T) {
	valueMap := map[string]string{
		// Single space
		" ": "",
		// Only spaces
		"     ": "",
		// Tab
		"\t": "",
		// Tab and space at the end
		"\t ": "",
		// Tab and space at the beginning
		" \t": "",
		// Tab and spaces at the end and beginning
		" \t ": "",
		// Tab and chars
		"A\tB":      "A B",
		"\tA B":     "A B",
		"\tA\tB":    "A B",
		"A B\t":     "A B",
		"A\tB\t":    "A B",
		"\tA\tB\tC": "A B C",
		// Spaces in the middle
		"A      B                 C D        E": "A B C D E",
		// Spaces at the beginning
		"C       A B": "C A B",
		// Spaces at the end
		"A B              C": "A B C",
		// Spaces at the beginning and end
		"A             B C             D": "A B C D",
		// Single space at the end
		"A B ": "A B",
		// Single space at the beginning
		" A B": "A B",
		// Single space at the beginning and end
		" A B ": "A B",
		// Single newline (which should be replaced with a space) and a single space
		" \nA":  "A",
		"A\n ":  "A",
		"A\n B": "A B",
		"A \nB": "A B",
		// Single newline should not break hyphenation
		"A-\nB": "A-B",
	}

	for input, expected := range valueMap {
		got := format(input)

		if got != expected {
			t.Errorf("got: '%s', wanted '%s'", got, expected)
		}
	}
}

func TestRemoveNewlineAndQuotePercentage(t *testing.T) {
	valueMap := map[string]string{
		"%\n":   "\\%",
		"\\\n%": "\\ \\%",
		"\n\\%": "\\%",
		"":      "",
	}

	for input, expected := range valueMap {
		got := format(input)

		if got != expected {
			t.Errorf("got: '%s', wanted '%s'", got, expected)
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

	for input, expected := range valueMap {
		got := format(input)

		if got != expected {
			t.Errorf("got: '%s', wanted '%s'", got, expected)
		}
	}
}

func TestRemoveSingleNewlineButKeepDouble(t *testing.T) {
	valueMap := map[string]string{
		"A\n\nB":      "A\n\nB",
		"\nA\n\nB":    "A\n\nB",
		"A\n\nB\n":    "A\n\nB",
		"A\n\nB\nC":   "A\n\nB C",
		"A\n\n-B":     "A\n\n-B",
		"A\n\nB\n\nC": "A\n\nB\n\nC",
		"A\nB\n\nC":   "A B\n\nC",
	}

	for input, expected := range valueMap {
		got := format(input)

		if got != expected {
			t.Errorf("got: '%s', wanted '%s'", got, expected)
		}
	}
}

func TestRemoveSingleNewline(t *testing.T) {
	valueMap := map[string]string{
		"This\ncontains\nmany\nnewlines":          "This contains many newlines",
		"\nThis started with a newline":           "This started with a newline",
		"This ends with a newline\n":              "This ends with a newline",
		"\nThis starts and ends with a newline\n": "This starts and ends with a newline",
		"\n": "",
	}

	for input, expected := range valueMap {
		got := format(input)

		if got != expected {
			t.Errorf("got: '%s', wanted '%s'", got, expected)
		}
	}
}

func TestNoChange(t *testing.T) {
	inputs := []string{
		"This line requires no modifications",
		"This line contains an escaped percentage sign: \\%",
		"This line contains two newlines (\n\n)",
		"This\n\nline\n\ncontains\n\nnewlines\n\ninstead\n\nof\n\nspaces",
		"",
	}

	for _, input := range inputs {
		got := format(input)

		if input != got {
			t.Errorf("got '%s', wanted '%s'", got, input)
		}
	}
}

func TestUnicode(t *testing.T) {
	inputs := []string{
		"“",
		"×",
		"🤒",
		"♫",
	}

	for _, input := range inputs {
		got := format(input)

		if input != got {
			t.Errorf("got '%s', wanted '%s'", got, input)
		}
	}
}

func TestUnicodeAndText(t *testing.T) {
	valueMap := map[string]string{
		"“Some quoted text“":                            "“Some quoted text“",
		"“Quote with %":                                 "“Quote with \\%",
		"Unicode before 圓%":                             "Unicode before 圓\\%",
		"Unicode before spaces Ģ                     A": "Unicode before spaces Ģ A",
		"Unicode after spaces                        Ħ": "Unicode after spaces Ħ",
	}

	for input, expected := range valueMap {
		got := format(input)

		if got != expected {
			t.Errorf("got: '%s', wanted '%s'", got, expected)
		}
	}
}
