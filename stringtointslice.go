package piscine

func StringToIntSlice(str string) []int {
	var arrayOfInt []int

	for _, elements := range str {
		arrayOfInt = append(arrayOfInt, int(elements))
	}
	return arrayOfInt
}
