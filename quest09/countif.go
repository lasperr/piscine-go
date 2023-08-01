package piscine

func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for _, element := range tab {
		if f(element) {
			count++
		}
	}
	return count
}
