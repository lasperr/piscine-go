package piscine

func ReverseMenuIndex(menu []string) []string {
	reverse := make([]string, len(menu))

	for i, element := range menu {
		reverse[len(menu)-i-1] = element
	}
	return reverse
}
