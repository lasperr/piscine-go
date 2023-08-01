package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n > 0 {
		var array []int
		eachValue := 0

		arrayCount := 0
		var minValue int
		for n != 0 {
			eachValue = n % 10
			n /= 10
			array = append(array, eachValue)
		}

		for count := range array {
			arrayCount = count + 1
		}
		for i := 0; i < arrayCount-1; i++ {
			for j := 0; j < arrayCount-i-1; j++ {
				if array[j] > array[j+1] {
					minValue = array[j]
					array[j] = array[j+1]
					array[j+1] = minValue
				}
			}
		}
		for i := 0; i < arrayCount; i++ {
			z01.PrintRune(rune(array[i] + 48))
		}
	}
}
