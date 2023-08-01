package piscine

func Join(strs []string, sep string) string {
	length := len(strs)
	if length == 0 {
		return ""
	}
	result := strs[0]
	for i := 1; i < length; i++ {
		result += sep + strs[i]
	}
	return result
}
