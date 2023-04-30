package main

import (
	"fmt"
)

func main() {
	abstract := parseFlags()

	fmt.Println(format(abstract))
}
