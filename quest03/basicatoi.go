package piscine

func BasicAtoi(s string) int {
	counter := 0
	number := 0
	sChangeable := []rune(s)
	for _, num := range sChangeable {
		for i := '0'; i < num; i++ {
			counter++
		}
		number = number*10 + counter
		counter = 0
	}
	return number
}
