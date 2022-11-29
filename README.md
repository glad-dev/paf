# paper-abstract-formatter
Formats the passed abstract to be a valid BibTeX value.

## Usage
```shell
$ paf "An exemplary


abstract %"
```

## Installation
Run `go install github.com/glad-dev/paper-abstract-formatter'`.

## Applied formatting:
- Escaping of unescaped percentage signs (`%`)
- Removal of tabs and repeated spaces
- Removal of single line breaks, while keeping double line breaks
