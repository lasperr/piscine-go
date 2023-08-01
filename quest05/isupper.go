package piscine

func IsUpper(s string) bool {
	for _, letter := range s {
		if letter < 65 || letter > 90 {
			return false
		}
	}
	return true
}
