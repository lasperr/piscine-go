package piscine

func MakeRange(min, max int) []int {
	if min < max {
		size := max - min
		answer := make([]int, size)
		for i := 0; i <= size-1; i++ {
			answer[i] = min
			min++
		}
		return answer
	}
	return nil
}
