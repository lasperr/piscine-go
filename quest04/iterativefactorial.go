package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 32 {
		return 0
	} else if nb == 0 {
		return 1
	}
	iterate := nb - 1
	for i := iterate; i >= 1; i-- {
		nb = nb * i
	}
	return nb
}
