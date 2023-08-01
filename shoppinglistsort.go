package piscine

func ShoppingListSort(slice []string) []string {
	for i := 0; i < len(slice); i++ {
		if i == len(slice)-1 {
			continue
		}

		if len(slice[i]) > len(slice[i+1]) {
			slice[i], slice[i+1] = slice[i+1], slice[i]
			return ShoppingListSort(slice)
		}
	}
	return slice
}
