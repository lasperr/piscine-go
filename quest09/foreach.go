package piscine

func ForEach(f func(int), a []int) {
	for _, res := range a {
		f(res)
	}
}
