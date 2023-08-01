package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	trigger1 := "01"
	trigger2 := "galaxy"
	trigger3 := "galaxy 01"

	for _, element := range args {
		if element == trigger1 || element == trigger2 || element == trigger3 {
			fmt.Println("Alert!!!")
			break
		}
	}
}
