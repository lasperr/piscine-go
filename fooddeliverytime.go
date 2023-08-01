package piscine

// Define the MenuItem struct
type MenuItem struct {
	ItemName string
	preptime int
}

func FoodDeliveryTime(order string) int {
	// Define the menu items
	menu := []MenuItem{
		{ItemName: "burger", preptime: 15},
		{ItemName: "chips", preptime: 10},
		{ItemName: "nuggets", preptime: 12},
	}

	// Split the order string into individual items
	items := splitOrder(order, ',')

	// Calculate the total cooking time
	totalTime := 0
	for _, item := range items {
		found := false
		for _, menuItem := range menu {
			if item == menuItem.ItemName {
				totalTime += menuItem.preptime
				found = true
				break
			}
		}
		if !found {
			return 404
		}
	}

	return totalTime
}

// Custom implementation of split function
func splitOrder(order string, delimiter rune) []string {
	var items []string
	currentItem := ""
	for _, char := range order {
		if char == delimiter {
			items = append(items, currentItem)
			currentItem = ""
		} else {
			currentItem += string(char)
		}
	}
	if currentItem != "" {
		items = append(items, currentItem)
	}
	return items
}

// Custom implementation of trim space function
// func trimSpace(s string) string {
// 	start := 0
// 	end := len(s) - 1
// 	for start <= end && (s[start] == ' ' || s[end] == ' ') {
// 		if s[start] == ' ' {
// 			start++
// 		}
// 		if s[end] == ' ' {
// 			end--
// 		}
// 	}
// 	return s[start : end+1]
// }

// func main() {
// 	order := "burger, chips, apple"
// 	totalTime := FoodDeliveryTime(order)
// 	if totalTime == 404 {
// 		fmt.Println(404)
// 	} else {
// 		fmt.Printf("Total cooking time for order '%s': %d minutes\n", order, totalTime)
// 	}
// }
