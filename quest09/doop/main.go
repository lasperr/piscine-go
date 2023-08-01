package main

import (
	"os"

	"github.com/01-edu/z01"
)

func CountDigits(i int) int {
	count := 0
	for i != 0 {
		i /= 10
		count = count + 1
	}
	return count
}

func IntToRunes(num int) []rune {
	count := CountDigits(num)
	p := num
	n := 0
	runes := make([]rune, count)
	for i := count - 1; i >= 0; i-- {
		n = p % 10
		p /= 10
		runes[i] = rune(n) + '0'
	}
	return runes
}

func PrintRunes(runes []rune) {
	for _, r := range runes {
		z01.PrintRune(r)
	}
}

func main() {
	arguments := os.Args
	importArgs := arguments[1:]
	validOp := []string{"+", "-", "/", "*", "%"}
	result := ""
	if len(importArgs) == 3 {
		for i := 0; i < len(validOp); i++ {
			if validOp[i] == importArgs[1] {
				if validOp[i] == "/" && importArgs[2] == "0" {
					result = "No division by 0"
				} else if validOp[i] == "%" && importArgs[2] == "0" {
					result = "No modulo by 0"
				} else if validOp[i] == "+" && importArgs[0] >= "0" && importArgs[0] <= "9" && importArgs[2] >= "0" && importArgs[2] <= "9" {
					a := 0
					for _, c := range importArgs[0] {
						a = a*10 + int(c-'0')
					}
					b := 0
					for _, c := range importArgs[2] {
						b = b*10 + int(c-'0')
					}
					resultRunes := IntToRunes(a + b)
					PrintRunes(resultRunes)
				} else if validOp[i] == "*" {
					a := 0
					for _, c := range importArgs[0] {
						a = a*10 + int(c-'0')
					}
					b := 0
					for _, c := range importArgs[2] {
						b = b*10 + int(c-'0')
					}
					resInt := a * b
					z01.PrintRune(rune(resInt))
				} else if validOp[i] == "-" {
					a := 0
					for _, c := range importArgs[0] {
						a = a*10 + int(c-'0')
					}
					b := 0
					for _, c := range importArgs[2] {
						b = b*10 + int(c-'0')
					}
					resultRunes := IntToRunes(a - b)
					PrintRunes(resultRunes)
				} else if validOp[i] == "/" {
					a := 0
					for _, c := range importArgs[0] {
						a = a*10 + int(c-'0')
					}
					b := 0
					for _, c := range importArgs[2] {
						b = b*10 + int(c-'0')
					}
					resultRunes := IntToRunes(a / b)
					PrintRunes(resultRunes)
				} else if validOp[i] == "%" {
					a := 0
					for _, c := range importArgs[0] {
						a = a*10 + int(c-'0')
					}
					b := 0
					for _, c := range importArgs[2] {
						b = b*10 + int(c-'0')
					}
					resultRunes := IntToRunes(a % b)
					PrintRunes(resultRunes)
				}
			}
		}
		z01.PrintRune('\n')
	}
}
