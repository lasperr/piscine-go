package piscine

func ToUpper(s string) string {
	sString := []byte(s)

	for i, element := range sString {
		if element >= 97 && element <= 122 {
			sString[i] = element - 32
		}
	}
	return string(sString)
}
