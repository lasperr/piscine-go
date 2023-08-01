package piscine

func AlphaCount(s string) int {
	sentences := []rune(s)
	count := 0
	for _, letter := range sentences {
		if (letter >= 65 && letter <= 90) || (letter >= 97 && letter <= 122) {
			count++
		}
	}
	return count
}
