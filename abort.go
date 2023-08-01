package piscine

func Abort(a, b, c, d, e int) int {
	array := []int{a, b, c, d, e}
	count := 5
	temp := 0
	for i := 0; i < count-1; i++ {
		temp = i
		for j := i + 1; j < count; j++ {
			if array[j] < array[temp] {
				temp = j
			}
		}
		array[i], array[temp] = array[temp], array[i]
	}
	return array[2]
}
