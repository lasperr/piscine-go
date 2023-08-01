package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	iter := l
	count := 0

	for iter != nil {
		if count == pos {
			return iter
		}
		iter = iter.Next
		count++
	}
	return nil
}
