package piscine

func Any(f func(string) bool, a []string) bool {
	for _, element := range a {
		if f(element) {
			return true
		}
	}
	return false
}
