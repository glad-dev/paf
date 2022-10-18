# paper-abstract-formatter
Formats the passed abstract to be a valid BibTeX value.

## Usage
```shell
$ paf "An exemplary


abstract %"
```

## Installation
1. Clone repo
2. Enter directory (`cd ./paper-abstract-formatter`)
3. Run `go install`

## Applied formatting:
- Escaping of unescaped percentage signs (`%`)
- Removal of tabs and repeated spaces
- Removal of single line breaks, while keeping double line breaks
