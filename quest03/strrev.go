package piscine

func StrRev(s string) string {
	var result string
	for _, v := range s {
		result = string(v) + result
	}
	return result
}
