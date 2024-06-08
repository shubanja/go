package main

import "fmt"

func main() {
	third := 1.0 / 3
	fmt.Println(third)            // Выводит: 0.3333333333333333
	fmt.Printf("%v\n", third)     // Выводит: 0.3333333333333333
	fmt.Printf("%f\n", third)     // Выводит: 0.333333
	fmt.Printf("%.3f\n", third)   // Выводит: 0.333
	fmt.Printf("%f\n", third)     // Выводит: 0.333333
	fmt.Printf("%4.4f\n", third)  // Выводит: 0.33
	fmt.Printf("%07.4f\n", third) // Выводит: 0.3333
	fmt.Printf("%6.2f\n", third)  // Выводит: 000.33
}
