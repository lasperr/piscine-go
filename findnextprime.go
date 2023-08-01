package piscine

func IsPrime(nb int) bool {
	if nb > 0 {
		if nb == 1 {
			return false
		}
		if nb == 2 || nb == 3 {
			return true
		}
		if nb%3 == 0 || nb%2 == 0 {
			return false
		}
		for divider := 4; divider < nb; divider++ {
			if nb%divider == 0 {
				return false
			}
		}
		return true
	}
	return false
}

func FindNextPrime(nb int) int {
	if IsPrime(nb) {
		return nb
	} else {
		n := nb + 1
		for IsPrime(n) == false {
			n++
		}
		return n
	}
}
