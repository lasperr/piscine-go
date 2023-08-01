package piscine

func ListSize(l *List) int {
	count := 0
	if l.Head == nil {
		return count
	} else {
		current := l.Head
		count = 1
		for current.Next != nil {
			current = current.Next
			count++
		}
		return count
	}
}
