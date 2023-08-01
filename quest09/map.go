package piscine

func Map(f func(int) bool, a []int) []bool {
	var arrBool []bool
	for _, element := range a {
		arrBool = append(arrBool, f(element))
	}
	return arrBool
}
