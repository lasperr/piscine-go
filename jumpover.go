package piscine

func JumpOver(str string) string {
	word := ""
	for i := 1; i <= len(str); i++ {
		if i%3 == 0 {
			word = word + string(str[i-1])
		}
	}
	word = word + "\n"
	return word
}
