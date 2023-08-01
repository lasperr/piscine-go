package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	arrayStr := []rune(s)

	for i := 0; i < len(arrayStr); i++ {
		z01.PrintRune(arrayStr[i])
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	lenArgs := 0

	args := os.Args
	for index := range args {
		lenArgs = index
	}

	if isEven(lenArgs) == true {
		printStr("I have an even number of arguments")
	} else {
		printStr("I have an odd number of arguments")
	}
}
