package piscine

func Unmatch(a []int) int {
	for _, res := range a {
		count := 0
		for _, element := range a {
			if element == res {
				count++
			}
		}
		if count == 1 || count%2 == 1 {
			return res
		}
	}
	return -1
}
