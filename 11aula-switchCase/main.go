package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Quando é sábado? ")
	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		fmt.Println("É hoje")
	case today + 1:
		fmt.Println("É amanhã")
	case today + 2:
		fmt.Println("É depois de amanhã")
	default:
		fmt.Println("Tá Longe ainda...")
	}

}
