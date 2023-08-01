package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	fmt.Printf("Player 1: ")
	fmt.Printf("%d, ", deck[0])
	fmt.Printf("%d, ", deck[1])
	fmt.Printf("%d\n", deck[2])

	fmt.Printf("Player 2: ")
	fmt.Printf("%d, ", deck[3])
	fmt.Printf("%d, ", deck[4])
	fmt.Printf("%d\n", deck[5])

	fmt.Printf("Player 3: ")
	fmt.Printf("%d, ", deck[6])
	fmt.Printf("%d, ", deck[7])
	fmt.Printf("%d\n", deck[8])

	fmt.Printf("Player 4: ")
	fmt.Printf("%d, ", deck[9])
	fmt.Printf("%d, ", deck[10])
	fmt.Printf("%d\n", deck[11])
}
