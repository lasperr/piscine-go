package piscine

func NRune(s string, n int) rune {
	sentences := []rune(s)
	for index, letter := range sentences {
		if index == n-1 {
			return letter
		}
	}
	return 0
}
