package piscine

func LastRune(s string) rune {
	sentences := []rune(s)
	for index, letter := range sentences {
		if index == len(sentences)-1 {
			return letter
		}
	}
	return 0
}
