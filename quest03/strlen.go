package piscine

func StrLen(s string) int {
	length := 0

	sStringChangeable := []rune(s)

	for i := 0; i < len(sStringChangeable); i++ {
		length = length + 1
	}
	return length
}
