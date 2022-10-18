# paper-abstract-formatter
Formats the passed abstract to be a valid BibTeX value.

## Usage
```shell
$ paf "An exemplary


abstract %"
```

## Installation
1. Build the program with `go build -o paf`
2. Copy the binary to a `bin` directory, e.g. `~/.local/bin/`

## Applied formatting:
- Escaping of unescaped percentage signs (`%`)
- Removal of tabs and repeated spaces
- Removal of single line breaks, while keeping double line breaks