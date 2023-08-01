package piscine

func Sqrt(nb int) int {
	if nb == 1 {
		return 1
	}

	if nb == 2 {
		return 0
	}

	result := 1
	sqrt := 0
	for a := 1; result <= nb; a++ {
		result = a * a
		sqrt++
		if result == nb {
			return sqrt
		}
	}
	return 0
}
