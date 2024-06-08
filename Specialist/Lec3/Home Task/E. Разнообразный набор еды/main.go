package main

import (
	"fmt"
)

func main() {
	var (
		drink   string
		meal    string
		subMeal string
		time1   int
		time2   int
	)

	fmt.Print("drink?\n")
	fmt.Scan(&drink)
	fmt.Print("meal?\n")
	fmt.Scan(&meal)
	fmt.Print("subMeal?\n")
	fmt.Scan(&subMeal)
	fmt.Print("time?\n")
	fmt.Scan(&time1, &time2)

	fmt.Printf("I wanna some '%s', my favorite meal - '%s'. Give me also '%s'. I will come soon... '%d:%.2d'", drink, meal, subMeal, time1, time2)
}
