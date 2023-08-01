package piscine

func FirstRune(s string) rune {
	sentences := []rune(s)
	for index, letter := range sentences {
		if index == 0 {
			return letter
		}
	}
	return 0
}
