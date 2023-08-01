package piscine

func ToLower(s string) string {
	sString := []byte(s)

	for i, element := range sString {
		if element >= 65 && element <= 90 {
			sString[i] = element + 32
		}
	}
	return string(sString)
}
