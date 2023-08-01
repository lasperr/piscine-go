package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := []rune(os.Args[0])[2:]
	for _, r := range args {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
