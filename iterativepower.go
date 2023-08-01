package piscine

func IterativePower(nb int, power int) int {
	result := 1

	if power < 0 || power > 32 {
		return 0
	}

	if power == 0 {
		return 1
	}

	for i := 0; i <= power-1; i++ {
		result = result * nb
	}
	return result
}
