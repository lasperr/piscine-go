package main

import "github.com/01-edu/z01"

func main() {
	var digits string = "0123456789"
	for i := 0; i < len(digits); i++ {
		z01.PrintRune(rune(digits[i]))
	}
	z01.PrintRune('\n')
}
