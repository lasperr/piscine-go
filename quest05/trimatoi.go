package piscine

func TrimAtoi(s string) int {
	i := 0
	for _, let := range s {
		if let >= '0' && let <= '9' {
			if let == '0' {
				i += 0
				i *= 10
			} else if let == '1' {
				i += 1
				i *= 10
			} else if let == '2' {
				i += 2
				i *= 10
			} else if let == '3' {
				i += 3
				i *= 10
			} else if let == '4' {
				i += 4
				i *= 10
			} else if let == '5' {
				i += 5
				i *= 10
			} else if let == '6' {
				i += 6
				i *= 10
			} else if let == '7' {
				i += 7
				i *= 10
			} else if let == '8' {
				i += 8
				i *= 10
			} else if let == '9' {
				i += 9
				i *= 10
			}
		}
	}
	num := i / 10
	mind := 100
	lin := 0
	for index, let := range s {
		if let == '-' {
			mind = index
		}
		if let == '1' || let == '2' || let == '3' || let == '4' || let == '5' || let == '6' || let == '7' || let == '8' || let == '9' {
			lin = index
			if index > mind {
				return -num
			} else if lin > 0 {
				return num
			}
		}
	}
	return num
}
